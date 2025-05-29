package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xf0xx0/axefetch/colors"
	"github.com/0xf0xx0/axefetch/icons"
	"github.com/0xf0xx0/axefetch/modules"
	"github.com/0xf0xx0/axefetch/paths"
	"github.com/0xf0xx0/axefetch/types"

	"github.com/go-andiamo/splitter"
	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v3"
)

var conf types.Config /// im not passing this stupid struct around
var testData = types.ApiInfo{
	AsicCount:              1,
	AsicModel:              "BM1370",
	BestDiff:               "210M",
	BestSessionDiff:        "21M",
	BoardFamily:            "Gamma",
	BoardVersion:           "601",
	StratumURL:             "not-so-public-pool.io",
	StratumPort:            3373,
	StratumUser:            "bc1qtesting.test-miner",
	FallbackStratumURL:     "closed-source-pool.evil",
	FallbackStratumPort:    666,
	FallbackStratumUser:    "bc1qfakefallbackaddress",
	IsUsingFallbackStratum: 0,
	Hostname:               "bitaxe",
	Version:                "v2.8.0",
	UptimeSeconds:          481824,
	SharesAccepted:         881,
	SharesRejected:         423,
	Hashrate:               1420,
	ExpectedHashrate:       1420,
	Power:                  20,
	FreeHeap:               8 * 1024 * 1024,
}

func main() {
	paths.MakeConfigDirTree()

	app := &cli.Command{
		Name:                   "axefetch",
		Version:                "0.0.1",
		Usage:                  "neofetch for *axes",
		UsageText:              "axefetch [options]",
		UseShortOptionHandling: true,
		EnableShellCompletion:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "conf",
				Usage: "config file `path`",
				Value: filepath.Join(paths.CONFIG_ROOT, "config.toml"),
			},
			&cli.StringFlag{
				Name:  "ip",
				Usage: "*axe ip `address`",
			},
			&cli.StringFlag{
				Name:  "icon",
				Usage: "ascii icon to use (name, path, or 'none')",
			},
			&cli.BoolFlag{
				Name:   "testing",
				Hidden: true,
			},
			&cli.StringFlag{
				Name:   "createdefaultconfig",
				Hidden: true,
			},
		},
		Action: func(_ context.Context, ctx *cli.Command) error {
			if path := ctx.String("createdefaultconfig"); path != "" {
				writeDefaultConfig(path)
				return nil
			}
			conf = loadConfig(ctx.String("conf"))
			ip := conf.General.IP
			if passedIP := ctx.String("ip"); passedIP != "" {
				ip = passedIP
			}

			/// start
			axeInfo := types.ApiInfo{}

			if !ctx.Bool("testing") {
				if ip == "" {
					return cli.Exit("no ip address given", 1)
				}
				infoReq, err := http.Get(fmt.Sprintf("http://%s/api/system/info", ip))
				if err != nil {
					return cli.Exit(fmt.Sprintf("error getting axe info: %s", err), 1)
				}
				body, err := io.ReadAll(infoReq.Body)
				if err != nil {
					return cli.Exit(fmt.Sprintf("error reading axe info: %s", err), 1)
				}
				if err := json.Unmarshal(body, &axeInfo); err != nil {
					return cli.Exit(fmt.Sprintf("error unmarshalling axe info: %s", err), 1)
				}
				/// this gets unmarshalled into the same struct to fill the rest of the data
				asicReq, err := http.Get(fmt.Sprintf("http://%s/api/system/asic", ip))
				if err != nil {
					return cli.Exit(fmt.Sprintf("error getting axe info: %s", err), 1)
				}
				body, err = io.ReadAll(asicReq.Body)
				if err != nil {
					return cli.Exit(fmt.Sprintf("error reading axe info: %s", err), 1)
				}
				if err := json.Unmarshal(body, &axeInfo); err != nil {
					return cli.Exit(fmt.Sprintf("error unmarshalling axe info: %s", err), 1)
				}
			} else {
				axeInfo = testData
			}
			info := processFormat(conf.Display.Format, axeInfo)

			/// select the icon
			var icon []string
			var iconname string
			if selectedicon := ctx.String("icon"); selectedicon != "" {
				iconname = selectedicon
				icon = icons.SearchAndLoadIcon(selectedicon)
			} else {
				switch conf.General.Icon {
				case "vendor":
					println("unimplemented, waiting for efuse")
					fallthrough
				case "family":
					fallthrough
				case "model":
					{
						iconname = testData.BoardVersion
						icon = icons.Models[iconname]
						break
					}
				case "asic":
					{
						iconname = testData.AsicModel
						icon = icons.Asics[iconname]
						break
					}
				case "none":
					{
						icon = []string{""}
					}
				default:
					{
						icon = icons.SearchAndLoadIcon(conf.General.Icon)
					}
				}
			}
			if icon == nil {
				return cli.Exit(fmt.Sprintf("couldnt load icon %q, does it exist?", iconname), 1)
			}
			/// print
			fmt.Println(strings.Join(stitchIconAndInfo(icon, info, conf.General.IconSpacing), "\n"))
			return nil
		},
	}
	if err := app.Run(context.TODO(), os.Args); err != nil {
		println(fmt.Sprint(err))
	}
}

