//! This module contains the Solidity templates for the unpacking methods.
//!
//! **Convention**: all generated unpacking functions receive their input buffer as a parameter
//! named `data` (`bytes calldata data` or `bytes memory data`). Generated code assumes this
//! name unconditionally — callers must not rename it.

use crate::unpack::hooks::{AssertDef, CaptureType, LengthSpec, UnpackArgs};
use solar::ast::{ElementaryType, StateMutability};
use solar::sema::Gcx;
use solar::sema::hir::{ContractId, Enum, ExprKind, ItemId, Struct, TypeKind};

/// Generates the code to unpack an elementary Solidity type from the front of `data`.
/// The result is stored in a freshly declared local named `output`.
///
/// The `calldata` flag selects whether `data` is `bytes calldata` or `bytes memory`.
///
/// Only the bytes consumed by this type are removed from `data`; the rest remains available
/// for subsequent unpacking calls.
pub fn unpack_elementary(
    output: &str,
    ty: ElementaryType,
    calldata: bool,
    length: Option<&LengthSpec>,
) -> String {
    let output = sanitize_local(output);
    match ty {
        ElementaryType::Bytes | ElementaryType::String => {
            let type_name = elementary_type_name(ty);
            if let Some(LengthSpec::Constant(expr)) = length {
                // No length prefix in the stream; the length is a known constant expression.
                if calldata {
                    format!(
                        "\n{type_name} memory {output};\
                    \n{{\
                    \n    {output} = {type_name}(data[0:{expr}]);\
                    \n    data = data[{expr}:];\
                    \n}}"
                    )
                } else {
                    let alloc = match ty {
                        ElementaryType::Bytes => format!("new bytes({expr})"),
                        _ => format!("string(new bytes({expr}))"),
                    };
                    format!(
                        "\n{type_name} memory {output};\
                    \n{{\
                    \n    {output} = {alloc};\
                    \n    assembly {{\
                    \n        mcopy(add({output}, 32), add(data, 32), {expr})\
                    \n        let _data_orig_len := mload(data)\
                    \n        data := add(data, {expr})\
                    \n        mstore(data, sub(_data_orig_len, {expr}))\
                    \n    }}\
                    \n}}"
                    )
                }
            } else {
                let prefix_size = length_prefix_size(length);
                if calldata {
                    let read_len = read_length_calldata(prefix_size);
                    format!(
                        "\n{type_name} memory {output};\
                    \n{{\
                    \n    uint256 length = {read_len};\
                    \n    {output} = {type_name}(data[{prefix_size}:length + {prefix_size}]);\
                    \n    data = data[{prefix_size} + length:];\
                    \n}}"
                    )
                } else {
                    let alloc = match ty {
                        ElementaryType::Bytes => "new bytes(data_length)".to_string(),
                        _ => "string(new bytes(data_length))".to_string(),
                    };
                    let read_len = read_length_memory(prefix_size);
                    let payload_offset = 32 + prefix_size;
                    format!(
                        "\n{type_name} memory {output};\
                    \n{{\
                    \n    uint256 data_length;\
                    \n    assembly {{ data_length := {read_len} }}\
                    \n    {output} = {alloc};\
                    \n    assembly {{\
                    \n        mcopy(add({output}, 32), add(data, {payload_offset}), data_length)\
                    \n        let _data_orig_len := mload(data)\
                    \n        data := add(data, add({prefix_size}, data_length))\
                    \n        mstore(data, sub(_data_orig_len, add({prefix_size}, data_length)))\
                    \n    }}\
                    \n}}"
                    )
                }
            }
        }
        _ => {
            let size = packed_size(ty);
            let shift = (32 - size) * 8;
            let type_name = elementary_type_name(ty);
            if calldata {
                let decode = calldata_decode("data", ty, size);
                format!(
                    "\n{type_name} {output} = {decode};\
                     \ndata = data[{size}:];"
                )
            } else {
                let read = match ty {
                    ElementaryType::Int(_) => {
                        format!("sar({shift}, mload(add(data, 32)))")
                    }
                    ElementaryType::FixedBytes(_) => {
                        format!("and(mload(add(data, 32)), shl({shift}, not(0)))")
                    }
                    _ => format!("shr({shift}, mload(add(data, 32)))"),
                };
                let slice = slice_memory(
                    "data",
                    &size.to_string(),
                    &format!("sub(mload(data), {size})"),
                );
                format!(
                    "\n{type_name} {output};\
                    \nassembly {{\
                    \n    {output} := {read}\
                    \n}}\
                    \n{slice}"
                )
            }
        }
    }
}

