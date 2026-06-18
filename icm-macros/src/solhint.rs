/// Returns `(open, close)` strings for wrapping generated Solidity code in
/// `/* solhint-disable */` / `/* solhint-enable */` comment blocks.
///
/// Each rule gets its own line. When `rules` is empty both strings are empty —
/// callers that know no rules fire can omit the guard entirely.
///
/// `open` ends with `"\n    "` so that code immediately following it is correctly
/// indented at the 4-space function-body level. `close` begins with `"\n    "` so
/// it lands on its own line at the same indentation level.
pub(crate) fn solhint_guards(rules: &[&str]) -> (String, String) {
    if rules.is_empty() {
        return (String::new(), String::new());
    }
    let disable = rules
        .iter()
        .map(|r| format!("/* solhint-disable {r} */"))
        .collect::<Vec<_>>()
        .join("\n    ");
    let enable = rules
        .iter()
        .map(|r| format!("/* solhint-enable {r} */"))
        .collect::<Vec<_>>()
        .join("\n    ");
    (format!("{disable}\n    "), format!("\n    {enable}"))
}