package colors

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// TODO
// vendor color themes
// model color themes
// model family color themes
var colorMap = map[string]func(a ...interface{}) string{
	"bold": color.New(color.Bold).SprintFunc(),
	"italic": color.New(color.Italic).SprintFunc(),
	"reset": color.New(color.Reset).SprintFunc(),

	"red":     color.New(color.FgRed).SprintFunc(),
	"green":   color.New(color.FgGreen).SprintFunc(),
	"blue":    color.New(color.FgBlue).SprintFunc(),
	"yellow":  color.New(color.FgYellow).SprintFunc(),
	"cyan":    color.New(color.FgCyan).SprintFunc(),
	"magenta": color.New(color.FgMagenta).SprintFunc(),
	"white":   color.New(color.FgWhite).SprintFunc(),
	"black":   color.New(color.FgBlack).SprintFunc(),

	"redbright":     color.New(color.FgHiRed).SprintFunc(),
	"greenbright":   color.New(color.FgHiGreen).SprintFunc(),
	"bluebright":    color.New(color.FgHiBlue).SprintFunc(),
	"yellowbright":  color.New(color.FgHiYellow).SprintFunc(),
	"cyanbright":    color.New(color.FgHiCyan).SprintFunc(),
	"magentabright": color.New(color.FgHiMagenta).SprintFunc(),
	"whitebright":   color.New(color.FgHiWhite).SprintFunc(),
	"blackbright":   color.New(color.FgHiBlack).SprintFunc(),
}
// processes a tagged string
func ProcessTags(s string) string {
	return formatLine(s)
}
func hex2RGB(hex string) []uint64 {
	hex = hex[1:]
	ret := make([]uint64, 3)
	ret[0],_ = strconv.ParseUint(hex[0:2], 16, 8)
	ret[1],_ = strconv.ParseUint(hex[2:4], 16, 8)
	ret[2],_ = strconv.ParseUint(hex[4:6], 16, 8)
	return ret
}
func format(s, col string) string {
	if strings.HasPrefix(col, "#") {
		/// hex code
		rgb := hex2RGB(col)
		return color.RGB(int(rgb[0]), int(rgb[1]), int(rgb[2])).Sprint(s)
	}
	if fn, ok := colorMap[col]; ok {
		return fn(s)
	}
	return s
}

type formatmatch struct {
	tagStart, tagEnd, targetEnd int
}
// finds format tags by regex, including their target
func selectFormats(line string) []formatmatch {
	reg := regexp.MustCompile(`\{.+?\}`)
	indexes := reg.FindAllStringIndex(line, -1)
	ret := make([]formatmatch, len(indexes))
	for i, element := range indexes {
		endIdx := 0
		if i+1 < len(indexes) {
			endIdx = indexes[i+1][0]
		} else {
			endIdx = len(line)
		}
		ret[i] = formatmatch{tagStart: element[0], tagEnd: element[1], targetEnd: endIdx}
	}
	return ret
}
// add a format tag to a string
func TagString(line, color string) string {
	return fmt.Sprintf("{%s}%s", color, line)
}
func formatLine(line string) string {
	formats := selectFormats(line)
	slices.Reverse(formats)
	/// find+replace in reverse to avoid indexes jumping around
	for _, fmt := range formats {
		color := line[fmt.tagStart+1 : fmt.tagEnd-1] /// {(color)}
		line = line[:fmt.tagStart] + format(line[fmt.tagEnd:fmt.targetEnd], color) + line[fmt.targetEnd:]
	}
	return line
}
// strip format tags (NOT ansi) from line
func StripLine(line string) string {
	ret := ""
	for _, format := range selectFormats(line) {
		ret += line[format.tagEnd:format.targetEnd]
	}
	return ret
}
// formats a string slice
func FormatIcon(icon []string) []string {
	for i, line := range icon {
		icon[i] = formatLine(line)
	}
	return icon
}
