use std::ops::Range;

/// Marker trait for the validation state of field args.
pub trait ValidationState {
    /// The type used to represent a length annotation in this state.
    type Length: Clone + std::fmt::Debug + PartialEq;
}

/// State marker: field args have not yet been validated against the HIR.
#[derive(Debug, Clone, PartialEq, Default)]
pub struct Unvalidated;

impl ValidationState for Unvalidated {
    type Length = String;
}

/// State marker: field args have been validated; `length` holds a [`LengthSpec`].
#[derive(Debug, Clone, PartialEq)]
pub struct Validated;

impl ValidationState for Validated {
    type Length = LengthSpec;
}

/// Represents a Solidity integer type used as a length/count prefix.
/// The default is `Uint(32)`, corresponding to the implicit `uint256` encoding.
#[derive(Clone, Copy, Debug, PartialEq)]
pub enum LengthType {
    /// An unsigned integer type (e.g. `uint32`), with the byte width as the inner value.
    Uint(usize),
    /// A signed integer type (e.g. `int32`), with the byte width as the inner value.
    Int(usize),
}

impl Default for LengthType {
    fn default() -> Self {
        Self::Uint(32) // uint256
    }
}

impl LengthType {
    /// Returns the byte width of this type (e.g. 4 for `uint32` / `int32`).
    pub fn byte_width(self) -> usize {
        match self {
            Self::Uint(w) | Self::Int(w) => w,
        }
    }

    /// Returns the canonical Solidity type name for this type (e.g. `"uint32"`, `"int64"`).
    pub fn name(self) -> String {
        match self {
            Self::Uint(w) => format!("uint{}", w * 8),
            Self::Int(w) => format!("int{}", w * 8),
        }
    }

    /// Parses and validates a Solidity integer type string (e.g. `"uint32"`, `"int64"`).
    /// Bare `uint` and `int` (without a suffix) are treated as aliases for `uint256`/`int256`.
    /// Returns an error if the string is not a valid Solidity integer type.
    /// Use [`eyre::WrapErr`] at the call site to attach field/struct context to the error.
    pub fn validate(ty: &str) -> eyre::Result<Self> {
        let parse_bits = |prefix: &str, ctor: fn(usize) -> Self| -> Option<eyre::Result<Self>> {
            let rest = ty.strip_prefix(prefix)?;
            let bits: u16 = if rest.is_empty() {
                256
            } else {
                rest.parse().ok().filter(|&n: &u16| {
                    (8..=256).contains(&n) && n % 8 == 0 && n.to_string() == rest
                })?
            };
            Some(Ok(ctor(bits as usize / 8)))
        };

        parse_bits("uint", Self::Uint)
            .or_else(|| parse_bits("int", Self::Int))
            .unwrap_or_else(|| Err(eyre::eyre!("`{ty}` is not a valid Solidity integer type")))
    }
}

/// How the length/count of a dynamically sized field is encoded on the wire.
#[derive(Debug, Clone, PartialEq)]
pub enum LengthSpec {
    /// The payload is prefixed with its length/count, encoded as this Solidity integer type.
    Type(LengthType),
    /// No prefix: the length/count is this compile-time Solidity expression (e.g. `32` or
    /// `BLST.SIG_LENGTH`). The pack side asserts that the value has exactly this length and the
    /// unpack side reads exactly this much, so both sides must agree on the constant.
    Constant(String),
}

impl Default for LengthSpec {
    /// The default is a `uint256` prefix.
    fn default() -> Self {
        Self::Type(LengthType::default())
    }
}

impl LengthSpec {
    /// Resolves the raw `length` value of a field annotation. A value naming a Solidity integer
    /// type selects a prefix of that type; any other non-empty value is taken as a constant
    /// length expression, which only `solc` can check.
    pub fn validate(value: String, field_name: &str, struct_name: &str) -> eyre::Result<Self> {
        if value.is_empty() {
            eyre::bail!(
                "empty `length` for field `{field_name}` of `{struct_name}`: expected a Solidity \
                 integer type or a constant length expression"
            );
        }
        if let Ok(ty) = LengthType::validate(&value) {
            return Ok(Self::Type(ty));
        }
        Ok(Self::Constant(value))
    }
}