pub fn unpack_enum(enum_def: &Enum, args: UnpackArgs, type_name: &str) -> eyre::Result<String> {
    if args
        .assert
        .iter()
        .any(|a| a.captured == CaptureType::Element)
    {
        eyre::bail!(
            "`each` assert on `{}` is not valid: enums are not container types",
            enum_def.name,
        );
    }
    let fn_name = args
        .name
        .unwrap_or_else(|| format!("unpack{}", enum_def.name));
    let input = if args.calldata {
        "bytes calldata data"
    } else {
        "bytes memory data"
    };
    let vis = vis_prefix(&args.visibility);
    let field_asserts: Vec<_> = args
        .assert
        .iter()
        .filter(|a| a.captured == CaptureType::Field)
        .collect();
    let body = if field_asserts.is_empty() {
        format!(
            "function {fn_name}({input}) {vis}pure returns (uint256, {type_name}) {{\
            \n    return (1, {type_name}(uint8(data[0])));\
            \n}}"
        )
    } else {
        let assert_lines: String = field_asserts
            .iter()
            .map(|a| {
                let expr = substitute_var(&a.expr, &a.var, "result");
                format!("\n    require({expr});")
            })
            .collect();
        {
            let (open, close) = crate::solhint::solhint_guards(&["reason-string"]);
            format!(
                "function {fn_name}({input}) {vis}pure returns (uint256, {type_name}) {{\
                \n    {open}{type_name} result = {type_name}(uint8(data[0]));{assert_lines}\
                \n    return (1, result);{close}\
                \n}}"
            )
        }
    };
    Ok(body)
}

