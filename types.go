package box

import (
	"os"

	"github.com/charmbracelet/colorprofile"
)

// BoxStyle defines a built‑in border style for a Box.
type BoxStyle string

const (
	// Single box style with single line borders.
	Single BoxStyle = "Single"
	// Double box style with double line borders.
	Double BoxStyle = "Double"
	// Round box style with rounded corners.
	Round BoxStyle = "Round"
	// Bold box style with bold lines.
	Bold BoxStyle = "Bold"
	// SingleDouble box style with single horizontal and double vertical lines.
	SingleDouble BoxStyle = "SingleDouble"
	// DoubleSingle box style with double horizontal and single vertical lines.
	DoubleSingle BoxStyle = "DoubleSingle"
	// Classic box style with plus and minus signs.
	Classic BoxStyle = "Classic"
	// Hidden box style with invisible borders.
	Hidden BoxStyle = "Hidden"
	// Block box style with block characters.
	Block BoxStyle = "Block"
)

// AlignType represents the horizontal alignment of content inside the box.
type AlignType string

const (
	// Center alignment
	Center AlignType = "Center"
	// Left alignment
	Left AlignType = "Left"
	// Right alignment
	Right AlignType = "Right"
)

// TitlePosition represents the position of the title relative to the box.
type TitlePosition string

const (
	// Inside title position
	Inside TitlePosition = "Inside"
	// Top title position
	Top TitlePosition = "Top"
	// Bottom title position
	Bottom TitlePosition = "Bottom"
)

// Color to support basic ANSI and bright ANSI colors.
type Color = string

const (
	Black   Color = "Black"
	Red     Color = "Red"
	Green   Color = "Green"
	Yellow  Color = "Yellow"
	Blue    Color = "Blue"
	Magenta Color = "Magenta"
	Cyan    Color = "Cyan"
	White   Color = "White"

	BrightBlack   Color = "BrightBlack"
	HiBlack       Color = "HiBlack"
	BrightRed     Color = "BrightRed"
	HiRed         Color = "HiRed"
	BrightGreen   Color = "BrightGreen"
	HiGreen       Color = "HiGreen"
	BrightYellow  Color = "BrightYellow"
	HiYellow      Color = "HiYellow"
	BrightBlue    Color = "BrightBlue"
	HiBlue        Color = "HiBlue"
	BrightMagenta Color = "BrightMagenta"
	HiMagenta     Color = "HiMagenta"
	BrightCyan    Color = "BrightCyan"
	HiCyan        Color = "HiCyan"
	BrightWhite   Color = "BrightWhite"
	HiWhite       Color = "HiWhite"
)

var (
	// boxes are inbuilt Box styles provided by the module
	boxes = map[BoxStyle]Box{
		Single: {
			topRight:    "┐",
			topLeft:     "┌",
			bottomRight: "┘",
			bottomLeft:  "└",
			horizontal:  "─",
			vertical:    "│",
		},
		Double: {
			topRight:    "╗",
			topLeft:     "╔",
			bottomRight: "╝",
			bottomLeft:  "╚",
			horizontal:  "═",
			vertical:    "║",
		},
		Round: {
			topRight:    "╮",
			topLeft:     "╭",
			bottomRight: "╯",
			bottomLeft:  "╰",
			horizontal:  "─",
			vertical:    "│",
		},
		Bold: {
			topRight:    "┓",
			topLeft:     "┏",
			bottomRight: "┛",
			bottomLeft:  "┗",
			horizontal:  "━",
			vertical:    "┃",
		},
		SingleDouble: {
			topRight:    "╖",
			topLeft:     "╓",
			bottomRight: "╜",
			bottomLeft:  "╙",
			horizontal:  "─",
			vertical:    "║",
		},
		DoubleSingle: {
			topRight:    "╕",
			topLeft:     "╒",
			bottomRight: "╛",
			bottomLeft:  "╘",
			horizontal:  "═",
			vertical:    "│",
		},
		Classic: {
			topRight:    "+",
			topLeft:     "+",
			bottomRight: "+",
			bottomLeft:  "+",
			horizontal:  "-",
			vertical:    "|",
		},
		Hidden: {
			topRight:    "+",
			topLeft:     "+",
			bottomRight: "+",
			bottomLeft:  "+",
			horizontal:  " ",
			vertical:    " ",
		},
		Block: {
			topRight:    "█",
			topLeft:     "█",
			bottomRight: "█",
			bottomLeft:  "█",
			horizontal:  "█",
			vertical:    "█",
		},
	}
)
var (
	// colorToHex maps color names to their hexadecimal codes.
	// This includes both standard and bright ANSI colors.
	colorToHex = map[Color]string{
		// Basic ANSI colors (0-7)
		Black:   "#000000",
		Red:     "#800000",
		Green:   "#008000",
		Yellow:  "#808000",
		Blue:    "#000080",
		Magenta: "#800080",
		Cyan:    "#008080",
		White:   "#C0C0C0",
		// Bright ANSI colors (8-15)
		BrightBlack:   "#808080",
		HiBlack:       "#808080",
		BrightRed:     "#FF0000",
		HiRed:         "#FF0000",
		BrightGreen:   "#00FF00",
		HiGreen:       "#00FF00",
		BrightYellow:  "#FFFF00",
		HiYellow:      "#FFFF00",
		BrightBlue:    "#0000FF",
		HiBlue:        "#0000FF",
		BrightMagenta: "#FF00FF",
		HiMagenta:     "#FF00FF",
		BrightCyan:    "#00FFFF",
		HiCyan:        "#00FFFF",
		BrightWhite:   "#FFFFFF",
		HiWhite:       "#FFFFFF",
	}

	profile = colorprofile.Detect(os.Stdout, os.Environ())
)
