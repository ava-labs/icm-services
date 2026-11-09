//! This module contains the tools the unpack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.

mod visitor;
use visitor::{ClosureVisitor, CommentVisitor, join_segments};

use reforge::PreprocessingData;
use solar::ast::{ElementaryType, Span};
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, ItemId, SourceId, Struct, Type, TypeKind};

use crate::common::{LengthSpec, Unvalidated, Validated, ValidationState, annotation_args};

pub struct UnpackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs<Validated>>,
    pub visibility: String,
    pub calldata: bool,
    pub solhint_disable: bool,
    /// Type-level post-condition. The capture name is replaced with `result`
    /// (the decoded struct) in the assertion body.
    pub assert: Vec<AssertDef>,
}

#[derive(Debug, Clone, PartialEq, Default)]
pub enum CaptureType {
    Element,
    #[default]
    Field,
}

#[derive(Debug, Clone, PartialEq)]
pub struct AssertDef {
    pub captured: CaptureType,
    pub var: String,
    /// Inline comment lines that precede this assertion's expression in the
    /// closure body. Empty when there is no leading comment.
    pub comment: String,
    pub expr: String,
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
    pub solhint_disable: bool,
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
    let solhint_disable = raw.solhint_disable;
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
        solhint_disable,
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
    let args_segments = CommentVisitor::new(&comment[args_start..args_end]).parse();
    let args_str = join_segments(&args_segments);

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();
    let mut calldata = false;
    let mut solhint_disable = false;
    let mut asserts: Vec<AssertDef> = Vec::new();

