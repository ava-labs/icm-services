//! This module contains the tools the unpack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.

use reforge::PreprocessingData;
use solar::ast::{ElementaryType, Span};
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, ItemId, SourceId, Struct, Type, TypeKind};

use crate::common::{LengthType, Unvalidated, Validated, ValidationState};

pub struct UnpackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Validated>>,
    pub visibility: String,
    pub calldata: bool,
    pub solhint_disable: bool,
}

/// Per-field unpack annotations, generic over their validation state.
///
/// - `FieldArgs<Unvalidated>`: raw parse output; `length` is `Option<String>`.
/// - `FieldArgs<Validated>`:   HIR-resolved;    `length` is `Option<LengthType>`.
#[derive(Debug, Default, Clone, PartialEq)]
pub struct FieldArgs<S: ValidationState = Unvalidated> {
    pub method: Option<String>,
    pub default: bool,
    pub memory: bool,
    pub length: Option<S::Length>,
}

impl FieldArgs<Unvalidated> {
    /// Validates the `length` string (if present), converting this
    /// `FieldArgs<Unvalidated>` into a `FieldArgs<Validated>`.
    pub fn validate(
        self,
        field_name: &str,
        struct_name: &str,
    ) -> eyre::Result<FieldArgs<Validated>> {
        use eyre::WrapErr;
        let length = self
            .length
            .map(|ty| {
                LengthType::validate(&ty).wrap_err(format!(
                    "length type `{ty}` for field `{field_name}` of `{struct_name}` \
                     must be a Solidity integer type"
                ))
            })
            .transpose()?;
        Ok(FieldArgs {
            method: self.method,
            default: self.default,
            memory: self.memory,
            length,
        })
    }
}

/// Intermediate parsed args before HIR resolution. Contract is stored as a name
/// string rather than a ContractId.
pub(crate) struct RawUnpackArgs {
    pub contract: Option<String>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Unvalidated>>,
    pub visibility: String,
    pub calldata: bool,
    pub solhint_disable: bool,
}

impl RawUnpackArgs {
    /// Validates the raw field args against the HIR struct definition (if any), consuming
    /// `self.fields` and returning the validated equivalents.
    pub fn validate_fields(
        self,
        ctx: &Gcx,
        struct_def: Option<&Struct<'_>>,
    ) -> eyre::Result<Vec<FieldArgs<Validated>>> {
        let struct_name = struct_def.map(|s| s.name.as_str()).unwrap_or("");
        let hir_fields: Vec<_> = struct_def
            .map(|s| s.fields.iter().map(|&id| ctx.hir.variable(id)).collect())
            .unwrap_or_default();
        self.fields
            .into_iter()
            .zip(hir_fields.iter())
            .map(|(raw_field, hir_field)| {
                let field_name = hir_field
                    .name
                    .map(|n| n.as_str().to_string())
                    .unwrap_or_default();
                raw_field.validate(&field_name, struct_name)
            })
            .collect()
    }
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
    let struct_def = ctx.hir.structs().find(|s| s.span == span);

    let field_meta: Vec<(Option<String>, bool)> = struct_def
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
        Some(name) => match ctx
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
        },
    };

    let calldata = raw.calldata;
    let solhint_disable = raw.solhint_disable;
    let name = raw.name.clone();
    let visibility = raw.visibility.clone();

    let fields = match raw.validate_fields(ctx, struct_def) {
        Ok(f) => f,
        Err(e) => return Some(Err(e)),
    };

    Some(Ok(UnpackArgs {
        contract,
        name,
        fields,
        visibility,
        calldata,
        solhint_disable,
    }))
}

/// Pure string-parsing layer. Extracts unpack args from a pre-fetched comment string
/// and a list of pre-fetched field comment strings (one per field, in order).
/// Returns None if no #[unpack(...)] annotation is present.
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

fn parse_field_args(comment: Option<&str>, memory: bool) -> FieldArgs<Unvalidated> {
    let Some(comment) = comment else {
        return FieldArgs {
            memory,
            ..Default::default()
        };
    };
    let Some(pack_start) = comment.find("#[unpack(") else {
        return FieldArgs {
            memory,
            ..Default::default()
        };
    };
    let args_start = pack_start + "#[unpack(".len();
    let Some(args_end_rel) = comment[args_start..].find(')') else {
        return FieldArgs {
            memory,
            ..Default::default()
        };
    };
    let args_str = comment[args_start..args_start + args_end_rel].trim();

    if args_str == "default" {
        return FieldArgs {
            default: true,
            memory,
            ..Default::default()
        };
    }

    let mut method = None;
    let mut default = false;
    let mut length = None;
    for arg in args_str.split(',') {
        if let Some((key, val)) = arg.trim().split_once('=') {
            match key.trim() {
                "method" => method = Some(val.trim().trim_matches('"').to_string()),
                "length" => length = Some(val.trim().trim_matches('"').to_string()),
                _ => {}
            }
        } else if arg.trim() == "default" {
            default = true;
        }
    }
    if default {
        FieldArgs {
            default: true,
            memory,
            ..Default::default()
        }
    } else {
        FieldArgs {
            method,
            default: false,
            memory,
            length,
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
        let mut expected = vec![FieldArgs::<Unvalidated>::default(); 4];
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
        let expected = vec![FieldArgs::<Unvalidated>::default(); 4];
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
    fn test_length_quoted() {
        // length = "uint32" (with quotes) must behave identically to length = uint32
        let raw = parse_comment(
            "#[unpack()]",
            &vec![
                (Some("#[unpack(length = \"uint32\")]".to_string()), false),
                (
                    Some("#[unpack(method=\"Foo.bar\", length = \"uint64\")]".to_string()),
                    false,
                ),
            ],
        )
        .unwrap();
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: None,
                default: false,
                memory: false,
                length: Some("uint32".to_string()),
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                default: false,
                memory: false,
                length: Some("uint64".to_string()),
            }
        );
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
                (Some("#[unpack(length = uint32)]".to_string()), true),
                (
                    Some("#[unpack(method=\"Foo.bar\", length = uint64)]".to_string()),
                    true,
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
                length: None,
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: None,
                default: true,
                memory: true,
                length: None,
            }
        );
        assert_eq!(
            raw.fields[2],
            FieldArgs {
                method: None,
                default: true,
                memory: false,
                length: None,
            }
        );
        assert_eq!(
            raw.fields[3],
            FieldArgs {
                method: None,
                default: false,
                memory: true,
                length: Some("uint32".to_string()),
            }
        );
        assert_eq!(
            raw.fields[4],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                default: false,
                memory: true,
                length: Some("uint64".to_string()),
            }
        );
    }
}
