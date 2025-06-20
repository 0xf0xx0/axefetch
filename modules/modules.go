package modules

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/0xf0xx0/axefetch/colors"
	"github.com/0xf0xx0/axefetch/types"
)

// these spit out some nice info text
var Modules = map[string]func(types.Config, types.ApiInfo, []string) string{
	/// special
	"title": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := make([]string, 0, 2)
		if conf.Title.Workername {
			workername := ""
			if ai.IsUsingFallbackStratum {
				workername = getWorkerFromUser(ai.FallbackStratumUser)
			} else {
				workername = getWorkerFromUser(ai.StratumUser)
			}
			workername = colors.TagString(workername, conf.ColorTheme.Title)
			if conf.Display.BoldTitles {
				workername = colors.TagString(workername, "bold")
			}
			ret = append(ret, workername)
		}
		if conf.Title.Hostname {
			hostname := colors.TagString(ai.Hostname, conf.ColorTheme.Title)
			if conf.Display.BoldTitles {
				hostname = colors.TagString(hostname, "bold")
			}
			ret = append(ret, hostname)
		}
		return strings.Join(filterEmptyStringsOut(ret), colors.TagString("@", conf.ColorTheme.At))
	},
	// this expects the title string (if any) to be passed in
	"underline": func(conf types.Config, _ types.ApiInfo, args []string) string {
		if len(args) == 0 || len(args[0]) == 0 {
			return ""
		}
		return colors.TagString(strings.Repeat(conf.Display.Underline, len(args[0])), conf.ColorTheme.Underline)
	},

	/// normal functions
	"tbd": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		return fmt.Sprintf("%s@%s", unitFormat(ai.Frequency, "mhz"), unitFormat(float64(ai.CoreVoltage), "mv"))
	},
	"asicmodel": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		/// this gets prepended
		if conf.Asicmodel.Asiccount {
			ret = append(ret, fmt.Sprintf("%dx", ai.AsicCount))
		}
		ret = append(ret, ai.AsicModel)
		return strings.Join(filterEmptyStringsOut(ret), " ")
	},
	"bestdiff": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		shortpawed := conf.Bestdiff.Shortpaw == "on"
		if conf.Bestdiff.Session {
			ret = append(ret, printWithShortpaw(unitFormat(float64(ai.BestSessionDiff), "binshort"), "session", shortpawed))
		}
		if conf.Bestdiff.Ath {
			ret = append(ret, printWithShortpaw(unitFormat(float64(ai.BestDiff), "binshort"), "best", shortpawed))
		}
		if shortpawed {
			return strings.Join(ret, "/")
		}
		return strings.Join(filterEmptyStringsOut(ret), ", ")
	},
	"efficiency": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		shortpawed := conf.Efficiency.Shortpaw == "on"
		if conf.Hashrate.Actual {
			actualEff := ai.Power / (ai.Hashrate / 1000)
			ret = append(ret, printWithShortpaw(unitFormat(actualEff, "j/th"), "(actual)", shortpawed))
		}
		if conf.Hashrate.Expected {
			expectedEff := ai.Power / (ai.ExpectedHashrate / 1000)
			ret = append(ret, printWithShortpaw(unitFormat(expectedEff, "j/th"), "(expected)", shortpawed))
		}
		return strings.Join(filterEmptyStringsOut(ret), ", ")
	},
	"firmware": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{"ESP-Miner"}
		if conf.Firmware.Version {
			ret = append(ret, ai.Version)
		}
		return strings.Join(filterEmptyStringsOut(ret), " ")
	},
	"hashrate": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		/// TODO: add "tiny" display
		shortpawed := conf.Hashrate.Shortpaw == "on"
		if conf.Hashrate.Actual {
			ret = append(ret, printWithShortpaw(unitFormat(ai.Hashrate, "gh/s"), "(actual)", shortpawed))
		}
		if conf.Hashrate.Expected {
			ret = append(ret, printWithShortpaw(unitFormat(ai.ExpectedHashrate, "gh/s"), "(expected)", shortpawed))
		}
		return strings.Join(filterEmptyStringsOut(ret), ", ")
	},
	"heap": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		return unitFormat(float64(ai.FreeHeap), "ib")
	},
	"model": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		if conf.Model.Vendor {
			/// TODO
			// ret = append(ret, ai.BoardVendor)
		}
		if conf.Model.Family {
			ret = append(ret, ai.BoardFamily)
		}
		if conf.Model.Boardversion {
			ret = append(ret, ai.BoardVersion)
		}
		return strings.Join(filterEmptyStringsOut(ret), " ")
	},
	"pool": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := ai.StratumURL
		port := ""
		if ai.IsUsingFallbackStratum {
			ret = ai.FallbackStratumURL
		}
		if conf.Pool.Port {
			port = ":"
			if ai.IsUsingFallbackStratum {
				port += strconv.FormatInt(int64(ai.FallbackStratumPort), 10)
			} else {
				port += strconv.FormatInt(int64(ai.StratumPort), 10)
			}
		}
		return ret + port
	},
	"shares": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := ""
		accepted := unitFormat(float64(ai.SharesAccepted), "short")
		rejected := unitFormat(float64(ai.SharesRejected), "short")
		switch conf.Shares.Shortpaw {
		case "on":
			{
				ret = fmt.Sprintf("%s/%s", accepted, rejected)
			}
		case "tiny":
			{
				ret = fmt.Sprintf("%s/%s (acc/rej)", accepted, rejected)
			}
		case "off":
			{
				ret = fmt.Sprintf("%s accepted, %s rejected", accepted, rejected)
			}
		}
		if conf.Shares.Ratio {
			return fmt.Sprintf("%s (%.2f%%)", ret, float32(ai.SharesRejected)/float32(ai.SharesAccepted)*100)
		}
		return ret
	},
	"temp": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		if conf.Temp.Asic {
			ret = append(ret, fmt.Sprintf("%s (asic)", unitFormat(ai.Temp, "c")))
		}
		if conf.Temp.Vreg {
			ret = append(ret, fmt.Sprintf("%s (vreg)", unitFormat(ai.VrTemp, "c")))
		}
		return strings.Join(filterEmptyStringsOut(ret), ", ")
	},
	"uptime": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		/// TODO: use date format strings?
		time := (time.Second * time.Duration(ai.UptimeSeconds))
		ret := conf.Uptime.Format
		replacer := strings.NewReplacer(
			"%d", strconv.Itoa(int(time.Hours())/24),
			"%h", strconv.Itoa(int(time.Hours())%24),
			"%m", strconv.Itoa(int(time.Minutes())%60),
			"%s", strconv.Itoa(int(time.Seconds())%60),
		)
		return replacer.Replace(ret)
	},
}

