package paths

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/0xf0xx0/axefetch/types"
)

// because i cant inject...
var CONFIG_ROOT = getConfigDir()
var ICON_ROOT = filepath.Join(CONFIG_ROOT, "icons/")
var ASIC_ROOT = filepath.Join(ICON_ROOT, "asics/")
var MODEL_ROOT = filepath.Join(ICON_ROOT, "models/")
var MISC_ROOT = filepath.Join(ICON_ROOT, "misc/")
var VENDOR_ROOT = filepath.Join(ICON_ROOT, "vendors/")
var PATHS = []string{
	CONFIG_ROOT,
	ICON_ROOT,
	ASIC_ROOT,
	MISC_ROOT,
	MODEL_ROOT,
	VENDOR_ROOT,
}

func getConfigDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		println(fmt.Sprintf("error getting config dir: %s", err))
		os.Exit(1)
	}
	return filepath.Join(userConfigDir, "./axefetch")
}

func MakeConfigDirTree(defaultConf types.Config) bool {
	if _, err := os.Stat(CONFIG_ROOT); err != nil {
		for _, path := range PATHS {
			mkdirIfNotExist(path, 0755)
		}
		return true
	}
	return false
}
func mkdirIfNotExist(path string, perm os.FileMode) {
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, perm)
		if err != nil {
			println(fmt.Sprintf("couldnt mkdir %s: %s", path, err))
		}
	}
	/// path exists, do nothing
}
