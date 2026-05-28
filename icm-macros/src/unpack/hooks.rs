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
    /// Type-level post-condition. The capture name is replaced with `result`
    /// (the decoded struct) in the assertion body.
    pub assert: Vec<AssertDef>,
}

#[derive(Debug, Clone, PartialEq)]
pub enum LengthSpec {
    /// Read a length/count prefix using this Solidity unsigned integer type (e.g. `"uint32"`).
    Type(String),
    /// The length/count is a compile-time Solidity expression; no prefix is read from the buffer.
    Constant(String),
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
    pub expr: String,
}

#[derive(Debug, Default, Clone, PartialEq)]
pub struct FieldArgs {
    pub method: Option<String>,
    pub default: bool,
    pub memory: bool,
    pub length: Option<LengthSpec>,
    /// Field-level post-conditions.
    pub assert: Vec<AssertDef>,
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
    pub assert: Vec<AssertDef>,
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
        assert: raw.assert,
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
    let args_end = args_start + find_closing_paren(&comment[args_start..])?;
    let args_str = strip_comment_prefixes(&comment[args_start..args_end]);

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();
    let mut calldata = false;
    let mut solhint_disable = false;
    let mut asserts: Vec<AssertDef> = Vec::new();

    for pair in split_args(&args_str) {
        let pair = pair.trim();
        if let Some((key, val)) = pair.split_once('=') {
            let val = val.trim().trim_matches('"');
            match key.trim() {
                "contract" => contract = Some(val.to_string()),
                "name" => name = Some(val.to_string()),
                "visibility" => visibility = val.to_string(),
                "assert" => {
                    if let Some((captured, var, expr)) = parse_closure(val) {
                        asserts.push(AssertDef {
                            captured,
                            var,
                            expr,
                        });
                    }
                }
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
        assert: asserts,
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
    let Some(args_end_rel) = find_closing_paren(&comment[args_start..]) else {
        return FieldArgs {
            memory,
            ..Default::default()
        };
    };
    let args_str = strip_comment_prefixes(comment[args_start..args_start + args_end_rel].trim());

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
    let mut asserts: Vec<AssertDef> = Vec::new();
    for arg in split_args(&args_str) {
        if let Some((key, val)) = arg.trim().split_once('=') {
            match key.trim() {
                "method" => method = Some(val.trim().trim_matches('"').to_string()),
                "length" => {
                    let v = val.trim();
                    let is_uint = v
                        .strip_prefix("uint")
                        .and_then(|s| s.parse::<u16>().ok())
                        .is_some_and(|n| (8..=256).contains(&n) && n % 8 == 0);
                    length = Some(if is_uint {
                        LengthSpec::Type(v.to_string())
                    } else {
                        LengthSpec::Constant(v.to_string())
                    });
                }
                "assert" => {
                    let inner = val.trim().trim_matches('"');
                    if let Some((captured, var, expr)) = parse_closure(inner) {
                        asserts.push(AssertDef {
                            captured,
                            var,
                            expr,
                        });
                    }
                }
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
            assert: asserts,
        }
    }
}

/// Strips comment-line prefixes (`//`, `///`, `*`) from a multi-line args string.
///
/// When `#[unpack(...)]` spans multiple lines each continuation line begins
/// with a comment prefix that must be removed before parsing keys and values.
/// Handles both formats that `reforge::get_comment` may produce:
///   - Newline-separated: `"\n//    arg1,\n//    arg2"`
///   - Concatenated (no newlines): `"//    arg1,//    arg2"`
///
/// Splits on `\n` and on `//` occurrences outside quoted strings, strips
/// leading whitespace and `*` block-comment markers from each segment, drops
/// empty segments, and joins the results with a single space.
fn strip_comment_prefixes(s: &str) -> String {
    let mut segments: Vec<&str> = Vec::new();
    let mut seg_start = 0;
    let mut in_quotes = false;
    let bytes = s.as_bytes();
    let mut i = 0;

    while i < bytes.len() {
        match bytes[i] {
            b'"' => {
                in_quotes = !in_quotes;
                i += 1;
            }
            b'\n' if !in_quotes => {
                segments.push(&s[seg_start..i]);
                i += 1;
                seg_start = i;
            }
            b'/' if !in_quotes && i + 1 < bytes.len() && bytes[i + 1] == b'/' => {
                segments.push(&s[seg_start..i]);
                i += 2;
                // skip optional third slash (///)
                if i < bytes.len() && bytes[i] == b'/' {
                    i += 1;
                }
                seg_start = i;
            }
            _ => i += 1,
        }
    }
    segments.push(&s[seg_start..]);

    segments
        .iter()
        .filter_map(|seg| {
            let t = seg.trim_start();
            let t = if let Some(r) = t.strip_prefix('*') {
                r.trim_start()
            } else {
                t
            };
            if t.is_empty() { None } else { Some(t) }
        })
        .collect::<Vec<_>>()
        .join(" ")
}

/// Finds the index of the `)` that closes the current paren level, skipping
/// over any nested `(...)` pairs and `"..."` quoted strings.
/// Returns `None` if no such `)` exists.
fn find_closing_paren(s: &str) -> Option<usize> {
    let mut depth = 0usize;
    let mut in_quotes = false;
    for (i, c) in s.char_indices() {
        match c {
            '"' => in_quotes = !in_quotes,
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
fn split_args(s: &str) -> Vec<&str> {
    let mut result = Vec::new();
    let mut start = 0;
    let mut in_quotes = false;
    for (i, c) in s.char_indices() {
        match c {
            '"' => in_quotes = !in_quotes,
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

/// Parses a closure of the form `|[each] capture| { body }`.
/// The `each` keyword (e.g. `|each x|`) sets `CaptureType::Element`; a plain name sets
/// `CaptureType::Field`. Returns `(CaptureType, var, body)`, or `None` if the syntax
/// doesn't match.
fn parse_closure(s: &str) -> Option<(CaptureType, String, String)> {
    let s = s.trim();
    let s = s.strip_prefix('|')?;
    let pipe_end = s.find('|')?;
    let capture_raw = s[..pipe_end].trim();
    let (captured, var) = match capture_raw
        .strip_prefix("each")
        .filter(|r| r.is_empty() || r.starts_with(char::is_whitespace))
    {
        Some(rest) => {
            let var = rest.trim();
            if var.is_empty() {
                return None;
            }
            (CaptureType::Element, var.to_string())
        }
        None => (CaptureType::Field, capture_raw.to_string()),
    };
    let rest = s[pipe_end + 1..].trim();
    let brace_start = rest.find('{')? + 1;
    // find the matching closing brace
    let inner = &rest[brace_start..];
    let mut depth = 0usize;
    let brace_end = inner.char_indices().find_map(|(i, c)| match c {
        '{' => {
            depth += 1;
            None
        }
        '}' if depth == 0 => Some(i),
        '}' => {
            depth -= 1;
            None
        }
        _ => None,
    })?;
    let body = inner[..brace_end].trim().to_string();
    Some((captured, var, body))
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
        assert!(raw.assert.is_empty());
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
        assert!(raw.assert.is_empty());
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
    fn test_assert() {
        // type-level assert, alone
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 }")]"#,
            &vec![(None, false); 1],
        )
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                expr: "r.version == 1".to_string(),
            }]
        );

        // assert can appear before other args (no longer required to be last)
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 }", calldata)]"#,
            &vec![(None, false); 1],
        )
        .unwrap();
        assert!(raw.calldata);
        assert_eq!(
            raw.assert,
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                expr: "r.version == 1".to_string(),
            }]
        );

        // multiple asserts
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { r.version == 1 }", assert = "|r| { r.x > 0 }")]"#,
            &vec![(None, false); 1],
        )
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![
                AssertDef {
                    captured: CaptureType::Field,
                    var: "r".to_string(),
                    expr: "r.version == 1".to_string(),
                },
                AssertDef {
                    captured: CaptureType::Field,
                    var: "r".to_string(),
                    expr: "r.x > 0".to_string(),
                },
            ]
        );

