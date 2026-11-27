mod common;
mod orchestrator;
mod pack;
mod solhint;
mod unpack;

use reforge::MacroRules;

use crate::orchestrator::derive_all;

fn main() -> eyre::Result<()> {
    let mut macros = MacroRules::default();
    // A single rule: `derive_all` runs the pack and unpack rules in order over shared state.
    // See the `orchestrator` module for why the rules are not registered individually.
    macros.rules.push(derive_all);
    macros.run()
}
