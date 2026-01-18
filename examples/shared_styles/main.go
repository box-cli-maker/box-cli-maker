package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {

	baseBox := box.NewBox().Padding(2, 5).Style(box.Single)

	greenBox := baseBox.Copy().Color(box.Green)
	redBox := baseBox.Copy().Color(box.Red)

	fmt.Println(baseBox.MustRender("Base box", "this is the base box"))

	fmt.Println(greenBox.MustRender("Green Box", "This is a green box"))

	fmt.Println(redBox.MustRender("Red Box", "This is a red box"))
}
