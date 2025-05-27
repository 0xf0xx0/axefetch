package paths

import (
	"fmt"
	"os"
	"path/filepath"
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
var ICON_PATHS = []string{
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

func MakeConfigDirTree() {
	mkdirIfNotExist(VENDOR_ROOT, 0755)
	mkdirIfNotExist(ASIC_ROOT, 0755)
	mkdirIfNotExist(MODEL_ROOT, 0755)
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
