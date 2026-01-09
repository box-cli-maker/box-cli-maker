package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	b := box.NewBox().Padding(2, 1).Style(box.Single).Color("Cyan").TitlePosition(box.Top)
	s, err := b.Render("\033[1mBold\033[0m, works", "Btw \033[1mit works here too\033[0m, very nice")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
