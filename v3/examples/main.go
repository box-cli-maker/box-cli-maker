package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	bx := box.NewBox().Width(0).Height(0).Color("#26552e").TitlePosition(box.Top).Style(box.Single).TitleColor("#427ef5").ContentColor("#f5b342")
	b, _ := bx.Render("Box CLI Maker", "Make Highly Customizable \tBoxes for CLI")
	fmt.Println(b)

	by := box.NewBox().
		Width(2).
		Height(3).TitlePosition(box.Inside).
		SetTopRight("xyxy").
		SetTopLeft("xyxy").
		SetBottomRight("xyxy").
		SetBottomLeft("xyxy").
		SetHorizontal("-").
		SetVertical("|")

	b, _ = by.Render("Box CLI Maker", "Another custom box type")
	fmt.Println(b)
}