pub fn unpack_struct(
    ctx: &Gcx,
    struct_def: &Struct,
    args: UnpackArgs,
    type_name: &str,
) -> eyre::Result<String> {
    let fn_name = args
        .name
        .unwrap_or_else(|| format!("unpack{}", struct_def.name));
    let input = if args.calldata {
        "bytes calldata data"
    } else {
        "bytes memory data"
    };
    let (length_tracking, pre_return, bytes_read) = if args.calldata {
        (
            "uint256 _initial_length = data.length;".to_string(),
            String::new(),
            "_initial_length - data.length".to_string(),
        )
    } else {
        (
            "uint256 _initial_length;\nassembly { _initial_length := mload(data) }".to_string(),
            "uint256 _final_length;\nassembly { _final_length := mload(data) }".to_string(),
            "_initial_length - _final_length".to_string(),
        )
    };
    let target_contract = args.contract.or(struct_def.contract);
    let mut body = String::new();
    for (&field_id, field_args) in struct_def.fields.iter().zip(args.fields.iter()) {
        if field_args.default {
            continue;
        }
        let field = ctx.hir.variable(field_id);
        let field_type = get_type_name(ctx, &field.ty.kind, target_contract)?;
        let field_name = field.name.unwrap();
        let local_name = sanitize_local(field_name.as_str());
        let declaration = if field_args.memory {
            format!("{field_type} memory {local_name};")
        } else {
            format!("{field_type} {local_name};")
        };
        let deserialize_field_code = if let Some(method) = &field_args.method {
            if let Some(LengthSpec::Constant(expr)) = &field_args.length {
                // method + constant: pass a pre-sliced buffer of exactly `expr` bytes.
                // The method returns just the value — no bytes-consumed count needed.
                if args.calldata {
                    format!(
                        "\n{declaration}\
                        \n{{\
                        \n    {local_name} = {method}(data[0:{expr}]);\
                        \n    data = data[{expr}:];\
                        \n}}"
                    )
                } else {
                    format!(
                        "\n{declaration}\
                        \n{{\
                        \n    bytes memory _slice_{local_name} = new bytes({expr});\
                        \n    assembly {{ mcopy(add(_slice_{local_name}, 32), add(data, 32), {expr}) }}\
                        \n    {local_name} = {method}(_slice_{local_name});\
                        \n    assembly {{\
                        \n        let _old_len := mload(data)\
                        \n        data := add(data, {expr})\
                        \n        mstore(data, sub(_old_len, {expr}))\
                        \n    }}\
                        \n}}"
                    )
                }
            } else if args.calldata {
                format!(
                    "\n{declaration}\
                    \n{{\
                    \n    uint256 read;\
                    \n    (read, {local_name}) = {method}(data);\
                    \n    data = data[read:];\
                    \n}}"
                )
            } else {
                format!(
                    "\n{declaration}\
                    \n{{\
                    \n    uint256 _len_before;\
                    \n    assembly {{ _len_before := mload(data) }}\
                    \n    uint256 read;\
                    \n    (read, {local_name}) = {method}(data);\
                    \n    assembly {{\
                    \n        data := add(data, read)\
                    \n        mstore(data, sub(_len_before, read))\
                    \n    }}\
                    \n}}"
                )
            }
        } else {
            if let Some(LengthSpec::Type(ty)) = &field_args.length {
                validate_length_type(ty, field_name.as_str(), struct_def.name.as_str())?;
            }
            let node = UnpackingTreeNode {
                depth: 0,
                output: local_name.clone(),
                ty: &field.ty.kind,
                length: field_args.length.clone(),
                element_asserts: field_args
                    .assert
                    .iter()
                    .filter(|a| a.captured == CaptureType::Element)
                    .cloned()
                    .collect(),
            };
            inline_unpack(ctx, args.calldata, target_contract, node)?
        };
        body.push_str(&deserialize_field_code);
        body.push_str(&format!("\nresult.{field_name} = {local_name};"));
        for assert_def in field_args
            .assert
            .iter()
            .filter(|a| a.captured == CaptureType::Field)
        {
            let expr = substitute_var(&assert_def.expr, &assert_def.var, &local_name);
            body.push_str(&format!("\nrequire({expr});"));
        }
    }
    let mut epilogue = pre_return;
    for assert_def in &args.assert {
        match assert_def.captured {
            CaptureType::Field => {
                let expr = substitute_var(&assert_def.expr, &assert_def.var, "result");
                if !epilogue.is_empty() {
                    epilogue.push('\n');
                }
                epilogue.push_str(&format!("require({expr});"));
            }
            CaptureType::Element => {
                eyre::bail!(
                    "`each` assert on `{}` is not valid at the type level: structs are not container types",
                    struct_def.name,
                );
            }
        }
    }
    let vis = vis_prefix(&args.visibility);
    let has_asserts = !args.assert.is_empty()
        || args.fields.iter().any(|fa| !fa.assert.is_empty());
    let mut rules: Vec<&str> = Vec::new();
    if !args.calldata {
        rules.push("no-inline-assembly");
    }
    rules.push("var-name-mixedcase");
    if has_asserts {
        rules.push("reason-string");
    }
    let (open, close) = crate::solhint::solhint_guards(&rules);
    Ok(format!(
        "function {fn_name}({input}) {vis}pure returns (uint256, {type_name} memory) {{\
        \n    {open}{length_tracking}\
        \n    {type_name} memory result;\
        \n    {body}\
        \n    {epilogue}\
        \n    return ({bytes_read}, result);{close}\
        \n}}"
    ))
}

/// A struct that tracks a recursive enumeration through a type to fully
/// expand its packing algorithm. Modelled as a node in a tree
struct UnpackingTreeNode<'a> {
    /// The level of recursion in the type
    depth: u8,
    /// The variable which will store the result at this
    /// level of the tree
    output: String,
    /// The type of this node
    ty: &'a TypeKind<'a>,
    /// Length/count spec for this node. `None` keeps the default `uint256` prefix.
    length: Option<LengthSpec>,
    /// `CaptureType::Element` asserts to inject into this node's array loop body,
    /// substituting the capture variable with the per-element local. Only consulted
    /// when `ty` is an array; never propagated to the recursive element node.
    element_asserts: Vec<AssertDef>,
}

