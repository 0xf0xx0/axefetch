package colors

import "axefetch/types"

// TODO
// vendor color themes
// model color themes
// model family color themes
// color settings
var Themes = map[string]types.ColorTheme{
	"gamma": {
		Title:     "green",
		At:        "white",
		Underline: "blackbright",
		Subtitle:  "greenbright",
		Separator: "blackbright",
		Info:      "green",
		Icon:      "green",
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

	"0xf0xx0": {
		Title:     "#afbbd9",
		At:        "#acb7b4",
		Underline: "#5f5a4c",
		Subtitle:  "#768b55",
		Separator: "#5f5a4c",
		Info:      "#acb7b4",
		Icon:      "#5f5a4c",
	},
}
