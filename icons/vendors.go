package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// bitaxe vendor list https://bitaxe.org/buy
// at least, the ones i know are legit
// TODO: icons for each vendor
var Vendors = map[string]string{
	"altair": filepath.Join(paths.VENDOR_ROOT, "altair.txt"),
	"d-central": filepath.Join(paths.VENDOR_ROOT, "d-central.txt"),
	"gekkoscience": filepath.Join(paths.VENDOR_ROOT, "gekkoscience.txt"),
	"solominingde": filepath.Join(paths.VENDOR_ROOT, "solominingde.txt"),
	"solominingco": filepath.Join(paths.VENDOR_ROOT, "solominingco.txt"),
	"solosatoshi": filepath.Join(paths.VENDOR_ROOT, "solosatoshi.txt"),
	"tinychiphub": filepath.Join(paths.VENDOR_ROOT, "tinychiphub.txt"),
}
