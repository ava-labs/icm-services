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

#[cfg(test)]
mod tests {
    #[test]
    fn test_e2e() {
        // Derive the path from the current test executable
        // (target/<profile>/deps/<bin> -> target/<profile>/).
        let mut bin = std::env::current_exe().expect("could not get test binary path");
        bin.pop(); // deps/
        bin.pop(); // profile dir (debug / release)
        bin.push("icm-macros");

        let output = std::process::Command::new(&bin)
            .args(["test", "--root", "testing/e2e"])
            .output()
            .unwrap_or_else(|e| panic!("failed to spawn {}: {e}", bin.display()));
        if !output.status.success() {
            eprintln!("stdout:\n{}", String::from_utf8_lossy(&output.stdout));
            eprintln!("stderr:\n{}", String::from_utf8_lossy(&output.stderr));
            panic!("e2e Solidity tests failed");
        }
    }
}
