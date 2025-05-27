package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// TODO: icons for each model
var Models = map[string][]string{
	"601": loadIcon(filepath.Join(paths.MODEL_ROOT, "gamma.txt")),
}
