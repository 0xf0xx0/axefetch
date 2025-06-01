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
			if ai.IsUsingFallbackStratum == 1 {
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
			ret = append(ret, printWithShortpaw(ai.BestSessionDiff, "session", shortpawed))
		}
		if conf.Bestdiff.Ath {
			ret = append(ret, printWithShortpaw(ai.BestDiff, "best", shortpawed))
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
			ret = append(ret, printWithShortpaw(fmt.Sprintf("%.2f J/TH", actualEff), "(actual)", shortpawed))
		}
		if conf.Hashrate.Expected {
			expectedEff := ai.Power / (float64(ai.ExpectedHashrate) / 1000)
			ret = append(ret, printWithShortpaw(fmt.Sprintf("%.2f J/TH", expectedEff), "(expected)", shortpawed))
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
			unit := "GH"
			hashrate := ai.Hashrate
			if hashrate > 1000 {
				unit = "TH"
				hashrate /= 1000
			}
			ret = append(ret, printWithShortpaw(fmt.Sprintf("%.2f %s/s", hashrate, unit), "(actual)", shortpawed))
		}
		if conf.Hashrate.Expected {
			unit := "GH"
			expected := ai.ExpectedHashrate
			if expected > 1000 {
				unit = "TH"
				expected /= 1000
			}
			ret = append(ret, printWithShortpaw(fmt.Sprintf("%d %s/s", expected, unit), "(expected)", shortpawed))
		}
		return strings.Join(filterEmptyStringsOut(ret), ", ")
	},
	"heap": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		/// MAYBE: use a unit format module?
		mib := float32(ai.FreeHeap) / (1024 * 1024)
		return fmt.Sprintf("%.2g MiB", mib)
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
		if ai.IsUsingFallbackStratum == 1 {
			ret = ai.FallbackStratumURL
		}
		if conf.Pool.Port {
			port = ":"
			if ai.IsUsingFallbackStratum == 1 {
				port += strconv.FormatInt(int64(ai.FallbackStratumPort), 10)
			} else {
				port += strconv.FormatInt(int64(ai.StratumPort), 10)
			}
		}
		return ret + port
	},
	"shares": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := ""
		switch conf.Shares.Shortpaw {
			case "on": {
				ret = fmt.Sprintf("%d/%d", ai.SharesAccepted, ai.SharesRejected)
			}
			case "tiny": {
				ret = fmt.Sprintf("%d/%d (acc/rej)", ai.SharesAccepted, ai.SharesRejected)
			}
			case "off": {
				ret = fmt.Sprintf("%d accepted, %d rejected", ai.SharesAccepted, ai.SharesRejected)
			}
		}
		if conf.Shares.Ratio {
			return fmt.Sprintf("%s (%.2f%%)", ret, float32(ai.SharesRejected)/float32(ai.SharesAccepted)*100)
		}
		return ret
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
