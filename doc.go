// Package box renders styled boxes around text for terminal applications.
//
// The core type is Box, constructed with NewBox and configured via a fluent API.
// Boxes support multiple built‑in styles, title positions, alignment, wrapping,
// and ANSI/truecolor output.
//
// Basic example:
//
//	b := box.NewBox().
//		Style(box.Single).
//		Padding(2, 1).
//		TitlePosition(box.Top).
//		ContentAlign(box.Center).
//		Color("Cyan").
//		TitleColor("BrightYellow").
//		ContentColor("White")
//
//	out, err := b.Render("Box CLI Maker", "Highly customizable terminal box maker")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(out)
//
// # Styles
//
// Box styles are selected with Style and the BoxStyle constants:
//
//	box.Single
//	box.Double
//	box.Round
//	box.Bold
//	box.SingleDouble
//	box.DoubleSingle
//	box.Classic
//	box.Hidden
//	box.Block
//
// You can further customize any style by overriding the corner and edge glyphs
// using TopRight, TopLeft, BottomRight, BottomLeft, Horizontal, and Vertical.
//
// # Titles and alignment
//
// Titles can be placed inside the box, on the top border, or on the bottom
// border using TitlePosition with the TitlePosition constants:
//
//	box.Inside
//	box.Top
//	box.Bottom
//
// Content alignment is controlled with ContentAlign and the AlignType
// constants:
//
//	box.Left
//	box.Center
//	box.Right
//
// # Wrapping
//
// WrapContent enables or disables automatic wrapping of the content. By
// default, when wrapping is enabled, the box width is based on two‑thirds of
// the terminal width. WrapLimit can be used to set an explicit maximum width.
//
// # Colors
//
// TitleColor, ContentColor, and Color accept either one of the first 16 ANSI
// color names (Black, Red, Green, Yellow, Blue, Magenta, Cyan, White and their
// bright variants) or a #RGB / #RRGGBB / rgb:RRRR/GGGG/BBBB /
// rgba:RRRR/GGGG/BBBB/AAAA value. Invalid colors cause Render to return an
// error.
//
// # Errors
//
// Render returns an error if the style or title position is invalid, the wrap
// limit or padding is negative, a multiline title is used with a non‑Inside
// title position, any configured colors are invalid, or the terminal width
// cannot be determined. MustRender is a convenience wrapper that panics on
// error.
//
// # Construction
//
// Box must be constructed with NewBox; the zero value is not usable.
//
// # Copying
//
// Copy returns a shallow copy of a Box so you can define a base style and
// derive variants without mutating the original:
//
//	base := box.NewBox().Style(box.Single).Padding(2, 1)
//	info := base.Copy().Color("Green")
//	warn := base.Copy().Color("Yellow")
//
// Each Copy can then be customized and rendered independently.
package box