fn inline_unpack(
    ctx: &Gcx,
    calldata: bool,
    target_contract: Option<ContractId>,
    node: UnpackingTreeNode<'_>,
) -> eyre::Result<String> {
    Ok(match node.ty {
        TypeKind::Elementary(ty) => {
            let mut code = unpack_elementary(&node.output, *ty, calldata, node.length.as_ref());
            if !node.element_asserts.is_empty() {
                code.push_str(&element_assert_loop(
                    *ty,
                    &node.output,
                    &node.element_asserts,
                )?);
            }
            code
        }
        TypeKind::Array(arr) => {
            let type_name = get_type_name(ctx, &arr.element.kind, target_contract)?;
            let next_output = format!("{}_{}", node.output, node.depth + 1);
            let next_node = UnpackingTreeNode {
                depth: node.depth + 1,
                output: next_output.clone(),
                ty: &arr.element.kind,
                length: None,
                element_asserts: vec![],
            };
            let output = sanitize_local(&node.output);
            let inner = inline_unpack(ctx, calldata, target_contract, next_node)?;
            let element_assert_code: String = node
                .element_asserts
                .iter()
                .map(|a| {
                    let expr = substitute_var(&a.expr, &a.var, &next_output);
                    format!("\n        require({expr});")
                })
                .collect();
            let (read_length, advance_past_count) = match node.length.as_ref() {
                Some(LengthSpec::Constant(expr)) => {
                    (format!("uint256 length = {expr};"), String::new())
                }
                other => {
                    let prefix_size = length_prefix_size(other);
                    if calldata {
                        (
                            format!("uint256 length = {};", read_length_calldata(prefix_size)),
                            format!("data = data[{prefix_size}:];"),
                        )
                    } else {
                        (
                            format!(
                                "uint256 length; assembly {{ length := {} }}",
                                read_length_memory(prefix_size)
                            ),
                            slice_memory(
                                "data",
                                &prefix_size.to_string(),
                                &format!("sub(mload(data), {prefix_size})"),
                            ),
                        )
                    }
                }
            };
            format!(
                "\n{type_name}[] memory {output};\
                \n{{\
                \n    {read_length}\
                \n    {advance_past_count}\
                \n    {output} = new {type_name}[](length);\
                \n    for (uint256 i = 0; i < length;){{\
                \n        {inner}\
                \n        {output}[i] = {next_output};{element_assert_code}\
                \n        unchecked {{ ++i;}}\
                \n    }}\
                \n}}"
            )
        }
        TypeKind::Custom(item_id) => {
            if !node.element_asserts.is_empty() {
                let type_name = custom_type_name(ctx, *item_id);
                eyre::bail!(
                    "`each` assert on `{}` requires a container type (array, bytes, string, or bytesN), not `{type_name}`",
                    node.output,
                );
            }
            let type_name = custom_type_name(ctx, *item_id);
            let output = sanitize_local(&node.output);
            let is_enum = matches!(item_id, ItemId::Enum(_));
            let decl = if is_enum {
                format!("{type_name} {output};")
            } else {
                format!("{type_name} memory {output};")
            };
            if calldata {
                format!(
                    "\n{decl}\
                    \n{{\
                    \n    uint256 read;\
                    \n    (read, {output}) = unpack{type_name}(data);\
                    \n    data = data[read:];\
                    \n}}"
                )
            } else {
                format!(
                    "\n{decl}\
                    \n{{\
                    \n    uint256 _len_before;\
                    \n    assembly {{ _len_before := mload(data) }}\
                    \n    uint256 read;\
                    \n    (read, {output}) = unpack{type_name}(data);\
                    \n    assembly {{\
                    \n        data := add(data, read)\
                    \n        mstore(data, sub(_len_before, read))\
                    \n    }}\
                    \n}}"
                )
            }
        }
        TypeKind::Function(_) => eyre::bail!("Cannot unpack function types without custom methods"),
        TypeKind::Mapping(_) => eyre::bail!("Cannot unpack mapping types without custom methods"),
        TypeKind::Err(_) => eyre::bail!("Cannot unpack error types without custom methods"),
    })
}

