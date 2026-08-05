//! Helpers for attaching macro attribution to generated code.
//!
//! When the derived pack/unpack code produced by these macros fails to compile,
//! `solc` reports an error against the *generated* source. Reforge can rewrite
//! such errors to point back at the annotation that triggered the generation,
//! provided each insertion is tagged with a macro name and the original source
//! location that triggered it (see `reforge::AdjustmentEntry::with`).
//!
//! [`trigger_location`] computes that original location for a `#[pack(...)]` /
//! `#[unpack(...)]` annotation.

use reforge::MacroOriginalLocation;
use solar::ast::Span;
use solar::sema::Gcx;
use solar::sema::hir::SourceId;

/// Computes the original-source location of the `marker` (e.g. `"#[pack("`) inside
/// the doc comment that triggered macro expansion for the item at `item_span`.
///
/// The returned location is expressed in coordinates of the *original*, unmodified
/// source, which is what reforge expects for error attribution. If the marker cannot
/// be found (it should always be present, since the macro only runs when it is), the
/// location falls back to the start of the item definition. Returns `None` only when
/// the source file has no real on-disk path.
pub(crate) fn trigger_location(
    ctx: &Gcx,
    source_id: SourceId,
    item_span: Span,
    marker: &str,
) -> Option<MacroOriginalLocation> {
    let source = ctx.sources.get(source_id)?;
    let file = source.file.name.as_real()?.to_path_buf();
    let src = source.file.src.as_str();
    let item_offset = (item_span.lo().0 - source.file.start_pos.0) as usize;
    // Prefer the exact position of the triggering annotation within the preceding
    // doc comment; fall back to the item definition itself.
    let offset = src
        .get(..item_offset)
        .and_then(|before| before.rfind(marker))
        .unwrap_or(item_offset);
    let (line, col) = line_col(src, offset);
    Some(MacroOriginalLocation { file, line, col })
}

/// Returns the 1-based line and column of the byte `offset` within `src`.
fn line_col(src: &str, offset: usize) -> (usize, usize) {
    let before = &src[..offset];
    let line = before.bytes().filter(|&b| b == b'\n').count() + 1;
    let line_start = before.rfind('\n').map(|i| i + 1).unwrap_or(0);
    let col = offset - line_start + 1;
    (line, col)
}

#[cfg(test)]
mod tests {
    use super::line_col;

    #[test]
    fn line_col_first_line() {
        assert_eq!(line_col("hello world", 0), (1, 1));
        assert_eq!(line_col("hello world", 6), (1, 7));
    }

    #[test]
    fn line_col_multi_line() {
        let src = "line one\nline two\n// #[pack()]\nstruct Foo {}";
        let offset = src.find("#[pack(").unwrap();
        // Third line, column of `#` (after the `// ` prefix).
        assert_eq!(line_col(src, offset), (3, 4));
    }
}
