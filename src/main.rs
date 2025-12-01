mod theme;

use serde_json::{Value, json};
use std::{self, io::Write};

fn main() {
    let colors = theme::Colors::new(theme::Variant {
        style: theme::Style::Material,
        strength: theme::Strength::Hard,
        mode: theme::Mode::Dark,
    });

    let scheme = include_str!("scheme.json");

    let mut processed = scheme.to_string();
    for (key, value) in colors.iter() {
        let placeholder = format!("{{{{{}}}}}", key);
        processed = processed.replace(&placeholder, value);
    }

    processed = processed.replace("{{style}}", &format!("{:#?}", theme::Style::Material));
    processed = processed.replace("{{strength}}", &format!("{:#?}", theme::Strength::Hard));
    processed = processed.replace("{{mode}}", &format!("{:#?}", theme::Mode::Dark));

    let themes: Value = serde_json::from_str(&processed).expect("Invalid JSON after substitution");

    let json = json!({
        "name": "Zedbox",
        "author": "isaiah hamilton <isaiah-hamilton@gmail.com>",
        "themes": [
            themes
        ]
    });

    let file_path = std::path::Path::new("./themes/zedbox.json");
    let pretty = serde_json::to_string_pretty(&json).unwrap();

    let mut file = std::fs::File::create(file_path).unwrap();
    file.write_all(pretty.as_bytes()).unwrap();

    println!(
        "'{}' has been created or overwritten successfully.",
        file_path.display()
    );
}