// prints a shortened data line
// TODO: rename shouldBeShort
func printWithShortpaw(str, shortpaw string, shouldBeShort bool) string {
	if shouldBeShort {
		return str
	}
	return fmt.Sprintf("%s %s", str, shortpaw)
}

// converts 123456 into 123.45K and 123456789876 into 123.45 G
func floatToBinshort(value float64) string {
	if value >= 1e12 {
		return fmt.Sprintf("%.5gT", value/1e9) /// Trillions
	} else if value >= 1e9 {
		return fmt.Sprintf("%.5gG", value/1e9) /// Billions
	} else if value >= 1e6 {
		return fmt.Sprintf("%.5gM", value/1e6) /// Millions
	} else if value >= 1e3 {
		return fmt.Sprintf("%.5gK", value/1e3) /// Thousands
	}
	return fmt.Sprintf("%.5g", value) /// Less than a thousand
}

// units: gh/s, j/th, mhz, mv, c, ib, short
func unitFormat(value float64, unit string) string {
	switch unit {
	case "gh/s":
		{
			unit = "GH"
			if value > 1000 {
				unit = "TH"
				value /= 1000
			}
			return fmt.Sprintf("%.2f %s/s", value, unit)
		}
	case "j/th":
		{
			/// expected to be precalced
			return fmt.Sprintf("%.3g J/TH", value)
		}
	case "mhz":
		{
			unit = "mHz"
			if value > 1000 {
				unit = "GHz"
				value /= 1000
			}
			return fmt.Sprintf("%.5g %s", value, unit)
		}
	case "mv":
		{
			return fmt.Sprintf("%g mV", value)
		}
	case "c":
		{
			return fmt.Sprintf("%.2f C", value)
		}
	/// wtf is this format even called?
	case "binshort":
		{
			return floatToBinshort(value)
		}
	case "short":
		{
			unit = ""
			if value > 1e12 {
				unit = "T"
				value /= 1e12
			} else if value > 1e9 {
				unit = "B"
				value /= 1e9
			} else if value > 1e6 {
				unit = "M"
				value /= 1e6
			} else if value > 1000 {
				unit = "k"
				value /= 1000
			}
			return fmt.Sprintf("%.3g%s", value, unit)
		}
	case "ib":
		{
			unit = "iB"
			if value > 0x100000 {
				unit = "MiB"
				value /= 0x100000
			} else if value > 1024 {
				unit = "KiB"
				value /= 1024
			}
			return fmt.Sprintf("%.3g %s", value, unit)
		}
	default:
		{
			return ""
		}
	}
}

// tries to get the worker name from the username string ('address.worker'),
// otherwise truncates the address
func getWorkerFromUser(username string) string {
	split := strings.Split(username, ".")
	if len(split) == 1 {
		return fmt.Sprintf("%s...%s", username[:4], username[len(username)-4:])
	}
	return split[1]
}
func filterEmptyStringsOut(s []string) []string {
	return slices.DeleteFunc(s, func(e string) bool {
		if e == "" {
			return true
		}
		return false
	})
}
