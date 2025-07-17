package box

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/huandu/xstrings"
	"github.com/mattn/go-runewidth"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

const (
	// 1 = separator, 2 = spacing, 3 = line; 4 = oddSpace; 5 = space; 6 = sideMargin
	centerAlign = "%[1]s%[2]s%[3]s%[4]s%[2]s%[1]s"
	leftAlign   = "%[1]s%[6]s%[3]s%[4]s%[2]s%[5]s%[1]s"
	rightAlign  = "%[1]s%[2]s%[4]s%[5]s%[3]s%[6]s%[1]s"
)

type Box struct {
	TopRight    string // TopRight Corner Symbols
	TopLeft     string // TopLeft Corner Symbols
	Vertical    string // Vertical Bar Symbols
	BottomRight string // BottomRight Corner Symbols
	BottomLeft  string // BottomLeft Corner Symbols
	Horizontal  string // Horizontal Bars Symbols
	config             // Box Config

}

// Config is the configuration needed for the Box to be designed
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
	styleSet      bool
}

func NewBox() *Box {
	return &Box{config: config{style: Single, styleSet: false}}
}

func (b *Box) Width(width int) *Box {
	b.px = width
	return b
}

func (b *Box) Height(height int) *Box {
	b.py = height
	return b
}

func (b *Box) Style(box BoxStyle) *Box {
	b.config.style = box
	b.styleSet = true
	return b
}

func (b *Box) SetTopRight(sym string) *Box {
	b.TopRight = sym
	return b
}

func (b *Box) SetTopLeft(sym string) *Box {
	b.TopLeft = sym
	return b
}

func (b *Box) SetBottomRight(sym string) *Box {
	b.BottomRight = sym
	return b
}

func (b *Box) SetBottomLeft(sym string) *Box {
	b.BottomLeft = sym
	return b
}

func (b *Box) SetHorizontal(sym string) *Box {
	b.Horizontal = sym
	return b
}

func (b *Box) SetVertical(sym string) *Box {
	b.Vertical = sym
	return b
}

func (b *Box) TitleColor(color string) *Box {
	b.titleColor = color
	return b
}

func (b *Box) ContentColor(color string) *Box {
	b.contentColor = color
	return b
}

func (b *Box) Color(color string) *Box {
	b.color = color
	return b
}

func (b *Box) TitlePosition(pos TitlePosition) *Box {
	b.titlePos = pos
	return b
}

func (b *Box) Render(title, content string) (string, error) {
	style, ok := boxes[b.config.style]

	if ok && b.styleSet {
		b.SetBottomLeft(style.BottomLeft).
			SetBottomRight(style.BottomRight).
			SetTopLeft(style.TopLeft).
			SetTopRight(style.TopRight).
			SetHorizontal(style.Horizontal).
			SetVertical(style.Vertical)
	}
	if !ok && b.styleSet {
		return "", fmt.Errorf("invalid Box type")
	}

	var content_ []string

	// Allow Wrapping according to the user
	if b.allowWrapping {
		// If limit not provided then use 2*TermWidth/3 as limit else
		// use the one provided
		if b.wrappingLimit != 0 {
			content = wordwrap.String(content, b.wrappingLimit)
		} else {
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				return "", fmt.Errorf("cannot get terminal size from the output")
			}
			content = wordwrap.String(content, 2*width/3)
		}
	}

	title = applyColor(title, b.titleColor)
	content = applyColor(content, b.contentColor)

	if b.titlePos == "" {
		b.titlePos = Inside
	}

	if title != "" {
		if b.titlePos != Inside && strings.Contains(title, "\n") {
			return "", fmt.Errorf("multiline titles are only supported inside only")
		}
		if b.titlePos == Inside {
			content_ = append(content_, strings.Split(title, "\n")...)
			content_ = append(content_, []string{""}...) // for empty line between title and content
		}
	}
	content_ = append(content_, strings.Split(content, "\n")...)

	titleLen := len(strings.Split(color.ClearCode(title), "\n"))
	sideMargin := strings.Repeat(" ", b.px)
	_longestLine, lines2 := longestLine(content_)

	// Get padding on one side
	paddingCount := b.px

	n := _longestLine + (paddingCount * 2) + 2

	if b.titlePos != Inside && runewidth.StringWidth(color.ClearCode(title)) > n-2 {
		return "", fmt.Errorf("title must be shorter than the Top and Bottom Bars")
	}

	// Create Top and Bottom Bars
	Bar := strings.Repeat(b.Horizontal, n-2)
	TopBar := b.TopLeft + Bar + b.TopRight
	BottomBar := b.BottomLeft + Bar + b.BottomRight

	var TitleBar string
	// If title has tabs then expand them accordingly.
	if strings.Contains(title, "\t") {
		TitleBar = repeatWithString(b.Horizontal, n-2, xstrings.ExpandTabs(title, 4))
	} else {
		TitleBar = repeatWithString(b.Horizontal, n-2, title)
	}

	// Check b.TitlePos if it is not Inside
	if b.titlePos != Inside {
		switch b.titlePos {
		case Top:
			TopBar = b.TopLeft + TitleBar + b.TopRight
		case Bottom:
			BottomBar = b.BottomLeft + TitleBar + b.BottomRight
		default:
			return "", fmt.Errorf("invalid TitlePos provided")
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

	if b.titlePos == Inside && runewidth.StringWidth(TopBar) != runewidth.StringWidth(BottomBar) {
		return "", fmt.Errorf("cannot create a Box with different sizes of Top and Bottom Bars")
	}

	// Create lines to print
	texts := b.addVertPadding(n, TopBar)
	texts = b.formatLine(lines2, _longestLine, titleLen, sideMargin, title, TopBar, texts)
	vertPadding := b.addVertPadding(n, TopBar)
	texts = append(texts, vertPadding...)

	// Using strings.Builder is more efficient and faster
	// than concatenating 6 times
	var sb strings.Builder

	sb.WriteString(TopBar)
	sb.WriteString("\n")
	sb.WriteString(strings.Join(texts, "\n"))
	sb.WriteString("\n")
	sb.WriteString(BottomBar)
	sb.WriteString("\n")

	return sb.String(), nil

}
