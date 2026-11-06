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
        if let Ok(ty) = LengthType::validate(&value) {
            return Ok(Self::Type(ty));
        }
        if value.is_empty() {
            eyre::bail!(
                "empty `length` for field `{field_name}` of `{struct_name}`: expected a Solidity \
                 integer type or a constant length expression"
            );
        }
        Ok(Self::Constant(value))
    }
}

/// Splits the argument list of the `{open}...)` annotation in `comment` (`open` being e.g.
/// `#[pack(`) on top-level commas, returning each argument trimmed. Parentheses are balanced and
/// quoted strings are skipped, so constant expressions such as `type(uint16).max` or
/// `Lib.f(a, b)` survive intact. Returns `None` if the annotation is absent or its parentheses
/// never close.
pub fn annotation_args<'a>(comment: &'a str, open: &str) -> Option<Vec<&'a str>> {
    let start = comment.find(open)? + open.len();
    let mut depth = 0usize;
    let mut quoted = false;
    let mut args = Vec::new();
    let mut arg_start = start;
    for (i, byte) in comment.bytes().enumerate().skip(start) {
        match byte {
            b'"' => quoted = !quoted,
            _ if quoted => {}
            b'(' => depth += 1,
            b')' if depth > 0 => depth -= 1,
            b')' => {
                args.push(comment[arg_start..i].trim());
                return Some(args);
            }
            b',' if depth == 0 => {
                args.push(comment[arg_start..i].trim());
                arg_start = i + 1;
            }
            _ => {}
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

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
        assert_eq!(annotation_args("#[pack()]", "#[pack("), Some(vec![""]));
        assert_eq!(
            annotation_args("#[pack(ignore)]", "#[pack("),
            Some(vec!["ignore"])
        );
        assert_eq!(
            annotation_args("/// #[unpack(contract = \"Foo\", calldata)]", "#[unpack("),
            Some(vec!["contract = \"Foo\"", "calldata"])
        );
        // parentheses in a constant expression do not end the argument list
        assert_eq!(
            annotation_args(
                "#[unpack(length = type(uint16).max, method = \"Lib.f\")]",
                "#[unpack("
            ),
            Some(vec!["length = type(uint16).max", "method = \"Lib.f\""])
        );
        // nor do commas nested inside them
        assert_eq!(
            annotation_args("#[pack(length = Lib.f(a, b) * (1 + 2))]", "#[pack("),
            Some(vec!["length = Lib.f(a, b) * (1 + 2)"])
        );
        // delimiters inside a quoted value are literal
        assert_eq!(
            annotation_args("#[pack(method = \"f(a,b)\", ignore)]", "#[pack("),
            Some(vec!["method = \"f(a,b)\"", "ignore"])
        );
        assert_eq!(annotation_args("#[pack]", "#[pack("), None);
        assert_eq!(annotation_args("#[unpack()]", "#[pack("), None);
        assert_eq!(
            annotation_args("#[pack(length = (1 + 2]", "#[pack("),
            None,
            "unbalanced parentheses never close the annotation"
        );
    }
}
