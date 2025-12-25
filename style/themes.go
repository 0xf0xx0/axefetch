package style

import "git.0xf0xx0.eth.limo/0xf0xx0/axefetch/types"

// TODO
// vendor color themes
// model color themes
// model family color themes
var Themes = map[string]types.ColorTheme{
	"gamma": {
		Title:     "greenbright",
		At:        "white",
		Underline: "blackbright",
		Subtitle:  "green",
		Separator: "blackbright",
		Info:      "cyan",
		Icon:      "greenbright",
	},
	"supra": {
		Title:     "blue",
		At:        "cyan",
		Underline: "white",
		Subtitle:  "blue",
		Separator: "white",
		Info:      "cyan",
		Icon:      "cyan",
	},
	"ultra": {
		Title:     "magenta",
		At:        "blue",
		Underline: "blackbright",
		Subtitle:  "magenta",
		Separator: "blackbright",
		Info:      "blue",
		Icon:      "bluebright",
	},
	"max": {
		Title:     "redbright",
		At:        "white",
		Underline: "white",
		Subtitle:  "redbright",
		Separator: "white",
		Info:      "red",
		Icon:      "red",
	},

	/// my fur
	"0xf0xx0": {
		Title:     "#deaf8e",
		At:        "#262638",
		Underline: "#262638",
		Subtitle:  "#f04651",
		Separator: "#262638",
		Info:      "white",
		Icon:      "#383631",
	},
}
