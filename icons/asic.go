package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// TODO: icons for each chip
var Asics = map[string]string{
	"BM1366": filepath.Join(paths.ASIC_ROOT, "BM1366.txt"),
	"BM1368": filepath.Join(paths.ASIC_ROOT, "BM1368.txt"),
	"BM1370": filepath.Join(paths.ASIC_ROOT, "BM1370.txt"),
	"BM1397": filepath.Join(paths.ASIC_ROOT, "BM1397.txt"),
}
