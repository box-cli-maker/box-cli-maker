package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	// An example from ksctl tool
	// It didn't work in v2 due to inconsistent padding
	// when title is bigger than content width
	title := "KUBECONFIG env var"
	content := "/use/ksctl"

	b := box.NewBox().Padding(4, 2).
		Style(box.Double).
		Color("#8B75FF").
		TitlePosition(box.Top).
		TitleColor("#00ffb2").
		ContentColor("#12c78f")

	fmt.Println(b.MustRender(title, content))
}
