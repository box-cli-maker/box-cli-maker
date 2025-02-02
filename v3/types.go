package box

import "github.com/muesli/termenv"

type BoxStyle string
type AlignType string
type TitlePosition string

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

const (
	Center AlignType = "Center"
	Left   AlignType = "Left"
	Right  AlignType = "Right"
)

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
	env = termenv.EnvColorProfile()
)
