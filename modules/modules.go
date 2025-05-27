package modules

import (
	"fmt"
	"strings"
	"time"

	"github.com/0xf0xx0/axefetch/colors"
	"github.com/0xf0xx0/axefetch/types"
)

// "github.com/0xf0xx0/axefetch/config"
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
			ret = append(ret, colors.TagString(workername, conf.Colors.Title))
		}
		if conf.Title.Hostname {
			ret = append(ret, colors.TagString(ai.Hostname, conf.Colors.Title))
		}
		return strings.Join(ret, colors.TagString("@", conf.Colors.At))
	},
	// this expects the title string (if any) to be passed in
	"underline": func(conf types.Config, _ types.ApiInfo, args []string) string {
		if len(args) == 0 || len(args[0]) == 0 {
			return ""
		}
		return colors.TagString(strings.Repeat(conf.General.Underline, len(args[0])), conf.Colors.Underline)
	},

	/// normal functions
	"model": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		if conf.Model.Boardversion {
			ret = append(ret, ai.BoardVersion)
		}
		return strings.Join(ret, " ")
	},
	"asicmodel": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		/// this gets prepended
		if conf.Asicmodel.Asiccount {
			ret = append(ret, fmt.Sprintf("%dx", ai.AsicCount))
		}
		ret = append(ret, ai.AsicModel)
		return strings.Join(ret, " ")
	},
	"firmware": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{"ESP-Miner"}
		if conf.Firmware.Version {
			ret = append(ret, ai.Version)
		}
		return strings.Join(ret, " ")
	},
	"uptime": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		/// TODO: use date format strings?
		return (time.Second * time.Duration(ai.UptimeSeconds)).String()
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
		return strings.Join(ret, ", ")
	},
	"shares": func(conf types.Config, ai types.ApiInfo, _ []string) string {
		ret := []string{}
		shortpawed := conf.Shares.Shortpaw == "on"
		ret = append(ret, printWithShortpaw(fmt.Sprintf("%d", ai.SharesAccepted), "accepted", shortpawed))
		ret = append(ret, printWithShortpaw(fmt.Sprintf("%d", ai.SharesRejected), "rejected", shortpawed))
		if conf.Shares.Ratio {
			return fmt.Sprintf("%s (%.2f%%)", strings.Join(ret, ", "), float32(ai.SharesRejected)/float32(ai.SharesAccepted)*100)
		}
		return strings.Join(ret, ", ")
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
func getWorkerFromUser(username string) string {
	split := strings.Split(username, ".")
	if len(split) == 1 {
		return username
	}
	return split[1]
}
