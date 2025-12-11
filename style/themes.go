package style

import "axefetch/types"

// TODO
// vendor color themes
// model color themes
// model family color themes
// color settings
var Themes = map[string]types.ColorTheme{
	"gamma": {
		Title:     "greenbright",
		At:        "white",
		Underline: "blackbright",
		Subtitle:  "greenbright",
		Separator: "blackbright",
		Info:      "greenbright",
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
		Icon:      "blue",
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
		At:        "#f04651",
		Underline: "#383631",
		Subtitle:  "#f04651",
		Separator: "#383631",
		Info:      "#deaf8e",
		Icon:      "#deaf8e",
	},
}
