/// Marker trait for the validation state of field args.
pub trait ValidationState {
    /// The type used to represent a length annotation in this state.
    type Length: Clone + std::fmt::Debug + PartialEq;
}

/// State marker: field args have not yet been validated against the HIR.
#[derive(Debug, Clone, PartialEq, Default)]
pub struct Unvalidated;

/// State marker: field args have been validated; `length` holds a `LengthType`.
#[derive(Debug, Clone, PartialEq)]
pub struct Validated;

impl ValidationState for Unvalidated {
    type Length = String;
}

impl ValidationState for Validated {
    type Length = LengthType;
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
                rest.parse()
                    .ok()
                    .filter(|&n: &u16| (8..=256).contains(&n) && n % 8 == 0 && n.to_string() == rest)?
            };
            Some(Ok(ctor(bits as usize / 8)))
        };

        parse_bits("uint", Self::Uint)
            .or_else(|| parse_bits("int", Self::Int))
            .unwrap_or_else(|| Err(eyre::eyre!("`{ty}` is not a valid Solidity integer type")))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_validate_canonical() {
        assert_eq!(LengthType::validate("uint32").unwrap(), LengthType::Uint(4));
        assert_eq!(LengthType::validate("uint256").unwrap(), LengthType::Uint(32));
        assert_eq!(LengthType::validate("uint").unwrap(), LengthType::Uint(32));
        assert_eq!(LengthType::validate("int64").unwrap(), LengthType::Int(8));
        assert_eq!(LengthType::validate("int").unwrap(), LengthType::Int(32));
    }

    #[test]
    fn test_validate_rejects_non_canonical() {
        assert!(LengthType::validate("uint032").is_err(), "leading zero");
        assert!(LengthType::validate("uint+32").is_err(), "leading plus");
        assert!(LengthType::validate("int032").is_err(), "leading zero (int)");
        assert!(LengthType::validate("int+32").is_err(), "leading plus (int)");
    }

    #[test]
    fn test_validate_rejects_invalid() {
        assert!(LengthType::validate("bool").is_err());
        assert!(LengthType::validate("uint0").is_err(), "zero bits");
        assert!(LengthType::validate("uint7").is_err(), "not multiple of 8");
        assert!(LengthType::validate("uint264").is_err(), "too large");
        assert!(LengthType::validate("uint257").is_err(), "not multiple of 8 and too large");
    }
}
