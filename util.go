package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"git.0xf0xx0.eth.limo/0xf0xx0/axefetch/types"

	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v3"
)

func loadConfig(path string, conf *types.Config) {
	configfile, err := os.Open(path)
	if err != nil {
		println(fmt.Sprintf("failed to load config at %s: %s", path, err))
		return
	}
	d := toml.NewDecoder(configfile)
	d.DisallowUnknownFields()
	if err := d.Decode(conf); err != nil {
		println(fmt.Sprintf("failed to decode config at %s: %s", path, err))
		os.Exit(1)
	}
}
func writeDefaultConfig(path string) error {
	conf, _ := toml.Marshal(types.DefaultConf)
	if err := os.WriteFile(resolvePath(path), conf, 0755); err != nil {
		return cli.Exit(fmt.Sprintf("couldnt create config file: %s", err), 1)
	}
	return nil
}

// resolves ~ and cleans path
// https://stackoverflow.com/a/17617721
func resolvePath(path string) string {
	if strings.HasPrefix(path, "~") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, path[1:])
	}
	return path
}
func getConfigDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		println(fmt.Sprintf("error getting config dir: %s", err))
		os.Exit(1)
	}
	return filepath.Join(userConfigDir, "./axefetch")
}
func copyConf(dest, src *types.Config) {
	dest.Bestdiff = src.Bestdiff
	dest.Asic = src.Asic
	dest.ColorTheme = src.ColorTheme
	dest.Display = src.Display
	dest.Efficiency = src.Efficiency
	// dest.Firmware = src.Firmware
	dest.General = src.General
	dest.Hashrate = src.Hashrate
	dest.Model = src.Model
	dest.Pool = src.Pool
	dest.Shares = src.Shares
	dest.Temp = src.Temp
	dest.Title = src.Title
	dest.Uptime = src.Uptime
}

// just http.Get but errors on non 200 response and wraps up the root path
func reqAPI(endpoint, ip string, unmarshalInto any) error {
	req, err := http.Get(fmt.Sprintf("http://%s/api/%s", ip, endpoint))
	if err != nil {
		return err
	}
	if req.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", req.Status)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, unmarshalInto); err != nil {
		return err
	}
	return nil
}

// converts 123456 into 123K and 123456789876 into 123 G
func floatToBinshort(value float64) string {
	if value >= 1e12 {
		return fmt.Sprintf("%.3gT", value/1e9) /// Trillions
	} else if value >= 1e9 {
		return fmt.Sprintf("%.3gG", value/1e9) /// Billions
	} else if value >= 1e6 {
		return fmt.Sprintf("%.3gM", value/1e6) /// Millions
	} else if value >= 1e3 {
		return fmt.Sprintf("%.3gK", value/1e3) /// Thousands
	}
	return fmt.Sprintf("%.3g", value) /// Less than a thousand
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
			unit = "MHz"
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
		if len(split[0]) > 16 {
			return fmt.Sprintf("%s...%s", username[:4], username[len(username)-4:])
		}
		return split[0]
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

// prints a shortened data line
// TODO: rename shouldBeShort
func printWithShortpaw(str, longpaw string, shouldBeShort bool) string {
	if shouldBeShort {
		return str
	}
	return fmt.Sprintf("%s %s", str, longpaw)
}
