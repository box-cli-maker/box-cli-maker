package main

import (
	"fmt"
	"strings"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 0).
		Style(box.Single).
		Color(box.Green).
		TitlePosition(box.Top).
		WrapContent(true)
		// Provide your limit with WrapLimit if needed

	s, err := b.Render("Content Wrappingg works!", strings.Repeat("\tBox CLI Maker 盒子製 造商,📦 ", 160))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
