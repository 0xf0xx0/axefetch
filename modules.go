package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"axefetch/types"

	"github.com/0xf0xx0/oigiki"
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
			workername = oigiki.TagString(workername, conf.ColorTheme.Title)
			if conf.Display.BoldTitles {
				workername = oigiki.TagString(workername, "bold")
			}
			ret = append(ret, workername)
		}
		if conf.Title.Hostname {
			hostname := oigiki.TagString(ai.Hostname, conf.ColorTheme.Title)
			if conf.Display.BoldTitles {
				hostname = oigiki.TagString(hostname, "bold")
			}
			ret = append(ret, hostname)
		}
		return strings.Join(filterEmptyStringsOut(ret), oigiki.TagString("@", conf.ColorTheme.At))
	},
	// this expects the title string (if any) to be passed in
	"underline": func(conf types.Config, _ types.ApiInfo, args []string) string {
		if len(args) == 0 || len(args[0]) == 0 {
			return ""
		}
		return oigiki.TagString(strings.Repeat(conf.Display.Underline, len(args[0])), conf.ColorTheme.Underline)
	},

	/// normal functions
	"clock": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		return fmt.Sprintf("%s@%s", unitFormat(ai.Frequency, "mhz"), unitFormat(float64(ai.CoreVoltage), "mv"))
	},
	"chip": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		/// this gets prepended
		if conf.Chip.Count {
			ret = append(ret, fmt.Sprintf("%dx", ai.AsicCount))
		}
		ret = append(ret, ai.AsicModel)
		return strings.Join(filterEmptyStringsOut(ret), " ")
	},
	"bestdiff": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		shortpawed := conf.Bestdiff.Shortpaw == "on"
		if conf.Bestdiff.Session {
			ret = append(ret, printWithShortpaw(floatToBinshort(ai.BestSessionDiff), "session", shortpawed))
		}
		if conf.Bestdiff.Ath {
			ret = append(ret, printWithShortpaw(floatToBinshort(ai.BestDiff), "best", shortpawed))
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
		ret := []string{}
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
		/// TODO: stratum+tcp://?
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
		shortpawed := conf.Bestdiff.Shortpaw == "on"
		if conf.Temp.Asic {
			ret = append(ret, printWithShortpaw(unitFormat(ai.Temp, "c"), "(asic)", shortpawed))
		}
		if conf.Temp.Vreg {
			ret = append(ret, printWithShortpaw(unitFormat(ai.VrTemp, "c"), "(vreg)", shortpawed))
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
