package box

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
)

// expandedLine stores a tab-expanded line, and its visible length.
type expandedLine struct {
	line string // tab-expanded line
	len  int    // line's visible length
}

// addVertPadding adds Vertical Padding
func (b *Box) addVertPadding(len int) []string {
	padding := strings.Repeat(" ", len-2)
	vertical := applyColor(b.Vertical, b.color)

	texts := make([]string, b.py)
	for i := range texts {
		texts[i] = fmt.Sprintf("%s%s%s", vertical, padding, vertical)
	}

	return texts
}

// longestLine expands tabs in lines and determines longest visible
// return longest length and array of expanded lines
func longestLine(lines []string) (int, []expandedLine) {
	longest := 0
	var expandedLines []expandedLine
	var tmpLine strings.Builder
	var lineLen int

	for _, line := range lines {
		tmpLine.Reset()
		for _, c := range line {
			lineLen = runewidth.StringWidth(tmpLine.String())

			if c == '\t' {
				tmpLine.WriteString(strings.Repeat(" ", 8-(lineLen&7)))
			} else {
				tmpLine.WriteRune(c)
			}
		}
		lineLen = runewidth.StringWidth(tmpLine.String())
		expandedLines = append(expandedLines, expandedLine{tmpLine.String(), lineLen})

		// Check if each line has ANSI Color Code then decrease the length accordingly
		if runewidth.StringWidth(ansi.Strip(tmpLine.String())) < runewidth.StringWidth(tmpLine.String()) {
			lineLen = runewidth.StringWidth(ansi.Strip(tmpLine.String()))
		}

		if lineLen > longest {
			longest = lineLen
		}
	}
	return longest, expandedLines
}

// formatLine formats the line according to the information passed
func (b *Box) formatLine(lines2 []expandedLine, longestLine, titleLen int, sideMargin, title string, texts []string) []string {
	for i, line := range lines2 {
		length := line.len

		// Use later
		var space, oddSpace string

		// Check if line.line has ANSI Color Code then decrease length accordingly
		if runewidth.StringWidth(ansi.Strip(line.line)) < runewidth.StringWidth(line.line) {
			length = runewidth.StringWidth(ansi.Strip(line.line))
		}

		// If current text is shorter than the longest one
		// center the text, so it looks better
		if length < longestLine {
			// Difference between longest and current one
			diff := longestLine - length

			// the spaces to add on each side
			toAdd := diff / 2
			space = strings.Repeat(" ", toAdd)

			// If difference between the longest and current one
			// is odd, we have to add one additional space before the last vertical separator
			if diff%2 != 0 {
				oddSpace = " "
			}
		}

		spacing := strings.Join([]string{space, sideMargin}, "")
		var format AlignType

		switch {
		case i < titleLen && title != "" && b.titlePos == Inside:
			format = centerAlign
		default:
			format = AlignType(b.findAlign())
		}

		sep := applyColor(b.Vertical, b.color)

		formatted := fmt.Sprintf(string(format), sep, spacing, line.line, oddSpace, space, sideMargin)
		texts = append(texts, formatted)
	}
	return texts
}

func (b *Box) findAlign() string {
	switch b.contentAlign {
	case Center:
		return centerAlign
	case Right:
		return rightAlign
	case Left, "":
		// If ContentAlign isn't provided then by default Alignment is Left
		return leftAlign
	default:
		return leftAlign
	}
}

func repeatWithString(c string, n int, str string) string {
	cstr := ansi.Strip(str)
	count := n - runewidth.StringWidth(cstr) - 2
	bar := strings.Repeat(c, count)
	return fmt.Sprintf(" %s %s", str, bar)
}

func applyColor(str string, colorStr string) string {
	colorValue := parseColorString(colorStr)
	convertedColor := profile.Convert(colorValue)
	return applyConvertedColor(str, convertedColor)
}

func stringColorToHex(colorName string) string {
	if hex, exists := colorNameToHex[colorName]; exists {
		return hex
	}
	// Return empty string for unknown colors to let ansi.XParseColor handle it
	return ""
}

// parseColorString converts a color string to color.Color using stringColorToHex and ansi.XParseColor
func parseColorString(colorStr string) color.Color {
	hexColor := stringColorToHex(colorStr)

	if hexColor == "" {
		hexColor = colorStr
	}

	colorValue := ansi.XParseColor(hexColor)
	if colorValue == nil {
		// Fallback to white if parsing fails
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}
	return colorValue
}

func applyConvertedColor(str string, c color.Color) string {
	if c == nil {
		return str
	}

	var style ansi.Style
	style = style.ForegroundColor(c)
	return style.Styled(str)
}

func (b *Box) applyColorBar(topBar, bottomBar, title string) (string, string) {
	if b.titleColor != "" {
		if b.titlePos == Top {
			bar_ := strings.Split(ansi.Strip(topBar), ansi.Strip(title))

			colorValue := parseColorString(b.color)
			convertedColor := profile.Convert(colorValue)

			b0 := applyConvertedColor(bar_[0], convertedColor)
			b1 := applyConvertedColor(bar_[1], convertedColor)

			topBar = b0 + applyColor(title, b.titleColor) + b1
		}

		if b.titlePos == Bottom {
			bar_ := strings.Split(ansi.Strip(bottomBar), ansi.Strip(title))

			colorValue := parseColorString(b.color)
			convertedColor := profile.Convert(colorValue)

			b0 := applyConvertedColor(bar_[0], convertedColor)
			b1 := applyConvertedColor(bar_[1], convertedColor)

			bottomBar = b0 + applyColor(title, b.titleColor) + b1
		}
	}

	return topBar, bottomBar
}
