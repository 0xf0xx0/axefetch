package types

type Config struct {
	General struct {
		IP          string `toml:"ip" comment:"IP address of your *axe"`
		Separator   string `toml:"separator" comment:"Separator between subtitle and info"`
		Underline   string `toml:"underline" comment:"Underline char"`
		Icon        bool   `toml:"icon" comment:"Selected icon name or path"`
		IconType    string `toml:"icon_type" comment:"Default: 'model'\nValues: 'vendor', 'family', 'model', path to ascii art in a plaintext file"`
		IconSpacing int    `toml:"icon_spacing" comment:"Spaces between the icon and the info"`
	} `toml:"general"`
	Display struct {
		Format     string `toml:"format" comment:"Neofetch-like, uses 'info' and 'prin'\nSupports 16 colors and hex colors, wrap them in {}"`
		Colors     string `toml:"colors" comment:"Default: 'board'\nBalues: 'vendor', 'family', 'manual'"`
		BoldTitles bool   `toml:"bold_titles"`
	} `toml:"display"`
	Colors struct {
		Title     string `toml:"title"`
		At        string `toml:"at"`
		Underline string `toml:"underline"`
		Subtitle  string `toml:"subtitle"`
		Separator string `toml:"separator"`
		Info      string `toml:"info"`
	} `toml:"colors"`
	Title struct {
		Workername bool `toml:"worker_name" comment:""`
		Hostname   bool `toml:"hostname"`
	} `toml:"title"`
	Model struct {
		Boardversion bool `toml:"board_version"`
		Family       bool `toml:"family"`
		Vendor       bool `toml:"vendor"`
	} `toml:"model"`
	Asicmodel struct {
		Asiccount      bool `toml:"asic_count"`
		Smallcorecount bool `toml:"small_core_count"`
	} `toml:"asicmodel"`
	Bestdiff struct {
		Ath      bool   `toml:"ath"`
		Session  bool   `toml:"session"`
		Shortpaw string `toml:"shortpaw" comment:""`
	} `toml:"bestdiff"`
	Efficiency struct {
		Expected bool   `toml:"expected"`
		Actual   bool   `toml:"actual"`
		Shortpaw string `toml:"shortpaw" comment:""`
	} `toml:"efficiency"`
	Firmware struct {
		Version bool `toml:"version"`
	} `toml:"firmware"`
	Hashrate struct {
		Expected bool   `toml:"expected"`
		Actual   bool   `toml:"actual"`
		Shortpaw string `toml:"shortpaw" comment:""`
	} `toml:"hashrate"`
	Pool struct {
		Port bool `toml:"port"`
	} `toml:"pool"`
	Shares struct {
		Shortpaw string `toml:"shortpaw" comment:""`
		Ratio    bool   `toml:"ratio" comment:""`
	} `toml:"shares"`
	Uptime struct {
		Format string `toml:"format"`
	} `toml:"uptime"`
}
