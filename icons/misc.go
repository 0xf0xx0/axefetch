package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

var Misc = map[string][]string{
	"osmu": loadIcon(filepath.Join(paths.MISC_ROOT, "osmu.txt")),
}
