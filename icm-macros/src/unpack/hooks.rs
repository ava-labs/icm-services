//! This module contains the tools the unpack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.

use reforge::PreprocessingData;
use solar::ast::{ElementaryType, Span};
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, ItemId, SourceId, Type, TypeKind};

pub struct UnpackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs>,
    pub visibility: String,
    pub calldata: bool,
    pub solhint_disable: bool,
}

#[derive(Debug, Default, Clone, PartialEq)]
pub struct FieldArgs {
    pub method: Option<String>,
    pub default: bool,
    pub memory: bool,
}

/// Intermediate parsed args before HIR resolution. Contract is stored as a name
/// string rather than a ContractId.
pub(crate) struct RawUnpackArgs {
    pub contract: Option<String>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs>,
    pub visibility: String,
    pub calldata: bool,
    pub solhint_disable: bool,
}

/// Fetches the comment preceding the given HIR item and parses any `#[unpack(...)]`
/// annotation from it. For structs, also fetches and parses `#[unpack(...)]` annotations
/// on each field. Resolves the optional `contract` name to a `ContractId` via the HIR.
/// Returns `None` if no annotation is present on the item itself, or
/// `Some(Err(...))` if a `contract` name was specified but not found in the HIR.
pub fn parse_args(
    ctx: &Gcx,
    source_id: SourceId,
    span: Span,
    data: &mut PreprocessingData<'_>,
) -> Option<eyre::Result<UnpackArgs>> {
    let doc_comment = reforge::get_comment(ctx, source_id, span, data)?;

    let field_meta: Vec<(Option<String>, bool)> = ctx
        .hir
        .structs()
        .find(|s| s.span == span)
        .map(|s| {
            s.fields
                .iter()
                .map(|&field_id| {
                    let field = ctx.hir.variable(field_id);
                    let comment = reforge::get_comment(ctx, source_id, field.span, data);
                    let memory = type_requires_memory(&field.ty);
                    (comment, memory)
                })
                .collect()
        })
        .unwrap_or_default();

    let raw = parse_comment(&doc_comment, &field_meta)?;

    let contract = match raw.contract.as_deref() {
        None => None,
        Some(name) => {
            match ctx
                .hir
                .contracts_enumerated()
                .find(|(_, c)| c.name.as_str() == name)
                .map(|(id, _)| id)
            {
                Some(id) => Some(id),
                None => {
                    return Some(Err(eyre::eyre!(
                        "contract `{name}` specified in #[unpack(contract=...)] was not found"
                    )));
                }
            }
        }
    };

    Some(Ok(UnpackArgs {
        contract,
        name: raw.name,
        fields: raw.fields,
        visibility: raw.visibility,
        calldata: raw.calldata,
        solhint_disable: raw.solhint_disable,
    }))
}

/// Pure string-parsing layer. Extracts pack args from a pre-fetched comment string
/// and a list of pre-fetched field comment strings (one per field, in order).
/// Returns None if no #[pack(...)] annotation is present.
pub(crate) fn parse_comment(
    comment: &str,
    field_meta: &[(Option<String>, bool)],
) -> Option<RawUnpackArgs> {
    let pack_start = comment.find("#[unpack(")?;
    let args_start = pack_start + "#[unpack(".len();
    let args_end = args_start + comment[args_start..].find(')')?;
    let args_str = &comment[args_start..args_end];

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();
    let mut calldata = false;
    let mut solhint_disable = false;

    for pair in args_str.split(',') {
        let pair = pair.trim();
        if let Some((key, val)) = pair.split_once('=') {
            let val = val.trim().trim_matches('"');
            match key.trim() {
                "contract" => contract = Some(val.to_string()),
                "name" => name = Some(val.to_string()),
                "visibility" => visibility = val.to_string(),
                _ => {}
            }
        } else {
            match pair {
                "calldata" => calldata = true,
                "solhint-disable" => solhint_disable = true,
                _ => {}
            }
        }
    }

    let fields = field_meta
        .iter()
        .map(|c| parse_field_args(c.0.as_deref(), c.1))
        .collect();

    Some(RawUnpackArgs {
        contract,
        name,
        fields,
        visibility,
        calldata,
        solhint_disable,
    })
}

