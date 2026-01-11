package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	// Reported from ksctl
	title := "KUBECONFIG env var"
	lines := "/jknc/csdc"

	px := 4
	// if len(title) >= 2*px+len(lines) {
	// 	px = int(math.Ceil(float64(len(title)-len(lines))/2)) + 1
	// }

	b := box.NewBox().Padding(px, 2).
		Style(box.Double).
		Color("Red").
		TitlePosition(box.Top)

	fmt.Println(b.MustRender(title, lines))
}
