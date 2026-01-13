package box

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
	"github.com/huandu/xstrings"
	"github.com/mattn/go-runewidth"
)

const (
	// 1 = separator, 2 = spacing, 3 = line; 4 = oddSpace; 5 = space; 6 = sideMargin
	centerAlign = "%[1]s%[2]s%[3]s%[4]s%[2]s%[1]s"
	leftAlign   = "%[1]s%[6]s%[3]s%[4]s%[2]s%[5]s%[1]s"
	rightAlign  = "%[1]s%[2]s%[4]s%[5]s%[3]s%[6]s%[1]s"

	defaultWrapDivisor = 3  // 2/3 of terminal width
	minWrapWidth       = 20 // Minimum width to wrap content
)

type Box struct {
	topRight    string // TopRight Corner Symbols
	topLeft     string // TopLeft Corner Symbols
	vertical    string // Vertical Bar Symbols
	bottomRight string // BottomRight Corner Symbols
	bottomLeft  string // BottomLeft Corner Symbols
	horizontal  string // Horizontal Bars Symbols
	config             // Box Config

}

// Config holds the configuration for the Box
type config struct {
	py            int           // Horizontal Padding
	px            int           // Vertical Padding
	contentAlign  AlignType     // Content Alignment inside Box
	style         BoxStyle      // Box Style
	titlePos      TitlePosition // Title Position
	titleColor    string        // Title Color
	contentColor  string        // Content Color
	color         string        // Box Color
	allowWrapping bool          // Flag to allow custom Content Wrapping
	wrappingLimit int           // Wrap the Content upto the Limit
	styleSet      bool          // Flag to check if inbuilt-Style is set
}

// NewBox creates a new Box with default configuration.
func NewBox() *Box {
	return &Box{config: config{style: Single}}
}

// Padding sets horizontal (px) and vertical (py) inner padding.
func (b *Box) Padding(px, py int) *Box {
	b.px = px
	b.py = py
	return b
}

// HPadding sets horizontal padding (left/right).
func (b *Box) HPadding(px int) *Box {
	b.px = px
	return b
}

// VPadding sets vertical padding (top/bottom).
func (b *Box) VPadding(py int) *Box {
	b.py = py
	return b
}

// Style sets the Box Style.
//
// Box Styles: box.Single, box.Double, box.Round, box.Bold, box.SingleDouble, box.DoubleSingle, box.Classic, box.Hidden, box.Block
//
// To make custom box styles, use the TopRight, TopLeft, BottomRight, BottomLeft, Horizontal, and Vertical methods.
//
// Example:
//
// b := box.NewBox()
// b.TopRight("+").TopLeft("+").BottomRight("+").BottomLeft("_").Horizontal("-").Vertical("|")
func (b *Box) Style(box BoxStyle) *Box {
	b.config.style = box
	b.styleSet = true
	return b
}

// TopRight sets the TopRight Corner Symbols.
func (b *Box) TopRight(sym string) *Box {
	b.topRight = sym
	return b
}

// TopLeft sets the TopLeft Corner Symbols.
func (b *Box) TopLeft(sym string) *Box {
	b.topLeft = sym
	return b
}

// BottomRight sets the BottomRight Corner Symbols.
func (b *Box) BottomRight(sym string) *Box {
	b.bottomRight = sym
	return b
}

// BottomLeft sets the BottomLeft Corner Symbols.
func (b *Box) BottomLeft(sym string) *Box {
	b.bottomLeft = sym
	return b
}

// Horizontal sets the Horizontal Bar Symbols.
func (b *Box) Horizontal(sym string) *Box {
	b.horizontal = sym
	return b
}

// Vertical sets the Vertical Bar Symbols.
func (b *Box) Vertical(sym string) *Box {
	b.vertical = sym
	return b
}

// TitleColor sets the Title Color.
func (b *Box) TitleColor(color string) *Box {
	b.titleColor = color
	return b
}

// ContentColor sets the Content Color.
func (b *Box) ContentColor(color string) *Box {
	b.contentColor = color
	return b
}

// Color sets the Box Color.
func (b *Box) Color(color string) *Box {
	b.color = color
	return b
}

// TitlePosition sets the Title Position.
//
// Title Positions: box.Inside, box.Top, box.Bottom
func (b *Box) TitlePosition(pos TitlePosition) *Box {
	b.titlePos = pos
	return b
}

// WrapContent enables or disables content wrapping.
//
// When enabled, the content will be wrapped to fit 2/3 width of the terminal by default.
// You can set a custom wrap limit using the WrapLimit method.
func (b *Box) WrapContent(allow bool) *Box {
	b.allowWrapping = allow
	return b
}

// WrapLimit sets the wrapping limit for content.
func (b *Box) WrapLimit(limit int) *Box {
	b.allowWrapping = true
	b.wrappingLimit = limit
	return b
}

// ContentAlign sets the content alignment inside the Box.
//
// Alignment Types: box.Left, box.Center, box.Right
func (b *Box) ContentAlign(align AlignType) *Box {
	b.contentAlign = align
	return b
}

