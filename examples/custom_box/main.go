package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 3).Color(box.Red).TitlePosition(box.Top)
	b.TopRight("📦").TopLeft("📦🚀").BottomRight("📦").BottomLeft("📦").Horizontal("📦").Vertical("📦")

	fmt.Println(b.MustRender("Box CLI Maker", "Make Highly Customizable Terminal Boxes"))

	b1 := box.NewBox().Padding(2, 3).Color(box.Green).TitlePosition(box.Bottom).TitleColor("#00ffb2").Color("#8B75FF").ContentColor("#12c78f")
	b1.TopRight("+").TopLeft("+").BottomRight("+").BottomLeft("++").Horizontal("-").Vertical("||")

	fmt.Println(b1.MustRender("Box CLI Maker", "Make Highly Customizable Terminal Boxes"))
}
