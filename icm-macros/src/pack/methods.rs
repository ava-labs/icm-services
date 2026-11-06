//! This module contains to solidity templates for the packing methods.

use eyre::eyre;
use solar::ast::ElementaryType;
use solar::sema::Gcx;
use solar::sema::hir::{Enum, ItemId, Struct, TypeKind};

use crate::common::{LengthSpec, LengthType};
use crate::pack::hooks::PackArgs;

/// Returns the expression that packs the elementary value `var`. Dynamically sized values are
/// prefixed with their length unless `length` is a constant, in which case only the payload is
/// written: the caller has already asserted the length with [`length_guard`].
pub fn pack_elementary(var: &str, ty: ElementaryType, length: &LengthSpec) -> String {
    let Some(len_expr) = dynamic_length(var, ty) else {
        return format!("abi.encodePacked({var})");
    };
    match length {
        LengthSpec::Constant(_) => format!("abi.encodePacked({var})"),
        LengthSpec::Type(length_type) => {
            format!(
                "abi.encodePacked({}, {var})",
                length_cast(&len_expr, *length_type)
            )
        }
    }
}

/// Returns the expression giving the byte length of `var`, or `None` if `ty` has a fixed size.
fn dynamic_length(var: &str, ty: ElementaryType) -> Option<String> {
    match ty {
        ElementaryType::Bytes => Some(format!("{var}.length")),
        ElementaryType::String => Some(format!("bytes({var}).length")),
        _ => None,
    }
}

/// Wraps a length expression in a cast to the specified type.
/// The default `uint256` type is a no-op cast so it is omitted.
fn length_cast(len_expr: &str, length_type: LengthType) -> String {
    match length_type {
        LengthType::Uint(32) => len_expr.to_string(),
        lt => format!("{}({len_expr})", lt.name()),
    }
}

/// Emits a `require` asserting that `len_expr` equals the constant `expr`. A value that does not
/// have the agreed fixed length must fail here: without a prefix on the wire, the receiver would
/// otherwise decode every later field out of alignment.
fn length_guard(fn_name: &str, field_name: &str, len_expr: &str, expr: &str) -> String {
    format!("\nrequire({len_expr} == {expr}, \"{fn_name}: invalid {field_name} length\");")
}

/// Generates a Solidity `pack` function for an enum type.
/// Enums are value types that are encoded directly via abi.encodePacked.
pub fn pack_enum(enum_def: &Enum, args: PackArgs, type_name: &str) -> String {
    let fn_name = args
        .name
        .unwrap_or_else(|| format!("pack{}", enum_def.name));
    let vis = vis_prefix(&args.visibility);
    format!(
        "function {fn_name}({type_name} obj) {vis}pure returns (bytes memory) {{\n    return abi.encodePacked(obj);\n}}"
    )
}

pub fn pack_struct(
    ctx: &Gcx,
    struct_def: &Struct,
    args: PackArgs,
    type_name: &str,
) -> eyre::Result<String> {
    let fn_name = args
        .name
        .unwrap_or_else(|| format!("pack{}", struct_def.name));
    // the arguments that will be passed to abi.encodePacked
    let mut encode_packed_args = Vec::new();
    // packing functions that need to be expanded recursively
    let mut recursive_packings = Vec::new();
    // the body of the packing function
    let mut body = String::new();

    for (&field_id, field_args) in struct_def.fields.iter().zip(args.fields.iter()) {
        if field_args.ignore {
            continue;
        }
        let field = ctx.hir.variable(field_id);
        let name = field.name.unwrap().to_string();
        let accessor = format!("obj.{name}");
        let length = field_args.length.clone().unwrap_or_default();
        let arg = if let Some(method) = &field_args.method {
            let packed = format!("{method}({accessor})");
            match &length {
                // The method's output goes on the wire verbatim, so a constant length
                // constrains that output rather than the field itself.
                LengthSpec::Constant(expr) => {
                    let output = format!("{name}_bytes");
                    body.push_str(&format!("\nbytes memory {output} = {packed};"));
                    body.push_str(&length_guard(
                        &fn_name,
                        &name,
                        &format!("{output}.length"),
                        expr,
                    ));
                    output
                }
                LengthSpec::Type(_) => packed,
            }
        } else {
            match field.ty.kind {
                TypeKind::Elementary(ty) => {
                    if let (LengthSpec::Constant(expr), Some(len_expr)) =
                        (&length, dynamic_length(&accessor, ty))
                    {
                        body.push_str(&length_guard(&fn_name, &name, &len_expr, expr));
                    }
                    pack_elementary(&accessor, ty, &length)
                }
                TypeKind::Array(_) => {
                    if let LengthSpec::Constant(expr) = &length {
                        body.push_str(&length_guard(
                            &fn_name,
                            &name,
                            &format!("{accessor}.length"),
                            expr,
                        ));
                    }
                    let output = format!("{name}_bytes");
                    // this will need to be handled recursively
                    recursive_packings.push(PackingTreeNode {
                        root: name,
                        depth: 0,
                        accessor,
                        output: output.clone(),
                        ty: &field.ty.kind,
                        length,
                    });
                    // instantiate the variable that will hold the bytes
                    body.push_str(&format!("\nbytes memory {output};"));
                    // add the bytes to the packing
                    output
                }
                TypeKind::Custom(item_id) => {
                    let type_name = custom_type_name(ctx, item_id);
                    format!("pack{type_name}({accessor})")
                }
                TypeKind::Mapping(_) => {
                    return Err(eyre!(
                        "field `{name}` of `{}` is a mapping and requires a user-provided packing method",
                        struct_def.name
                    ));
                }
                TypeKind::Function(_) => {
                    return Err(eyre!(
                        "field `{name}` of `{}` is a function type and requires a user-provided packing method",
                        struct_def.name
                    ));
                }
                TypeKind::Err(_) => unreachable!(),
            }
        };
        encode_packed_args.push(arg);
    }
    // collect the arguments for abi.encodePacked into a comma-separated list
    let encode_packed_args = encode_packed_args.join(", ");
    // recursively expand the packing methods for the composite types and add them to the body
    for node in recursive_packings {
        body.push_str(&expand_packings_recursively(ctx, node)?);
    }
    let vis = vis_prefix(&args.visibility);
    Ok(format!(
        "\nfunction {fn_name}({type_name} memory obj) {vis}pure returns (bytes memory)\
        {{\n    {body}\n    return abi.encodePacked({encode_packed_args});\n}}"
    ))
}