/// Splits the argument list of the `{open}...)` annotation in `comment` (`open` being e.g.
/// `#[pack(`) on top-level commas, returning each argument trimmed. Parentheses are balanced and
/// quoted strings are skipped, so constant expressions such as `type(uint16).max` or
/// `Lib.f(a, b)` survive intact. Returns `Ok(None)` if the annotation is absent, and an error if
/// it is present but a quoted string or its parentheses never close: a mistyped annotation must
/// not be silently treated as a missing one.
pub fn annotation_args<'a>(comment: &'a str, open: &str) -> eyre::Result<Option<Vec<&'a str>>> {
    let Some(annotation_start) = comment.find(open) else {
        return Ok(None);
    };
    let start = annotation_start + open.len();
    let mut depth = 0usize;
    let mut quoted = false;
    // Tracks `{...}` depth inside a quoted value. While a closure body like `"|b| { expr }"`
    // is inside its outer quotes, inner `"..."` Solidity string literals must not toggle
    // `quoted`. The rule mirrors `split_args` and `find_closing_paren`: quote-toggling is
    // suppressed while `brace_depth > 0`, and braces only count while already quoted.
    let mut brace_depth = 0usize;
    let mut args = Vec::new();
    let mut arg_start = start;
    for (i, byte) in comment.bytes().enumerate().skip(start) {
        match byte {
            b'"' if brace_depth == 0 => quoted = !quoted,
            b'{' if quoted => brace_depth += 1,
            b'}' if quoted && brace_depth > 0 => brace_depth -= 1,
            _ if quoted => {}
            b'(' => depth += 1,
            b')' if depth > 0 => depth -= 1,
            b')' => {
                args.push(comment[arg_start..i].trim());
                return Ok(Some(args));
            }
            b',' if depth == 0 => {
                args.push(comment[arg_start..i].trim());
                arg_start = i + 1;
            }
            _ => {}
        }
    }
    // Quote the offending line rather than the whole comment block.
    let annotation = comment[annotation_start..]
        .lines()
        .next()
        .unwrap_or("")
        .trim_end();
    let problem = if quoted || brace_depth > 0 {
        "unterminated string"
    } else {
        "unbalanced parentheses"
    };
    eyre::bail!("{problem} in annotation `{annotation}`")
}

/// Whether the closure variable captures a container element or the whole decoded value.
#[derive(Debug, Clone, PartialEq, Default)]
pub enum CaptureType {
    /// `|each var|` — variable is bound to each element inside the decode loop.
    Element,
    /// `|var|` — variable is bound to the whole field or decoded struct.
    #[default]
    Field,
}

#[derive(Debug, Clone, PartialEq)]
pub struct AssertDef {
    pub captured: CaptureType,
    /// Byte ranges within `expr` where the capture variable should be substituted.
    pub var_spans: Vec<Range<usize>>,
    /// Inline comment lines that precede this assertion's expression in the
    /// closure body. Empty when there is no leading comment.
    pub comment: String,
    pub expr: String,
}

/// Tracks whether the visitor is currently inside a Solidity comment inside a
/// closure body. Used by both `CommentVisitor` (to avoid mis-stripping `//`
/// that is content inside a string literal) and `ClosureVisitor` (to classify
/// lines as comment vs expression).
#[derive(Default)]
pub(crate) enum CommentType {
    /// Inside a `//` line comment.
    Comment,
    /// Inside a `/* ... */` block comment.
    Multiline,
    /// Not inside any comment.
    #[default]
    None,
}

impl CommentType {
    pub(crate) fn in_comment(&self) -> bool {
        !matches!(self, CommentType::None)
    }

    pub(crate) fn is_multiline(&self) -> bool {
        matches!(self, CommentType::Multiline)
    }

    pub(crate) fn set_comment(&mut self) {
        *self = CommentType::Comment;
    }

    pub(crate) fn set_multiline(&mut self) {
        *self = CommentType::Multiline;
    }

    pub(crate) fn exit_comment(&mut self) {
        *self = CommentType::None;
    }
}

/// Validates a capture-variable name: must be non-empty, start with a letter or
/// underscore, and contain only letters, digits, and underscores.
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
struct ClosureVisitor<'a> {
    raw: &'a str,
    /// The capture variable name, set by `parse_capture` and used by `push_line`.
    var_name: Option<&'a str>,
    pos: usize,
    comment_tracker: CommentType,
    comment_lines: Vec<&'a str>,
    expr_lines: Vec<&'a str>,
    /// Running byte offset of the next line's start in the eventual `expr_lines.join(" ")`.
    expr_offset: usize,
    /// Byte ranges within `expr_lines.join(" ")` where the capture variable should be
    /// substituted. Populated by `push_line`.
    var_spans: Vec<Range<usize>>,
}

