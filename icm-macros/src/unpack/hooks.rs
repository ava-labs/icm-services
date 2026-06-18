//! This module contains the tools the unpack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.
//!
//! ## Assert closure bodies
//!
//! The `assert = "|var| { body }"` annotation value uses a restricted mini-language.
//! Each `assert = "..."` key produces exactly one [`AssertDef`]. Use multiple `assert`
//! attributes for multiple postconditions. Only boolean expressions are valid — the
//! parser does not understand Solidity statement structure, so:
//!
//! - **Variable declarations and assignments** are not detected at parse time and will
//!   produce a Solidity compiler error downstream.
//! - **Nested control flow** (`if`, `for`, `while`, block scopes) is not supported.
//! - **Multi-line expressions** are supported: text lines are flattened with spaces,
//!   so line breaks are purely cosmetic.

use reforge::PreprocessingData;
use solar::ast::{ElementaryType, Span};
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, ItemId, SourceId, Struct, Type, TypeKind};

use crate::common::{
    AssertDef, CommentVisitor, LengthSpec, Unvalidated, Validated, ValidationState,
    annotation_args, find_closing_paren, join_segments, parse_closure, split_args,
};

pub struct UnpackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Validated>>,
    pub visibility: String,
    pub calldata: bool,
    /// Type-level post-condition. The capture name is replaced with `result`
    /// (the decoded struct) in the assertion body.
    pub assert: Vec<AssertDef>,
}

/// Per-field unpack annotations, generic over their validation state.
///
/// - `FieldArgs<Unvalidated>`: raw parse output; `length` is `Option<String>`.
/// - `FieldArgs<Validated>`:   HIR-resolved;    `length` is `Option<LengthSpec>`.
#[derive(Debug, Default, Clone, PartialEq)]
pub struct FieldArgs<S: ValidationState = Unvalidated> {
    pub method: Option<String>,
    pub default: bool,
    pub memory: bool,
    pub length: Option<S::Length>,
    /// Field-level post-conditions.
    pub assert: Vec<AssertDef>,
}

impl FieldArgs<Unvalidated> {
    /// Resolves the `length` string (if present) into a [`LengthSpec`], converting this
    /// `FieldArgs<Unvalidated>` into a `FieldArgs<Validated>`.
    pub fn validate(
        self,
        field_name: &str,
        struct_name: &str,
    ) -> eyre::Result<FieldArgs<Validated>> {
        if self.default && !self.assert.is_empty() {
            eyre::bail!(
                "Cannot use `assert` on field `{field_name}` of `{struct_name}` with `default`: \
                 default fields are skipped during unpacking"
            );
        }
        let length = self
            .length
            .map(|value| LengthSpec::validate(value, field_name, struct_name))
            .transpose()?;
        Ok(FieldArgs {
            method: self.method,
            default: self.default,
            memory: self.memory,
            length,
            assert: self.assert,
        })
    }
}

/// Intermediate parsed args before HIR resolution. Contract is stored as a name
/// string rather than a ContractId.
#[derive(Debug)]
pub(crate) struct RawUnpackArgs {
    pub contract: Option<String>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Unvalidated>>,
    pub visibility: String,
    pub calldata: bool,
    pub assert: Vec<AssertDef>,
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
/// Returns `None` if no annotation is present on the item itself, or `Some(Err(...))` if an
/// annotation is malformed, a field's `length` is invalid, or a `contract` name was specified
/// but not found in the HIR.
pub fn parse_args(
    ctx: &Gcx,
    source_id: SourceId,
    span: Span,
    data: &mut PreprocessingData<'_>,
) -> Option<eyre::Result<UnpackArgs>> {
    let doc_comment = reforge::get_comment(ctx, source_id, span, data)?;
    let struct_def = ctx.hir.structs().find(|s| s.span == span);

    let field_meta: Vec<(Option<String>, bool, String)> = struct_def
        .map(|s| {
            s.fields
                .iter()
                .map(|&field_id| {
                    let field = ctx.hir.variable(field_id);
                    let comment = reforge::get_comment(ctx, source_id, field.span, data);
                    let memory = type_requires_memory(&field.ty);
                    let name = field
                        .name
                        .map(|n| n.as_str().to_string())
                        .unwrap_or_default();
                    (comment, memory, name)
                })
                .collect()
        })
        .unwrap_or_default();
    let struct_name = struct_def.map(|s| s.name.as_str()).unwrap_or("");

    let mut raw = match parse_comment(&doc_comment, &field_meta, struct_name) {
        Ok(Some(raw)) => raw,
        Ok(None) => return None,
        Err(e) => return Some(Err(e)),
    };

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
    let name = raw.name.clone();
    let visibility = raw.visibility.clone();
    let assert = std::mem::take(&mut raw.assert);

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
        assert,
    }))
}

