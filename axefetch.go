package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	BoardVersion:           "601",
	StratumURL:             "not-so-public-pool.io",
	StratumPort:            3373,
	StratumUser:            "bc1qtesting.test-miner",
	FallbackStratumURL:     "closed-source-pool.evil",
	FallbackStratumPort:    666,
	FallbackStratumUser:    "bc1qfallback",
	IsUsingFallbackStratum: 0,
	Hostname:               "bitaxe",
	Version:                "v2.8.0",
	UptimeSeconds:          481824,
	SharesAccepted:         881,
	SharesRejected:         423,
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
				Usage: "config file path",
				Value: filepath.Join(paths.CONFIG_ROOT, "config.toml"),
			},
			&cli.StringFlag{
				Name:  "ip",
				Usage: "*axe ip address",
			},
			&cli.StringFlag{
				Name:  "icon",
				Usage: "ascii icon to use (name, path, or 'none')",
			},
			&cli.BoolFlag{
				Name: "testing",
			},
		},
		Action: func(_ context.Context, ctx *cli.Command) error {
			conf = loadConfig(ctx.String("conf"))
			ip := ctx.String("ip")
			selectedicon := ctx.String("icon")
			/// start

			axeInfo := types.ApiInfo{}

			if !ctx.Bool("testing") {
				if ip == "" {
					println("no ip address given")
					os.Exit(1)
				}
				infoReq, err := http.Get(fmt.Sprintf("http://%s/api/system/info", ip))
				if err != nil {
					println(fmt.Sprintf("error getting axe info: %s", err))
					os.Exit(1)
				}
				body, err := io.ReadAll(infoReq.Body)
				if err != nil {
					println(fmt.Sprintf("error reading axe info: %s", err))
					os.Exit(1)
				}
				if err := json.Unmarshal(body, &axeInfo); err != nil {
					println(fmt.Sprintf("error unmarshalling axe info: %s", err))
					os.Exit(1)
				}
				/// this gets unmarshalled into the same struct to fill the rest of the data
				asicReq, err := http.Get(fmt.Sprintf("http://%s/api/system/asic", ip))
				if err != nil {
					println(fmt.Sprintf("error getting axe info: %s", err))
					os.Exit(1)
				}
				body, err = io.ReadAll(asicReq.Body)
				if err != nil {
					println(fmt.Sprintf("error reading axe info: %s", err))
					os.Exit(1)
				}
				if err := json.Unmarshal(body, &axeInfo); err != nil {
					println(fmt.Sprintf("error unmarshalling axe info: %s", err))
					os.Exit(1)
				}

			} else {
				axeInfo = testData
			}
			info := processFormat(conf.Display.Format, axeInfo)

			/// select the icon
			var icon []string
			var iconname string
			if selectedicon != "" {
				iconname = selectedicon
				icon = icons.SearchAndLoadIcon(selectedicon)
			} else {
				switch conf.General.IconType {
				case "vendor":
					println("unimplemented, waiting for efuse")
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
				default:
					{
						icon = icons.SearchAndLoadIcon(conf.General.IconType)
					}
				}
			}
			if icon == nil {
				println(fmt.Sprintf("couldnt load icon %q, does it exist?", iconname))
				os.Exit(1)
			}
			/// print
			fmt.Println(strings.Join(stitchIconAndInfo(icon, info, conf.General.IconSpacing), "\n"))
			return nil
		},
	}
	if err := app.Run(context.TODO(), os.Args); err != nil {
		println(err)
		os.Exit(1)
	}
}

func stitchIconAndInfo(icon, info []string, spacing int) []string {
	res := []string{}

	if len(icon) < len(info) {
		/// strip color tags to get print length
		padLen := len(colors.StripLine(icon[len(icon)-1]))
		/// pad the icon slice
		for diff := len(info) - len(icon); diff > 0; diff-- {
			icon = append(icon, strings.Repeat(" ", padLen))
		}
	} else if len(icon) > len(info) {
		/// lazily pad the info slice
		for diff := len(icon) - len(info); diff > 0; diff-- {
			info = append(info, "")
		}
	}
	for i := range icon {
		res = append(res, fmt.Sprintf("%s%s%s", colors.ProcessTags(icon[i]), strings.Repeat(" ", spacing), colors.ProcessTags(info[i])))
	}
	return res
}

// processes the format string and returns a slice of the (valid) lines
func processFormat(format string, data types.ApiInfo) []string {
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
