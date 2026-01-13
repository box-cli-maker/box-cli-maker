package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	b := box.NewBox().Padding(2, 3).Color("Red").TitlePosition(box.Top)
	b.TopRight("📦").TopLeft("📦🚀").BottomRight("📦").BottomLeft("📦").Horizontal("📦").Vertical("📦")

	s, err := b.Render("Box CLI 	Maker", "Make Highly Customizable Terminal Boxes")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	b1 := box.NewBox().Padding(2, 3).Color("Green").TitlePosition(box.Bottom)
	b1.TopRight("+").TopLeft("+").BottomRight("+").BottomLeft("++").Horizontal("-").Vertical("|")

	s1, err := b1.Render("Box CLI Maker", "Make Highly Customizable Terminal Boxes")
	if err != nil {
		panic(err)
	}
	fmt.Println(s1)
}
