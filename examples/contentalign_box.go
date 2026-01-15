package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 0).
		Style(box.Single).
		Color("Green").
		ContentAlign(box.Left) // Change to box.Right or box.Center to see the difference

	s, err := b.Render("Lorem Ipsum", "LoremIpsum\nfoo bar hello world\n123456 abcdefghijk")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