impl<'a> ClosureVisitor<'a> {
    fn new(s: &'a str) -> Self {
        Self {
            raw: s,
            pos: 0,
            var_name: None,
            comment_tracker: CommentType::default(),
            comment_lines: Vec::new(),
            expr_lines: Vec::new(),
            expr_offset: 0,
            var_spans: Vec::new(),
        }
    }

    /// Parses `|[each] name|` from the current position, advances `pos` past the
    /// closing `|`, and returns `(CaptureType, var_name)`.
    fn parse_capture(&mut self) -> Option<eyre::Result<(CaptureType, String)>> {
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
                self.var_name = Some(var);
                (CaptureType::Element, var.to_string())
            }
            None => {
                if let Err(e) = validate_capture_name(capture_raw) {
                    return Some(Err(e));
                }
                self.var_name = Some(capture_raw);
                (CaptureType::Field, capture_raw.to_string())
            }
        };
        self.pos += trim_len + 1 + pipe_end + 1; // trim + '|' + capture + '|'
        Some(Ok(result))
    }

    /// Returns the unprocessed tail of `raw` after the last advance.
    fn remaining(&self) -> &'a str {
        &self.raw[self.pos..]
    }

    /// Finds the `{ ... }` body from the current position, advances `pos` past the
    /// closing `}`, and returns the trimmed inner content.
    ///
    /// Returns an error if there is any non-whitespace text between the current
    /// position and the opening `{`, or if no `{` is found at all.
    fn extract_body(&mut self) -> eyre::Result<&'a str> {
        let tail = &self.raw[self.pos..];
        let s = tail.trim_start();
        let trim_len = tail.len() - s.len();
        if !s.starts_with('{') {
            eyre::bail!(
                "unexpected text before closure body `{{`; \
                 expected `|name| {{ expr }}`, found `{s}`"
            );
        }
        let brace_start = 1; // s starts with '{'
        let inner = &s[brace_start..];
        let mut depth = 0usize;
        let brace_end = inner
            .char_indices()
            .find_map(|(i, c)| match c {
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
            })
            .ok_or_else(|| eyre::eyre!("closure body is missing closing `}}`"))?;
        self.pos += trim_len + brace_start + brace_end + 1;
        Ok(inner[..brace_end].trim())
    }

    /// Strips a trailing `//`-style line comment, returning only the code that
    /// precedes it. `//` inside a Solidity string literal (`"..."` or `'...'`,
    /// with `\` escape handling) or inside a `/* ... */` block comment is not
    /// treated as a comment start.
    fn strip_line_comment(l: &'a str) -> &'a str {
        let mut quote: Option<char> = None;
        let mut escaped = false;
        let mut in_block = false;
        let mut skip_next = false;
        for (i, c) in l.char_indices() {
            if skip_next {
                skip_next = false;
                continue;
            }
            if escaped {
                escaped = false;
                continue;
            }
            if in_block {
                if c == '*' && l.as_bytes().get(i + 1) == Some(&b'/') {
                    in_block = false;
                    skip_next = true;
                }
                continue;
            }
            match quote {
                None => match c {
                    '"' | '\'' => quote = Some(c),
                    '/' if l.as_bytes().get(i + 1) == Some(&b'*') => {
                        in_block = true;
                        skip_next = true;
                    }
                    '/' if l.as_bytes().get(i + 1) == Some(&b'/') => {
                        return l[..i].trim_end();
                    }
                    _ => {}
                },
                Some(q) if c == q => quote = None,
                Some(_) if c == '\\' => escaped = true,
                _ => {}
            }
        }
        l
    }

    /// Handles one line while already inside a `/* ... */` block comment.
    /// Exits the block on `*/`, and captures any expression text that follows on the same line.
    fn mark_in_block_comment(&mut self, l: &'a str) {
        if let Some(close) = l.find("*/") {
            self.comment_lines.push(&l[..close + 2]);
            self.comment_tracker.exit_comment();
            let after = Self::strip_line_comment(l[close + 2..].trim());
            if !after.is_empty() {
                self.push_line(after);
            }
        } else {
            self.comment_lines.push(l);
        }
    }

    /// Classifies a single trimmed, non-empty line into `comment_lines` or `expr_lines`.
    /// Leading `//` and `/* */` lines before the first expression become `comment_lines`.
    /// Once expression lines have started, standalone `//` and `/* */` comment lines are
    /// silently dropped; inline `//` comments are stripped from expression lines.
    fn handle_line(&mut self, l: &'a str) {
        if self.comment_tracker.is_multiline() {
            self.mark_in_block_comment(l);
        } else if !self.expr_lines.is_empty() {
            if l.starts_with("//") {
                // trailing standalone line comment — drop it
            } else if l.starts_with("/*") {
                if let Some(close) = l.find("*/") {
                    // inline block comment on its own line; grab any code after `*/`
                    let after = Self::strip_line_comment(l[close + 2..].trim());
                    if !after.is_empty() {
                        self.push_line(after);
                    }
                } else {
                    self.comment_tracker.set_multiline();
                }
            } else {
                self.push_line(Self::strip_line_comment(l));
            }
        } else if l.starts_with("//") {
            self.comment_lines.push(l);
        } else if l.starts_with("/*") {
            if let Some(close) = l.find("*/") {
                self.comment_lines.push(&l[..close + 2]);
                let after = Self::strip_line_comment(l[close + 2..].trim());
                if !after.is_empty() {
                    self.push_line(after);
                }
            } else {
                self.comment_lines.push(l);
                self.comment_tracker.set_multiline();
            }
        } else {
            self.push_line(Self::strip_line_comment(l));
        }
    }

    /// Splits `body` into trimmed non-empty lines and classifies each into
    /// `comment_lines` or `expr_lines`. Leading `//` and `/* */`
    /// lines before the first expression are treated as doc comments; everything
    /// else is treated as expression content.
    fn classify_lines(&mut self, body: &'a str) {
        for line in body.lines() {
            let l = line.trim();
            if !l.is_empty() {
                self.handle_line(l);
            }
        }
    }

    /// Pushes `line` onto `expr_lines` and records byte ranges within the eventual
    /// `expr_lines.join(" ")` where `var_name` should be substituted.
    ///
    /// Occurrences inside Solidity string literals (`"..."` / `'...'`, with `\` escape
    /// handling) and non-word-boundary hits are skipped. Solidity strings cannot span
    /// lines, so string-literal state resets fresh for each call.
    fn push_line(&mut self, line: &'a str) {
        if let Some(var) = self.var_name {
            let line_start = self.expr_offset;
            let mut remaining = line;
            let mut inner_offset = 0usize;
            let mut quote: Option<char> = None;
            let mut escaped = false;
            let mut in_block = false;
            // Before the first expr line there is no leading separator; for subsequent
            // lines `join(" ")` inserts a space, which counts as a word boundary.
            let mut last_char: Option<char> = if self.expr_lines.is_empty() {
                None
            } else {
                Some(' ')
            };

            while let Some(pos) = remaining.find(var) {
                let before = &remaining[..pos];
                let after = &remaining[pos + var.len()..];

                // Advance string-literal and block-comment state over `before`.
                let in_non_code = {
                    let mut skip_next = false;
                    for (i, c) in before.char_indices() {
                        if skip_next {
                            skip_next = false;
                            continue;
                        }
                        if escaped {
                            escaped = false;
                            continue;
                        }
                        if in_block {
                            if c == '*' && before.as_bytes().get(i + 1) == Some(&b'/') {
                                in_block = false;
                                skip_next = true;
                            }
                            continue;
                        }
                        match quote {
                            None => {
                                if c == '"' || c == '\'' {
                                    quote = Some(c);
                                } else if c == '/'
                                    && before.as_bytes().get(i + 1) == Some(&b'*')
                                {
                                    in_block = true;
                                    skip_next = true;
                                }
                            }
                            Some(q) if c == q => quote = None,
                            Some(_) if c == '\\' => escaped = true,
                            _ => {}
                        }
                    }
                    quote.is_some() || in_block
                };

                let left_char = if before.is_empty() {
                    last_char
                } else {
                    before.chars().last()
                };
                let is_ident = |c: char| c.is_alphanumeric() || c == '_';
                let left_ok = left_char.is_none_or(|c| !is_ident(c) && c != '.');
                let right_ok = after.chars().next().is_none_or(|c| !is_ident(c));

                if left_ok && right_ok && !in_non_code {
                    let abs_start = line_start + inner_offset + pos;
                    self.var_spans.push(abs_start..abs_start + var.len());
                }

                last_char = var.chars().last();
                inner_offset += pos + var.len();
                remaining = after;
            }
        }
        // +1 accounts for the space `join(" ")` inserts before the next line.
        self.expr_offset += line.len() + 1;
        self.expr_lines.push(line);
    }
}

