package types

type Config struct {
	General struct {
		Separator   string `toml:"separator" comment:"Separator between subtitle and info"`
		Underline   string `toml:"underline" comment:"Underline char"`
		Icon        bool   `toml:"icon" comment:"selected icon name or path"`
		IconType    string `toml:"icon_type" comment:"default: 'model'\nvalues: 'vendor', 'family', 'model', path to ascii art in a plaintext file"`
		IconSpacing int    `toml:"icon_spacing" comment:"spaces between the icon and the info"`
	} `toml:"general"`
	Display struct {
		Format string `toml:"format" comment:"neofetch-like, uses 'info' and 'prin'\nsupports 16 colors and hex colors, wrap them in {}"`
		Colors string `toml:"colors" comment:"default: 'board'\nvalues: 'vendor', 'family', 'manual'"`
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
		Workername bool `toml:"workername" comment:""`
		Hostname   bool `toml:"hostname"`
	} `toml:"title"`
	Model struct {
		Boardversion bool `toml:"boardversion"`
	} `toml:"model"`
	Asicmodel struct {
		Asiccount      bool `toml:"asiccount"`
		Smallcorecount bool `toml:"smallcorecount"`
	} `toml:"asicmodel"`
	Bestdiff struct {
		Ath      bool   `toml:"ath"`
		Session  bool   `toml:"session"`
		Shortpaw string `toml:"shortpaw" comment:""`
	} `toml:"bestdiff"`
	Firmware struct {
		Version bool `toml:"version"`
	} `toml:"firmware"`
	Shares struct {
		Shortpaw string `toml:"shortpaw" comment:""`
		Ratio    bool   `toml:"ratio" comment:""`
	} `toml:"shares"`
}
