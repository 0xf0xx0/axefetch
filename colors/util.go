package colors

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// processes a tagged string
func ProcessTags(s string) string {
	return formatLine(s)
}
func hex2RGB(hex string) (int, int, int) {
	hex = hex[1:]
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)
	return int(r), int(g), int(b)
}

func format(s, col string) string {
	if strings.HasPrefix(col, "#") {
		/// hex code
		return color.RGB(hex2RGB(col)).Sprint(s)
	}
	if strings.HasPrefix(col, "bg#") {
		return color.BgRGB(hex2RGB(col[2:])).Sprint(s)
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
	/// matches groups of tags, eg `{bold}{green}foobar`
	taggroupreg := regexp.MustCompile(`(\{[#\w\d]+\})+`)
	/// matches each individual tag
	tagreg := regexp.MustCompile(`(\{[#\w\d]+\})`)
	indices := taggroupreg.FindAllStringIndex(line, -1)
	indicesLen := len(indices)
	ret := make([]formatmatch, indicesLen)
	for i, element := range indices {
		endIdx := 0
		if i+1 < indicesLen {
			endIdx = indices[i+1][0]
		} else {
			endIdx = len(line)
		}
		/// this is either '{tag}' or '{tag}{tag}...'
		tagGroup := line[element[0]:element[1]]
		tagIndices := tagreg.FindAllStringIndex(tagGroup, -1)
		tags := make([]formattag, len(tagIndices))
		for ii, tag := range tagIndices {
			/// tag[] is relative to the outer match
			/// make it relative to the whole line
			tags[ii] = formattag{Start: element[0] + tag[0], End: element[0] + tag[1]}
		}
		ret[i] = formatmatch{tags: tags, targetEnd: endIdx}
	}
	return ret
}

// add a format tag to a string
func TagString(line, color string) string {
	if line == "" {
		return line
	}
	return fmt.Sprintf("{%s}%s", color, line)
}
func formatLine(line string) string {
	formats := selectFormats(line)
	slices.Reverse(formats)
	/// find+replace in reverse to avoid indexes jumping around
	for _, fmt := range formats {
		tags := fmt.tags
		/// save tags before flippin the array
		furstTag := tags[0]
		lastTag := tags[len(tags)-1]
		/// same with tags
		slices.Reverse(tags)
		/// save the target line and update it before substring-replacing
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