/// Whether the annotation lives inside a `//`-style or `/* */`-style comment.
/// Used by [`CommentVisitor`] to decide which line-start prefix to strip inside
/// a quoted closure body.
pub(crate) enum AnnotationStyle {
    /// Annotation is in a `//`-style comment, or inside a single-line `/* ... */`
    /// block that closes before the first `\n`. In either case, `//` at line-start
    /// inside a quoted closure body is an annotation-continuation prefix and should
    /// be stripped.
    LineComment,
    /// Annotation is in a multi-line `/* */`-style comment; `//` at line-start
    /// inside a quoted closure body is a Solidity comment, not a prefix.
    BlockComment,
}

impl AnnotationStyle {
    /// Detects the style from the full comment string. Returns
    /// [`AnnotationStyle::BlockComment`] when `comment` starts with `/*` and
    /// the closing `*/` does not appear before the first `\n` (i.e. the block
    /// spans multiple lines). Everything else — `//` comments, single-line
    /// `/* */`, or bare annotations — is [`AnnotationStyle::LineComment`].
    fn new(comment: &str) -> Self {
        let trimmed = comment.trim_start();
        if trimmed.starts_with("/*") {
            let first_newline = trimmed.find('\n').unwrap_or(trimmed.len());
            if !trimmed[..first_newline].contains("*/") {
                return AnnotationStyle::BlockComment;
            }
        }
        AnnotationStyle::LineComment
    }

