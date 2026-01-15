package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 3).
		Style(box.Single).
		Color(box.Cyan).
		ContentColor("#d73a4a")

	content := `Lorem ipsum dolor sit amet, 
	ボックスメーカー
	consectetur adipiscing elit. Integer nec odio. Praesent libero.
	 Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh 
	 elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec 
	 tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.
	Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.`

	s, err := b.Render("Lorem ipsum", content)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
