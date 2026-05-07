//! This contains the macro definitions for deriving a pack method on Solidity
//! structs.
//!
//! The macro is triggered by decorating the type (with name DEF) with #[pack(...)] in its doc
//! comment. This will produce a function with the signature
//! `function pack{DEF}(DEF memory def) public pure returns (bytes memory)`
//!
//! For each non-primitive type appearing as a field of DEF, it is assumed to possess a pack method
//! with the above-specified signature. Furthermore, there is not a canonical way to pack mappings as
//! they cannot be iterated over. A user-defined method must always be provided for mappings and
//! functions, or it must be skipped (more on how to do this below).
//!
//! _Optional arguments_
//! This macro takes several optional arguments to modify the above behavior. This should be given
//! as a comma-separated list of key-value pairs, e.g.
//! `#[pack(contract = "MyContract", name = "serialize")]`.
//!
//! `contract`: If provided, the derived code will be placed inside the contract/library with
//! this name. This is not necessary if the type is defined in this contract already. Useful for
//! free-standing types, in which case the default behavior of the macro is to create a
//! free-standing function.
//!
//! `name`: If provided, the name of the function will be changed to this value. Useful for cases
//! when name collisions might otherwise occur.
//!
//! `visibility`: If provided, the visibility of the function will be set to this value. The default
//! is `public`.
//!
//! _Fields_
//! As mentioned above, each non-primitive field is assumed to also possess a pack method. This is
//! assumed to be in the scope in which the macro-expanded code will be placed. If not, or it is
//! desired to use a different serialization method, the field can be annotated with a comment of
//! the form `#[pack(method = "...")]`. Note that this is only relevant for structs, not enums.
//!
//! It is also possible to annotate a field with `#[pack(ignore)]` and its serialization will be
//! skipped by the macro-generated code.
//!
//! _Algorithm_
//! For structs, the macro will walk the fields and call the pack method for each field. There is a
//! default method for primitive types: For fixed size types, encodedPacked is used. For dynamically
//! sized types, the length (as uint256) and followed by the data is passed to encodePacked. The
//! exception, as mentioned above, is mappings and functions.
//!
//! The result of all these calls is then passed to encodePacked, concatenating their bytes in order
//! of their definition.
mod hooks;
mod methods;

use reforge::PreprocessingData;
use solar::ast::Span;
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, SourceId};

use crate::pack::methods::{pack_enum, pack_struct};

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

pub fn derive_pack(
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
        let code = pack_struct(ctx, struct_def, arg, &type_name)
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
        let code = pack_enum(enum_def, arg, &type_name);

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
        reforge::testing::test_macros(
            "testing/pack/source",
            "testing/pack/expected",
            "testing/pack/mismatched",
            &[derive_pack],
        )
        .expect("Test suite failed");
    }

    #[test]
    fn test_errors() {
        let cases = [
            (
                "testing/pack/errors/MappingField.sol",
                "field `data` of `HasMapping` is a mapping and requires a user-provided packing method",
            ),
            (
                "testing/pack/errors/FunctionField.sol",
                "field `handler` of `HasFunction` is a function type and requires a user-provided packing method",
            ),
            (
                "testing/pack/errors/MappingInArray.sol",
                "could not derive pack for `data` as it contains a mapping",
            ),
            (
                "testing/pack/errors/FunctionInArray.sol",
                "could not derive pack for `handlers` as it contains a function type",
            ),
            (
                "testing/pack/errors/BadContract.sol",
                "contract `NonExistent` specified in #[pack(contract=...)] was not found",
            ),
        ];
        for (path, expected) in cases {
            let err = reforge::testing::test_macro_err(path, derive_pack)
                .unwrap_or_else(|e| panic!("test_macro_err failed for {path}: {e}"));
            assert!(
                err.to_string().contains(expected),
                "wrong error for {path}:\n  expected to contain: {expected:?}\n  got: {:?}",
                err.to_string()
            );
        }
    }
}