    for arg in split_args(&args_str) {
        let arg = arg.trim();
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
                _ => {}
            }
        } else {
            match arg {
                "calldata" => calldata = true,
                "solhint-disable" => solhint_disable = true,
                _ => {}
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
        solhint_disable,
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
    let raw_args = comment[args_start..args_start + args_end_rel].trim();
    let args_segments = CommentVisitor::new(raw_args).parse();
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
        if let Some((key, val)) = arg.trim().split_once('=') {
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
                _ => {}
            }
        } else if arg.trim() == "default" {
            default = true;
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

/// Finds the index of the `)` that closes the current paren level, skipping
/// over any nested `(...)` pairs and `"..."` quoted strings.
/// Returns `None` if no such `)` exists.
///
/// See [`strip_comment_prefixes`] for the quote / brace nesting rules that
/// allow Solidity string literals inside assert closure bodies.
fn find_closing_paren(s: &str) -> Option<usize> {
    let mut depth = 0usize;
    let mut in_quotes = false;
    let mut brace_depth = 0usize;
    for (i, c) in s.char_indices() {
        match c {
            '"' if brace_depth == 0 => in_quotes = !in_quotes,
            '{' if in_quotes => brace_depth += 1,
            '}' if in_quotes && brace_depth > 0 => brace_depth -= 1,
            '(' if !in_quotes => depth += 1,
            ')' if !in_quotes => {
                if depth == 0 {
                    return Some(i);
                }
                depth -= 1;
            }
            _ => {}
        }
    }
    None
}

/// Splits `s` on commas that are outside of `"..."` quoted strings.
///
/// See [`strip_comment_prefixes`] for the quote / brace nesting rules that
/// allow Solidity string literals inside assert closure bodies.
fn split_args(s: &str) -> Vec<&str> {
    let mut result = Vec::new();
    let mut start = 0;
    let mut in_quotes = false;
    let mut brace_depth = 0usize;
    for (i, c) in s.char_indices() {
        match c {
            '"' if brace_depth == 0 => in_quotes = !in_quotes,
            '{' if in_quotes => brace_depth += 1,
            '}' if in_quotes && brace_depth > 0 => brace_depth -= 1,
            ',' if !in_quotes => {
                result.push(&s[start..i]);
                start = i + 1;
            }
            _ => {}
        }
    }
    result.push(&s[start..]);
    result
}

/// Parses a closure of the form `|[each] capture| { body }` and returns a single
/// `AssertDef` for the body expression. Any leading `//` or `/* */` Solidity comment
/// lines are extracted into `AssertDef::comment`; everything else is the expression.
///
/// Each `assert = "..."` annotation key produces exactly one `AssertDef`. Use multiple
/// `assert = ...` attributes for multiple postconditions.
///
/// - Returns `None` if the string doesn't look like a closure (no leading `|` or no body).
/// - Returns `Some(Err(...))` if the capture name is invalid or the body is empty.
/// - Returns `Some(Ok(...))` with a single-element vec on success.
fn parse_closure(s: &str) -> Option<eyre::Result<Vec<AssertDef>>> {
    let mut visitor = ClosureVisitor::new(s.trim());
    let (captured, var) = match visitor.parse_capture() {
        Some(Ok(t)) => t,
        Some(Err(e)) => return Some(Err(e)),
        None => return None,
    };
    let body = visitor.extract_body()?;
    visitor.classify_lines(body);
    let expr = visitor.expr_lines.join(" ");
    if expr.is_empty() {
        return Some(Err(eyre::eyre!(
            "assert closure body for `{var}` contains no boolean expressions"
        )));
    }
    Some(Ok(vec![AssertDef {
        captured,
        var,
        comment: visitor.comment_lines.join("\n"),
        expr,
    }]))
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::common::LengthType;

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
        assert!(!raw.solhint_disable);
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
        assert!(!raw.solhint_disable);
        assert!(raw.assert.is_empty());
        let expected = vec![FieldArgs::<Unvalidated>::default(); 4];
        assert_eq!(raw.fields, expected);
    }

    #[test]
    fn test_flags() {
let raw = parse_comment(
            "#[unpack(solhint-disable)]",
            &vec![(None, false, String::new()); 2],
            "",
        )
        .unwrap()
        .unwrap();
        assert!(raw.solhint_disable);
        assert!(!raw.calldata);

        let raw = parse_comment(
            "#[unpack(calldata, solhint-disable)]",
            &vec![(None, false, String::new()); 2],
            "",
        )
        .unwrap()
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: "r.version == 1".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: "r.version == 1".to_string(),
            }]
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
                AssertDef {
                    captured: CaptureType::Field,
                    var: "r".to_string(),
                    comment: String::new(),
                    expr: "r.version == 1".to_string(),
                },
                AssertDef {
                    captured: CaptureType::Field,
                    var: "r".to_string(),
                    comment: String::new(),
                    expr: "r.x > 0".to_string(),
                },
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: "isValid(r, 0)".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "b".to_string(),
                comment: String::new(),
                expr: "b.length == 48".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Element,
                var: "x".to_string(),
                comment: String::new(),
                expr: "x > 0".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "eachItem".to_string(),
                comment: String::new(),
                expr: "eachItem > 0".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Element,
                var: "c".to_string(),
                comment: String::new(),
                expr: r#"c != ",""#.to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: r#"keccak256(")") == hash"#.to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: r#"r.url != "http://""#.to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Element,
                var: "c".to_string(),
                comment: String::new(),
                expr: r#"c != ",""#.to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: r#"r.url != "http://""#.to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: "r.version == 1".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: String::new(),
                expr: "r.value == 0".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: "// this value\n// should always be zero".to_string(),
                expr: "r.value == 0".to_string(),
            }]
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
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                comment: "// this comment".to_string(),
                expr: "r.value == 0".to_string(),
            }]
        );
    }

    #[test]
    fn test_assert_empty_body() {
        // An empty closure body must be an error, not a silent no-op.
        let err = parse_comment(
            r#"#[unpack(assert = "|r| { }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("no boolean expressions"), "wrong error: {err}");

        // A body containing only a comment must also be an error.
        let err = parse_comment(
            r#"#[unpack(assert = "|r| { // just a comment }")]"#,
            &vec![(None, false, String::new())],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("no boolean expressions"), "wrong error: {err}");
    }

    #[test]
    fn test_assert_capture_errors() {
        // empty capture in a type-level assert
        let err = parse_comment(
            r#"#[unpack(assert = "|| { true }")]"#,
            &vec![(None, false, String::new()); 1],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("empty capture name"), "wrong error: {err}");

        // empty capture in a field-level assert
        let err = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|| { true }")]"#.to_string()),
                false,
                "myField".to_string(),
            )],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("empty capture name"), "wrong error: {err}");

        // non-identifier capture name
        let err = parse_comment(
            r#"#[unpack(assert = "|2bad| { true }")]"#,
            &vec![(None, false, String::new()); 1],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("not a valid identifier"), "wrong error: {err}");

        // empty capture after `each`
        let err = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|each | { true }")]"#.to_string()),
                false,
                "myField".to_string(),
            )],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("empty capture name"), "wrong error: {err}");

        // malformed assert (missing pipe) on a type-level assert — None path
        let err = parse_comment(
            r#"#[unpack(assert = "x > 0")]"#,
            &vec![(None, false, String::new()); 1],
            "MyStruct",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("not a valid closure"), "wrong error: {err}");
        assert!(err.contains("MyStruct"), "error should name the struct: {err}");

        // malformed assert (missing braces) on a field-level assert — None path
        let err = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|x| x > 0")]"#.to_string()),
                false,
                "myField".to_string(),
            )],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(err.contains("not a valid closure"), "wrong error: {err}");
        assert!(err.contains("myField"), "error should name the field: {err}");
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
            &[(Some("#[unpack(length = (1 + 2]".to_string()), true, String::new())],
            "",
        )
        .unwrap_err()
        .to_string();
        assert!(
            err.contains("unbalanced parentheses in annotation `#[unpack(length = (1 + 2]`"),
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