fn type_requires_memory(ty: &Type) -> bool {
    match &ty.kind {
        TypeKind::Elementary(e) => matches!(e, ElementaryType::Bytes | ElementaryType::String),
        TypeKind::Array(_) => true,
        TypeKind::Custom(ItemId::Struct(_)) => true,
        _ => false,
    }
}

fn parse_field_args(comment: Option<&str>, memory: bool) -> FieldArgs {
    let Some(comment) = comment else {
        return FieldArgs {
            method: None,
            default: false,
            memory,
        };
    };
    let Some(pack_start) = comment.find("#[unpack(") else {
        return FieldArgs {
            method: None,
            default: false,
            memory,
        };
    };
    let args_start = pack_start + "#[unpack(".len();
    let Some(args_end_rel) = comment[args_start..].find(')') else {
        return FieldArgs {
            method: None,
            default: false,
            memory,
        };
    };
    let args_str = comment[args_start..args_start + args_end_rel].trim();

    if args_str == "default" {
        return FieldArgs {
            method: None,
            default: true,
            memory,
        };
    }

    let mut method = None;
    let mut default = false;
    for arg in args_str.split(',') {
        if let Some((key, val)) = arg.trim().split_once('=') {
            if key.trim() == "method" {
                method = Some(val.trim().trim_matches('"').to_string());
            }
        } else if arg.trim() == "default" {
            default = true;
        }
    }
    if default {
        FieldArgs {
            method: None,
            default: true,
            memory,
        }
    } else {
        FieldArgs {
            method,
            default: false,
            memory,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_basic() {
        let field_meta = vec![(None, false), (None, true), (None, false), (None, false)];
        let raw = parse_comment("#[unpack]", &field_meta);
        assert!(raw.is_none());
        let raw = parse_comment("#[unpack()]", &field_meta).expect("Test failed");
        assert_eq!(raw.contract, None);
        assert_eq!(raw.name, None);
        assert_eq!(raw.visibility, "public".to_string());
        assert!(!raw.calldata);
        assert!(!raw.solhint_disable);
        let mut expected = vec![FieldArgs::default(); 4];
        expected[1].memory = true;
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_all_keys() {
        let raw = parse_comment(
            "#[unpack(contract=\"Foo\", name=\"Bar\", visibility=\"private\", calldata)]",
            &vec![(None, false); 4],
        )
        .unwrap();
        assert_eq!(raw.contract, Some("Foo".to_string()));
        assert_eq!(raw.name, Some("Bar".to_string()));
        assert_eq!(raw.visibility, "private".to_string());
        assert!(raw.calldata);
        assert!(!raw.solhint_disable);
        let expected = vec![FieldArgs::default(); 4];
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_flags() {
        let raw = parse_comment("#[unpack(solhint-disable)]", &vec![(None, false); 2]).unwrap();
        assert!(raw.solhint_disable);
        assert!(!raw.calldata);

        let raw = parse_comment(
            "#[unpack(calldata, solhint-disable)]",
            &vec![(None, false); 2],
        )
        .unwrap();
        assert!(raw.calldata);
        assert!(raw.solhint_disable);
    }

    #[test]
    fn test_field_args() {
        let raw = parse_comment(
            "#[unpack(visibility=\"private\", name=\"Foo\")]",
            &vec![
                (Some("#[unpack(method=\"Foo.bar\")]".to_string()), false),
                (Some("#[unpack(default)]".to_string()), true),
                (
                    Some("#[unpack(default, method=\"Foo.bar\")]".to_string()),
                    false,
                ),
            ],
        )
        .unwrap();
        assert_eq!(raw.name, Some("Foo".to_string()));
        assert_eq!(raw.visibility, "private".to_string());
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                default: false,
                memory: false,
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: None,
                default: true,
                memory: true,
            }
        );
        assert_eq!(
            raw.fields[2],
            FieldArgs {
                method: None,
                default: true,
                memory: false,
            }
        );
    }
}
