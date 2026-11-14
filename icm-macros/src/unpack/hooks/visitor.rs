use super::CaptureType;

/// How two adjacent segments should be joined when reconstructing the args string.
#[derive(Clone, Copy)]
pub(super) enum SegSep {
    /// Outside a quoted value — a newline is just a line-continuation marker.
    Space,
    /// Inside a quoted value — the newline is content-significant and must be
    /// preserved so that multi-line assert bodies keep their line structure.
    Newline,
}

impl SegSep {
    fn as_str(self) -> &'static str {
        match self {
            SegSep::Space => " ",
            SegSep::Newline => "\n",
        }
    }
}

/// One piece of the args string after comment prefixes have been stripped, together
/// with the separator that should precede it when the segments are joined.
pub(super) struct Segment<'a> {
    /// Separator to insert before this segment when joining (ignored for the
    /// first non-empty segment).
    pub(super) sep: SegSep,
    /// The text of this segment (annotation comment prefixes already excluded).
    pub(super) statement: &'a str,
}

/// Tracks whether the visitor is currently inside a Solidity comment inside a
/// closure body. Used by both `CommentVisitor` (to avoid mis-stripping `//`
/// that is content inside a string literal) and `ClosureVisitor` (to classify
/// lines as comment vs expression).
#[derive(Default)]
pub(super) enum CommentType {
    /// Inside a `//` line comment.
    Comment,
    /// Inside a `/* ... */` block comment.
    Multiline,
    /// Not inside any comment.
    #[default]
    None,
}

impl CommentType {
    pub(super) fn in_comment(&self) -> bool {
        !matches!(self, CommentType::None)
    }

    pub(super) fn is_multiline(&self) -> bool {
        matches!(self, CommentType::Multiline)
    }

    pub(super) fn set_comment(&mut self) {
        *self = CommentType::Comment;
    }

    pub(super) fn set_multiline(&mut self) {
        *self = CommentType::Multiline;
    }

    pub(super) fn exit_comment(&mut self) {
        *self = CommentType::None;
    }
}

/// Strips annotation comment prefixes (`//`, `///`, `* `) from a raw multi-line annotation
/// argument string, producing a flat list of [`Segment`]s ready to be joined into a single
/// expression string.
///
/// Lines inside a double-quoted value (e.g. a closure body) are treated as content rather
/// than comment syntax, so Solidity `//` and `/* */` inside those values are left intact.
pub(super) struct CommentVisitor<'a> {
    raw: &'a str,
    pos: usize,
    segments: Vec<Segment<'a>>,
    seg_start: usize,
    /// Separator to use for the *next* segment: `Space` outside a quoted value,
    /// `Newline` inside (preserves closure body line structure).
    seg_sep: SegSep,
    /// True at the logical start of a line (after `\n` or at position 0, possibly
    /// after whitespace). Used to distinguish a `//` annotation prefix from
    /// mid-line content such as `http://`.
    newline: bool,
    in_quotes: bool,
    brace_depth: usize,
    comment_tracker: CommentType,
}

impl<'a> CommentVisitor<'a> {
    pub(super) fn new(s: &'a str) -> Self {
        Self {
            raw: s,
            pos: 0,
            segments: Vec::new(),
            seg_start: 0,
            seg_sep: SegSep::Space,
            in_quotes: false,
            newline: true,
            brace_depth: 0,
            comment_tracker: CommentType::default(),
        }
    }