/// Pure string-parsing layer. Extracts unpack args from a pre-fetched comment string
/// and a list of pre-fetched field comment strings (one per field, in order).
/// Returns `Ok(None)` if no #[unpack(...)] annotation is present, and an error if the item's or
/// a field's annotation is malformed.
pub(crate) fn parse_comment(
    comment: &str,
    field_meta: &[(Option<String>, bool, String)],
    struct_name: &str,
) -> eyre::Result<Option<RawUnpackArgs>> {
    // annotation_args handles error detection for malformed annotations (unterminated
    // strings, unbalanced parens) and tells us whether the annotation is present.
    // CommentVisitor is then used for the actual content extraction, since annotation_args
    // does not strip the `//` / `* ` comment-line prefixes needed for multi-line annotations.
    let Some(_) = annotation_args(comment, "#[unpack(")? else {
        return Ok(None);
    };
    let pack_start = comment.find("#[unpack(").unwrap(); // safe: annotation_args confirmed it
    let args_start = pack_start + "#[unpack(".len();
    let Some(closing) = find_closing_paren(&comment[args_start..]) else {
        return Err(eyre::eyre!(
            "#[unpack(...)] annotation on `{struct_name}` has an unclosed `(` \
             — check for an unbalanced or stray `\"`"
        ));
    };
    let args_end = args_start + closing;
    let args_segments = CommentVisitor::new(comment, args_start, args_end).parse();
    let args_str = join_segments(&args_segments);

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();
    let mut calldata = false;
    let mut asserts: Vec<AssertDef> = Vec::new();

    for arg in split_args(&args_str) {
        let arg = arg.trim();
        if arg.is_empty() {
            continue;
        }
        if let Some((key, val)) = arg.split_once('=') {
            let val = val.trim().trim_matches('"');
            match key.trim() {
                "contract" => contract = Some(val.to_string()),
                "name" => name = Some(val.to_string()),
                "visibility" => visibility = val.to_string(),
                "assert" => match parse_closure(val) {
                    Some(Ok(defs)) => asserts.extend(defs),
                    Some(Err(e)) => return Err(e),
                    None => {
                        return Err(eyre::eyre!(
                            "assert value {val:?} on `{struct_name}` is not a valid closure; \
                             expected `|name| {{ expr }}` or `|each name| {{ expr }}`"
                        ));
                    }
                },
                key => {
                    return Err(eyre::eyre!(
                        "unknown key `{key}` in #[unpack(...)] on `{struct_name}`"
                    ));
                }
            }
        } else {
            match arg {
                "calldata" => calldata = true,
                "solhint-disable" => {}
                _ => {
                    return Err(eyre::eyre!(
                        "unknown flag `{arg}` in #[unpack(...)] on `{struct_name}`"
                    ));
                }
            }
        }
    }

    let fields = field_meta
        .iter()
        .map(|(c, memory, field_name)| parse_field_args(c.as_deref(), *memory, field_name))
        .collect::<eyre::Result<_>>()?;

    Ok(Some(RawUnpackArgs {
        contract,
        name,
        fields,
        visibility,
        calldata,
        assert: asserts,
    }))
}