        // body containing a function call with commas and parens (safe inside quotes)
        let raw = parse_comment(
            r#"#[unpack(assert = "|r| { isValid(r, 0) }")]"#,
            &vec![(None, false); 1],
        )
        .unwrap();
        assert_eq!(
            raw.assert,
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "r".to_string(),
                expr: "isValid(r, 0)".to_string(),
            }]
        );

        // field-level assert
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|b| { b.length == 48 }")]"#.to_string()),
                true,
            )],
        )
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "b".to_string(),
                expr: "b.length == 48".to_string(),
            }]
        );

        // element-level assert using `each` keyword
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|each x| { x > 0 }")]"#.to_string()),
                false,
            )],
        )
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![AssertDef {
                captured: CaptureType::Element,
                var: "x".to_string(),
                expr: "x > 0".to_string(),
            }]
        );

        // `each` prefix on a variable name that starts with "each" is not treated as the keyword
        let raw = parse_comment(
            "#[unpack()]",
            &vec![(
                Some(r#"#[unpack(assert = "|eachItem| { eachItem > 0 }")]"#.to_string()),
                false,
            )],
        )
        .unwrap();
        assert_eq!(
            raw.fields[0].assert,
            vec![AssertDef {
                captured: CaptureType::Field,
                var: "eachItem".to_string(),
                expr: "eachItem > 0".to_string(),
            }]
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
                (
                    Some("#[unpack(length = BLST.BLS_SIGNATURE_LENGTH)]".to_string()),
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
                length: Some(LengthSpec::Type("uint32".to_string())),
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[4],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                default: false,
                memory: true,
                length: Some(LengthSpec::Type("uint64".to_string())),
                assert: vec![],
            }
        );
        assert_eq!(
            raw.fields[5],
            FieldArgs {
                method: None,
                default: false,
                memory: true,
                length: Some(LengthSpec::Constant(
                    "BLST.BLS_SIGNATURE_LENGTH".to_string()
                )),
                assert: vec![],
            }
        );
    }
}
