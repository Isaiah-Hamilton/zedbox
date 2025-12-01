use serde::Deserialize;
use std;

#[derive(Debug, Deserialize)]
pub struct Colors {
    pub bg_dim: String,
    pub bg0: String,
    pub bg1: String,
    pub bg2: String,
    pub bg3: String,
    pub bg4: String,
    pub bg5: String,
    pub bg_statusline1: String,
    pub bg_statusline2: String,
    pub bg_statusline3: String,
    pub bg_visual_red: String,
    pub bg_visual_yellow: String,
    pub bg_visual_green: String,
    pub bg_visual_blue: String,
    pub bg_visual_purple: String,
    pub bg_diff_red: String,
    pub bg_diff_green: String,
    pub bg_diff_blue: String,
    pub bg_current_word: String,
    pub fg0: String,
    pub fg1: String,
    pub red: String,
    pub orange: String,
    pub yellow: String,
    pub green: String,
    pub aqua: String,
    pub blue: String,
    pub purple: String,
    pub bg_red: String,
    pub bg_green: String,
    pub bg_yellow: String,
    pub grey0: String,
    pub grey1: String,
    pub grey2: String,
}

#[derive(Debug)]
pub enum Style {
    Material,
    Mix,
    Original,
}

#[derive(Debug)]
pub enum Strength {
    Soft,
    Medium,
    Hard,
}

#[derive(Debug)]
pub enum Mode {
    Light,
    Dark,
}

#[derive(Debug)]
pub struct Variant {
    pub style: Style,
    pub strength: Strength,
    pub mode: Mode,
}

impl Colors {
    pub fn new(variant: Variant) -> Self {
        match (variant.style, variant.strength, variant.mode) {
            (Style::Material, Strength::Hard, Mode::Dark) => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Hard,
                mode: Mode::Dark,
            }),

            (Style::Material, Strength::Medium, Mode::Dark) => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Medium,
                mode: Mode::Dark,
            }),

            (Style::Material, Strength::Soft, Mode::Dark) => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Soft,
                mode: Mode::Dark,
            }),

            (Style::Material, Strength::Hard, Mode::Light) => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Hard,
                mode: Mode::Light,
            }),

            (Style::Material, Strength::Medium, Mode::Light) => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Medium,
                mode: Mode::Light,
            }),

            (Style::Material, Strength::Soft, Mode::Light) => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Soft,
                mode: Mode::Light,
            }),

            _ => Self::read(Variant {
                style: Style::Material,
                strength: Strength::Soft,
                mode: Mode::Light,
            }),
        }
    }

    fn read(variant: Variant) -> Colors {
        let file_name = format!(
            "./src/colors/{:#?}_{:#?}_{:#?}.json",
            variant.style, variant.strength, variant.mode
        );

        let file = std::fs::File::open(file_name).unwrap();
        let reader = std::io::BufReader::new(file);

        let colors: Colors =
            serde_json::from_reader(reader).expect("Failed to parse JSON from file");

        colors
    }

    pub fn iter(&self) -> impl Iterator<Item = (&'static str, &String)> {
        [
            ("bg_dim", &self.bg_dim),
            ("bg0", &self.bg0),
            ("bg1", &self.bg1),
            ("bg2", &self.bg2),
            ("bg3", &self.bg3),
            ("bg4", &self.bg4),
            ("bg5", &self.bg5),
            ("bg_statusline1", &self.bg_statusline1),
            ("bg_statusline2", &self.bg_statusline2),
            ("bg_statusline3", &self.bg_statusline3),
            ("bg_visual_red", &self.bg_visual_red),
            ("bg_visual_yellow", &self.bg_visual_yellow),
            ("bg_visual_green", &self.bg_visual_green),
            ("bg_visual_blue", &self.bg_visual_blue),
            ("bg_visual_purple", &self.bg_visual_purple),
            ("bg_diff_red", &self.bg_diff_red),
            ("bg_diff_green", &self.bg_diff_green),
            ("bg_diff_blue", &self.bg_diff_blue),
            ("bg_current_word", &self.bg_current_word),
            ("fg0", &self.fg0),
            ("fg1", &self.fg1),
            ("red", &self.red),
            ("orange", &self.orange),
            ("yellow", &self.yellow),
            ("green", &self.green),
            ("aqua", &self.aqua),
            ("blue", &self.blue),
            ("purple", &self.purple),
            ("bg_red", &self.bg_red),
            ("bg_green", &self.bg_green),
            ("bg_yellow", &self.bg_yellow),
            ("grey0", &self.grey0),
            ("grey1", &self.grey1),
            ("grey2", &self.grey2),
        ]
        .into_iter()
    }
}
