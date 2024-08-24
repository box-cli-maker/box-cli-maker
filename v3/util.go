package box

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
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
	vertical := b.applyColor(b.Vertical, b.color)

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
		if runewidth.StringWidth(color.ClearCode(tmpLine.String())) < runewidth.StringWidth(tmpLine.String()) {
			lineLen = runewidth.StringWidth(color.ClearCode(tmpLine.String()))
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
		if runewidth.StringWidth(color.ClearCode(line.line)) < runewidth.StringWidth(line.line) {
			length = runewidth.StringWidth(color.ClearCode(line.line))
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
			format = AlignType(centerAlign)
		default:
			format = AlignType(b.findAlign())
		}

		sep := b.applyColor(b.Vertical, b.color)

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
	cstr := color.ClearCode(str)
	count := n - runewidth.StringWidth(cstr) - 2
	bar := strings.Repeat(c, count)
	return fmt.Sprintf(" %s %s", str, bar)
}

func (b *Box) applyColor(str string, color string) string {
	o := env.String(str)
	return o.Foreground(env.Color(color)).String()
}
