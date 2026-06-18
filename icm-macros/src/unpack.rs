//! This contains the macro definitions for deriving an unpack method on Solidity
//! structs.
//!
//! This macro is triggered by decorating the type (with name DEF) with #[unpack(...)] in its doc
//! comment. This will produce a function with the signature
//!  `function unpack{DEF}(bytes MEM_TYPE data) public pure returns (uint256, DEF memory)`
//! where MEM_TYPE is `memory` by default. The uint256 returned is the number of bytes consumed.
//!
//! **Convention**: the input buffer parameter is always named `data`. All generated unpacking
//! code assumes this name unconditionally.
//!
//! For each non-primitive type appearing as a field of DEF, it is assumed to possess an unpack
//! method with the above-specified signature. There is no canonical way to unpack mappings or
//! function types; a user-defined method must always be provided for these, or the field must be
//! skipped (see field annotations below).
//!
//! _Optional arguments_
//! This macro takes several optional arguments. Key-value arguments use `key = "value"` syntax;
//! flags are bare identifiers. All are comma-separated, e.g.
//! `#[unpack(contract = "MyContract", calldata)]`.
//!
//! `contract`: Place the generated function inside the named contract or library. Required for
//! free-standing types; otherwise the function is emitted adjacent to the type definition.
//!
//! `name`: Override the generated function name. Useful to avoid name collisions.
//!
//! `visibility`: Set the visibility of the generated function. Defaults to `public`.
//!
//! `assert = "|var| { expr }"`: Assert a post-condition on the fully-deserialized struct. The
//! closure value must be a quoted string. `var` is replaced with `result` (the decoded struct) in
//! the emitted `require` expression. Multiple `assert` keys may appear; each generates a separate
//! `require` call emitted just before the function's `return` statement. Not valid on enums.
//!
//! `calldata` (flag): Accept `bytes calldata data` instead of `bytes memory data`. The generated
//! code uses calldata array slices throughout, avoiding any memory allocation for the buffer itself.
//!
//! Generated function bodies are wrapped with specific `/* solhint-disable <rule> */` /
//! `/* solhint-enable <rule> */` guards for only the lint rules that actually fire.
//! Functions with no suppressible violations are emitted with no guards at all.
//!
//! _Fields_
//! Each non-primitive field is decoded by calling its own unpack function, which must be in scope
//! at the insertion point. To override this or handle special cases, annotate the field:
//!
//! `#[unpack(method = "expr")]`: Use `(uint256 read, field) = expr(data)` instead of the default
//! inline decoder. The method must accept `data` and return `(uint256, FieldType)`.
//!
//! `#[unpack(default)]`: Skip this field; the struct is returned with a zero value for it.
//! Intended for use alongside `#[pack(ignore)]`.
//!
//! `#[unpack(length=type|constant)]`: Specify the solidity type used to serialize the length of a dynamically
//! sized type. This must be an unsigned integer solidity type or a constant value. If unspecified,
//! the default is `uint256`.
//!
//! `#[unpack(method = "expr", length = constant)]`: Pass a pre-sliced buffer of exactly `constant`
//! bytes to `expr`. The method returns just the field value (no bytes-consumed count); the macro
//! advances `data` by `constant` bytes. Use this for fixed-size fields decoded by a helper that
//! does not implement the `(uint256, T)` unpack convention.
//!
//! `assert = "|var| { expr }"`: Assert a post-condition on the decoded field value. `var` is
//! replaced with the field's local variable name. Emitted as `require(expr)` immediately after
//! `result.field = local`. Multiple `assert` keys may appear on a single field.
//!
//! `assert = "|each var| { expr }"`: Assert a post-condition on each element of a container field
//! (array, `bytes`, `string`, or `bytesN`). For arrays, `var` is replaced with the per-element
//! local inside the decode loop. For `bytes`/`string`/`bytesN`, a post-decode loop iterates the
//! bytes and binds `bytes1 _elem`; `var` is replaced with `_elem`. Not valid on non-container
//! types or at the type level.
//!
//! _Arrays_
//! Arrays are supported by the macro. The macro will walk the array's elements and call the unpack
//! method for each element. The result of all these calls is then passed to `mcopy`, concatenating
//! their bytes in order of their definition.
//!
//! _Mappings_
//!
//! _Algorithm_
//!
//! _Algorithm and buffer handling_
//! The generated code is **zero-copy on the input buffer**: rather than maintaining a separate
//! read cursor, `data`'s memory pointer and length word are updated in place via assembly after
//! each field is consumed. This means `data` must not be read after the unpack call — its pointer
//! and length will have been mutated to reflect whatever bytes were not consumed.
//!
//! Field decoding per type:
//!   - Fixed-size elementary types: read their exact packed byte width via assembly shift/mask,
//!     then advance `data` by that width.
//!   - `bytes` / `string`: the first 32 bytes of the stream encode the payload length; the
//!     payload is copied into a fresh allocation via `mcopy`. This is the only field type that
//!     allocates; the resulting value is independent of `data`.
//!   - Arrays: the first 32 bytes of the stream encode the element count; each element is then
//!     decoded inline in a loop.
//!   - Custom struct/enum/UDVT: delegates to `unpack{TypeName}(data)`.
//!   - Mappings / function types: error — must use `#[unpack(method = "...")]` or
//!     `#[unpack(default)]`.

mod hooks;
mod methods;

use reforge::PreprocessingData;
use solar::ast::Span;
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, SourceId};

