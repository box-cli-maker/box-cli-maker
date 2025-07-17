package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	bx := box.NewBox().Width(12).Height(15).Color("#26552e").TitlePosition(box.Top).Style(box.Single).TitleColor("#427ef5").ContentColor("#f5b342")
	b, _ := bx.Render("Box CLI Maker", "Make Highly Customizable Boxes for CLI in Go")
	fmt.Println(b)

	// by := box.NewBox().
	// 	Width(10).
	// 	Height(5).TitlePosition(box.Inside).
	// 	SetTopRight("xyxy").
	// 	SetTopLeft("xyxt").
	// 	SetBottomRight("xyxt").
	// 	SetBottomLeft("xyxy").
	// 	SetHorizontal("-").
	// 	SetVertical("|")

	// b, _ = by.Render("Box CLI Maker", "Another custom box type")
	// fmt.Println(b)
}
