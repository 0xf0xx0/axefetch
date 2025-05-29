package icons

import (
	"path/filepath"

	"github.com/0xf0xx0/axefetch/paths"
)

// bitaxe vendor list https://bitaxe.org/buy
// at least, the ones i know are legit
// TODO: icons for each vendor
var Vendors = map[string][]string{
	"altair": loadIcon(filepath.Join(paths.VENDOR_ROOT, "altair.txt")),
	"d-central": loadIcon(filepath.Join(paths.VENDOR_ROOT, "d-central.txt")),
	"gekkoscience": loadIcon(filepath.Join(paths.VENDOR_ROOT, "gekkoscience.txt")),
	"solomining": loadIcon(filepath.Join(paths.VENDOR_ROOT, "solomining.txt")),
	"solominingco": loadIcon(filepath.Join(paths.VENDOR_ROOT, "solominingco.txt")),
	"solosatoshi": loadIcon(filepath.Join(paths.VENDOR_ROOT, "solosatoshi.txt")),
	"tinychiphub": loadIcon(filepath.Join(paths.VENDOR_ROOT, "tinychiphub.txt")),
}
