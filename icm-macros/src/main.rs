mod solhint;
mod pack;
mod unpack;

use reforge::MacroRules;

use crate::pack::derive_pack;
use crate::unpack::derive_unpack;

fn main() -> eyre::Result<()> {
    let mut macros = MacroRules::default();
    macros.rules.push(derive_pack);
    macros.rules.push(derive_unpack);
    macros.run()
}
