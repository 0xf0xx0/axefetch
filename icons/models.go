package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// TODO: icons for each model and family
var Models = map[string]string{
	"102": ModelFamilies["max"],

	"201": ModelFamilies["ultra"],
	"202": ModelFamilies["ultra"],
	"203": ModelFamilies["ultra"],
	"204": ModelFamilies["ultra"],
	"205": ModelFamilies["ultra"],

	"400": ModelFamilies["supra"],
	"401": ModelFamilies["supra"],
	"402": ModelFamilies["supra"],
	"403": ModelFamilies["supra"],

	"600": ModelFamilies["gamma"],
	"601": ModelFamilies["gamma"],
	"602": ModelFamilies["gamma"],
}
var ModelFamilies = map[string]string{
	"gamma": filepath.Join(paths.MODEL_ROOT, "gamma.txt"),
	"max":   filepath.Join(paths.MODEL_ROOT, "max.txt"),
	"supra": filepath.Join(paths.MODEL_ROOT, "supra.txt"),
	"ultra": filepath.Join(paths.MODEL_ROOT, "ultra.txt"),
}