    fn is_line(&self) -> bool {
        matches!(self, AnnotationStyle::LineComment)
    }

    fn is_block(&self) -> bool {
        matches!(self, AnnotationStyle::BlockComment)
    }
}

/// How two adjacent segments should be joined when reconstructing the args string.
#[derive(Clone, Copy)]
pub(crate) enum SegSep {
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
pub(crate) struct Segment<'a> {
    /// Separator to insert before this segment when joining (ignored for the
    /// first non-empty segment).
    pub(crate) sep: SegSep,
    /// The text of this segment (annotation comment prefixes already excluded).
    pub(crate) statement: &'a str,
}

/// Strips annotation comment prefixes (`//`, `///`, `* `) from a raw multi-line annotation
/// argument string, producing a flat list of [`Segment`]s ready to be joined into a single
/// expression string.
///
/// Lines inside a double-quoted value (e.g. a closure body) are treated as content rather
/// than comment syntax, so Solidity `//` and `/* */` inside those values are left intact.
pub(crate) struct CommentVisitor<'a> {
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
    annotation_style: AnnotationStyle,
}

impl<'a> CommentVisitor<'a> {
    pub(crate) fn new(full_comment: &'a str, start: usize, end: usize) -> Self {
        Self {
            raw: &full_comment[start..end],
            pos: 0,
            segments: Vec::new(),
            seg_start: 0,
            seg_sep: SegSep::Space,
            in_quotes: false,
            newline: true,
            brace_depth: 0,
            comment_tracker: CommentType::default(),
            annotation_style: AnnotationStyle::new(full_comment),
        }
    }

