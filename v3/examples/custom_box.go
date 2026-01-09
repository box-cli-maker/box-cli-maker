package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	b := box.NewBox().Padding(2, 3).Color("Red").TitlePosition(box.Inside)
	b.WithTopRight("█").WithTopLeft("█").WithBottomRight("█").WithBottomLeft("█").WithHorizontal("█").WithVertical("█")

	s, err := b.Render("Box CLI Maker", "Make Highly Customized Terminal Boxes")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