// Render generates the box with the given title and content.
func (b *Box) Render(title, content string) (string, error) {
	style, ok := boxes[b.config.style]

	if ok && b.styleSet {
		b.BottomLeft(style.bottomLeft).
			BottomRight(style.bottomRight).
			TopLeft(style.topLeft).
			TopRight(style.topRight).
			Horizontal(style.horizontal).
			Vertical(style.vertical)
	}
	if !ok && b.styleSet {
		return "", fmt.Errorf("invalid Box style %s", b.config.style)
	}

	var content_ []string

	// Allow Wrapping according to the user
	if b.allowWrapping {
		// If limit not provided then use 2*TermWidth/3 as limit else
		// use the one provided
		if b.wrappingLimit != 0 {
			content = ansi.Wrap(content, b.wrappingLimit, "")
		} else {
			width, _, err := term.GetSize(os.Stdout.Fd())
			if err != nil {
				return "", fmt.Errorf("cannot get terminal size from the output, provide own wrapping limit using .WrapLimit(limit int) method")
			}
			// Use 2/3 of terminal width as default wrapping limit
			wrapWidth := 2 * width / defaultWrapDivisor
			if wrapWidth < minWrapWidth {
				wrapWidth = minWrapWidth
			}
			content = ansi.Wrap(content, wrapWidth, "")
		}
	}

	title = applyColor(title, b.titleColor)
	content = applyColor(content, b.contentColor)

	if b.titlePos == "" {
		b.titlePos = Inside
	}

	if title != "" {
		if b.titlePos != Inside && strings.Contains(title, "\n") {
			return "", fmt.Errorf("multiline titles are only supported Inside title position only")
		}
		if b.titlePos == Inside {
			content_ = append(content_, strings.Split(title, "\n")...)
			content_ = append(content_, []string{""}...) // for empty line between title and content
		}
	}
	content_ = append(content_, strings.Split(content, "\n")...)

	titleLen := 0
	if title != "" {
		titleLen = len(strings.Split(ansi.Strip(title), "\n"))
	}

	sideMargin := strings.Repeat(" ", b.px)
	_longestLine, lines2 := longestLine(content_)

	// Compute desired inner width (between the vertical borders, excluding them).
	contentInnerWidth := _longestLine + 2*b.px
	innerWidth := contentInnerWidth

	// Make sure the box is wide enough to fit the title when it's on Top/Bottom.
	if b.titlePos != Inside && title != "" {
		titleWidth := runewidth.StringWidth(ansi.Strip(title))
		minTitleInnerWidth := titleWidth + 2 // title + left/right padding

		if minTitleInnerWidth > innerWidth {
			innerWidth = minTitleInnerWidth
		}
	}

	// If we enlarged the inner width to fit the title, reflect that in longestLine.
	if innerWidth > contentInnerWidth {
		_longestLine = max(innerWidth-2*b.px, 0)
	}

	// Visible widths of box characters; fall back to 1 so we always make progress.
	verticalWidth := charWidth(b.vertical)
	horizontalWidth := charWidth(b.horizontal)
	topLeftWidth := charWidth(b.topLeft)
	topRightWidth := charWidth(b.topRight)
	bottomLeftWidth := charWidth(b.bottomLeft)
	bottomRightWidth := charWidth(b.bottomRight)

	// Ensure the inner width is a multiple of the horizontal glyph width when
	// drawing horizontal bars (e.g. emoji) so we don't need to pad with extra
	// spaces before the corner. This keeps the bar visually uniform.
	if horizontalWidth > 1 && innerWidth%horizontalWidth != 0 {
		innerWidth += horizontalWidth - (innerWidth % horizontalWidth)
		_longestLine = max(innerWidth-2*b.px, 0)
	}

	// Total visible width of a rendered line (including vertical borders).
	lineWidth := innerWidth + 2*verticalWidth

	TopBar := buildPlainBar(b.topLeft, b.horizontal, b.topRight, topLeftWidth, topRightWidth, lineWidth, horizontalWidth)
	BottomBar := buildPlainBar(b.bottomLeft, b.horizontal, b.bottomRight, bottomLeftWidth, bottomRightWidth, lineWidth, horizontalWidth)

	// Check b.TitlePos if it is not Inside
	if b.titlePos != Inside {
		switch b.titlePos {
		case Top:
			TopBar = buildTitledBar(b.topLeft, b.horizontal, b.topRight, topLeftWidth, topRightWidth, lineWidth, horizontalWidth, title)
		case Bottom:
			BottomBar = buildTitledBar(b.bottomLeft, b.horizontal, b.bottomRight, bottomLeftWidth, bottomRightWidth, lineWidth, horizontalWidth, title)
		default:
			return "", fmt.Errorf("invalid TitlePosition %s", b.titlePos)
		}
	}
	TopBar, BottomBar = applyColor(TopBar, b.color), applyColor(BottomBar, b.color)

	// Check type of b.Color then assign the Colors to TopBar and BottomBar accordingly
	// If title has tabs then expand them accordingly.
	if strings.Contains(title, "\t") {
		TopBar, BottomBar = b.applyColorBar(TopBar, BottomBar, xstrings.ExpandTabs(title, 4))
	} else {
		TopBar, BottomBar = b.applyColorBar(TopBar, BottomBar, title)
	}

	// Create lines to print
	texts := b.addVertPadding(innerWidth)
	texts, err := b.formatLine(lines2, _longestLine, titleLen, sideMargin, title, texts)
	if err != nil {
		return "", err
	}
	vertPadding := b.addVertPadding(innerWidth)
	texts = append(texts, vertPadding...)

	var sb strings.Builder

	sb.WriteString(TopBar)
	sb.WriteString("\n")
	sb.WriteString(strings.Join(texts, "\n"))
	sb.WriteString("\n")
	sb.WriteString(BottomBar)
	sb.WriteString("\n")

	return sb.String(), nil

}

// MustRender is like Render but panics if an error occurs.
//
// Useful to generate boxes without having to handle the error.
func (b *Box) MustRender(title, content string) string {
	s, err := b.Render(title, content)
	if err != nil {
		panic(err)
	}
	return s
}
