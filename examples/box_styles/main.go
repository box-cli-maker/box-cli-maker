package main

import (
	"fmt"
	"strings"

	box "github.com/Delta456/box-cli-maker/v3"
)

func indentBox(s string, spaces int) string {
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		lines[i] = pad + l
	}
	return strings.Join(lines, "\n")
}

func main() {
	styles := []box.BoxStyle{
		box.Single,
		box.SingleDouble,
		box.Double,
		box.DoubleSingle,
		box.Bold,
		box.Round,
		box.Hidden,
		box.Classic,
		box.Block,
	}

	for _, style := range styles {
		b := box.NewBox().
			Padding(4, 3).
			Style(style).
			TitleColor("#00ffb2").
			Color("#8B75FF").
			ContentColor("#12c78f").ContentAlign(box.Center)

		out, err := b.Render("Box CLI Maker",
			"Render highly customizable boxes\nin the terminal")
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%s", indentBox(out, 4))
	}
}