/// Returns the Solidity type name for an elementary type.
fn elementary_type_name(ty: ElementaryType) -> String {
    match ty {
        ElementaryType::Bool => "bool".to_string(),
        ElementaryType::Address(false) => "address".to_string(),
        ElementaryType::Address(true) => "address payable".to_string(),
        ElementaryType::UInt(s) => format!("uint{}", s.bits()),
        ElementaryType::Int(s) => format!("int{}", s.bits()),
        ElementaryType::FixedBytes(s) => format!("bytes{}", s.bytes()),
        ElementaryType::Bytes => "bytes".to_string(),
        ElementaryType::String => "string".to_string(),
        ElementaryType::Fixed(s, f) => format!("fixed{}x{}", s.bytes(), f.get()),
        ElementaryType::UFixed(s, f) => format!("ufixed{}x{}", s.bytes(), f.get()),
    }
}

fn custom_type_name(ctx: &Gcx, item_id: ItemId) -> String {
    match item_id {
        ItemId::Struct(id) => ctx.hir.strukt(id).name.to_string(),
        ItemId::Enum(id) => ctx.hir.enumm(id).name.to_string(),
        ItemId::Udvt(id) => ctx.hir.udvt(id).name.to_string(),
        _ => unreachable!("unexpected ItemId in custom type position"),
    }
}

/// Qualifies `name` with its owning contract's name when that contract differs from
/// `target_contract`. Returns the bare name when no qualification is needed.
fn qualify_type_name(
    ctx: &Gcx,
    name: &str,
    item_contract: Option<ContractId>,
    target_contract: Option<ContractId>,
) -> String {
    match item_contract {
        Some(c) if Some(c) != target_contract => {
            format!("{}.{}", ctx.hir.contract(c).name, name)
        }
        _ => name.to_string(),
    }
}

/// Returns the Solidity type name for a HIR type kind, qualifying cross-contract custom types
/// relative to `target_contract`.
fn get_type_name(
    ctx: &Gcx<'_>,
    kind: &TypeKind<'_>,
    target_contract: Option<ContractId>,
) -> eyre::Result<String> {
    Ok(match kind {
        TypeKind::Elementary(e) => elementary_type_name(*e),
        TypeKind::Array(arr) => {
            let elem = get_type_name(ctx, &arr.element.kind, target_contract)?;
            match arr.size {
                None => format!("{elem}[]"),
                Some(expr) => {
                    let size_str = match &expr.kind {
                        ExprKind::Lit(lit) => lit.symbol.as_str().to_string(),
                        _ => eyre::bail!(
                            "array size is not a literal and cannot be used in a type declaration"
                        ),
                    };
                    format!("{elem}[{size_str}]")
                }
            }
        }
        TypeKind::Custom(ItemId::Struct(id)) => {
            let s = ctx.hir.strukt(*id);
            qualify_type_name(ctx, s.name.as_str(), s.contract, target_contract)
        }
        TypeKind::Custom(ItemId::Enum(id)) => {
            let e = ctx.hir.enumm(*id);
            qualify_type_name(ctx, e.name.as_str(), e.contract, target_contract)
        }
        TypeKind::Custom(ItemId::Udvt(id)) => {
            let u = ctx.hir.udvt(*id);
            qualify_type_name(ctx, u.name.as_str(), u.contract, target_contract)
        }
        TypeKind::Custom(_) => unreachable!("unexpected ItemId in custom type position"),
        TypeKind::Mapping(m) => {
            let key = get_type_name(ctx, &m.key.kind, target_contract)?;
            let value = get_type_name(ctx, &m.value.kind, target_contract)?;
            format!("mapping({key} => {value})")
        }
        TypeKind::Function(f) => {
            let params = f
                .parameters
                .iter()
                .map(|&vid| get_type_name(ctx, &ctx.hir.variable(vid).ty.kind, target_contract))
                .collect::<eyre::Result<Vec<_>>>()?
                .join(", ");
            let returns = f
                .returns
                .iter()
                .map(|&vid| get_type_name(ctx, &ctx.hir.variable(vid).ty.kind, target_contract))
                .collect::<eyre::Result<Vec<_>>>()?;
            let mutability = match f.state_mutability {
                StateMutability::NonPayable => String::new(),
                m => format!(" {m}"),
            };
            if returns.is_empty() {
                format!("function({params}) {}{mutability}", f.visibility)
            } else {
                format!(
                    "function({params}) {}{mutability} returns ({})",
                    f.visibility,
                    returns.join(", ")
                )
            }
        }
        TypeKind::Err(_) => unreachable!("Cannot unpacks errors"),
    })
}

