package icons

import (
	"path/filepath"

	"axefetch/paths"
)

// TODO: icons
var Icons = map[string]string{
	/// families
	"gamma": filepath.Join(paths.MODEL_ROOT, "gamma.txt"),
	"max":   filepath.Join(paths.MODEL_ROOT, "max.txt"),
	"supra": filepath.Join(paths.MODEL_ROOT, "supra.txt"),
	"ultra": filepath.Join(paths.MODEL_ROOT, "ultra.txt"),

	/// chips
	"BM1366": filepath.Join(paths.ASIC_ROOT, "BM1366.txt"),
	"BM1368": filepath.Join(paths.ASIC_ROOT, "BM1368.txt"),
	"BM1370": filepath.Join(paths.ASIC_ROOT, "BM1370.txt"),
	"BM1397": filepath.Join(paths.ASIC_ROOT, "BM1397.txt"),

	/// vendors
	"altair": filepath.Join(paths.VENDOR_ROOT, "altair.txt"),
	"d-central": filepath.Join(paths.VENDOR_ROOT, "d-central.txt"),
	"gekkoscience": filepath.Join(paths.VENDOR_ROOT, "gekkoscience.txt"),
	"solominingde": filepath.Join(paths.VENDOR_ROOT, "solominingde.txt"),
	"solominingco": filepath.Join(paths.VENDOR_ROOT, "solominingco.txt"),
	"solosatoshi": filepath.Join(paths.VENDOR_ROOT, "solosatoshi.txt"),
	"tinychiphub": filepath.Join(paths.VENDOR_ROOT, "tinychiphub.txt"),

	/// misc
	"bitcoin": filepath.Join(paths.MISC_ROOT, "bitcoin.txt"),
	"osmu": filepath.Join(paths.MISC_ROOT, "osmu.txt"),
}
