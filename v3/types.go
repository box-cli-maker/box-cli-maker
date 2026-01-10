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
