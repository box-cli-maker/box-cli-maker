package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	bx := box.NewBox().Width(12).Height(14)
	b, _ := bx.Render("Box CLI Maker", "Make")
	fmt.Println(b)
}
