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
	"bold":   color.New(color.Bold).SprintFunc(),
	"italic": color.New(color.Italic).SprintFunc(),
	"reset":  color.New(color.Reset).SprintFunc(),

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
// TODO:
func ProcessTags(s string) string {
	return formatLine(s)
}
func hex2RGB(hex string) []uint64 {
	hex = hex[1:]
	ret := make([]uint64, 3)
	ret[0], _ = strconv.ParseUint(hex[0:2], 16, 8)
	ret[1], _ = strconv.ParseUint(hex[2:4], 16, 8)
	ret[2], _ = strconv.ParseUint(hex[4:6], 16, 8)
	return ret
}

// TODO: tag stacking '{bold}{italic}'
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

type formattag struct {
	Start, End int
}
type formatmatch struct {
	tags      []formattag
	targetEnd int
}

// finds format tags by regex, including their target
func selectFormats(line string) []formatmatch {
	taggroupreg := regexp.MustCompile(`(\{[#\w\d]+\})+`)
	tagreg := regexp.MustCompile(`(\{[#\w\d]+\})`)
	indices := taggroupreg.FindAllStringIndex(line, -1)
	ret := make([]formatmatch, len(indices))
	for i, element := range indices {
		endIdx := 0
		if i+1 < len(indices) {
			endIdx = indices[i+1][0]
		} else {
			endIdx = len(line)
		}
		tagGroup := line[element[0]:element[1]]
		//println("g:", tagGroup)
		tagIndices := tagreg.FindAllStringIndex(tagGroup, -1)
		// fmt.Println(tagIndices)
		tags := make([]formattag, len(tagIndices))
		for ii, tag := range tagIndices {
			//println(tagGroup[tag[0]:tag[1]])
			tags[ii] = formattag{Start: element[0]+tag[0], End: element[0]+tag[1]}
		}
		ret[i] = formatmatch{tags: tags, targetEnd: endIdx}
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
		tags := fmt.tags
		furstTag := tags[0]
		lastTag := tags[len(tags)-1]
		slices.Reverse(tags)
		formatted := line[lastTag.End:fmt.targetEnd]
		for _, tag := range tags {
			color := line[tag.Start+1 : tag.End-1] /// {(color)}
			formatted = format(formatted, color)
		}
		line = line[:furstTag.Start] + formatted + line[fmt.targetEnd:]
	}
	return line
}

// strip format tags (NOT ansi) from line
func StripLine(line string) string {
	ret := ""
	for _, format := range selectFormats(line) {
		lastTag := format.tags[len(format.tags)-1]
		ret += line[lastTag.End:format.targetEnd]
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
