//! This module contains the tools the pack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.

use reforge::PreprocessingData;
use solar::ast::Span;
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, SourceId, Struct};

use crate::common::{
    AssertDef, CommentVisitor, LengthSpec, Unvalidated, Validated, ValidationState,
    annotation_args, find_closing_paren, join_segments, parse_closure, split_args,
};

pub struct PackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Validated>>,
    pub visibility: String,
    pub assert: Vec<AssertDef>,
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
    /// Resolves the `length` string (if present) into a [`LengthSpec`], converting this
    /// `FieldArgs<Unvalidated>` into a `FieldArgs<Validated>`.
    pub fn validate(
        self,
        field_name: &str,
        struct_name: &str,
    ) -> eyre::Result<FieldArgs<Validated>> {
        let length = self
            .length
            .map(|value| LengthSpec::validate(value, field_name, struct_name))
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
#[derive(Debug)]
pub(crate) struct RawPackArgs {
    pub contract: Option<String>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Unvalidated>>,
    pub visibility: String,
    pub assert: Vec<AssertDef>,
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
/// Returns `None` if no annotation is present on the item itself, or `Some(Err(...))` if an
/// annotation is malformed, a field's `length` is invalid, or a `contract` name was specified
/// but not found in the HIR.
pub fn parse_args(
    ctx: &Gcx,
    source_id: SourceId,
    span: Span,
    data: &mut PreprocessingData<'_>,
) -> Option<eyre::Result<PackArgs>> {
    let doc_comment = reforge::get_comment(ctx, source_id, span, data)?;
    let struct_def = ctx.hir.structs().find(|s| s.span == span);
    let item_name = struct_def
        .map(|s| s.name.as_str())
        .or_else(|| {
            ctx.hir
                .enums()
                .find(|e| e.span == span)
                .map(|e| e.name.as_str())
        })
        .unwrap_or("");

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

    let mut raw = match parse_comment(&doc_comment, &field_comments, item_name) {
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
                    "contract `{name}` specified in #[pack(contract=...)] was not found"
                )));
            }
        },
    };

    let assert = std::mem::take(&mut raw.assert);

    let fields = match raw.validate_fields(ctx, struct_def) {
        Ok(f) => f,
        Err(e) => return Some(Err(e)),
    };

    Some(Ok(PackArgs {
        contract,
        name: raw.name,
        fields,
        visibility: raw.visibility,
        assert,
    }))
}

/// Pure string-parsing layer. Extracts pack args from a pre-fetched comment string
/// and a list of pre-fetched field comment strings (one per field, in order).
/// Returns `Ok(None)` if no #[pack(...)] annotation is present, and an error if the item's or a
/// field's annotation is malformed.
pub(crate) fn parse_comment(
    comment: &str,
    field_comments: &[Option<String>],
    item_name: &str,
) -> eyre::Result<Option<RawPackArgs>> {
    // annotation_args handles error detection for malformed annotations (unterminated
    // strings, unbalanced parens) and tells us whether the annotation is present.
    // CommentVisitor is then used for the actual content extraction, since annotation_args
    // does not strip the `//` / `* ` comment-line prefixes needed for multi-line annotations.
    let Some(_) = annotation_args(comment, "#[pack(")? else {
        return Ok(None);
    };
    let pack_start = comment.find("#[pack(").unwrap(); // safe: annotation_args confirmed it
    let args_start = pack_start + "#[pack(".len();
    let Some(closing) = find_closing_paren(&comment[args_start..]) else {
        return Err(eyre::eyre!(
            "#[pack(...)] annotation on `{item_name}` has an unclosed `(`"
        ));
    };
    let args_end = args_start + closing;
    let args_segments = CommentVisitor::new(comment, args_start, args_end).parse();
    let args_str = join_segments(&args_segments);

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();
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
                            "assert value {val:?} on `{item_name}` is not a valid closure; \
                             expected `|name| {{ expr }}` or `|each name| {{ expr }}`"
                        ));
                    }
                },
                key => {
                    return Err(eyre::eyre!(
                        "unknown key `{key}` in #[pack(...)] on `{item_name}`"
                    ));
                }
            }
        } else {
            return Err(eyre::eyre!(
                "unknown flag `{arg}` in #[pack(...)] on `{item_name}`"
            ));
        }
    }

    let fields = field_comments
        .iter()
        .map(|c| parse_field_args(c.as_deref(), item_name))
        .collect::<eyre::Result<_>>()?;

    Ok(Some(RawPackArgs {
        contract,
        name,
        fields,
        visibility,
        assert: asserts,
    }))
}

