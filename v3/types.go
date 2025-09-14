package box

import (
	"os"

	"github.com/charmbracelet/colorprofile"
)

type BoxStyle string

const (
	Single       BoxStyle = "Single"
	Double       BoxStyle = "Double"
	Round        BoxStyle = "Round"
	Bold         BoxStyle = "Bold"
	SingleDouble BoxStyle = "SingleDouble"
	DoubleSingle BoxStyle = "DoubleSingle"
	Classic      BoxStyle = "Classic"
	Hidden       BoxStyle = "Hidden"
	Block        BoxStyle = "Block"
)

type AlignType string

const (
	Center AlignType = "Center"
	Left   AlignType = "Left"
	Right  AlignType = "Right"
)

type TitlePosition string

const (
	Inside TitlePosition = "Inside"
	Top    TitlePosition = "Top"
	Bottom TitlePosition = "Bottom"
)

var (
	boxes = map[BoxStyle]Box{
		Single: {
			TopRight:    "┐",
			TopLeft:     "┌",
			BottomRight: "┘",
			BottomLeft:  "└",
			Horizontal:  "─",
			Vertical:    "│",
		},
		Double: {
			TopRight:    "╗",
			TopLeft:     "╔",
			BottomRight: "╝",
			BottomLeft:  "╚",
			Horizontal:  "═",
			Vertical:    "║",
		},
		Round: {
			TopRight:    "╮",
			TopLeft:     "╭",
			BottomRight: "╯",
			BottomLeft:  "╰",
			Horizontal:  "─",
			Vertical:    "│",
		},
		Bold: {
			TopRight:    "┓",
			TopLeft:     "┏",
			BottomRight: "┛",
			BottomLeft:  "┗",
			Horizontal:  "━",
			Vertical:    "┃",
		},
		SingleDouble: {
			TopRight:    "╖",
			TopLeft:     "╓",
			BottomRight: "╜",
			BottomLeft:  "╙",
			Horizontal:  "─",
			Vertical:    "║",
		},
		DoubleSingle: {
			TopRight:    "╕",
			TopLeft:     "╒",
			BottomRight: "╛",
			BottomLeft:  "╘",
			Horizontal:  "═",
			Vertical:    "│",
		},
		Classic: {
			TopRight:    "+",
			TopLeft:     "+",
			BottomRight: "+",
			BottomLeft:  "+",
			Horizontal:  "-",
			Vertical:    "|",
		},
		Hidden: {
			TopRight:    "+",
			TopLeft:     "+",
			BottomRight: "+",
			BottomLeft:  "+",
			Horizontal:  " ",
			Vertical:    " ",
		},
		Block: {
			TopRight:    "█",
			TopLeft:     "█",
			BottomRight: "█",
			BottomLeft:  "█",
			Horizontal:  "█",
			Vertical:    "█",
		},
	}

	colorNameToHex = map[string]string{
		// Basic ANSI colors (0-7)
		"Black":   "#000000",
		"Red":     "#800000",
		"Green":   "#008000",
		"Yellow":  "#808000",
		"Blue":    "#000080",
		"Magenta": "#800080",
		"Cyan":    "#008080",
		"White":   "#C0C0C0",
		// Bright ANSI colors (8-15)
		"BrightBlack":   "#808080",
		"DarkGray":      "#808080",
		"BrightRed":     "#FF0000",
		"HiRed":         "#FF0000",
		"BrightGreen":   "#00FF00",
		"HiGreen":       "#00FF00",
		"BrightYellow":  "#FFFF00",
		"HiYellow":      "#FFFF00",
		"BrightBlue":    "#0000FF",
		"HiBlue":        "#0000FF",
		"BrightMagenta": "#FF00FF",
		"HiMagenta":     "#FF00FF",
		"BrightCyan":    "#00FFFF",
		"HiCyan":        "#00FFFF",
		"BrightWhite":   "#FFFFFF",
	}

	profile = colorprofile.Detect(os.Stdout, os.Environ())
)