/// Returns the byte size of an elementary type as encoded by `abi.encodePacked`.
fn packed_size(ty: ElementaryType) -> usize {
    match ty {
        ElementaryType::Bool => 1,
        ElementaryType::Address(_) => 20,
        ElementaryType::UInt(s) | ElementaryType::Int(s) => s.bytes() as usize,
        ElementaryType::FixedBytes(s) => s.bytes() as usize,
        ElementaryType::Fixed(s, _) | ElementaryType::UFixed(s, _) => s.bytes() as usize,
        ElementaryType::Bytes | ElementaryType::String => {
            unreachable!("dynamic types have no fixed packed size")
        }
    }
}

/// Returns a Solidity expression that decodes `size` bytes from the front of the calldata
/// slice `var` into the appropriate type.
fn calldata_decode(var: &str, ty: ElementaryType, size: usize) -> String {
    let bytes_type = format!("bytes{size}");
    match ty {
        ElementaryType::Bool => format!("bytes1({var}[0:1]) != 0x00"),
        ElementaryType::FixedBytes(_) => format!("{bytes_type}({var}[0:{size}])"),
        ElementaryType::Address(_) => format!("address(bytes20({var}[0:20]))"),
        ElementaryType::UInt(s) => format!("uint{}({bytes_type}({var}[0:{size}]))", s.bits()),
        ElementaryType::Int(s) => format!(
            "int{}(uint{}({bytes_type}({var}[0:{size}])))",
            s.bits(),
            s.bits()
        ),
        ElementaryType::Fixed(s, f) => {
            format!(
                "fixed{}x{}({bytes_type}({var}[0:{size}]))",
                s.bytes(),
                f.get()
            )
        }
        ElementaryType::UFixed(s, f) => {
            format!(
                "ufixed{}x{}({bytes_type}({var}[0:{size}]))",
                s.bytes(),
                f.get()
            )
        }
        ElementaryType::Bytes | ElementaryType::String => unreachable!(),
    }
}

/// Generates assembly that zero-copy slices a `bytes memory` variable in place: advances its
/// pointer by `bytes_read` bytes and writes `new_length` as the new length word.
/// `new_length` is evaluated to a temporary before the pointer moves, so expressions that
/// read `mload(array)` in `new_length` correctly capture the old length word.
fn slice_memory(array: &str, bytes_read: &str, new_length: &str) -> String {
    format!(
        "assembly {{\
    \n    let _{array}_new_len := {new_length}\
    \n    {array} := add({array}, {bytes_read})\
    \n    mstore({array}, _{array}_new_len)\
    \n}}"
    )
}

/// Returns a local variable name that does not clash with any names introduced
/// by the generated unpacking function body. Appends trailing underscores until the name is free.
///
/// Reserved names: `data`, `result`, `read`, `length`, `i`, `_initial_length`, `_final_length`,
/// `data_length`, `data_orig_length`, `_len_before`.
fn sanitize_local(name: &str) -> String {
    const RESERVED: &[&str] = &[
        "data",
        "result",
        "read",
        "length",
        "i",
        "_initial_length",
        "_final_length",
        "data_length",
        "data_orig_length",
        "_len_before",
    ];
    let mut out = name.to_string();
    while RESERVED.contains(&out.as_str()) {
        out.push('_');
    }
    out
}

/// Returns the byte width of the length/count prefix for the given spec.
/// Defaults to 32 (uint256) when no override is specified.
/// Callers must not pass `Constant` — that case has no prefix and must be handled before calling.
fn length_prefix_size(length: Option<&LengthSpec>) -> usize {
    match length {
        None => 32,
        Some(LengthSpec::Type(ty)) => {
            ty.strip_prefix("uint").unwrap().parse::<usize>().unwrap() / 8
        }
        Some(LengthSpec::Constant(_)) => {
            unreachable!("Constant length has no prefix size")
        }
    }
}

