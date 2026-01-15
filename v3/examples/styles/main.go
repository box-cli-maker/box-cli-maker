package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

// This example mirrors the old TestInbuiltStyles by rendering all built-in
// styles with the same padding, title, and content.
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
	}

	for _, style := range styles {
		b := box.NewBox().
			Padding(2, 5).
			Style(style)

		out, err := b.Render("Box CLI Maker", "Render highly customizable boxes for terminal")
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%s\n\n", out)
	}
}
