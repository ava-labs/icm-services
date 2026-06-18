//! This module contains the tools the pack macros use to know when they should
//! be run and extract optional arguments that can be passed to them.
//! This is done by analyzing comment strings.

use reforge::PreprocessingData;
use solar::ast::Span;
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, SourceId};

pub struct PackArgs {
    pub contract: Option<ContractId>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs>,
    pub visibility: String,
}

#[derive(Debug, Clone, PartialEq)]
pub enum LengthSpec {
    /// Encode the length/count using the given Solidity unsigned integer type.
    Type(String),
    /// Omit the length/count prefix entirely.
    Drop,
}

#[derive(Debug, Default, Clone, PartialEq)]
pub struct FieldArgs {
    pub method: Option<String>,
    pub ignore: bool,
    pub length: Option<LengthSpec>,
}

/// Intermediate parsed args before HIR resolution. Contract is stored as a name
/// string rather than a ContractId.
pub(crate) struct RawPackArgs {
    pub contract: Option<String>,
    pub name: Option<String>,
    pub fields: Vec<FieldArgs>,
    pub visibility: String,
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

    let field_comments: Vec<Option<String>> = ctx
        .hir
        .structs()
        .find(|s| s.span == span)
        .map(|s| {
            s.fields
                .iter()
                .map(|&field_id| {
                    let field = ctx.hir.variable(field_id);
                    reforge::get_comment(ctx, source_id, field.span, data)
                })
                .collect()
        })
        .unwrap_or_default();

    let raw = parse_comment(&doc_comment, &field_comments)?;

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
                        "contract `{name}` specified in #[pack(contract=...)] was not found"
                    )));
                }
            }
        }
    };

    Some(Ok(PackArgs {
        contract,
        name: raw.name,
        fields: raw.fields,
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
    let args_end = args_start + find_closing_paren(&comment[args_start..])?;
    let args_str = strip_comment_prefixes(&comment[args_start..args_end]);

    let mut contract = None;
    let mut name = None;
    let mut visibility = "public".to_string();

    for pair in split_args(&args_str) {
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

fn parse_field_args(comment: Option<&str>) -> FieldArgs {
    let Some(comment) = comment else {
        return FieldArgs::default();
    };
    let Some(pack_start) = comment.find("#[pack(") else {
        return FieldArgs::default();
    };
    let args_start = pack_start + "#[pack(".len();
    let Some(args_end_rel) = find_closing_paren(&comment[args_start..]) else {
        return FieldArgs::default();
    };
    let args_str = strip_comment_prefixes(comment[args_start..args_start + args_end_rel].trim());

    if args_str == "ignore" {
        return FieldArgs {
            ignore: true,
            ..Default::default()
        };
    }

    let mut method = None;
    let mut ignore = false;
    let mut length = None;
    for arg in split_args(&args_str) {
        if let Some((key, val)) = arg.trim().split_once('=') {
            match key.trim() {
                "method" => method = Some(val.trim().trim_matches('"').to_string()),
                "length" => {
                    let v = val.trim();
                    length = Some(if v == "drop" {
                        LengthSpec::Drop
                    } else {
                        LengthSpec::Type(v.to_string())
                    });
                }
                _ => {}
            }
        } else if arg.trim() == "ignore" {
            ignore = true;
        }
    }
    if ignore {
        FieldArgs {
            ignore: true,
            ..Default::default()
        }
    } else {
        FieldArgs {
            method,
            ignore: false,
            length,
        }
    }
}

/// Strips comment-line prefixes (`//`, `///`, `*`) from a multi-line args string.
///
/// When `#[pack(...)]` spans multiple lines each continuation line begins
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
        let expected = vec![FieldArgs::default(); 4];
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
        let expected = vec![FieldArgs::default(); 4];
        assert_eq!(raw.fields, expected);
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
                length: None,
            }
        );
        assert_eq!(
            raw.fields[2],
            FieldArgs {
                method: None,
                ignore: true,
                length: None,
            }
        );
        assert_eq!(
            raw.fields[3],
            FieldArgs {
                method: None,
                ignore: false,
                length: Some(LengthSpec::Type("uint32".to_string())),
            }
        );
        assert_eq!(
            raw.fields[4],
            FieldArgs {
                method: Some("Foo.bar".to_string()),
                ignore: false,
                length: Some(LengthSpec::Type("uint64".to_string())),
            }
        );
        assert_eq!(
            raw.fields[5],
            FieldArgs {
                method: None,
                ignore: false,
                length: Some(LengthSpec::Drop),
            }
        );
    }
}
