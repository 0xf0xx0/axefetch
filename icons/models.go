package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// TODO: icons for each model and family
var Models = map[string][]string{
	"601": ModelFamilies["Gamma"],
}
var ModelFamilies = map[string][]string{
	"Gamma": loadIcon(filepath.Join(paths.MODEL_ROOT, "gamma.txt")),
}