/// Returns the `(SourceId, byte_offset)` at which derived code should be inserted.
/// If a target contract is specified (either via `arg_contract` or the item's own `item_contract`),
/// the offset is just before the contract's closing `}`. Otherwise it falls back to just after
/// the item's own span.
fn insertion_point(
    ctx: &Gcx,
    arg_contract: Option<ContractId>,
    item_contract: Option<ContractId>,
    item_source: SourceId,
    item_span: Span,
) -> (SourceId, usize) {
    if let Some(contract_id) = arg_contract.or(item_contract) {
        let contract = ctx.hir.contract(contract_id);
        let source = ctx.sources.get(contract.source).unwrap();
        let offset = (contract.span.hi().0 - source.file.start_pos.0) as usize - 1;
        (contract.source, offset)
    } else {
        let source = ctx.sources.get(item_source).unwrap();
        let offset = (item_span.hi().0 - source.file.start_pos.0) as usize;
        (item_source, offset)
    }
}

fn qualified_type_name(
    ctx: &Gcx,
    type_name: &str,
    item_contract: Option<ContractId>,
    target_contract: Option<ContractId>,
) -> String {
    match (item_contract, target_contract) {
        (Some(item), Some(target)) if item != target => {
            format!("{}.{}", ctx.hir.contract(item).name, type_name)
        }
        _ => type_name.to_string(),
    }
}

pub fn derive_unpack(
    ctx: &Gcx,
    data: &mut PreprocessingData<'_>,
) -> foundry_compilers::error::Result<()> {
    for struct_def in ctx.hir.structs() {
        let Some(arg) = hooks::parse_args(ctx, struct_def.source, struct_def.span, data) else {
            continue;
        };
        let mut arg = arg.map_err(foundry_compilers::error::SolcError::msg)?;
        if arg.contract.is_none() && struct_def.contract.is_none() {
            arg.visibility = String::new();
        }

        let (target_source_id, insertion_offset) = insertion_point(
            ctx,
            arg.contract,
            struct_def.contract,
            struct_def.source,
            struct_def.span,
        );

        let type_name = qualified_type_name(
            ctx,
            struct_def.name.as_str(),
            struct_def.contract,
            arg.contract,
        );
        let code = methods::unpack_struct(ctx, struct_def, arg, &type_name)
            .map_err(foundry_compilers::error::SolcError::msg)?;

        let path = ctx
            .sources
            .get(target_source_id)
            .unwrap()
            .file
            .name
            .as_real()
            .unwrap();
        data.insert(path, insertion_offset, format!("\n\n{code}\n"));
    }

    for enum_def in ctx.hir.enums() {
        let Some(arg) = hooks::parse_args(ctx, enum_def.source, enum_def.span, data) else {
            continue;
        };
        let mut arg = arg.map_err(foundry_compilers::error::SolcError::msg)?;
        if arg.contract.is_none() && enum_def.contract.is_none() {
            arg.visibility = String::new();
        }

        let (target_source_id, insertion_offset) = insertion_point(
            ctx,
            arg.contract,
            enum_def.contract,
            enum_def.source,
            enum_def.span,
        );

        let type_name =
            qualified_type_name(ctx, enum_def.name.as_str(), enum_def.contract, arg.contract);
        let code = methods::unpack_enum(enum_def, arg, &type_name)
            .map_err(foundry_compilers::error::SolcError::msg)?;

        let path = ctx
            .sources
            .get(target_source_id)
            .unwrap()
            .file
            .name
            .as_real()
            .unwrap();
        data.insert(path, insertion_offset, format!("\n\n{code}\n"));
    }
    Ok(())
}

#[cfg(test)]
mod tests {

    use super::*;
    #[test]
    fn test_suite() {
        let mismatched = "testing/unpack/mismatched";
        let _ = std::fs::remove_dir_all(mismatched);
        std::fs::create_dir_all(mismatched).expect("Failed to create mismatched dir");
        reforge::testing::test_macros(
            "testing/unpack/source",
            "testing/unpack/expected",
            mismatched,
            &[derive_unpack],
        )
        .expect("Test suite failed");
    }

    #[test]
    fn test_errors() {
        let cases = [
            (
                "testing/unpack/errors/MappingField.sol",
                "Cannot unpack mapping types without custom methods",
            ),
            (
                "testing/unpack/errors/FunctionField.sol",
                "Cannot unpack function types without custom methods",
            ),
            (
                "testing/unpack/errors/MappingInArray.sol",
                "Cannot unpack mapping types without custom methods",
            ),
            (
                "testing/unpack/errors/FunctionInArray.sol",
                "Cannot unpack function types without custom methods",
            ),
            (
                "testing/unpack/errors/BadContract.sol",
                "contract `NonExistent` specified in #[unpack(contract=...)] was not found",
            ),
            (
                "testing/unpack/errors/EachOnStruct.sol",
                "`each` assert on `HasEachAssert` is not valid at the type level: structs are not container types",
            ),
            (
                "testing/unpack/errors/EachOnEnum.sol",
                "`each` assert on `HasEachAssert` is not valid: enums are not container types",
            ),
            (
                "testing/unpack/errors/EachOnNonContainer.sol",
                "`each` assert on `value` requires a container type (array, bytes, string, or bytesN), not `uint256`",
            ),
            (
                "testing/unpack/errors/EachOnCustom.sol",
                "`each` assert on `inner` requires a container type (array, bytes, string, or bytesN), not `Inner`",
            ),
        ];
        for (path, expected) in cases {
            let err = reforge::testing::test_macro_err(path, derive_unpack)
                .unwrap_or_else(|e| panic!("test_macro_err failed for {path}: {e}"));
            assert!(
                err.to_string().contains(expected),
                "wrong error for {path}:\n  expected to contain: {expected:?}\n  got: {:?}",
                err.to_string()
            );
        }
    }
}