fn type_requires_memory(ty: &Type) -> bool {
    match &ty.kind {
        TypeKind::Elementary(e) => matches!(e, ElementaryType::Bytes | ElementaryType::String),
        TypeKind::Array(_) => true,
        TypeKind::Custom(ItemId::Struct(_)) => true,
        _ => false,
    }
}

fn parse_field_args(
    comment: Option<&str>,
    memory: bool,
    field_name: &str,
) -> eyre::Result<FieldArgs<Unvalidated>> {
    let annotation = comment
        .map(|c| annotation_args(c, "#[unpack("))
        .transpose()?
        .flatten();
    let Some(_) = annotation else {
        return Ok(FieldArgs {
            memory,
            ..Default::default()
        });
    };
    let comment = comment.unwrap(); // safe: annotation_args confirmed the annotation is present
    let pack_start = comment.find("#[unpack(").unwrap();
    let args_start = pack_start + "#[unpack(".len();
    let Some(args_end_rel) = find_closing_paren(&comment[args_start..]) else {
        return Err(eyre::eyre!(
            "#[unpack(...)] annotation on field `{field_name}` has an unclosed `(` \
             — check for an unbalanced or stray `\"`"
        ));
    };
    let args_end = args_start + args_end_rel;
    let args_segments = CommentVisitor::new(comment, args_start, args_end).parse();
    let args_str = join_segments(&args_segments);

    if args_str.trim() == "default" {
        return Ok(FieldArgs {
            default: true,
            memory,
            ..Default::default()
        });
    }

    let mut method = None;
    let mut default = false;
    let mut length = None;
    let mut asserts: Vec<AssertDef> = Vec::new();
    for arg in split_args(&args_str) {
        let arg = arg.trim();
        if arg.is_empty() {
            continue;
        }
        if let Some((key, val)) = arg.split_once('=') {
            match key.trim() {
                "method" => method = Some(val.trim().trim_matches('"').to_string()),
                "length" => length = Some(val.trim().trim_matches('"').to_string()),
                "assert" => {
                    let inner = val.trim().trim_matches('"');
                    match parse_closure(inner) {
                        Some(Ok(defs)) => asserts.extend(defs),
                        Some(Err(e)) => return Err(e),
                        None => {
                            return Err(eyre::eyre!(
                                "assert value {inner:?} on field `{field_name}` is not a valid \
                                 closure; expected `|name| {{ expr }}` or `|each name| {{ expr }}`"
                            ));
                        }
                    }
                }
                key => {
                    return Err(eyre::eyre!(
                        "unknown key `{key}` in #[unpack(...)] field annotation on `{field_name}`"
                    ));
                }
            }
        } else {
            match arg {
                "default" => default = true,
                _ => {
                    return Err(eyre::eyre!(
                        "unknown flag `{arg}` in #[unpack(...)] field annotation on `{field_name}`"
                    ));
                }
            }
        }
    }
    Ok(if default {
        FieldArgs {
            default: true,
            memory,
            assert: asserts,
            ..Default::default()
        }
    } else {
        FieldArgs {
            method,
            default: false,
            memory,
            length,
            assert: asserts,
        }
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::common::tests::assert_def;
    use crate::common::{CaptureType, LengthType};

    #[test]
    fn test_basic() {
        let field_meta = vec![
            (None, false, String::new()),
            (None, true, String::new()),
            (None, false, String::new()),
            (None, false, String::new()),
        ];
        let raw = parse_comment("#[unpack]", &field_meta, "").unwrap();
        assert!(raw.is_none());
        let raw = parse_comment("#[unpack()]", &field_meta, "")
            .unwrap()
            .expect("Test failed");
        assert_eq!(raw.contract, None);
        assert_eq!(raw.name, None);
        assert_eq!(raw.visibility, "public".to_string());
        assert!(!raw.calldata);
        assert!(raw.assert.is_empty());
        let mut expected = vec![FieldArgs::<Unvalidated>::default(); 4];
        expected[1].memory = true;
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_all_keys() {
        let raw = parse_comment(
            "#[unpack(contract=\"Foo\", name=\"Bar\", visibility=\"private\", calldata)]",
            &vec![(None, false, String::new()); 4],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(raw.contract, Some("Foo".to_string()));
        assert_eq!(raw.name, Some("Bar".to_string()));
        assert_eq!(raw.visibility, "private".to_string());
        assert!(raw.calldata);
        assert!(raw.assert.is_empty());
        let expected = vec![FieldArgs::<Unvalidated>::default(); 4];
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_flags() {
        // solhint-disable is accepted (backward-compat) but does nothing; calldata is still parsed.
        let raw = parse_comment(
            "#[unpack(solhint-disable)]",
            &vec![(None, false, String::new()); 2],
            "",
        )
        .unwrap()
        .unwrap();
        assert!(!raw.calldata);

        let raw = parse_comment(
            "#[unpack(calldata, solhint-disable)]",
            &vec![(None, false, String::new()); 2],
            "",
        )
        .unwrap()
        .unwrap();
        assert!(raw.calldata);
    }

    #[test]
    fn test_length_quoted() {
        // length = "uint32" (with quotes) must behave identically to length = uint32
        let raw = parse_comment(
            "#[unpack()]",
            &vec![
                (
                    Some("#[unpack(length = \"uint32\")]".to_string()),
                    false,
                    String::new(),
                ),
                (
                    Some("#[unpack(method=\"Foo.bar\", length = \"uint64\")]".to_string()),
                    false,
                    String::new(),
                ),
            ],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: None,
                default: false,
                memory: false,
                length: Some("uint32".to_string()),
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                default: false,
                memory: false,
                length: Some("uint64".to_string()),
                assert: vec![],
            }
        );
    }

    #[test]
    fn test_assert() {
        // type-level assert, alone
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 }")]"#,
            &vec![(None, false, String::new()); 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { r.version == 1 }",
                "",
                "r.version == 1"
            )]
        );

        // assert can appear before other args (no longer required to be last)
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 }", calldata)]"#,
            &vec![(None, false, String::new()); 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert!(raw.calldata);
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { r.version == 1 }",
                "",
                "r.version == 1"
            )]
        );

        // multiple asserts
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 }", assert = "|r| { r.x > 0 }")]"#,
            &vec![(None, false, String::new()); 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![
                assert_def(
                    CaptureType::Field,
                    "|r| { r.version == 1 }",
                    "",
                    "r.version == 1"
                ),
                assert_def(CaptureType::Field, "|r| { r.x > 0 }", "", "r.x > 0"),
            ]
        );

        // body containing a function call with commas and parens (safe inside quotes)
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { isValid(r, 0) }")]"#,
            &vec![(None, false, String::new()); 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { isValid(r, 0) }",
                "",
                "isValid(r, 0)"
            )]
        );

        // field-level assert
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|b| { b.length == 48 }")]"#.to_string()),
                true,
                String::new(),
            )],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { b.length == 48 }",
                "",
                "b.length == 48"
            )]
        );

        // element-level assert using `each` keyword
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|each x| { x > 0 }")]"#.to_string()),
                false,
                String::new(),
            )],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![assert_def(
                CaptureType::Element,
                "|each x| { x > 0 }",
                "",
                "x > 0"
            )]
        );

        // `each` prefix on a variable name that starts with "each" is not treated as the keyword
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|eachItem| { eachItem > 0 }")]"#.to_string()),
                false,
                String::new(),
            )],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![assert_def(
                CaptureType::Field,
                "|eachItem| { eachItem > 0 }",
                "",
                "eachItem > 0"
            )]
        );
    }

    #[test]
    fn test_assert_solidity_string_literals() {
        // Assert body contains a Solidity string literal with a comma.
        let raw = parse_comment(
            r#"#[unpack(assert = "|each c| { c != "," }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Element,
                r#"|each c| { c != "," }"#,
                "",
                r#"c != ",""#
            )]
        );

        // Assert body contains a Solidity string literal with a `)`.
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { keccak256(")") == hash }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                r#"|r| { keccak256(")") == hash }"#,
                "",
                r#"keccak256(")") == hash"#
            )]
        );

        // Assert body contains a Solidity string literal with `//`.
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.url != "http://" }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                r#"|r| { r.url != "http://" }"#,
                "",
                r#"r.url != "http://""#
            )]
        );

        // Other args after a string-literal assert are still parsed.
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.sep != "," }", calldata)]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert!(raw.calldata);
        assert_eq!(raw.assert[0].expr, r#"r.sep != ",""#.to_string());

        // Field-level: same handling.
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|each c| { c != "," }")]"#.to_string()),
                false,
                String::new(),
            )],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![assert_def(
                CaptureType::Element,
                r#"|each c| { c != "," }"#,
                "",
                r#"c != ",""#
            )]
        );
    }

    #[test]
    fn test_assert_multiline() {
        // Single-line assert with a Solidity string literal containing `//`.
        let raw = parse_comment(
            "#[unpack(assert = \"|r| { r.url != \"http://\" }\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                r#"|r| { r.url != "http://" }"#,
                "",
                r#"r.url != "http://""#
            )]
        );

        // Multi-line assert: `//` comment prefixes on continuation lines must be stripped.
        let raw = parse_comment(
            "#[unpack(assert = \"|r| {\n//     r.version == 1\n// }\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { r.version == 1 }",
                "",
                "r.version == 1"
            )]
        );

        // Multi-line assert with a Solidity string literal containing `//` —
        // the comment prefixes on continuation lines are stripped, but `//`
        // inside the string literal on its own line is preserved.
        let raw = parse_comment(
            "#[unpack(assert = \"|r| {\n//     r.url != \"http://\"\n// }\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                r#"|r| { r.url != "http://" }"#,
                "",
                r#"r.url != "http://""#
            )]
        );

        // Block comment style with `*` prefixes on continuation lines.
        let raw = parse_comment(
            "/**\n * #[unpack(assert = \"|r| {\n *    r.value == 0\n *  }\")]\n */",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { r.value == 0 }",
                "",
                "r.value == 0"
            )]
        );

        // Star-less `/* */` outer with no inner comment — body lines have no prefix to strip.
        let raw = parse_comment(
            "/* #[unpack(assert = \"|r| {\n   r.version == 1\n}\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { r.version == 1 }",
                "",
                "r.version == 1"
            )]
        );
    }

    #[test]
    fn test_assert_inline_comments() {
        // `//` comment lines inside a multi-line closure body must be preserved.
        let raw = parse_comment(
            "#[unpack(assert = \"|r| {\n//    // this value\n//    // should always be zero\n//    r.value == 0\n//  }\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { // this value\n// should always be zero\nr.value == 0 }",
                "// this value\n// should always be zero",
                "r.value == 0",
            )]
        );

        // An empty continuation line (`//` with nothing after it) must not
        // interfere with the comment-preserving on the lines around it.
        let raw = parse_comment(
            "#[unpack(assert = \"|r| {\n//    // this value\n//\n//    r.value == 0\n//  }\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { // this value\nr.value == 0 }",
                "// this value",
                "r.value == 0",
            )]
        );

        // /* */ outer + // inner.
        let raw = parse_comment(
            "/**\n * #[unpack(assert = \"|r| {\n *    // this comment\n *    r.value == 0\n *  }\")]\n */",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { // this comment\nr.value == 0 }",
                "// this comment",
                "r.value == 0",
            )]
        );

        // Star-less /* */ outer + // inner: the // must be treated as a Solidity comment,
        // not stripped as an annotation prefix.
        let raw = parse_comment(
            "/* #[unpack(assert = \"|r| {\n   // version check\n   r.version == 1\n}\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { // version check\nr.version == 1 }",
                "// version check",
                "r.version == 1",
            )]
        );
    }

    #[test]
    fn test_assert_block_comment_plus_expr() {
        // An inline `/* ... */` block comment on the same line as an expression must not
        // consume the expression.  Previously the whole line was pushed to comment_lines and
        // the expression was silently dropped, emitting no require().
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { /* v1 only */ r.version == 1 }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { /* v1 only */ r.version == 1 }",
                "/* v1 only */",
                "r.version == 1",
            )]
        );

        // A closing `*/` on the same line as the expression (multi-line block comment):
        // `/* \n */ expr` — the expression after `*/` must not be dropped.
        let raw = parse_comment(
            "#[unpack(assert = \"|r| {\n//    /*\n//     */ r.version == 1\n//  }\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { /*\n*/ r.version == 1 }",
                "/*\n*/",
                "r.version == 1",
            )]
        );

        // A trailing `/* */` after the expression on the same line stays in `expr` —
        // Solidity accepts block comments inside require() arguments, so this is valid.
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 /* v1 only */ }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { r.version == 1 /* v1 only */ }",
                "",
                "r.version == 1 /* v1 only */",
            )]
        );

        // `/* */` outer with `* ` prefixes + inline `/* */` inner on a continuation line.
        let raw = parse_comment(
            "/**\n * #[unpack(assert = \"|r| {\n *    /* v1 only */ r.version == 1\n *  }\")]\n */",
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { /* v1 only */ r.version == 1 }",
                "/* v1 only */",
                "r.version == 1",
            )]
        );

        // Star-less `/* */` outer + inline `/* */` inner on a continuation line.
        let raw = parse_comment(
            "/* #[unpack(assert = \"|r| {\n   /* v1 only */ r.version == 1\n}\")]",
            &vec![(None, false, String::new())],
            "",
        )
        .expect("annotation present")
        .expect("valid parse");
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|r| { /* v1 only */ r.version == 1 }",
                "/* v1 only */",
                "r.version == 1",
            )]
        );
    }

    #[test]
    fn test_field_args() {
        let raw = parse_comment(
            "#[unpack(visibility=\"private\", name=\"Foo\")]",
            &vec![
                (
                    Some("#[unpack(method=\"Foo.bar\")]".to_string()),
                    false,
                    String::new(),
                ),
                (Some("#[unpack(default)]".to_string()), true, String::new()),
                (
                    Some("#[unpack(default, method=\"Foo.bar\")]".to_string()),
                    false,
                    String::new(),
                ),
                (
                    Some("#[unpack(length = uint32)]".to_string()),
                    true,
                    String::new(),
                ),
                (
                    Some("#[unpack(method=\"Foo.bar\", length = uint64)]".to_string()),
                    true,
                    String::new(),
                ),
                (
                    Some("#[unpack(length = BLST.BLS_SIGNATURE_LENGTH)]".to_string()),
                    true,
                    String::new(),
                ),
            ],
            "",
        )
        .unwrap()
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
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: None,
                default: true,
                memory: true,
                length: None,
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[2],
            FieldArgs {
                method: None,
                default: true,
                memory: false,
                length: None,
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[3],
            FieldArgs {
                method: None,
                default: false,
                memory: true,
                length: Some("uint32".to_string()),
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[4],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                default: false,
                memory: true,
                length: Some("uint64".to_string()),
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[5],
            FieldArgs {
                method: None,
                default: false,
                memory: true,
                length: Some("BLST.BLS_SIGNATURE_LENGTH".to_string()),
                assert: vec![],
            }
        );
    }

    #[test]
    fn test_length_expressions() {
        // A constant length may contain parentheses and commas without ending the annotation
        // or swallowing the arguments after it.
        let raw = parse_comment(
            "#[unpack()]",
            &vec![
                (
                    Some("#[unpack(length = uint256(Lengths.HASH_LENGTH))]".to_string()),
                    true,
                    String::new(),
                ),
                (
                    Some("#[unpack(length = (2 * 16), method=\"parseHash\")]".to_string()),
                    true,
                    String::new(),
                ),
            ],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: None,
                default: false,
                memory: true,
                length: Some("uint256(Lengths.HASH_LENGTH)".to_string()),
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: Some("parseHash".to_string()),
                default: false,
                memory: true,
                length: Some("(2 * 16)".to_string()),
                assert: vec![],
            }
        );

        // A function call with multiple arguments must not be split on the inner comma.
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some("#[unpack(length = Lib.f(a, b) * 2, method = \"g\")]".to_string()),
                true,
                String::new(),
            )],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: Some("g".to_string()),
                default: false,
                memory: true,
                length: Some("Lib.f(a, b) * 2".to_string()),
                assert: vec![],
            }
        );
    }

    #[test]
    fn test_malformed_annotation() {
        // a malformed annotation is an error, not an absent annotation
        let err = parse_comment("#[unpack(name = \"Foo)]", &[], "")
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unterminated string in annotation `#[unpack(name = \"Foo)]`"),
            "{err}"
        );
        let err = parse_comment(
            "#[unpack()]",
            &[(
                Some("#[unpack(length = (1 + 2]".to_string()),
                true,
                String::new(),
            )],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(
            err.contains("unbalanced parentheses in annotation `#[unpack(length = (1 + 2]`"),
            "{err}"
        );
        let err = parse_comment("#[unpack(bogus = \"x\")]", &[], "MyStruct")
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unknown key `bogus` in #[unpack(...)] on `MyStruct`"),
            "{err}"
        );
        let err = parse_comment("#[unpack(bogus)]", &[], "MyStruct")
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unknown flag `bogus` in #[unpack(...)] on `MyStruct`"),
            "{err}"
        ); // Field-level unknown key (e.g. misspelled `asserts`).
        let err = parse_comment(
            "#[unpack()]",
            &[(
                Some("#[unpack(asserts = \"|x| { x > 0 }\")]".to_string()),
                false,
                "MyStruct".to_string(),
            )],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(
            err.contains("unknown key `asserts` in #[unpack(...)] field annotation on `MyStruct`"),
            "{err}"
        );
        // Field-level unknown flag.
        let err = parse_comment(
            "#[unpack()]",
            &[(
                Some("#[unpack(bogus)]".to_string()),
                false,
                "MyStruct".to_string(),
            )],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(
            err.contains("unknown flag `bogus` in #[unpack(...)] field annotation on `MyStruct`"),
            "{err}"
        );
    }

    #[test]
    fn test_validate_length() {
        let validate = |length: &str| {
            FieldArgs::<Unvalidated> {
                method: None,
                default: false,
                memory: true,
                length: Some(length.to_string()),
                assert: vec![],
            }
            .validate("data", "Foo")
            .map(|f| f.length)
        };
        assert_eq!(
            validate("uint32").unwrap(),
            Some(LengthSpec::Type(LengthType::Uint(4)))
        );
        assert_eq!(
            validate("32").unwrap(),
            Some(LengthSpec::Constant("32".to_string()))
        );
        assert_eq!(
            validate("BLST.SIG_LENGTH").unwrap(),
            Some(LengthSpec::Constant("BLST.SIG_LENGTH".to_string()))
        );
        let err = validate("").unwrap_err().to_string();
        assert!(
            err.contains("empty `length` for field `data` of `Foo`"),
            "{err}"
        );
    }

    #[test]
    fn test_unclosed_paren() {
        // Stray " in a struct-level annotation: annotation_args catches the unterminated string.
        // Previously this silently produced None (no unpack function generated).
        let err = parse_comment(
            r#"#[unpack(method = "parseX)]"#,
            &vec![(None, false, String::new()); 1],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("unterminated string"), "wrong error: {err}");

        // Stray " in a field-level annotation: annotation_args catches the unterminated string.
        // Previously this silently returned default FieldArgs, dropping method/length/assert.
        let err = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(method = "parseField)]"#.to_string()),
                false,
                "myField".to_string(),
            )],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("unterminated string"), "wrong error: {err}");
    }
}