    /// Runs the visitor to completion and returns the collected segments.
    pub(super) fn parse(mut self) -> Vec<Segment<'a>> {
        while self.pos < self.raw.len() {
            self.handle_byte();
        }
        let end = self.raw.len();
        self.push_segment(end, 0);
        self.segments
    }

    /// Processes one logical unit starting at `self.pos`. Calls each `mark_*` helper in
    /// order, then advances `pos` by the number of bytes consumed (1–3 depending on the
    /// token at the current position).
    fn handle_byte(&mut self) {
        let pos = self.pos;
        let advance_by = self.advance_by(pos);
        let byte = self.raw.as_bytes()[pos];

        self.mark_in_quotes(byte);
        self.mark_brace_depth(byte);
        self.mark_segment(pos, advance_by);
        self.mark_comments(pos);
        // this should be done last so that it doesn't interfere with previous checks
        self.mark_newline(byte);

        self.pos += advance_by;
    }

    /// Returns how many bytes `handle_byte` should consume at `pos`. Multi-byte tokens
    /// (`*/`, `//`, `///`, `* `) are consumed in one step so that their constituent bytes
    /// are never individually re-examined as potential comment or segment boundaries.
    fn advance_by(&self, pos: usize) -> usize {
        let b = self.raw.as_bytes()[pos];
        match b {
            b'*' if self.comment_tracker.is_multiline() => match self.raw.as_bytes().get(pos + 1) {
                Some(b'/') => 2,
                _ => 1,
            },
            // Annotation block-comment continuation prefix `* ` or bare `*` at line start.
            b'*' if self.newline && !self.comment_tracker.is_multiline() => {
                match self.raw.as_bytes().get(pos + 1) {
                    Some(&b' ') => 2,
                    _ => 1,
                }
            }
            b'/' if (!self.in_quotes || self.newline)
                && self.raw.as_bytes().get(pos + 1) == Some(&b'/') =>
            {
                if self.raw.as_bytes().get(pos + 2) == Some(&b'/') {
                    3
                } else {
                    2
                }
            }
            b'/' if !self.comment_tracker.in_comment() => match self.raw.as_bytes().get(pos + 1) {
                Some(&b'/') | Some(&b'*') => 2,
                _ => 1,
            },
            _ => 1,
        }
    }

    /// Toggles `in_quotes` on each `"` that appears outside a brace block. Braces inside
    /// a quoted value are not Rust string delimiters, so the toggle is suppressed while
    /// `brace_depth > 0`.
    fn mark_in_quotes(&mut self, byte: u8) {
        if byte == b'"' && self.brace_depth == 0 {
            self.in_quotes = !self.in_quotes;
        }
    }

    /// Tracks the nesting depth of `{ }` pairs inside a quoted value. This depth is used
    /// by `mark_in_quotes` to suppress quote-toggle on a `"` that opens a Solidity string
    /// literal inside a closure body, and by `mark_comments` to gate Solidity comment
    /// tracking to lines that are actually inside a closure body.
    fn mark_brace_depth(&mut self, b: u8) {
        if b == b'{' && self.in_quotes {
            self.brace_depth += 1;
        } else if b == b'}' && self.in_quotes && self.brace_depth > 0 {
            self.brace_depth -= 1;
        }
    }

    /// Updates `comment_tracker` as the visitor crosses comment boundaries. Outside a
    /// quoted closure value this is a no-op — annotation `//` prefixes are handled by
    /// `mark_segment`, not here. Inside a quoted closure value (`in_quotes &&
    /// brace_depth > 0`) it detects `//` and `/* */` so that `mark_segment` does not
    /// misidentify Solidity comment syntax as an annotation prefix.
    fn mark_comments(&mut self, pos: usize) {
        match self.raw.as_bytes()[pos] {
            b'*' if self.comment_tracker.is_multiline() => {
                if let Some(b'/') = self.raw.as_bytes().get(pos + 1) {
                    self.push_comment();
                }
            }
            // `\n` ends a line comment (`//`) but NOT a block comment (`/* */`).
            b'\n' if !self.comment_tracker.is_multiline() => {
                self.push_comment();
            }
            // Track Solidity `//` and `/* */` inside a closure body (in_quotes && brace_depth > 0).
            // `!newline` ensures annotation `//` line-start prefixes are not mis-tracked —
            // they are stripped by mark_segment; a Solidity `//` only appears mid-line
            // (newline = false) after the annotation prefix has been consumed.
            b'/' if self.in_quotes
                && self.brace_depth > 0
                && !self.newline
                && !self.comment_tracker.in_comment() =>
            {
                match self.raw.as_bytes().get(pos + 1) {
                    Some(&b'/') => self.comment_tracker.set_comment(),
                    Some(&b'*') => self.comment_tracker.set_multiline(),
                    _ => {}
                }
            }
            _ => {}
        }
    }

    /// Clears the active Solidity comment state. Called when a comment terminator is
    /// reached (`\n` for `//`, `*/` for `/* */`).
    fn push_comment(&mut self) {
        self.comment_tracker.exit_comment();
    }

    /// Decides whether the byte at `pos` starts a segment boundary and, if so, closes
    /// the current segment and begins a new one. Boundaries are:
    /// - `* ` / `*` at line-start — block-comment continuation prefix, stripped entirely.
    /// - `\n` — ends the current line; the newline itself is excluded from both segments.
    ///   Sets `seg_sep` to `Newline` inside a quoted value and `Space` outside.
    /// - `//` / `///` at line-start (or inside quotes where `newline` is true) — annotation
    ///   comment prefix, stripped by advancing past the slashes.
    fn mark_segment(&mut self, pos: usize, advance_by: usize) {
        match self.raw.as_bytes()[pos] {
            // `* ` or `*` block-comment continuation prefix at start of line.
            b'*' if self.newline => self.push_segment(pos, advance_by),
            b'\n' => {
                // Skip the newline itself so it doesn't appear at the head of the next segment.
                self.push_segment(pos, 1);
                // Inside a quoted value the newline is content-significant:
                // subsequent segments on new lines should be joined with `\n`
                // so the multi-line structure survives into the expression.
                self.seg_sep = if self.in_quotes {
                    SegSep::Newline
                } else {
                    SegSep::Space
                };
            }
            // `//` or `///` annotation comment prefix: strip the leading slashes.
            b'/' if (!self.in_quotes || self.newline)
                && self.raw.as_bytes().get(pos + 1) == Some(&b'/') =>
            {
                self.push_segment(pos, advance_by);
            }
            _ => {}
        }
    }

    /// Closes the segment `[seg_start, pos)` and sets `seg_start` to `pos + skip`,
    /// effectively consuming `skip` bytes of delimiter that should not appear in either
    /// the current or the next segment.
    fn push_segment(&mut self, pos: usize, skip: usize) {
        self.segments.push(Segment {
            sep: self.seg_sep,
            statement: &self.raw[self.seg_start..pos],
        });
        self.seg_start = pos + skip;
    }

    /// Maintains the `newline` flag: set to `true` after `\n` and kept `true` through
    /// subsequent whitespace, cleared to `false` on the first non-whitespace byte. This
    /// distinguishes a `//` annotation prefix at the logical line-start from mid-line `//`
    /// such as in `http://`.
    fn mark_newline(&mut self, b: u8) {
        if b == b'\n' {
            self.newline = true;
        } else if !matches!(b, b' ' | b'\t') {
            self.newline = false;
        }
    }
}

