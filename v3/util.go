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
	vertical := applyColor(b.vertical, b.color)

	texts := make([]string, b.py)
	for i := range texts {
		texts[i] = vertical + padding + vertical
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
func (b *Box) formatLine(lines2 []expandedLine, longestLine, titleLen int, sideMargin, title string, texts []string) ([]string, error) {
	for i, line := range lines2 {
		length := line.len

		// Use later
		var space, oddSpace string

		// compute stripped width once
		strippedWidth := runewidth.StringWidth(ansi.Strip(line.line))
		if strippedWidth < runewidth.StringWidth(line.line) {
			length = strippedWidth
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

		spacing := space + sideMargin
		var format AlignType

		switch {
		case i < titleLen && title != "" && b.titlePos == Inside:
			format = centerAlign
		default:
			align, err := b.findAlign()
			if err != nil {
				return nil, err
			}
			format = AlignType(align)
		}

		sep := applyColor(b.vertical, b.color)

		formatted := fmt.Sprintf(string(format), sep, spacing, line.line, oddSpace, space, sideMargin)
		texts = append(texts, formatted)
	}
	return texts, nil
}

func (b *Box) findAlign() (string, error) {
	switch b.contentAlign {
	case Center:
		return centerAlign, nil
	case Right:
		return rightAlign, nil
	case Left, "":
		// If ContentAlign isn't provided then by default Alignment is Left
		return leftAlign, nil
	default:
		return "", fmt.Errorf("invalid Content Alignment %s", b.contentAlign)
	}
}

func repeatWithString(c string, n int, str string) string {
	cstr := ansi.Strip(str)
	count := n - runewidth.StringWidth(cstr) - 2
	if count < 0 {
		count = 0
	}
	bar := strings.Repeat(c, count)
	return " " + str + " " + bar
}

func getConvertedColor(colorStr string) color.Color {
	cv := parseColorString(colorStr)
	// If profile conversion results in nil, fall back to the
	// parsed color so we always emit color.
	converted := profile.Convert(cv)
	if converted == nil {
		return cv
	}
	return converted
}

func applyColor(str string, colorStr string) string {
	// Empty color string means: do not apply any styling.
	if colorStr == "" {
		return str
	}
	convertedColor := getConvertedColor(colorStr)
	return applyConvertedColor(str, convertedColor)
}

func stringColorToHex(colorName string) string {
	if hex, exists := colorNameToHex[colorName]; exists {
		return hex
	}
	// Return empty string for unknown colors to let ansi.XParseColor handle it
	return ""
}

// addStylePreservingOriginalFormat allows to add style around the orginal formating
func addStylePreservingOriginalFormat(s string, f func(a string) string) string {
	const reset = "\033[0m"
	if !strings.Contains(s, reset) {
		return f(s)
	}

	var sb strings.Builder
	start := 0
	for {
		idx := strings.Index(s[start:], reset)
		if idx == -1 {
			sb.WriteString(f(s[start:]))
			break
		}
		sb.WriteString(f(s[start : start+idx]))
		// skip the reset sequence (preserve original behavior of removing it)
		start += idx + len(reset)
	}
	return sb.String()
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

	style := ansi.Style{}.ForegroundColor(c)
	styled := style.Styled

	// Fast path: no newlines
	if !strings.Contains(str, "\n") {
		return addStylePreservingOriginalFormat(str, styled)
	}

	var sb strings.Builder
	start := 0
	for {
		idx := strings.IndexByte(str[start:], '\n')
		if idx == -1 {
			sb.WriteString(addStylePreservingOriginalFormat(str[start:], styled))
			break
		}
		sb.WriteString(addStylePreservingOriginalFormat(str[start:start+idx], styled))
		sb.WriteByte('\n')
		start += idx + 1
	}
	return sb.String()
}

func (b *Box) applyColorBar(topBar, bottomBar, title string) (string, string) {
	if b.titleColor == "" || title == "" {
		return topBar, bottomBar
	}

	converted := getConvertedColor(b.color)

	if b.titlePos == Top {
		strippedBar := ansi.Strip(topBar)
		strippedTitle := ansi.Strip(title)
		if idx := strings.Index(strippedBar, strippedTitle); idx != -1 {
			// split around first occurrence to preserve any other repeats
			b0 := applyConvertedColor(strippedBar[:idx], converted)
			b1 := applyConvertedColor(strippedBar[idx+len(strippedTitle):], converted)
			topBar = b0 + applyColor(title, b.titleColor) + b1
		}
	}

	if b.titlePos == Bottom {
		strippedBar := ansi.Strip(bottomBar)
		strippedTitle := ansi.Strip(title)
		if idx := strings.Index(strippedBar, strippedTitle); idx != -1 {
			// split around first occurrence to preserve any other repeats
			b0 := applyConvertedColor(strippedBar[:idx], converted)
			b1 := applyConvertedColor(strippedBar[idx+len(strippedTitle):], converted)
			bottomBar = b0 + applyColor(title, b.titleColor) + b1
		}
	}

	return topBar, bottomBar
}
