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

// Box renders styled borders around text content.
type Box struct {
	// topRight renders the glyph used in the upper-right corner.
	topRight string
	// topLeft renders the glyph used in the upper-left corner.
	topLeft string
	// vertical renders the glyph used for the left and right walls.
	vertical string
	// bottomRight renders the glyph used in the lower-right corner.
	bottomRight string
	// bottomLeft renders the glyph used in the lower-left corner.
	bottomLeft string
	// horizontal renders the glyph used for the top and bottom edges.
	horizontal string
	config
}

// config contains configuration options for the Box.
type config struct {
	py            int           // Vertical padding.
	px            int           // Horizontal padding.
	contentAlign  AlignType     // Alignment for content inside the box.
	style         BoxStyle      // Active box style preset.
	titlePos      TitlePosition // Where the title, if any, is rendered.
	titleColor    string        // ANSI color (or hex code) for the title.
	contentColor  string        // ANSI color (or hex code) for the content.
	color         string        // ANSI color (or hex code) for the box chrome.
	allowWrapping bool          // Whether long content may wrap.
	wrappingLimit int           // Custom wrap width when wrapping is enabled.
	styleSet      bool          // Tracks if a style preset has already been applied.
}

// NewBox creates a new Box with Single box style.
func NewBox() *Box {
	b := &Box{}
	b.Style(Single)
	return b
}

// Copy returns a shallow copy of the Box so further mutations do not affect the original.
//
// Useful for creating base styles and deriving multiple boxes from them.
func (b *Box) Copy() *Box {
	if b == nil {
		return nil
	}
	clone := *b
	return &clone
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
	// Set the box style characters from predefined styles
	// This also allows manual overrides after setting style
	// and have a standard base.
	if styleDef, ok := boxes[box]; ok {
		b.BottomLeft(styleDef.bottomLeft).
			BottomRight(styleDef.bottomRight).
			TopLeft(styleDef.topLeft).
			TopRight(styleDef.topRight).
			Horizontal(styleDef.horizontal).
			Vertical(styleDef.vertical)
	}
	return b
}

// TopRight sets the glyph used in the upper-right corner.
func (b *Box) TopRight(glyph string) *Box {
	b.topRight = glyph
	return b
}

// TopLeft sets the glyph used in the upper-left corner.
func (b *Box) TopLeft(glyph string) *Box {
	b.topLeft = glyph
	return b
}

// BottomRight sets the glyph used in the lower-right corner.
func (b *Box) BottomRight(glyph string) *Box {
	b.bottomRight = glyph
	return b
}

// BottomLeft sets the glyph used in the lower-left corner.
func (b *Box) BottomLeft(glyph string) *Box {
	b.bottomLeft = glyph
	return b
}

// Horizontal sets the glyph used for the horizontal edges.
func (b *Box) Horizontal(glyph string) *Box {
	b.horizontal = glyph
	return b
}

// Vertical sets the glyph used for the vertical edges.
func (b *Box) Vertical(glyph string) *Box {
	b.vertical = glyph
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
//
// Not suitable for non-TTY outputs.
//
// For custom wrap limit and non-TTY outputs, use the WrapLimit method.
func (b *Box) WrapContent(allow bool) *Box {
	b.allowWrapping = allow
	return b
}

// WrapLimit sets the wrapping limit for content.
//
// When set wrapping will be done according to the limit provided.
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

// Render generates the box with the given title and content.
func (b *Box) Render(title, content string) (string, error) {
	if b.styleSet {
		if _, ok := boxes[b.config.style]; !ok {
			return "", fmt.Errorf("invalid Box style %s", b.config.style)
		}
	}

	var content_ []string

	// Allow wrapping according to the user
	if b.allowWrapping {
		if b.wrappingLimit < 0 {
			return "", fmt.Errorf("wrapping limit cannot be negative")
		}
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
			wrapWidth := max(2*width/defaultWrapDivisor, minWrapWidth)
			content = ansi.Wrap(content, wrapWidth, "")
		}
	}

	title, err := applyColor(title, b.titleColor)
	if err != nil {
		return "", err
	}
	content, err = applyColor(content, b.contentColor)
	if err != nil {
		return "", err
	}

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
	if TopBar, err = applyColor(TopBar, b.color); err != nil {
		return "", err
	}
	if BottomBar, err = applyColor(BottomBar, b.color); err != nil {
		return "", err
	}

	// Apply title coloring to the bars once, expanding tabs in the title if needed.
	titleForBar := title
	if strings.Contains(titleForBar, "\t") {
		titleForBar = xstrings.ExpandTabs(titleForBar, 4)
	}
	if TopBar, BottomBar, err = b.applyColorBar(TopBar, BottomBar, titleForBar); err != nil {
		return "", err
	}

	// Create lines to print
	texts, err := b.addVertPadding(innerWidth)
	if err != nil {
		return "", err
	}
	texts, err = b.formatLine(lines2, _longestLine, titleLen, sideMargin, title, texts)
	if err != nil {
		return "", err
	}
	vertPadding, err := b.addVertPadding(innerWidth)
	if err != nil {
		return "", err
	}
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