func stitchIconAndInfo(icon, info []string, spacing int) []string {
	res := []string{}
	iconLen := len(icon)
	infoLen := len(info)
	if iconLen < infoLen {
		/// strip color tags to get print length
		repeat := strings.Repeat(" ", len(colors.StripLine(icon[iconLen-1])))
		/// pad the icon slice
		for diff := infoLen - iconLen; diff > 0; diff-- {
			icon = append(icon, repeat)
		}
	} else if iconLen > infoLen {
		/// lazily pad the info slice
		for diff := iconLen - infoLen; diff > 0; diff-- {
			info = append(info, "")
		}
	}
	spaces := strings.Repeat(" ", spacing)
	for i := range icon {
		res = append(res, fmt.Sprintf("%s%s%s%s", spaces, colors.ProcessTags(icon[i]), spaces, colors.ProcessTags(info[i])))
	}
	return res
}

// processes the display format string and returns a slice of the (valid) lines
func processFormat(format string, data types.ApiInfo) []string {
	/// all of this to handle some QUOTES
	/// I HATE ESCAPING
	split, _ := splitter.NewSplitter(' ', splitter.DoubleQuotesBackSlashEscaped)
	split.AddDefaultOptions(splitter.IgnoreEmpties, splitter.StripQuotes)
	res := []string{}
	lastline := "" /// store the last printed line and pass it in
	/// this is only used for the underline icl, prolly needs to be redone

	for line := range strings.Lines(format) {
		splitline, _ := split.Split(strings.TrimSpace(line))

		/// skip empty lines
		if len(splitline) == 0 {
			continue
		}
		args := splitline[1:]

		switch splitline[0] {
		case "info":
			{
				if v := info(args, colors.StripLine(lastline), data); v != "" {
					lastline = v
					res = append(res, v)
				}
				break
			}
		case "prin":
			{
				lastline = strings.Join(args, " ")
				res = append(res, lastline)
				break
			}
		default:
			{
				/// ignore
				continue
			}
		}
	}
	return res
}
func info(args []string, lastline string, data types.ApiInfo) string {
	ret := ""
	/// two formats: 'info <func>' and 'info <subtitle> <func>'
	switch len(args) {
	/// <func>
	case 1:
		{
			/// coloring for these is handled in each func
			ret = modules.Modules[args[0]](conf, data, []string{lastline})
			break
		}
	/// <subtitle> <func>
	case 2:
		{
			fn := modules.Modules[args[1]]
			if fn == nil {
				return ret
			}
			ret = fn(conf, data, []string{})
			if ret == "" {
				return ret
			}
			subtitle := colors.TagString(args[0], conf.Colors.Subtitle)
			if conf.Display.BoldTitles {
				subtitle = colors.TagString(subtitle, "bold")
			}
			ret = fmt.Sprintf("%s%s %s", subtitle,
				colors.TagString(conf.General.Separator, conf.Colors.Separator),
				colors.TagString(ret, conf.Colors.Info))
			break
		}
	}
	return ret
}

func loadConfig(path string) types.Config {
	var conf types.Config
	// configfile, err := os.ReadFile(path)
	configfile, err := os.Open(path)
	if err != nil {
		println(fmt.Sprintf("failed to load config at %s: %s", path, err))
		os.Exit(1)
	}
	d := toml.NewDecoder(configfile)
	d.DisallowUnknownFields()
	if err := d.Decode(&conf); err != nil {
		println(fmt.Sprintf("failed to decode config at %s: %s", path, err))
		os.Exit(1)
	}
	return conf
}
func writeDefaultConfig(path string) error {
	defaultDisplayFormat := strings.Join([]string{
		`this is an invalid line, so its not printed :3`,
		`info title`,
		`info underline`,
		`info "Model" model`,
		`info "ASIC(s)" asicmodel`,
		`info "Firmware" firmware`,
		`info "Uptime" uptime`,
		`info "Best Difficulty" bestdiff`,
		`info "Shares" shares`,
		`info "Pool" pool`,
		`info "Hashrate" hashrate`,
		`info "Efficiency" efficiency`,
		`info "Heap" heap`,
		``,
		`prin circlejerking into open source`,
	}, "\n")

	defaultConf := types.Config{
		General: types.General{
			IP:          "replace me",
			Separator:   ":",
			Underline:   "-",
			Icon:        "model",
			IconSpacing: 3,
		},
		Display: types.Display{
			Format:     defaultDisplayFormat,
			Colors:     "family",
			BoldTitles: true,
		},
		Colors: types.Colors{
			Title:     "green",
			At:        "white",
			Underline: "blackbright",
			Subtitle:  "green",
			Separator: "blackbright",
			Info:      "white",
		},
		Title: types.Title{
			Workername: true,
			Hostname:   true,
		},
		Model: types.Model{
			Boardversion: true,
			Family:       false,
			Vendor:       false,
		},
		Asicmodel: types.Asicmodel{
			Asiccount:      true,
			Smallcorecount: true,
		},
		Efficiency: types.Efficiency{
			Expected: true,
			Actual:   true,
			Shortpaw: "off",
		},
		Firmware: types.Firmware{
			Version: true,
		},
		Hashrate: types.Hashrate{
			Expected: true,
			Actual:   true,
			Shortpaw: "off",
		},
		Pool: types.Pool{
			Port: true,
		},
		Shares: types.Shares{
			Ratio:    true,
			Shortpaw: "off",
		},
		Uptime: types.Uptime{
			Format: "%dd %hh %mm %ss",
		},
	}
	conf, _ := toml.Marshal(defaultConf)
	if err := os.WriteFile(path, conf, 0755); err != nil {
		return cli.Exit(fmt.Sprintf("couldnt create config file: %s", err), 1)
	}
	return nil
}
