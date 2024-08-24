package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	bx := box.NewBox().Width(0).Height(0).Color("#26552e")
	b, _ := bx.Render("Box CLI Maker", "Make Highly Customizable Boxes for CLI")
	fmt.Println(b)
}
