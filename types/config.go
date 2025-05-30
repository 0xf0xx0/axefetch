package types

type Config struct {
	General    `toml:"general"`
	Display    `toml:"display"`
	ColorTheme `toml:"theme" comment:"Supports everything display.format does"`
	Title      `toml:"title"`
	Model      `toml:"model"`
	Asicmodel  `toml:"asicmodel"`
	Bestdiff   `toml:"bestdiff"`
	Efficiency `toml:"efficiency"`
	Firmware   `toml:"firmware"`
	Hashrate   `toml:"hashrate"`
	Pool       `toml:"pool"`
	Shares     `toml:"shares"`
	Uptime     `toml:"uptime"`
}
type General struct {
	IP string `toml:"ip" comment:"IP address of your *axe"`
}
type Display struct {
	Format      string `toml:"format" comment:"Neofetch-like, uses 'info' and 'prin'\nSupports 16 and hex colors, bg coloring, and bold/italic/underline with chainable color tags\n'{white}', '{bg#ff00ff}', '{italic}{bgmagentabright}'\ninvalid lines are ignored"`
	Icon        string `toml:"icon" comment:"Selected icon name or path\nDefault: 'model'\nValues: 'none', 'vendor', 'family', 'model', path to ascii art in a plaintext file"`
	IconSpacing int    `toml:"icon_spacing" comment:"Spaces between the icon and the info"`
	Theme       string `toml:"theme" comment:"Default: 'family'\nValues: 'vendor', 'family', 'manual', or theme name"`
	BoldTitles  bool   `toml:"bold_titles"`
	Separator   string `toml:"separator" comment:"Separator between subtitle and info"`
	Underline   string `toml:"underline" comment:"Underline char"`
}
type ColorTheme struct {
	Title     string `toml:"title"`
	At        string `toml:"at"`
	Underline string `toml:"underline"`
	Subtitle  string `toml:"subtitle"`
	Separator string `toml:"separator"`
	Info      string `toml:"info"`
}
type Title struct {
	Workername bool `toml:"worker_name" comment:""`
	Hostname   bool `toml:"hostname"`
}
type Model struct {
	Boardversion bool `toml:"board_version"`
	Family       bool `toml:"family"`
	Vendor       bool `toml:"vendor"`
}
type Asicmodel struct {
	Asiccount      bool `toml:"asic_count"`
	Smallcorecount bool `toml:"small_core_count"`
}
type Bestdiff struct {
	Ath      bool   `toml:"ath"`
	Session  bool   `toml:"session"`
	Shortpaw string `toml:"shortpaw" comment:""`
}
type Efficiency struct {
	Expected bool   `toml:"expected"`
	Actual   bool   `toml:"actual"`
	Shortpaw string `toml:"shortpaw" comment:""`
}
type Firmware struct {
	Version bool `toml:"version"`
}
type Hashrate struct {
	Expected bool   `toml:"expected"`
	Actual   bool   `toml:"actual"`
	Shortpaw string `toml:"shortpaw" comment:""`
}
type Pool struct {
	Port bool `toml:"port"`
}
type Shares struct {
	Shortpaw string `toml:"shortpaw" comment:""`
	Ratio    bool   `toml:"ratio" comment:""`
}
type Uptime struct {
	Format string `toml:"format"`
}
