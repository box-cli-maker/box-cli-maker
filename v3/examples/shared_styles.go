package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {

	baseBox := box.NewBox().Padding(2, 5).Style(box.Single)

	greenBox := baseBox.Copy().Color("Green")
	redBox := baseBox.Copy().Color("Red")

	fmt.Println(baseBox.MustRender("Base box", "content"))

	fmt.Println(greenBox.MustRender("Green Box", "This is a green box"))

	fmt.Println(redBox.MustRender("Red Box", "This is a red box"))
}