    /// Runs the visitor to completion and returns the collected segments.
    pub(crate) fn parse(mut self) -> Vec<Segment<'a>> {
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
            // Inside a quoted value, only consume as a prefix unit when the annotation uses
            // block-comment style; in a line-comment annotation a `*` here is Solidity content.
            b'*' if self.newline
                && !self.comment_tracker.is_multiline()
                && (!self.in_quotes || self.annotation_style.is_block()) =>
            {
                match self.raw.as_bytes().get(pos + 1) {
                    Some(&b' ') => 2,
                    _ => 1,
                }
            }
            // Annotation line-comment prefix `//` / `///` at line-start, or outside quotes.
            // Inside a quoted value, only consume as a prefix unit when the annotation uses
            // line-comment style; in a block-comment annotation a `//` here is a Solidity comment.
            b'/' if (!self.in_quotes || (self.newline && self.annotation_style.is_line()))
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
            // Inside a quoted value, only strip when the annotation uses block-comment style;
            // in a line-comment annotation a `*` at line-start is Solidity content.
            b'*' if self.newline && (!self.in_quotes || self.annotation_style.is_block()) => {
                self.push_segment(pos, advance_by)
            }
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
            // Inside a quoted value, only strip when the annotation uses line-comment style;
            // in a block-comment annotation a `//` at line-start is a Solidity comment.
            b'/' if (!self.in_quotes || (self.newline && self.annotation_style.is_line()))
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
pub(crate) fn join_segments(segments: &[Segment<'_>]) -> String {
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

/// Finds the index of the `)` that closes the current paren level, skipping
/// over any nested `(...)` pairs and `"..."` quoted strings.
/// Returns `None` if no such `)` exists.
pub(crate) fn find_closing_paren(s: &str) -> Option<usize> {
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

/// Splits `s` on commas that are outside of `"..."` quoted strings and outside
/// unquoted parentheses (so `Lib.f(a, b)` is not split on the inner comma).
pub(crate) fn split_args(s: &str) -> Vec<&str> {
    let mut result = Vec::new();
    let mut start = 0;
    let mut in_quotes = false;
    let mut brace_depth = 0usize;
    let mut paren_depth = 0usize;
    for (i, c) in s.char_indices() {
        match c {
            '"' if brace_depth == 0 => in_quotes = !in_quotes,
            '{' if in_quotes => brace_depth += 1,
            '}' if in_quotes && brace_depth > 0 => brace_depth -= 1,
            '(' if !in_quotes => paren_depth += 1,
            ')' if !in_quotes && paren_depth > 0 => paren_depth -= 1,
            ',' if !in_quotes && paren_depth == 0 => {
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
pub(crate) fn parse_closure(s: &str) -> Option<eyre::Result<Vec<AssertDef>>> {
    let mut visitor = ClosureVisitor::new(s.trim());
    let (captured, var) = match visitor.parse_capture() {
        Some(Ok(t)) => t,
        Some(Err(e)) => return Some(Err(e)),
        None => return None,
    };
    let body = match visitor.extract_body() {
        Ok(b) => b,
        Err(e) => return Some(Err(e)),
    };
    let tail = visitor.remaining().trim();
    if !tail.is_empty() {
        return Some(Err(eyre::eyre!(
            "unexpected text after closure body `}}`; \
             expected `|name| {{ expr }}`, found `{tail}`"
        )));
    }
    visitor.classify_lines(body);
    // Strip a trailing `;` so that `|r| { r.a > 0; }` does not emit `require(r.a > 0;)`.
    // var_spans are unaffected: they are byte ranges into the joined string that all precede
    // the semicolon.
    let expr = visitor
        .expr_lines
        .join(" ")
        .trim_end_matches(|c: char| c == ';' || c.is_whitespace())
        .to_string();
    if expr.is_empty() {
        return Some(Err(eyre::eyre!(
            "assert closure body for `{var}` contains no boolean expressions"
        )));
    }
    Some(Ok(vec![AssertDef {
        captured,
        comment: visitor.comment_lines.join("\n"),
        var_spans: visitor.var_spans,
        expr,
    }]))
}

#[cfg(test)]
pub mod tests {
    use super::*;

    /// Convenience constructor: builds an `AssertDef` with `var_spans` derived from
    /// parsing `closure_src` through the visitor, so tests don't hardcode byte offsets.
    pub fn assert_def(
        captured: CaptureType,
        closure_src: &str,
        comment: &str,
        expr: &str,
    ) -> AssertDef {
        let defs = parse_closure(closure_src)
            .expect("closure present")
            .expect("closure valid");
        assert_eq!(defs.len(), 1);
        let d = defs.into_iter().next().unwrap();
        assert_eq!(d.captured, captured);
        assert_eq!(d.comment, comment);
        assert_eq!(d.expr, expr);
        d
    }

    #[test]
    fn test_assert_capture_errors() {
        // empty capture
        let err = parse_closure(r#"|| { true }"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(err.contains("empty capture name"), "wrong error: {err}");

        // non-identifier capture name
        let err = parse_closure(r#"|2bad| { true }"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(err.contains("not a valid identifier"), "wrong error: {err}");

        // empty capture after `each`
        let err = parse_closure(r#"|each | { true }"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(err.contains("empty capture name"), "wrong error: {err}");

        // missing braces
        let err = parse_closure(r#"|x| x > 0"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unexpected text before closure body"),
            "wrong error: {err}"
        );

        // non-whitespace after closing `}`
        let err = parse_closure(r#"|r| { r.a > 0 } && r.b > 0"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(
            err.contains("unexpected text after closure body"),
            "wrong error: {err}"
        );
    }

    #[test]
    fn test_assert_empty_body() {
        // An empty closure body must be an error, not a silent no-op.
        let err = parse_closure(r#"|r| { }"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(err.contains("no boolean expressions"), "wrong error: {err}");

        // A body containing only a comment must also be an error.
        let err = parse_closure(r#"|r| { // just a comment }"#)
            .unwrap()
            .unwrap_err()
            .to_string();
        assert!(err.contains("no boolean expressions"), "wrong error: {err}");
    }

    #[test]
    fn test_assert_solidity_string_literals_via_closure() {
        // Collision: var name appears both as a real reference and inside a string
        // literal. Only the real reference should become a span.
        let defs = parse_closure(r#"|s| { keccak256(bytes(s.name)) != keccak256("s") }"#)
            .unwrap()
            .unwrap();
        let def = &defs[0];
        assert_eq!(def.expr, r#"keccak256(bytes(s.name)) != keccak256("s")"#);
        assert_eq!(
            def.var_spans.len(),
            1,
            "expected exactly one span, got {:?}",
            def.var_spans
        );
        assert_eq!(&def.expr[def.var_spans[0].clone()], "s");

        // Escaped quote: `"\""` must not close the string early.
        let defs = parse_closure(r#"|s| { s.ok && keccak256("\"s\"") == h }"#)
            .unwrap()
            .unwrap();
        let def = &defs[0];
        assert_eq!(
            def.var_spans.len(),
            1,
            "expected exactly one span, got {:?}",
            def.var_spans
        );
        assert_eq!(&def.expr[def.var_spans[0].clone()], "s");

        // `'` inside a `/* ... */` block comment must not open a string literal:
        // capture names after the comment must still be detected.
        let defs = parse_closure("|r| { r.ok && /* r's note */ r.value > 0 }")
            .unwrap()
            .unwrap();
        let def = &defs[0];
        assert_eq!(def.expr, "r.ok && /* r's note */ r.value > 0");
        assert_eq!(
            def.var_spans.len(),
            2,
            "expected two spans, got {:?}",
            def.var_spans
        );
        assert_eq!(&def.expr[def.var_spans[0].clone()], "r");
        assert_eq!(&def.expr[def.var_spans[1].clone()], "r");
    }

    #[test]
    fn test_assert_trailing_semicolon() {
        // A trailing `;` on the expression must be stripped so that the generated
        // `require(expr)` is valid Solidity (not `require(expr;)`).
        let defs = parse_closure("|r| { r.a > 0; }").unwrap().unwrap();
        assert_eq!(defs[0].expr, "r.a > 0");

        // Semicolons in the middle of an expression are left alone.
        let defs = parse_closure("|r| { r.a > 0 }").unwrap().unwrap();
        assert_eq!(defs[0].expr, "r.a > 0");
    }

    #[test]
    fn test_assert_trailing_line_comments() {
        // Inline `//` comment at the end of the expression line must be stripped so
        // that the closing `);` of the generated `require()` is not swallowed.
        let defs = parse_closure("|b| { b.length > 0 // non-empty }")
            .unwrap()
            .unwrap();
        assert_eq!(defs[0].expr, "b.length > 0");

        // Standalone `//` comment line after the expression must be dropped.
        let defs = parse_closure("|b| {\n    b.length > 0\n    // trailing comment\n}")
            .unwrap()
            .unwrap();
        assert_eq!(defs[0].expr, "b.length > 0");

        // Continuation lines after a trailing `//` are still captured.
        let defs =
            parse_closure("|b| {\n    b.length > 0\n    // mid comment\n    && b.length < 100\n}")
                .unwrap()
                .unwrap();
        assert_eq!(defs[0].expr, "b.length > 0 && b.length < 100");

        // `//` inside a string literal must not be treated as a comment start.
        let defs = parse_closure(r#"|b| { b == "http://x" // trailing }"#)
            .unwrap()
            .unwrap();
        assert_eq!(defs[0].expr, r#"b == "http://x""#);

        // `//` inside a `/* ... */` inline block comment must not cut the expression:
        // `/* see https://x */` should be preserved in the expr, not truncated at `//`.
        let defs = parse_closure("|r| { r.value > 0 /* see https://x.com/path */ }")
            .unwrap()
            .unwrap();
        assert_eq!(defs[0].expr, "r.value > 0 /* see https://x.com/path */");
    }

    #[test]
    fn test_validate_canonical() {
        assert_eq!(LengthType::validate("uint32").unwrap(), LengthType::Uint(4));
        assert_eq!(
            LengthType::validate("uint256").unwrap(),
            LengthType::Uint(32)
        );
        assert_eq!(LengthType::validate("uint").unwrap(), LengthType::Uint(32));
        assert_eq!(LengthType::validate("int64").unwrap(), LengthType::Int(8));
        assert_eq!(LengthType::validate("int").unwrap(), LengthType::Int(32));
    }

    #[test]
    fn test_validate_rejects_non_canonical() {
        assert!(LengthType::validate("uint032").is_err(), "leading zero");
        assert!(LengthType::validate("uint+32").is_err(), "leading plus");
        assert!(
            LengthType::validate("int032").is_err(),
            "leading zero (int)"
        );
        assert!(
            LengthType::validate("int+32").is_err(),
            "leading plus (int)"
        );
    }

    #[test]
    fn test_validate_rejects_invalid() {
        assert!(LengthType::validate("bool").is_err());
        assert!(LengthType::validate("uint0").is_err(), "zero bits");
        assert!(LengthType::validate("uint7").is_err(), "not multiple of 8");
        assert!(LengthType::validate("uint264").is_err(), "too large");
        assert!(
            LengthType::validate("uint257").is_err(),
            "not multiple of 8 and too large"
        );
    }

    #[test]
    fn test_length_spec_validate() {
        let validate = |value: &str| LengthSpec::validate(value.to_string(), "data", "Foo");
        assert_eq!(
            validate("uint32").unwrap(),
            LengthSpec::Type(LengthType::Uint(4))
        );
        assert_eq!(
            validate("32").unwrap(),
            LengthSpec::Constant("32".to_string())
        );
        assert_eq!(
            validate("BLST.SIG_LENGTH").unwrap(),
            LengthSpec::Constant("BLST.SIG_LENGTH".to_string())
        );
        assert_eq!(
            validate("type(uint16).max").unwrap(),
            LengthSpec::Constant("type(uint16).max".to_string())
        );
        let err = validate("").unwrap_err().to_string();
        assert!(
            err.contains("empty `length` for field `data` of `Foo`"),
            "{err}"
        );
    }

    #[test]
    fn test_annotation_args() {
        let args = |comment: &'static str, open: &str| annotation_args(comment, open).unwrap();
        assert_eq!(args("#[pack()]", "#[pack("), Some(vec![""]));
        assert_eq!(args("#[pack(ignore)]", "#[pack("), Some(vec!["ignore"]));
        assert_eq!(
            args("/// #[unpack(contract = \"Foo\", calldata)]", "#[unpack("),
            Some(vec!["contract = \"Foo\"", "calldata"])
        );
        // parentheses in a constant expression do not end the argument list
        assert_eq!(
            args(
                "#[unpack(length = type(uint16).max, method = \"Lib.f\")]",
                "#[unpack("
            ),
            Some(vec!["length = type(uint16).max", "method = \"Lib.f\""])
        );
        // nor do commas nested inside them
        assert_eq!(
            args("#[pack(length = Lib.f(a, b) * (1 + 2))]", "#[pack("),
            Some(vec!["length = Lib.f(a, b) * (1 + 2)"])
        );
        // delimiters inside a quoted value are literal
        assert_eq!(
            args("#[pack(method = \"f(a,b)\", ignore)]", "#[pack("),
            Some(vec!["method = \"f(a,b)\"", "ignore"])
        );
        // a Solidity string literal inside a closure value must not desync the scanner:
        // the `"` that opens/closes the inner literal must not be treated as the end of
        // the outer quoted annotation value, so a `,` inside the literal must not split.
        assert_eq!(
            args(
                r#"#[pack(assert = "|b| { b == "foo, bar" }", name = "X")]"#,
                "#[pack("
            ),
            Some(vec![
                r#"assert = "|b| { b == "foo, bar" }""#,
                r#"name = "X""#
            ])
        );
        // `{` and `}` in an unquoted expression do not confuse the scanner
        assert_eq!(
            args("#[pack(length = foo({bar}))]", "#[pack("),
            Some(vec!["length = foo({bar})"])
        );
        // an absent annotation is not an error
        assert_eq!(args("#[pack]", "#[pack("), None);
        assert_eq!(args("#[unpack()]", "#[pack("), None);
    }

    #[test]
    fn test_annotation_args_malformed() {
        // a present but malformed annotation is an error rather than an absent annotation, and
        // the message quotes only the offending line of the comment
        let err = annotation_args("/// #[pack(length = (1 + 2]\n/// more docs", "#[pack(")
            .unwrap_err()
            .to_string();
        assert_eq!(
            err,
            "unbalanced parentheses in annotation `#[pack(length = (1 + 2]`"
        );
        let err = annotation_args("#[unpack(method = \"Foo.bar)]", "#[unpack(")
            .unwrap_err()
            .to_string();
        assert_eq!(
            err,
            "unterminated string in annotation `#[unpack(method = \"Foo.bar)]`"
        );
    }
}