/// Generates a calldata expression that reads a length/count prefix of `size` bytes.
fn read_length_calldata(size: usize) -> String {
    if size == 32 {
        "uint256(bytes32(data[0:32]))".to_string()
    } else {
        let bits = size * 8;
        format!("uint256(uint{bits}(bytes{size}(data[0:{size}])))")
    }
}

/// Generates an assembly expression that reads a length/count prefix of `size` bytes
/// from the front of a memory `bytes` buffer (after the EVM length word).
fn read_length_memory(size: usize) -> String {
    if size == 32 {
        "mload(add(data, 32))".to_string()
    } else {
        let shift = (32 - size) * 8;
        format!("shr({shift}, mload(add(data, 32)))")
    }
}

/// Validates that `ty` is a Solidity unsigned integer type (uint8 through uint256, in steps of 8).
fn validate_length_type(ty: &str, field_name: &str, struct_name: &str) -> eyre::Result<()> {
    let bits: Option<u16> = ty.strip_prefix("uint").and_then(|s| s.parse().ok());
    let valid = bits.is_some_and(|n| (8..=256).contains(&n) && n % 8 == 0);
    if !valid {
        return Err(eyre::eyre!(
            "length type `{ty}` for field `{field_name}` of `{struct_name}` \
             must be an unsigned integer type (uint8 to uint256)"
        ));
    }
    Ok(())
}

/// Returns the visibility keyword with a trailing space, or an empty string when
/// `visibility` is empty (i.e. for file-level functions that take no visibility specifier).
fn vis_prefix(visibility: &str) -> String {
    if visibility.is_empty() {
        String::new()
    } else {
        format!("{visibility} ")
    }
}

/// Generates a post-decode element-assertion loop for container elementary types
/// (`bytes`, `string`, `bytesN`). Each iteration binds `bytes1 _elem` to the current
/// element and evaluates the assert expressions with the capture variable substituted
/// by `_elem`. Returns an empty string for non-container types.
fn element_assert_loop(
    ty: ElementaryType,
    output: &str,
    asserts: &[AssertDef],
) -> eyre::Result<String> {
    let (length_expr, access_expr) = match ty {
        ElementaryType::Bytes => (format!("{output}.length"), format!("{output}[_j]")),
        ElementaryType::String => (
            format!("bytes({output}).length"),
            format!("bytes({output})[_j]"),
        ),
        ElementaryType::FixedBytes(s) => (s.bytes().to_string(), format!("{output}[_j]")),
        _ => eyre::bail!(
            "`each` assert on `{output}` requires a container type (array, bytes, string, or bytesN), \
             not `{}`",
            elementary_type_name(ty),
        ),
    };
    let assert_lines: String = asserts
        .iter()
        .map(|a| {
            let expr = substitute_var(&a.expr, &a.var, "_elem");
            format!("\n    require({expr});")
        })
        .collect();
    Ok(format!(
        "\nfor (uint256 _j = 0; _j < {length_expr}; _j++){{\
        \n    bytes1 _elem = {access_expr};{assert_lines}\
        \n}}"
    ))
}

/// Replaces whole-word occurrences of `var` in `expr` with `replacement`.
/// A word boundary is any position adjacent to a non-identifier character
/// (`[^a-zA-Z0-9_]`) or the start/end of the string.
fn substitute_var(expr: &str, var: &str, replacement: &str) -> String {
    fn is_ident(c: char) -> bool {
        c.is_alphanumeric() || c == '_'
    }
    let mut result = String::new();
    let mut remaining = expr;
    while let Some(pos) = remaining.find(var) {
        let before = &remaining[..pos];
        let after = &remaining[pos + var.len()..];
        let left_ok = before.chars().last().is_none_or(|c| !is_ident(c));
        let right_ok = after.chars().next().is_none_or(|c| !is_ident(c));
        result.push_str(before);
        if left_ok && right_ok {
            result.push_str(replacement);
        } else {
            result.push_str(var);
        }
        remaining = after;
    }
    result.push_str(remaining);
    result
}
