//! Runs the macro rules as a single orchestrated pass.
//!
//! A reforge macro rule is a plain function pointer, so two registered rules cannot share any
//! state: each sees the same HIR of the original source and nothing about what the other
//! generated. [`derive_all`] is the one registered rule. It owns a [`PackState`], runs the
//! rules in a fixed order, and each rule records what it generated as it goes. A rule can
//! therefore read whatever the rules before it recorded, and that order is spelled out in
//! [`derive_all`] rather than implied by registration order.
//!
//! The individual rules (`derive_pack`, `derive_unpack`) remain callable on their own, with a
//! throwaway [`PackState`], so they can still be tested in isolation.

use std::collections::BTreeMap;

use reforge::PreprocessingData;
use solar::ast::Span;
use solar::interface::BytePos;
use solar::sema::Gcx;
use solar::sema::hir::ContractId;

use crate::pack::derive_pack_with;
use crate::unpack::derive_unpack_with;

/// How one direction (pack or unpack) of the generated code treats a field.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum FieldMode {
    /// Handled by the macro's own encoding or decoding.
    Auto,
    /// Delegated to a user-provided `method`.
    Method,
    /// Not written (`pack(ignore)`) or not read (`unpack(default)`).
    Skipped,
}

/// What a rule generated for one item: enough to emit a correct call to the function and to
/// tell how strictly it can be tested.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct GeneratedFn {
    /// The resolved function name, after any `name` override.
    pub fn_name: String,
    /// The contract or library the function was placed in, if the annotation named one.
    pub contract: Option<ContractId>,
    /// The visibility keyword as annotated. Empty for file-level functions.
    pub visibility: String,
    /// Whether the function takes `bytes calldata data` instead of `bytes memory data`. Always
    /// `false` for pack. A caller has to know: a calldata parameter can only be passed calldata,
    /// and the memory form overwrites the buffer it is given as it reads, so a caller that still
    /// needs the original bytes must pass a copy.
    pub calldata: bool,
    /// One entry per struct field in declaration order. Empty for enums.
    pub fields: Vec<FieldMode>,
}

/// Everything recorded about one annotated item.
#[derive(Debug, Default, Clone, PartialEq, Eq)]
pub struct ItemRecord {
    pub pack: Option<GeneratedFn>,
    pub unpack: Option<GeneratedFn>,
}

/// Identifies an item by where it starts and ends. Byte positions are global across the source
/// map, so a key is unique across files, and keys order by position in the source.
#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
struct ItemKey {
    lo: BytePos,
    hi: BytePos,
}

impl From<Span> for ItemKey {
    fn from(span: Span) -> Self {
        Self {
            lo: span.lo(),
            hi: span.hi(),
        }
    }
}

/// State that the rules write to as they run, keyed by the span of the annotated item.
///
/// Ordered by source position, so a rule that iterates it emits output in a stable order. The
/// expanded contracts are committed and CI fails when they change, so iteration order must not
/// vary from run to run.
#[derive(Debug, Default)]
pub struct PackState {
    items: BTreeMap<ItemKey, ItemRecord>,
}

impl PackState {
    pub fn record_pack(&mut self, item: Span, generated: GeneratedFn) {
        self.items.entry(item.into()).or_default().pack = Some(generated);
    }

    pub fn record_unpack(&mut self, item: Span, generated: GeneratedFn) {
        self.items.entry(item.into()).or_default().unpack = Some(generated);
    }

    /// What the rules that have run so far recorded for the item at `item`, if anything.
    // Nothing reads this yet: it is the interface later rules will use.
    #[cfg_attr(not(test), allow(dead_code))]
    pub fn get(&self, item: Span) -> Option<&ItemRecord> {
        self.items.get(&item.into())
    }
}

/// Runs every rule against `state`. Rules run in the order below, and a rule may read what the
/// ones above it recorded.
fn run_rules(
    ctx: &Gcx,
    data: &mut PreprocessingData<'_>,
    state: &mut PackState,
) -> foundry_compilers::error::Result<()> {
    derive_pack_with(ctx, data, state)?;
    derive_unpack_with(ctx, data, state)
}

/// The single registered macro rule.
pub fn derive_all(
    ctx: &Gcx,
    data: &mut PreprocessingData<'_>,
) -> foundry_compilers::error::Result<()> {
    run_rules(ctx, data, &mut PackState::default())
}

#[cfg(test)]
mod tests {
    use std::sync::Mutex;

    use super::*;

