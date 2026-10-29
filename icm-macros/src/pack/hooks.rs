//! This module contains the tools the pack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.

use reforge::PreprocessingData;
use solar::ast::Span;
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, SourceId, Struct};

use crate::common::{LengthType, Unvalidated, ValidationState};

pub struct PackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Validated>>,
    pub visibility: String,
}

/// How the length/count of a dynamically sized field is encoded.
#[derive(Debug, Clone, Copy, PartialEq)]
pub enum LengthSpec {
    /// Prefix the payload with its length/count, encoded as this Solidity integer type.
    Type(LengthType),
    /// Omit the length/count prefix entirely. The receiver must know the length
    /// (pairs with `#[unpack(length = constant)]`).
    Drop,
}

impl Default for LengthSpec {
    /// The default is a `uint256` prefix.
    fn default() -> Self {
        Self::Type(LengthType::default())
    }
}

/// State marker: field args have been validated; `length` holds a `LengthSpec`.
#[derive(Debug, Clone, PartialEq)]
pub struct Validated;

impl ValidationState for Validated {
    type Length = LengthSpec;
}

/// Per-field pack annotations, generic over their validation state.
///
/// - `FieldArgs<Unvalidated>`: raw parse output; `length` is `Option<String>`.
/// - `FieldArgs<Validated>`:   HIR-resolved;    `length` is `Option<LengthSpec>`.
#[derive(Debug, Default, Clone, PartialEq)]
pub struct FieldArgs<S: ValidationState = Unvalidated> {
    pub method: Option<String>,
    pub ignore: bool,
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
                if ty == "drop" {
                    return Ok(LengthSpec::Drop);
                }
                LengthType::validate(&ty)
                    .map(LengthSpec::Type)
                    .wrap_err(format!(
                        "length type `{ty}` for field `{field_name}` of `{struct_name}` \
                         must be a Solidity integer type or `drop`"
                    ))
            })
            .transpose()?;
        Ok(FieldArgs {
            method: self.method,
            ignore: self.ignore,
            length,
        })
    }
}

/// Intermediate parsed args before HIR resolution. Contract is stored as a name
/// string rather than a ContractId.
pub(crate) struct RawPackArgs {
    pub contract: Option<String>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Unvalidated>>,
    pub visibility: String,
}

impl RawPackArgs {
    /// Validates the raw field args against the HIR struct definition (if any), consuming
    /// `self.fields` and returning the validated equivalents.
    pub fn validate_fields(
        &mut self,
        ctx: &Gcx,
        struct_def: Option<&Struct<'_>>,
    ) -> eyre::Result<Vec<FieldArgs<Validated>>> {
        let struct_name = struct_def.map(|s| s.name.as_str()).unwrap_or("");
        let hir_fields: Vec<_> = struct_def
            .map(|s| s.fields.iter().map(|&id| ctx.hir.variable(id)).collect())
            .unwrap_or_default();
        let fields = std::mem::take(&mut self.fields);
        fields
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

/// Fetches the comment preceding the given HIR item and parses any `#[pack(...)]`
/// annotation from it. For structs, also fetches and parses `#[pack(...)]` annotations
/// on each field. Resolves the optional `contract` name to a `ContractId` via the HIR.
/// Returns `None` if no annotation is present on the item itself, or
/// `Some(Err(...))` if a `contract` name was specified but not found in the HIR.
pub fn parse_args(
    ctx: &Gcx,
    source_id: SourceId,
    span: Span,
    data: &mut PreprocessingData<'_>,
) -> Option<eyre::Result<PackArgs>> {
    let doc_comment = reforge::get_comment(ctx, source_id, span, data)?;
    let struct_def = ctx.hir.structs().find(|s| s.span == span);

    let field_comments: Vec<Option<String>> = struct_def
        .map(|s| {
            s.fields
                .iter()
                .map(|&field_id| {
                    reforge::get_comment(ctx, source_id, ctx.hir.variable(field_id).span, data)
                })
                .collect()
        })
        .unwrap_or_default();

    let mut raw = parse_comment(&doc_comment, &field_comments)?;

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
                    "contract `{name}` specified in #[pack(contract=...)] was not found"
                )));
            }
        },
    };

    let fields = match raw.validate_fields(ctx, struct_def) {
        Ok(f) => f,
        Err(e) => return Some(Err(e)),
    };

    Some(Ok(PackArgs {
        contract,
        name: raw.name,
        fields,
        visibility: raw.visibility,
    }))
}

/// Pure string-parsing layer. Extracts pack args from a pre-fetched comment string
/// and a list of pre-fetched field comment strings (one per field, in order).
/// Returns None if no #[pack(...)] annotation is present.
pub(crate) fn parse_comment(
    comment: &str,
    field_comments: &[Option<String>],
) -> Option<RawPackArgs> {
    let pack_start = comment.find("#[pack(")?;
    let args_start = pack_start + "#[pack(".len();
    let args_end = args_start + comment[args_start..].find(')')?;
    let args_str = &comment[args_start..args_end];

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();

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
        }
    }

    let fields = field_comments
        .iter()
        .map(|c| parse_field_args(c.as_deref()))
        .collect();

    Some(RawPackArgs {
        contract,
        name,
        fields,
        visibility,
    })
}