fn parse_field_args(
    comment: Option<&str>,
    item_name: &str,
) -> eyre::Result<FieldArgs<Unvalidated>> {
    let annotation = comment
        .map(|c| annotation_args(c, "#[pack("))
        .transpose()?
        .flatten();
    let Some(_) = annotation else {
        return Ok(FieldArgs::default());
    };
    let comment = comment.unwrap(); // safe: annotation_args confirmed the annotation is present
    let pack_start = comment.find("#[pack(").unwrap();
    let args_start = pack_start + "#[pack(".len();
    let Some(args_end_rel) = find_closing_paren(&comment[args_start..]) else {
        return Err(eyre::eyre!(
            "#[pack(...)] annotation on field of `{item_name}` has an unclosed `(`"
        ));
    };
    let args_end = args_start + args_end_rel;
    let args_segments = CommentVisitor::new(comment, args_start, args_end).parse();
    let args_str = join_segments(&args_segments);

    let mut method = None;
    let mut ignore = false;
    let mut length = None;
    for arg in split_args(&args_str) {
        let arg = arg.trim();
        if arg.is_empty() {
            continue;
        }
        if let Some((key, val)) = arg.split_once('=') {
            match key.trim() {
                "method" => method = Some(val.trim().trim_matches('"').to_string()),
                "length" => length = Some(val.trim().trim_matches('"').to_string()),
                key => {
                    return Err(eyre::eyre!(
                        "unknown key `{key}` in #[pack(...)] field annotation on `{item_name}`"
                    ));
                }
            }
        } else {
            match arg {
                "ignore" => ignore = true,
                _ => {
                    return Err(eyre::eyre!(
                        "unknown flag `{arg}` in #[pack(...)] field annotation on `{item_name}`"
                    ));
                }
            }
        }
    }
    Ok(if ignore {
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
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::common::tests::assert_def;
    use crate::common::{CaptureType, LengthType};

    #[test]
    fn test_basic() {
        let raw = parse_comment("#[pack]", &vec![None, None, None, None], "").unwrap();
        assert!(raw.is_none());
        let raw = parse_comment("#[pack()]", &vec![None, None, None, None], "")
            .unwrap()
            .expect("Test failed");
        assert_eq!(raw.contract, None);
        assert_eq!(raw.name, None);
        assert_eq!(raw.visibility, "public".to_string());
        assert!(raw.assert.is_empty());
        let expected = vec![FieldArgs::<Unvalidated>::default(); 4];
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_all_keys() {
        let raw = parse_comment(
            "#[pack(contract=\"Foo\", name=\"Bar\", visibility=\"private\")]",
            &vec![None, None, None, None],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(raw.contract, Some("Foo".to_string()));
        assert_eq!(raw.name, Some("Bar".to_string()));
        assert_eq!(raw.visibility, "private".to_string());
        assert!(raw.assert.is_empty());
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
            "",
        )
        .unwrap()
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
                Some("#[pack(length = 32)]".to_string()),
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
                length: Some("32".to_string()),
            }
        );
    }

    #[test]
    fn test_length_expressions() {
        // A constant length may contain parentheses and commas without ending the annotation
        // or swallowing the arguments after it.
        let raw = parse_comment(
            "#[pack()]",
            &vec![
                Some("#[pack(length = type(uint16).max)]".to_string()),
                Some("#[pack(length = Lib.f(a, b) * 2, method = \"g\")]".to_string()),
            ],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.fields[0],
            FieldArgs {
                method: None,
                ignore: false,
                length: Some("type(uint16).max".to_string()),
            }
        );
        assert_eq!(
            raw.fields[1],
            FieldArgs {
                method: Some("g".to_string()),
                ignore: false,
                length: Some("Lib.f(a, b) * 2".to_string()),
            }
        );
    }

    #[test]
    fn test_malformed_annotation() {
        // a malformed annotation is an error, not an absent annotation
        let err = parse_comment("#[pack(name = \"Foo)]", &[], "")
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unterminated string in annotation `#[pack(name = \"Foo)]`"),
            "{err}"
        );
        let err = parse_comment(
            "#[pack()]",
            &[Some("#[pack(length = (1 + 2]".to_string())],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(
            err.contains("unbalanced parentheses in annotation `#[pack(length = (1 + 2]`"),
            "{err}"
        );
        let err = parse_comment("#[pack(bogus = \"x\")]", &[], "MyStruct")
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unknown key `bogus` in #[pack(...)] on `MyStruct`"),
            "{err}"
        );
        let err = parse_comment("#[pack(bogus)]", &[], "MyStruct")
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unknown flag `bogus` in #[pack(...)] on `MyStruct`"),
            "{err}"
        );
        let err = parse_comment(
            "#[pack()]",
            &[Some("#[pack(bogus)]".to_string())],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(
            err.contains("unknown flag `bogus` in #[pack(...)] field annotation on `MyStruct`"),
            "{err}"
        );
    }

    #[test]
    fn test_assert() {
        // type-level assert
        let raw = parse_comment(
            r#"#[pack(assert = "|b| { b.length > 0 }")]"#,
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { b.length > 0 }",
                "",
                "b.length > 0"
            )]
        );

        // multiple asserts
        let raw = parse_comment(
            r#"#[pack(assert = "|b| { b.length > 0 }", assert = "|b| { b.length % 2 == 0 }")]"#,
            &[],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![
                assert_def(
                    CaptureType::Field,
                    "|b| { b.length > 0 }",
                    "",
                    "b.length > 0"
                ),
                assert_def(
                    CaptureType::Field,
                    "|b| { b.length % 2 == 0 }",
                    "",
                    "b.length % 2 == 0"
                ),
            ]
        );

        // assert with other args
        let raw = parse_comment(
            r#"#[pack(assert = "|b| { b.length == 48 }", name = "serialize")]"#,
            &[],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(raw.name, Some("serialize".to_string()));
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { b.length == 48 }",
                "",
                "b.length == 48"
            )]
        );

        // malformed assert
        let err = parse_comment(r#"#[pack(assert = "x > 0")]"#, &[], "MyStruct")
            .unwrap_err()
            .to_string();
        assert!(err.contains("not a valid closure"), "wrong error: {err}");
        assert!(
            err.contains("MyStruct"),
            "error should name the type: {err}"
        );
    }

    #[test]
    fn test_assert_block_comment_annotations() {
        // `/* */` block comment style with `* ` continuation prefixes — the prefixes must be
        // stripped so that the closure body is correctly extracted.
        let raw = parse_comment(
            "/**\n * #[pack(assert = \"|b| {\n *     b.length > 0\n * }\")]\n */",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { b.length > 0 }",
                "",
                "b.length > 0"
            )]
        );

        // Multi-line expression in a block comment: `&&` on a continuation line must not
        // lose its operator due to the `* ` prefix strip.
        let raw = parse_comment(
            "/**\n * #[pack(assert = \"|b| {\n *     b.length > 0\n *     && b.length <= 100\n * }\")]\n */",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(raw.assert[0].expr, "b.length > 0 && b.length <= 100");

        // Star-less `/* */` annotation — no `* ` on continuation lines, body still correct.
        let raw = parse_comment(
            "/* #[pack(assert = \"|b| {\n    b.length > 0\n}\")]",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(raw.assert[0].expr, "b.length > 0");

        // `/* */` outer with `* ` prefixes + `//` inner comment: `//` must be preserved,
        // not stripped as an annotation prefix.
        let raw = parse_comment(
            "/**\n * #[pack(assert = \"|b| {\n *    // length check\n *    b.length > 0\n * }\")]\n */",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { // length check\nb.length > 0 }",
                "// length check",
                "b.length > 0",
            )]
        );

        // Star-less `/* */` outer + `//` inner: the `//` must be treated as a Solidity
        // comment, not stripped as an annotation prefix.
        let raw = parse_comment(
            "/* #[pack(assert = \"|b| {\n   // length check\n   b.length > 0\n}\")]",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { // length check\nb.length > 0 }",
                "// length check",
                "b.length > 0",
            )]
        );

        // `/* */` outer with `* ` prefixes + inline `/* */` inner on a continuation line.
        let raw = parse_comment(
            "/**\n * #[pack(assert = \"|b| {\n *    /* non-empty */ b.length > 0\n * }\")]\n */",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { /* non-empty */ b.length > 0 }",
                "/* non-empty */",
                "b.length > 0",
            )]
        );

        // Star-less `/* */` outer + inline `/* */` inner on a continuation line.
        let raw = parse_comment(
            "/* #[pack(assert = \"|b| {\n   /* non-empty */ b.length > 0\n}\")]",
            &vec![None; 1],
            "",
        )
        .unwrap()
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![assert_def(
                CaptureType::Field,
                "|b| { /* non-empty */ b.length > 0 }",
                "/* non-empty */",
                "b.length > 0",
            )]
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
        assert_eq!(
            validate("32").unwrap(),
            Some(LengthSpec::Constant("32".to_string()))
        );
        assert_eq!(
            validate("Lengths.HASH_LENGTH").unwrap(),
            Some(LengthSpec::Constant("Lengths.HASH_LENGTH".to_string()))
        );
        let err = validate("").unwrap_err().to_string();
        assert!(
            err.contains("empty `length` for field `data` of `Foo`"),
            "{err}"
        );
    }
}