/// Assembles a flat string from visitor segments, joining with the per-segment separator
/// and skipping empty statements.
pub(super) fn join_segments(segments: &[Segment<'_>]) -> String {
    let mut result = String::new();
    for seg in segments {
        let s = seg.statement;
        if s.is_empty() {
            continue;
        }
        if !result.is_empty() {
            result.push_str(seg.sep.as_str());
        }
        result.push_str(s);
    }
    result
}

/// Validates a capture-variable name: must be non-empty, start with a letter or
/// underscore, and contain only letters, digits, and underscores. Consistent with the
/// word-boundary definition in `substitute_var`.
fn validate_capture_name(name: &str) -> eyre::Result<()> {
    if name.is_empty() {
        eyre::bail!(
            "assert closure has an empty capture name; \
             expected `|name| {{ ... }}` or `|each name| {{ ... }}`"
        );
    }
    let mut chars = name.chars();
    let first_ok = chars.next().is_some_and(|c| c.is_alphabetic() || c == '_');
    if !first_ok || !chars.all(|c| c.is_alphanumeric() || c == '_') {
        eyre::bail!(
            "assert capture `{name}` is not a valid identifier; \
             must start with a letter or underscore and contain only \
             letters, digits, and underscores"
        );
    }
    Ok(())
}

/// Visitor for parsing a `|[each] capture| { body }` closure string.
///
/// The annotation comment prefixes have already been stripped by `CommentVisitor`, so
/// `//` and `/* */` encountered here are Solidity comments inside the closure body.
/// Lines before the first expression line are collected as `comment_lines`; all
/// remaining non-empty lines are `expr_lines`.
pub(super) struct ClosureVisitor<'a> {
    raw: &'a str,
    pos: usize,
    comment_tracker: CommentType,
    pub(super) comment_lines: Vec<&'a str>,
    pub(super) expr_lines: Vec<&'a str>,
}

