//! This module contains to solidity templates for the packing methods.

use eyre::eyre;
use solar::ast::ElementaryType;
use solar::sema::Gcx;
use solar::sema::hir::{Enum, ItemId, Struct, TypeKind};

use crate::pack::hooks::{LengthSpec, PackArgs};

pub fn pack_elementary(var: &str, ty: ElementaryType, length: Option<&LengthSpec>) -> String {
    match ty {
        ElementaryType::Bytes => match length {
            Some(LengthSpec::Drop) => format!("abi.encodePacked({var})"),
            other => {
                let len = length_cast(&format!("{var}.length"), other);
                format!("abi.encodePacked({len}, {var})")
            }
        },
        ElementaryType::String => match length {
            Some(LengthSpec::Drop) => format!("abi.encodePacked({var})"),
            other => {
                let len = length_cast(&format!("bytes({var}).length"), other);
                format!("abi.encodePacked({len}, {var})")
            }
        },
        _ => format!("abi.encodePacked({var})"),
    }
}

/// Wraps a length expression in a cast to the specified type, or returns it unchanged
/// when no override is specified (leaving the implicit uint256 encoding in place).
/// `Drop` is treated as no-op here; callers that care about drop must handle it before calling.
fn length_cast(len_expr: &str, length: Option<&LengthSpec>) -> String {
    match length {
        None | Some(LengthSpec::Drop) => len_expr.to_string(),
        Some(LengthSpec::Type(ty)) => format!("{ty}({len_expr})"),
    }
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
        if let Some(LengthSpec::Type(ty)) = &field_args.length {
            validate_length_type(ty, field.name.unwrap().as_str(), struct_def.name.as_str())?;
        }
        let arg = if let Some(method) = &field_args.method {
            format!("{method}(obj.{})", field.name.unwrap())
        } else {
            match field.ty.kind {
                TypeKind::Elementary(ty) => pack_elementary(
                    &format!("obj.{}", field.name.unwrap()),
                    ty,
                    field_args.length.as_ref(),
                ),
                TypeKind::Array(_) => {
                    let name = field.name.unwrap().to_string();
                    let output = format!("{name}_bytes");
                    // this will need to be handled recursively
                    recursive_packings.push(PackingTreeNode {
                        root: name.clone(),
                        depth: 0,
                        accessor: format!("obj.{name}"),
                        output: output.clone(),
                        ty: &field.ty.kind,
                        length: field_args.length.clone(),
                    });
                    // instantiate the variable that will hold the bytes
                    body.push_str(&format!("bytes memory {output};"));
                    // add the bytes to the packing
                    output
                }
                TypeKind::Custom(item_id) => {
                    let type_name = custom_type_name(ctx, item_id);
                    format!("pack{type_name}(obj.{})", field.name.unwrap())
                }
                TypeKind::Mapping(_) => {
                    return Err(eyre!(
                        "field `{}` of `{}` is a mapping and requires a user-provided packing method",
                        field.name.unwrap(),
                        struct_def.name
                    ));
                }
                TypeKind::Function(_) => {
                    return Err(eyre!(
                        "field `{}` of `{}` is a function type and requires a user-provided packing method",
                        field.name.unwrap(),
                        struct_def.name
                    ));
                }
                TypeKind::Err(_) => unreachable!(),
            }
        };
        encode_packed_args.push(arg);
    }
    // collect the arguments for abi.encodePacked into a comma-separated list
    let has_array_fields = !recursive_packings.is_empty();
    let encode_packed_args = encode_packed_args.join(", ");
    // recursively expand the packing methods for the composite types and add them to the body
    for node in recursive_packings {
        body.push_str(&expand_packings_recursively(ctx, node)?);
    }
    let fn_name = args
        .name
        .unwrap_or_else(|| format!("pack{}", struct_def.name));
    let vis = vis_prefix(&args.visibility);
    let (open, close) = crate::solhint::solhint_guards(
        if has_array_fields { &["var-name-mixedcase"] } else { &[] },
    );
    let body_prefix = if body.is_empty() {
        String::new()
    } else {
        format!("{body}\n    ")
    };
    Ok(format!(
        "\nfunction {fn_name}({type_name} memory obj) {vis}pure returns (bytes memory)\
        {{\n    {open}{body_prefix}return abi.encodePacked({encode_packed_args});{close}\n}}"
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
    /// Length/count prefix spec for this node. `None` keeps the default `uint256` encoding.
    length: Option<LengthSpec>,
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
                pack_elementary(&node.accessor, *elem, node.length.as_ref())
            )
        }
        TypeKind::Array(arr) => {
            let loop_var = format!("i_{}", node.depth);
            let next_output = format!("{}_{}_bytes", node.root, node.depth + 1);
            let next_node = PackingTreeNode {
                root: node.root.clone(),
                depth: node.depth + 1,
                accessor: format!("{}[{loop_var}]", node.accessor),
                output: next_output.clone(),
                ty: &arr.element.kind,
                length: None,
            };
            let inner = expand_packings_recursively(ctx, next_node)?;
            let init = if matches!(node.length, Some(LengthSpec::Drop)) {
                format!("{} = abi.encodePacked();", node.output)
            } else {
                let count = length_cast(&format!("{}.length", node.accessor), node.length.as_ref());
                format!("{} = abi.encodePacked({count});", node.output)
            };

            format!(
                "\n{init}\
                \nfor (uint256 {loop_var} = 0; {loop_var} < {}.length;){{\
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

/// Validates that `ty` is a Solidity unsigned integer type (uint8 through uint256, in steps of 8).
fn validate_length_type(ty: &str, field_name: &str, struct_name: &str) -> eyre::Result<()> {
    let bits: Option<u16> = ty.strip_prefix("uint").and_then(|s| s.parse().ok());
    let valid = bits.is_some_and(|n| (8..=256).contains(&n) && (n % 8) == 0);
    if !valid {
        return Err(eyre!(
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
