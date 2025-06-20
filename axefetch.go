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
	"github.com/fatih/color"
	"github.com/tiendc/go-deepcopy"

	"github.com/go-andiamo/splitter"
	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v3"
)

var conf types.Config /// im not passing this stupid struct around
var testData = types.ApiInfo{
	AsicCount:              1,
	AsicModel:              "BM1370",
	BestDiff:               210_210_000_000,
	BestSessionDiff:        330_730_000,
	BoardFamily:            "Gamma",
	BoardVersion:           "601",
	BoardVendor:            "Fluffy Inc.",
	StratumURL:             "pooblic-pool.io",
	StratumPort:            3333,
	StratumUser:            "bc1qfakeaddress.bitaxuh",
	FallbackStratumURL:     "closed-source-pool.evil",
	FallbackStratumPort:    666,
	Frequency:              42069,
	CoreVoltage:            42069,
	FallbackStratumUser:    "bc1qfakefallbackaddress",
	IsUsingFallbackStratum: false,
	Hostname:               "bitaxe",
	Version:                "v2.8.0",
	UptimeSeconds:          481824,
	SharesAccepted:         881_435_387_204,
	SharesRejected:         423_482_465,
	Hashrate:               1420,
	ExpectedHashrate:       1420,
	Power:                  20,
	FreeHeap:               8 * 1024 * 1024,
	Temp:                   55.329478,
	VrTemp:                 66,
}

func main() {
	if paths.MakeConfigDirTree(types.DefaultConf) {
		writeDefaultConfig(filepath.Join(paths.CONFIG_ROOT, "config.toml"))
	}

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
				Usage: "config file `path` (or 'none')",
				Value: filepath.Join(paths.CONFIG_ROOT, "config.toml"),
			},
			&cli.StringFlag{
				Name:  "ip",
				Usage: "*axe ip `address`",
			},
			&cli.StringFlag{
				Name:  "icon",
				Usage: "ascii icon to use (`name`, path, or 'none')",
			},
			&cli.StringFlag{
				Name:  "theme",
				Usage: "color `theme` (name, 'manual')",
			},
			&cli.BoolFlag{
				Name:  "force-color",
				Usage: "force color output",
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
			if ctx.Bool("force-color") {
				color.NoColor = false
			}
			/// set defaults
			deepcopy.Copy(&conf, &types.DefaultConf)

			if passedConfig := ctx.String("conf"); passedConfig != "" && passedConfig != "none" {
				loadConfig(passedConfig, &conf)
			}
			/// config overrides
			if passedIP := ctx.String("ip"); passedIP != "" {
				conf.General.IP = passedIP
			}
			if passedIcon := ctx.String("icon"); passedIcon != "" {
				conf.Display.Icon = strings.ToLower(passedIcon)
			}
			if passedTheme := ctx.String("theme"); passedTheme != "" {
				conf.Display.Theme = strings.ToLower(passedTheme)
			}

			/// start
			axeInfo := types.ApiInfo{}

			if !ctx.Bool("testing") {
				if conf.General.IP == "" {
					return cli.Exit("no ip address given", 1)
				}
				statusReq, err := http.Get(fmt.Sprintf("http://%s/api/v2/system/status", conf.General.IP))
				if err != nil {
					return cli.Exit(fmt.Sprintf("error getting axe info: %s", err), 1)
				}
				body, err := io.ReadAll(statusReq.Body)
				if err != nil {
					return cli.Exit(fmt.Sprintf("error reading axe info: %s", err), 1)
				}
				if err := json.Unmarshal(body, &axeInfo); err != nil {
					return cli.Exit(fmt.Sprintf("error unmarshalling axe info: %s", err), 1)
				}
				/// this gets unmarshalled into the same struct to fill the rest of the asic info
				asicReq, err := http.Get(fmt.Sprintf("http://%s/api/v2/system/board", conf.General.IP))
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
				confReq, err := http.Get(fmt.Sprintf("http://%s/api/v2/system/config", conf.General.IP))
				if err != nil {
					return cli.Exit(fmt.Sprintf("error getting axe info: %s", err), 1)
				}
				body, err = io.ReadAll(confReq.Body)
				if err != nil {
					return cli.Exit(fmt.Sprintf("error reading axe info: %s", err), 1)
				}
				if err := json.Unmarshal(body, &axeInfo); err != nil {
					return cli.Exit(fmt.Sprintf("error unmarshalling axe info: %s", err), 1)
				}
			} else {
				axeInfo = testData
			}

			/// select the icon
			var icon []string
			switch conf.Display.Icon {
			case "vendor":
				println("unimplemented, waiting for efuse")
				fallthrough
			case "family":
				{
					conf.Display.Icon = strings.ToLower(axeInfo.BoardFamily)
				}
			case "asic":
				{
					conf.Display.Icon = axeInfo.AsicModel
				}
			case "none":
				{
					icon = []string{""}
					conf.Display.IconSpacing = 0
				}
			}

			switch conf.Display.Theme {
			case "manual":
				{
					break
				}
			case "vendor":
				fallthrough
			case "family":
				{
					conf.Display.Theme = axeInfo.BoardFamily
				}
				/// no default case, we assume its a theme name and do a lookup
			}

			if conf.Display.Icon != "none" {
				icon = icons.SearchAndLoadIcon(conf.Display.Icon)
				if icon == nil {
					icon = []string{""} /// just print no icon
					println(fmt.Sprintf("couldnt load icon %q, does it exist?", conf.Display.Icon))
				}
			}
			if conf.Display.Theme != "manual" {
				if theme, ok := colors.Themes[strings.ToLower(conf.Display.Theme)]; ok {
					conf.ColorTheme = theme
				} else {
					println(fmt.Sprintf("unknown theme %q", conf.Display.Theme))
				}
			}
			/// print
			info := processFormat(conf.Display.Format, axeInfo)
			printIconAndInfo(icon, info, conf.Display.IconSpacing)
			return nil
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		println(fmt.Sprint(err))
	}
}

// merges icon and info slices and processes tags
func printIconAndInfo(icon, info []string, spacing int) []string {
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
	for i := range icon {
		trench := strings.Repeat(" ", spacing)
		fmt.Printf("%s%s%s\n",
			colors.ProcessTags(colors.TagString(icon[i], conf.ColorTheme.Icon)),
			trench, colors.ProcessTags(info[i]))
	}
	return icon
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
				lastline = colors.TagString(strings.Join(args, " "), conf.ColorTheme.Info)
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
	/// left loose on purpose
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
			subtitle := colors.TagString(args[0], conf.ColorTheme.Subtitle)
			if conf.Display.BoldTitles {
				subtitle = colors.TagString(subtitle, "bold")
			}
			ret = fmt.Sprintf("%s%s %s", subtitle,
				colors.TagString(conf.Display.Separator, conf.ColorTheme.Separator),
				colors.TagString(ret, conf.ColorTheme.Info))
			break
		}
	}
	return ret
}

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
	if err := os.WriteFile(path, conf, 0755); err != nil {
		return cli.Exit(fmt.Sprintf("couldnt create config file: %s", err), 1)
	}
	return nil
}