fn parse_field_args(comment: Option<&str>) -> FieldArgs<Unvalidated> {
    let Some(comment) = comment else {
        return FieldArgs::default();
    };
    let Some(pack_start) = comment.find("#[pack(") else {
        return FieldArgs::default();
    };
    let args_start = pack_start + "#[pack(".len();
    let Some(args_end_rel) = comment[args_start..].find(')') else {
        return FieldArgs::default();
    };
    let args_str = comment[args_start..args_start + args_end_rel].trim();

    if args_str == "ignore" {
        return FieldArgs {
            method: None,
            ignore: true,
            length: None,
        };
    }

    let mut method = None;
    let mut ignore = false;
    let mut length = None;
    for arg in args_str.split(',') {
        if let Some((key, val)) = arg.trim().split_once('=') {
            match key.trim() {
                "method" => method = Some(val.trim().trim_matches('"').to_string()),
                "length" => length = Some(val.trim().trim_matches('"').to_string()),
                _ => {}
            }
        } else if arg.trim() == "ignore" {
            ignore = true;
        }
    }
    if ignore {
        FieldArgs {
            method: None,
            ignore: true,
            length: None,
        }
    } else {
        FieldArgs {
            method,
            ignore: false,
            length,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_basic() {
        let raw = parse_comment("#[pack]", &vec![None, None, None, None]);
        assert!(raw.is_none());
        let raw = parse_comment("#[pack()]", &vec![None, None, None, None]).expect("Test failed");
        assert_eq!(raw.contract, None);
        assert_eq!(raw.name, None);
        assert_eq!(raw.visibility, "public".to_string());
        let expected = vec![FieldArgs::<Unvalidated>::default(); 4];
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_all_keys() {
        let raw = parse_comment(
            "#[pack(contract=\"Foo\", name=\"Bar\", visibility=\"private\")]",
            &vec![None, None, None, None],
        )
        .unwrap();
        assert_eq!(raw.contract, Some("Foo".to_string()));
        assert_eq!(raw.name, Some("Bar".to_string()));
        assert_eq!(raw.visibility, "private".to_string());
        let expected = vec![FieldArgs::<Unvalidated>::default(); 4];
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_length_quoted() {
        // length = "uint32" (with quotes) must behave identically to length = uint32
        let raw = parse_comment(
            "#[pack()]",
            &vec![
                Some("#[pack(length = \"uint32\")]".to_string()),
                Some("#[pack(method=\"Foo.bar\", length = \"uint64\")]".to_string()),
            ],
        )
        .unwrap();
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: None,
                ignore: false,
                length: Some("uint32".to_string()),
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                ignore: false,
                length: Some("uint64".to_string()),
            }
        );
    }

    #[test]
    fn test_field_args() {
        let raw = parse_comment(
            "#[pack(visibility=\"private\", name=\"Foo\")]",
            &vec![
                Some("#[pack(method=\"Foo.bar\")]".to_string()),
                Some("#[pack(ignore)]".to_string()),
                Some("#[pack(ignore, method=\"Foo.bar\")]".to_string()),
                Some("#[pack(length = uint32)]".to_string()),
                Some("#[pack(method=\"Foo.bar\", length = uint64)]".to_string()),
                Some("#[pack(length = drop)]".to_string()),
            ],
        )
        .unwrap();
        assert_eq!(raw.name, Some("Foo".to_string()));
        assert_eq!(raw.visibility, "private".to_string());
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                ignore: false,
                length: None,
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: None,
                ignore: true,
                length: None
            }
        );
        assert_eq!(
            raw.fields[2],
            FieldArgs {
                method: None,
                ignore: true,
                length: None
            }
        );
        assert_eq!(
            raw.fields[3],
            FieldArgs {
                method: None,
                ignore: false,
                length: Some("uint32".to_string()),
            }
        );
        assert_eq!(
            raw.fields[4],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                ignore: false,
                length: Some("uint64".to_string()),
            }
        );
        assert_eq!(
            raw.fields[5],
            FieldArgs {
                method: None,
                ignore: false,
                length: Some("drop".to_string()),
            }
        );
    }

    #[test]
    fn test_validate_length() {
        let validate = |length: &str| {
            FieldArgs::<Unvalidated> {
                method: None,
                ignore: false,
                length: Some(length.to_string()),
            }
            .validate("data", "Foo")
            .map(|f| f.length)
        };
        assert_eq!(
            validate("uint32").unwrap(),
            Some(LengthSpec::Type(LengthType::Uint(4)))
        );
        assert_eq!(validate("drop").unwrap(), Some(LengthSpec::Drop));
        let err = validate("bool").unwrap_err().to_string();
        assert!(
            err.contains("length type `bool` for field `data` of `Foo`"),
            "{err}"
        );
        assert!(validate("32").is_err(), "a constant is not a pack length");
    }
}
