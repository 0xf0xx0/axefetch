package types

import "strings"

var DefaultConf = Config{
	General: General{
		IP: "replace me",
	},
	Display: Display{
		Format: strings.Join([]string{
			`this is an invalid line, so its not printed :3`,
			`info title`,
			`info underline`,
			`info "Model" model`,
			`info "Asic" asic`,
			// `info "Clock" power`,
			`info "Version" version`,
			`info "Uptime" uptime`,
			`info "Best Difficulty" bestdiff`,
			`info "Shares" shares`,
			`info "Pool" pool`,
			`info "Hashrate" hashrate`,
			`info "Efficiency" efficiency`,
			`info "Temperature" temp`,
			`info "Heap" heap`,
			``,
			`prin circlejerking into open source`,
			``,
		}, "\n"),
		Theme:       "family",
		BoldTitles:  true,
		Separator:   ":",
		Underline:   "-",
		Icon:        "family",
		IconSpacing: 3,
	},
	ColorTheme: ColorTheme{
		Title:     "green",
		At:        "green",
		Underline: "white",
		Subtitle:  "green",
		Separator: "blue",
		Info:      "white",
		Icon:      "green",
	},
	Title: Title{
		Workername: true,
		Hostname:   true,
	},
	Model: Model{
		Boardversion: true,
		Family:       true,
		Vendor:       false,
	},
	Asic: Asic{
		Count:     true,
		Frequency: true,
	},
	Bestdiff: Bestdiff{
		Ath:      true,
		Session:  true,
		Shortpaw: "off",
	},
	Efficiency: Efficiency{
		Expected: true,
		Actual:   true,
		Shortpaw: "off",
	},
	// Firmware: Firmware{},
	Hashrate: Hashrate{
		Expected: true,
		Actual:   true,
		Shortpaw: "off",
	},
	Pool: Pool{
		Port: true,
	},
	Shares: Shares{
		Ratio:    true,
		Shortpaw: "off",
	},
	Temp: Temp{
		Shortpaw: "off",
		Asic:     true,
		Vreg:     true,
	},
	Uptime: Uptime{
		Format: "%dd %hh %mm %ss",
	},
}

type Config struct {
	/// avoid embedding, just because it clutters intellisense
	/// maybe omitempty? unneccesary tho, unmarshalling doesnt touch keys that arent in the conf
	General    `toml:"general"`
	Display    `toml:"display"`
	ColorTheme `toml:"theme" comment:"Supports everything display.format does"`
	Title      `toml:"title"`
	Model      `toml:"model"`
	Asic       `toml:"asic"`
	Bestdiff   `toml:"bestdiff"`
	Efficiency `toml:"efficiency"`
	Hashrate   `toml:"hashrate"`
	Pool       `toml:"pool"`
	Shares     `toml:"shares"`
	Temp       `toml:"temp"`
	Uptime     `toml:"uptime"`
}
type General struct {
	IP string `toml:"ip" comment:"IP address of your *axe"`
}
type Display struct {
	Format      string `toml:"format,multiline" comment:"Neofetch-like\nSee https://pkg.go.dev/git.0xf0xx0.eth.limo/0xf0xx0/oigiki for color tag docs\ninvalid lines are ignored"`
	Icon        string `toml:"icon" comment:"Selected icon name or path\nDefault: 'family'\nValues: 'asic', 'family', 'none', or path to ascii art in a plaintext file"`
	IconSpacing int    `toml:"icon_spacing" comment:"Spaces between the icon and the info"`
	Theme       string `toml:"theme" comment:"Default: 'family'\nValues: 'family', 'manual', or theme name"`
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
	Icon      string `toml:"icon" comment:"Default icon color, if the icon doesn't specify it's own"`
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
type Asic struct {
	Count     bool `toml:"count"`
	Frequency bool `toml:"frequency"`
}
type Bestdiff struct {
	Ath      bool   `toml:"ath"`
	Session  bool   `toml:"session"`
	Shortpaw string `toml:"shortpaw" comment:"'on' or 'off'"`
}
type Efficiency struct {
	Expected bool   `toml:"expected"`
	Actual   bool   `toml:"actual"`
	Shortpaw string `toml:"shortpaw" comment:"'on' or 'off'"`
}
type Hashrate struct {
	Expected bool   `toml:"expected"`
	Actual   bool   `toml:"actual"`
	Shortpaw string `toml:"shortpaw" comment:"'on' or 'off'"`
}
type Pool struct {
	Port bool `toml:"port"`
}
type Shares struct {
	Shortpaw string `toml:"shortpaw" comment:"'on', 'tiny' or 'off'"`
	Ratio    bool   `toml:"ratio"`
}
type Temp struct {
	Shortpaw string `toml:"shortpaw" comment:"'on', 'tiny' or 'off'"`
	Asic     bool   `toml:"asic"`
	Vreg     bool   `toml:"vreg"`
}
type Uptime struct {
	Format string `toml:"format"`
}

//	type Firmware struct {
//		Version bool `toml:"version"`
//	}
