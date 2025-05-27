package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)


// TODO: icons for each chip
var Asics = map[string][]string{
	"BM1370": loadIcon(filepath.Join(paths.ASIC_ROOT, "BM1370.txt")),
}
