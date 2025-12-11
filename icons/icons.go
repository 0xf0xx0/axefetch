package icons

import (
	_ "embed"
)

// embedding
var (
	//go:embed embeds/models/max.txt
	max string
	//go:embed embeds/models/ultra.txt
	ultra string
	//go:embed embeds/models/supra.txt
	supra string
	//go:embed embeds/models/gamma.txt
	gamma string
	// go:embed embeds/models/gammaturbo.txt
	// gammaturbo string
	// go:embed embeds/models/naja.txt
	// naja string
)
var (
	//go:embed embeds/asics/BM1366.txt
	BM1366 string
	//go:embed embeds/asics/BM1368.txt
	BM1368 string
	//go:embed embeds/asics/BM1370.txt
	BM1370 string
	//go:embed embeds/asics/BM1397.txt
	BM1397 string
)
/*var (
	//go:embed embeds/vendors/altair.txt
	altair string
	//go:embed embeds/vendors/d_central.txt
	d_central string
	//go:embed embeds/vendors/gekkoscience.txt
	gekkoscience string
	//go:embed embeds/vendors/solominingde.txt
	solominingde string
	//go:embed embeds/vendors/solominingco.txt
	solominingco string
	//go:embed embeds/vendors/solosatoshi.txt
	solosatoshi string
	//go:embed embeds/vendors/tinychiphub.txt
	tinychiphub string
)*/

var (
	//go:embed embeds/misc/bitcoin.txt
	bitcoin string
	//go:embed embeds/misc/osmu.txt
	osmu string
)

// TODO: icons
// must be lowercase
var Icons = map[string]string{
	/// families
	"max":   max,
	"ultra": ultra,
	"supra": supra,
	"gamma": gamma,

	/// chips
	"bm1366": BM1366,
	"bm1368": BM1368,
	"bm1370": BM1370,
	"bm1397": BM1397,

	/// vendors
	/*"altair":       altair,
	"d-central":    d_central,
	"gekkoscience": gekkoscience,
	"solominingde": solominingde,
	"solominingco": solominingco,
	"solosatoshi":  solosatoshi,
	"tinychiphub":  tinychiphub,
	*/

	/// misc
	"bitcoin": bitcoin,
	"osmu":    osmu,
}