    /// Runs the rules like [`derive_all`], then captures a description of what was recorded for
    /// each annotated item. A rule is a plain function pointer, so the only way to get the state
    /// out of one is a static.
    fn derive_and_capture(
        ctx: &Gcx,
        data: &mut PreprocessingData<'_>,
    ) -> foundry_compilers::error::Result<()> {
        let mut state = PackState::default();
        run_rules(ctx, data, &mut state)?;

        let describe = |direction: &str, generated: &Option<GeneratedFn>| {
            generated.as_ref().map(|g| {
                let contract = g
                    .contract
                    .map_or("-".to_string(), |id| ctx.hir.contract(id).name.to_string());
                format!(
                    "{direction} {} in {contract} [{}] calldata={} {:?}",
                    g.fn_name, g.visibility, g.calldata, g.fields
                )
            })
        };
        let mut captured = Vec::new();
        let items = ctx
            .hir
            .enums()
            .map(|e| (e.name.to_string(), e.span))
            .chain(ctx.hir.structs().map(|s| (s.name.to_string(), s.span)));
        for (name, span) in items {
            if let Some(record) = state.get(span) {
                captured.push(format!(
                    "{name}: {} | {}",
                    describe("pack", &record.pack).unwrap_or_default(),
                    describe("unpack", &record.unpack).unwrap_or_default(),
                ));
            }
        }
        *CAPTURED.lock().unwrap() = captured;
        Ok(())
    }

    /// What the last [`derive_and_capture`] run recorded. Only `records_generated_functions`
    /// uses it, so tests running in parallel do not interfere.
    static CAPTURED: Mutex<Vec<String>> = Mutex::new(Vec::new());

    /// With only one kind of annotation present, the orchestrator must emit exactly what the
    /// individual rule did.
    #[test]
    fn matches_pack_fixtures() {
        let mismatched = "testing/orchestrator/mismatched/pack";
        let _ = std::fs::remove_dir_all(mismatched);
        std::fs::create_dir_all(mismatched).expect("Failed to create mismatched dir");
        reforge::testing::test_macros(
            "testing/pack/source",
            "testing/pack/expected",
            mismatched,
            &[derive_all],
        )
        .expect("derive_all differs from derive_pack on the pack fixtures");
    }

    #[test]
    fn matches_unpack_fixtures() {
        let mismatched = "testing/orchestrator/mismatched/unpack";
        let _ = std::fs::remove_dir_all(mismatched);
        std::fs::create_dir_all(mismatched).expect("Failed to create mismatched dir");
        reforge::testing::test_macros(
            "testing/unpack/source",
            "testing/unpack/expected",
            mismatched,
            &[derive_all],
        )
        .expect("derive_all differs from derive_unpack on the unpack fixtures");
    }

    /// Items carrying both `#[pack]` and `#[unpack]` get both functions generated in one pass.
    #[test]
    fn generates_both_directions() {
        let mismatched = "testing/orchestrator/mismatched/both";
        let _ = std::fs::remove_dir_all(mismatched);
        std::fs::create_dir_all(mismatched).expect("Failed to create mismatched dir");
        reforge::testing::test_macros(
            "testing/orchestrator/source",
            "testing/orchestrator/expected",
            mismatched,
            &[derive_all],
        )
        .expect("derive_all output does not match testing/orchestrator/expected");
    }

    /// The state records the resolved name, placement and per-field handling of each function.
    #[test]
    fn records_generated_functions() {
        let mismatched = "testing/orchestrator/mismatched/state";
        let _ = std::fs::remove_dir_all(mismatched);
        std::fs::create_dir_all(mismatched).expect("Failed to create mismatched dir");
        CAPTURED.lock().unwrap().clear();
        reforge::testing::test_macros(
            "testing/orchestrator/source",
            "testing/orchestrator/expected",
            mismatched,
            &[derive_and_capture],
        )
        .expect("derive_and_capture output does not match testing/orchestrator/expected");

        let captured = CAPTURED.lock().unwrap().clone();
        assert_eq!(
            captured,
            vec![
                "Kind: pack packKind in Messages [public] calldata=false [] \
                 | unpack unpackKind in Messages [public] calldata=false []",
                "Header: pack encodeHeader in Messages [internal] calldata=false \
                 [Auto, Auto, Skipped, Method] \
                 | unpack decodeHeader in Messages [internal] calldata=true \
                 [Auto, Auto, Skipped, Method]",
            ]
        );
    }
}