/// A struct that tracks a recursive enumeration through a type to fully
/// expand its packing algorithm. Modelled as a node in a tree
struct PackingTreeNode<'a> {
    /// The name of the variable in the prelude that will
    /// ultimately contain the entire serialized type
    root: String,
    /// The level of recursion in the type
    depth: u8,
    /// The method to access the child of this node
    accessor: String,
    /// The variable which will store the result at this
    /// level of the tree
    output: String,
    /// The type of this node
    ty: &'a TypeKind<'a>,
    /// How this node's length/count prefix is encoded. Defaults to a `uint256` prefix.
    length: LengthSpec,
}

fn custom_type_name(ctx: &Gcx, item_id: ItemId) -> String {
    match item_id {
        ItemId::Struct(id) => ctx.hir.strukt(id).name.to_string(),
        ItemId::Enum(id) => ctx.hir.enumm(id).name.to_string(),
        ItemId::Udvt(id) => ctx.hir.udvt(id).name.to_string(),
        _ => unreachable!("unexpected ItemId in custom type position"),
    }
}

fn expand_packings_recursively(ctx: &Gcx, node: PackingTreeNode) -> eyre::Result<String> {
    Ok(match node.ty {
        TypeKind::Elementary(elem) => {
            format!(
                "{} = {};",
                node.output,
                pack_elementary(&node.accessor, *elem, &node.length)
            )
        }
        TypeKind::Array(arr) => {
            let loop_var = format!("i_{}", node.depth);
            // The element count is read into a local once rather than on every iteration.
            let count_var = format!("{}_{}_count", node.root, node.depth);
            let next_output = format!("{}_{}_bytes", node.root, node.depth + 1);
            let next_node = PackingTreeNode {
                root: node.root.clone(),
                depth: node.depth + 1,
                accessor: format!("{}[{loop_var}]", node.accessor),
                output: next_output.clone(),
                ty: &arr.element.kind,
                length: LengthSpec::default(),
            };
            let inner = expand_packings_recursively(ctx, next_node)?;
            let init = match &node.length {
                // No count prefix: the caller asserted the count against the constant, and the
                // output variable starts out empty.
                LengthSpec::Constant(_) => String::new(),
                LengthSpec::Type(length_type) => {
                    let count = length_cast(&count_var, *length_type);
                    format!("\n{} = abi.encodePacked({count});", node.output)
                }
            };

            format!(
                "\nuint256 {count_var} = {}.length;\
                {init}\
                \nfor (uint256 {loop_var} = 0; {loop_var} < {count_var};){{\
                \n    bytes memory {next_output};\
                \n    {inner}\
                \n    {} = abi.encodePacked({}, {next_output});\
                \n    unchecked {{ ++{loop_var};}}\
                \n}}",
                node.accessor, node.output, node.output
            )
        }
        TypeKind::Custom(item_id) => {
            let type_name = custom_type_name(ctx, *item_id);
            format!("{} = pack{type_name}({});", node.output, node.accessor)
        }
        TypeKind::Mapping(_) => {
            return Err(eyre!(
                "could not derive pack for `{}` as it contains a mapping",
                node.root
            ));
        }
        TypeKind::Function(_) => {
            return Err(eyre!(
                "could not derive pack for `{}` as it contains a function type",
                node.root
            ));
        }
        TypeKind::Err(_) => unreachable!(),
    })
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
