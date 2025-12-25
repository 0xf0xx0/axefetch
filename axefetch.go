package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.0xf0xx0.eth.limo/0xf0xx0/axefetch/style"
	"git.0xf0xx0.eth.limo/0xf0xx0/axefetch/types"

	"git.0xf0xx0.eth.limo/0xf0xx0/oigiki"
	"github.com/urfave/cli/v3"
)

var conf types.Config /// im not passing this stupid struct around
var testData = types.ApiInfo{
	AsicCount:              1,
	AsicModel:              "BM1370",
	BestDiff:               210_000_000,
	BestSessionDiff:        330_730_700,
	BoardFamily:            "Gamma",
	BoardVersion:           "621",
	BoardVendor:            "Fluffy Inc.",
	StratumURL:             "pogolo.local",
	StratumPort:            5621,
	StratumUser:            "bc1qfakeaddress.bitaxuh",
	FallbackStratumURL:     "closed-source-pool.evil",
	FallbackStratumPort:    666,
	Frequency:              42069.69,
	CoreVoltage:            42069,
	FallbackStratumUser:    "bc1qfakefallbackaddress",
	IsUsingFallbackStratum: 0,
	Hostname:               "bitaxe",
	Version:                "v4.2.0",
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

const commandHelpTemplate = `Name:
   {bold}{blue}{{.Name}} - {{.Usage}}{/}

Usage:
   {green}{{.Name}} {blue}[options]{/}

Options:{blue}
   {{range .VisibleFlags}}{{.String}}
   {{end}}{/}
Version:
   {green}v{{.Version}}
`

func main() {
	cli.RootCommandHelpTemplate = oigiki.ProcessTags(commandHelpTemplate)
	app := &cli.Command{
		Name:                   "axefetch",
		Version:                "0.0.1",
		Usage:                  "neofetch for *axes",
		UsageText:              "axefetch [options]",
		UseShortOptionHandling: true,
		MutuallyExclusiveFlags: []cli.MutuallyExclusiveFlags{
			{
				Flags: [][]cli.Flag{
					{
						&cli.BoolFlag{
							Name:  "color",
							Usage: "force color output",
						},
					},
					{
						&cli.BoolFlag{
							Name:    "nocolor",
							Aliases: []string{"stdout"},
							Usage:   "disable color output",
						},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "writedefaultconf",
				Usage: "write default config to `path`",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c", "conf"},
				Usage:   "config file `path` (or 'none')",
				Value:   filepath.Join(getConfigDir(), "config.toml"),
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
				Name:   "testing",
				Hidden: true,
			},
		},
		Action: func(_ context.Context, ctx *cli.Command) error {
			if path := ctx.String("writedefaultconfig"); path != "" {
				writeDefaultConfig(path)
				return nil
			}
			if ctx.Bool("color") {
				oigiki.NoColor = false
			}
			if ctx.Bool("nocolor") {
				oigiki.NoColor = true
			}
			/// set defaults
			copyConf(&conf, &types.DefaultConf)

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
				err := reqAPI("system/info", conf.General.IP, &axeInfo)
				if err != nil {
					return cli.Exit(fmt.Sprintf("error getting axe status: %s", err), 1)
				}

				/// this gets unmarshalled into the same struct to fill the rest of the board info
				err = reqAPI("system/asic", conf.General.IP, &axeInfo)
				if err != nil {
					return cli.Exit(fmt.Sprintf("error getting axe info: %s", err), 1)
				}
			} else {
				axeInfo = testData
			}

			/// select the icon
			var icon []string
			switch conf.Display.Icon {
			case "vendor":
				println("vendor unimplemented, waiting for efuse")
				fallthrough
			case "family":
				{
					conf.Display.Icon = strings.ToLower(axeInfo.BoardFamily)
				}
			case "chip":
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
				potentialIcon, ok := style.Icons[conf.Display.Icon]
				if ok {
					icon = filterEmptyStringsOut(strings.Split(potentialIcon, "\n"))
				} else {
					icon = []string{""} /// just print no icon
					println(fmt.Sprintf("unknown icon %q", conf.Display.Icon))
					conf.Display.IconSpacing = 0
				}
			}
			if conf.Display.Theme != "manual" {
				if theme, ok := style.Themes[strings.ToLower(conf.Display.Theme)]; ok {
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
func printIconAndInfo(icon, info []string, spacing int) {
	iconLen := len(icon)
	infoLen := len(info)
	if iconLen < infoLen {
		/// strip color tags to get print length
		repeat := strings.Repeat(" ", len([]rune(oigiki.StripTags(icon[iconLen-1]))))
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
			oigiki.ProcessTags(oigiki.TagString(icon[i], conf.ColorTheme.Icon)),
			trench, oigiki.ProcessTags(info[i]))
	}
}

// splits string, preserving quotes
func splitFormatLine(line string) []string {
	out := make([]string, 0, 4)
	b := strings.Builder{}
	quoteActive := false
	escaped := false
	for _, c := range line {
		if c == '\\' {
			escaped = true
		} else if c == '"' && !escaped {
			quoteActive = !quoteActive
		} else if !quoteActive && c == ' ' {
			out = append(out, b.String())
			b.Reset()
		} else {
			b.WriteRune(c)
			if escaped {
				escaped = false
			}
		}
	}
	if b.Len() > 0 {
		out = append(out, b.String())
	}
	return out
}

// processes the display format string and returns a slice of the (valid) lines
func processFormat(format string, data types.ApiInfo) []string {
	res := []string{}
	lastline := "" /// store the last printed line and pass it in
	/// this is only used for the underline icl, prolly needs to be redone

	for line := range strings.Lines(format) {
		splitline := splitFormatLine(strings.TrimSpace(line))

		/// skip empty lines
		if len(splitline) == 0 {
			continue
		}
		args := splitline[1:]

		switch splitline[0] {
		case "info":
			{
				if v := info(args, oigiki.StripTags(lastline), data); v != "" {
					lastline = v
					res = append(res, v)
				}
				break
			}
		case "prin":
			{
				if v := prin(args, oigiki.StripTags(lastline), data); v != "" {
					lastline = v
					res = append(res, v)
				}
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
			ret = Modules[args[0]](conf, data, []string{lastline})
			break
		}
	/// <subtitle> <func>
	case 2:
		{
			fn := Modules[args[1]]
			if fn == nil {
				return ret
			}
			ret = fn(conf, data, []string{})
			if ret == "" {
				return ret
			}
			subtitle := oigiki.TagString(args[0], conf.ColorTheme.Subtitle)
			if conf.Display.BoldTitles {
				subtitle = oigiki.TagString(subtitle, "bold")
			}
			ret = fmt.Sprintf("%s{/bold}%s %s", subtitle,
				oigiki.TagString(conf.Display.Separator, conf.ColorTheme.Separator),
				oigiki.TagString(ret, conf.ColorTheme.Info))
			break
		}
	}
	return ret
}
func prin(args []string, lastline string, data types.ApiInfo) string {
	return oigiki.TagString(strings.Join(args, " "), conf.ColorTheme.Info)
}
