package box

import "fmt"

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
	return &Box{config: config{style: Single}}
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

func (b *Box) Render(title, content string) (string, error) {
	_, ok := boxes[b.config.style]
	if !ok && b.styleSet {
		return "", fmt.Errorf("invalid Box type")
	}
	return title, nil

}
