package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// TODO: icons for each model and family
var Models = map[string][]string{
	"102": ModelFamilies["Max"],

	"201": ModelFamilies["Ultra"],
	"202": ModelFamilies["Ultra"],
	"203": ModelFamilies["Ultra"],
	"204": ModelFamilies["Ultra"],
	"205": ModelFamilies["Ultra"],

	"400": ModelFamilies["Supra"],
	"401": ModelFamilies["Supra"],
	"402": ModelFamilies["Supra"],
	"403": ModelFamilies["Supra"],

	"600": ModelFamilies["Gamma"],
	"601": ModelFamilies["Gamma"],
	"602": ModelFamilies["Gamma"],
}
var ModelFamilies = map[string][]string{
	"Gamma": loadIcon(filepath.Join(paths.MODEL_ROOT, "gamma.txt")),
	"Max":   loadIcon(filepath.Join(paths.MODEL_ROOT, "max.txt")),
	"Supra": loadIcon(filepath.Join(paths.MODEL_ROOT, "supra.txt")),
	"Ultra": loadIcon(filepath.Join(paths.MODEL_ROOT, "ultra.txt")),
}