impl<'a> ClosureVisitor<'a> {
    pub(super) fn new(s: &'a str) -> Self {
        Self {
            raw: s,
            pos: 0,
            comment_tracker: CommentType::default(),
            comment_lines: Vec::new(),
            expr_lines: Vec::new(),
        }
    }

    /// Parses `|[each] name|` from the current position, advances `pos` past the
    /// closing `|`, and returns `(CaptureType, var_name)`.
    pub(super) fn parse_capture(&mut self) -> Option<eyre::Result<(CaptureType, String)>> {
        let tail = &self.raw[self.pos..];
        let s = tail.trim_start();
        let trim_len = tail.len() - s.len();
        let s = s.strip_prefix('|')?;
        let pipe_end = s.find('|')?;
        let capture_raw = s[..pipe_end].trim();
        let result = match capture_raw
            .strip_prefix("each")
            .filter(|r| r.is_empty() || r.starts_with(char::is_whitespace))
        {
            Some(rest) => {
                let var = rest.trim();
                if let Err(e) = validate_capture_name(var) {
                    return Some(Err(e));
                }
                (CaptureType::Element, var.to_string())
            }
            None => {
                if let Err(e) = validate_capture_name(capture_raw) {
                    return Some(Err(e));
                }
                (CaptureType::Field, capture_raw.to_string())
            }
        };
        self.pos += trim_len + 1 + pipe_end + 1; // trim + '|' + capture + '|'
        Some(Ok(result))
    }

    /// Finds the `{ ... }` body from the current position, advances `pos` past the
    /// closing `}`, and returns the trimmed inner content.
    pub(super) fn extract_body(&mut self) -> Option<&'a str> {
        let tail = &self.raw[self.pos..];
        let s = tail.trim_start();
        let trim_len = tail.len() - s.len();
        let brace_start = s.find('{')? + 1;
        let inner = &s[brace_start..];
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
        self.pos += trim_len + brace_start + brace_end + 1;
        Some(inner[..brace_end].trim())
    }

    /// Handles one line while already inside a `/* ... */` block comment.
    /// Exits the block on `*/`, and captures any expression text that follows on the same line.
    fn mark_in_block_comment(&mut self, l: &'a str) {
        if let Some(close) = l.find("*/") {
            self.comment_lines.push(&l[..close + 2]);
            self.comment_tracker.exit_comment();
            let after = l[close + 2..].trim();
            if !after.is_empty() {
                self.expr_lines.push(after);
            }
        } else {
            self.comment_lines.push(l);
        }
    }

    /// Classifies a single trimmed, non-empty line into `comment_lines` or `expr_lines`.
    fn handle_line(&mut self, l: &'a str) {
        if self.comment_tracker.is_multiline() {
            self.mark_in_block_comment(l);
        } else if !self.expr_lines.is_empty() {
            self.expr_lines.push(l);
        } else if l.starts_with("//") {
            self.comment_lines.push(l);
        } else if l.starts_with("/*") {
            if let Some(close) = l.find("*/") {
                self.comment_lines.push(&l[..close + 2]);
                let after = l[close + 2..].trim();
                if !after.is_empty() {
                    self.expr_lines.push(after);
                }
            } else {
                self.comment_lines.push(l);
                self.comment_tracker.set_multiline();
            }
        } else {
            self.expr_lines.push(l);
        }
    }

    /// Splits `body` into trimmed non-empty lines and classifies each into
    /// [`Self::comment_lines`] or [`Self::expr_lines`]. Leading `//` and `/* */`
    /// lines before the first expression are treated as doc comments; everything
    /// else is treated as expression content.
    pub(super) fn classify_lines(&mut self, body: &'a str) {
        for line in body.lines() {
            let l = line.trim();
            if !l.is_empty() {
                self.handle_line(l);
            }
        }
    }
}
