# icm-macros

A [Reforge](../reforge)-based macro preprocessor for ICM Solidity contracts. It is a drop-in `forge` replacement that runs macro expansion over source files before handing them to `solc`.

## Macros

### `#[pack(...)]`

Derives a `pack` function for a struct or enum. The generated function serializes the type to `bytes`.

**Struct signature:**
```solidity
function pack{TypeName}({TypeName} memory def) public pure returns (bytes memory)
```

**Enum signature:**
```solidity
function pack{TypeName}({TypeName} obj) public pure returns (bytes memory)
```

#### Type-level annotation

Place a `#[pack(...)]` annotation in the NatSpec comment above the type declaration:

```solidity
/// #[pack()]
struct MyStruct {
    uint256 a;
    address b;
}
```

**Optional arguments** (comma-separated key=value pairs):

| Argument | Description | Default |
|---|---|---|
| `contract` | Name of the contract/library to inject the generated function into. Required for free-standing types; otherwise defaults to the type's own contract. | *(none — free-standing)* |
| `name` | Override the generated function name. | `pack{TypeName}` |
| `visibility` | Visibility of the generated function. | `public` |

```solidity
/// #[pack(contract = "MyLibrary", name = "serialize", visibility = "internal")]
struct MyStruct { ... }
```

#### Field-level annotations

Place a `#[pack(...)]` annotation in the NatSpec comment above a struct field to control how it is serialised.

| Annotation | Description |
|---|---|
| `#[pack(ignore)]` | Skip this field entirely. |
| `#[pack(method = "expr")]` | Use `expr(obj.field)` instead of the default pack call. |

```solidity
/// #[pack()]
struct Example {
    uint256 a;
    /// #[pack(ignore)]
    uint256 ignored;
    /// #[pack(method = "MyLib.encodeB")]
    MyType b;
}
```

#### Primitive field encoding

| Type | Encoding |
|---|---|
| Fixed-size elementary (`uint*`, `int*`, `address`, `bool`, `bytes1`–`bytes32`, …) | `abi.encodePacked(field)` |
| `bytes` | `abi.encodePacked(field.length, field)` |
| `string` | `abi.encodePacked(bytes(field).length, field)` |
| Array | Loop over elements, recursively packed, concatenated with `abi.encodePacked` |
| Custom struct/enum/UDVT | `pack{TypeName}(field)` (must be in scope) |
| Mapping / function | **Error** — must provide `#[pack(method = "...")]` or `#[pack(ignore)]` |

### `#[unpack(...)]`

Derives an `unpack` function for a struct or enum. The generated function deserializes the type from a packed byte buffer produced by `#[pack(...)]`.

**Struct signature:**
```solidity
function unpack{TypeName}(bytes memory data) public pure returns (uint256, TypeName memory)
```

The returned `uint256` is the number of bytes consumed. The `data` parameter must be named exactly `data` — generated code references it by name unconditionally.

**Enum signature:**
```solidity
function unpack{TypeName}(bytes memory data) public pure returns (uint256, TypeName)
```

#### Type-level annotation

```solidity
/// #[unpack()]
struct MyStruct {
    uint256 a;
    address b;
}
```

**Optional arguments** — key-value pairs use `key = "value"` syntax; flags are bare identifiers. All comma-separated:

| Argument | Kind | Description | Default |
|---|---|---|---|
| `contract` | key-value | Name of the contract/library to inject the generated function into. | *(none — free-standing)* |
| `name` | key-value | Override the generated function name. | `unpack{TypeName}` |
| `visibility` | key-value | Visibility of the generated function. | `public` |
| `calldata` | flag | Accept `bytes calldata data` instead of `bytes memory data`. | *(off)* |
| `solhint-disable` | flag | Wrap the generated function with `// solhint-disable no-inline-assembly` / `// solhint-enable no-inline-assembly`. | *(off)* |

#### Field-level annotations

| Annotation | Description |
|---|---|
| `#[unpack(default)]` | Skip this field; the struct is returned with its default (zero) value for this field. Useful alongside `#[pack(ignore)]`. |
| `#[unpack(method = "expr")]` | Use `(uint256 read, field) = expr(data)` instead of the inline decoder. The method must have the same signature as a generated unpack function. |

#### Primitive field decoding

| Type | Decoding |
|---|---|
| Fixed-size elementary (`uint*`, `int*`, `address`, `bool`, `bytes1`–`bytes32`, …) | Read exact packed byte width; shift/mask as needed |
| `bytes` / `string` | Read 32-byte length prefix, then copy payload |
| Array | Read 32-byte element count, then decode each element in a loop |
| Custom struct/enum/UDVT | `unpack{TypeName}(data)` (must be in scope) |
| Mapping / function | **Error** — must provide `#[unpack(method = "...")]` or `#[unpack(default)]` |

#### Zero-copy buffer handling — read carefully

The generated code operates **zero-copy on the input buffer**: rather than allocating a separate read cursor, it mutates `data` in place as each field is consumed. Internally, assembly is used to advance `data`'s memory pointer and overwrite its length word after every field read.

**Consequences:**
- `data` is **not safe to read after** calling an unpack function — its pointer and length have been modified to point at whatever bytes were not consumed.
- If you need the original buffer preserved, copy it before passing it in.`
- The one exception is `bytes`/`string` fields in the memory path: their payloads **are** copied into freshly allocated memory (`mcopy`), so the returned field values are independent of `data`.

> **EVM version requirement:** The memory path uses the `mcopy` opcode, which requires **Cancun** or later. Projects deriving `#[unpack()]` must set `evm_version = "cancun"` (or newer) in their `foundry.toml`.

## Usage

Build `icm-macros` and point your Forge project at the resulting binary instead of `forge`:

```sh
cargo build --release
./target/release/icm-macros build --root path/to/project
./target/release/icm-macros test  --root path/to/project
```

All standard `forge` flags are supported. Two additional flags are available:

| Flag | Description |
|---|---|
| `--disable-macros` | Skip macro expansion and behave as a plain `forge` wrapper. |
| `--display <GLOB>` | Print macro-expanded sources matching `GLOB` to stdout and exit. `build` only. |

## Testing

### Unit tests

Tests use the `reforge::testing` utilities. Each macro has its own subdirectory under `testing/`:

```
testing/{macro}/
  source/      ← input .sol files with macro annotations
  expected/    ← pre-expanded .sol files to compare against
  mismatched/  ← written on failure; copy to expected/ to accept new output (gitignored)
  errors/      ← .sol files expected to cause the macro to error, each tested for a specific error message
```

Run the unit tests with:

```sh
cargo test -- --skip tests::test_e2e
```

When a snapshot test fails, the actual output is written to `mismatched/`. Inspect the diff, and if the new output is correct copy it to `expected/`:

```sh
mv testing/{macro}/mismatched/Foo.sol testing/{macro}/expected/Foo.sol
```

### End-to-end tests

`testing/e2e/` is a Foundry project that exercises the full pack/unpack round-trip against a real EVM. It uses `forge-std` (checked out as a git submodule under `testing/e2e/lib/forge-std`).

Initialise the submodule if needed:

```sh
git submodule update --init --recursive
```

Run the end-to-end tests with:

```sh
cargo test tests::test_e2e
```

This builds the `icm-macros` binary and invokes it as `icm-macros test --root testing/e2e`, running all Solidity fuzz tests under `testing/e2e/tests/`.