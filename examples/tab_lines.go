package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	styleCases := []box.BoxStyle{box.Single, box.Double, box.SingleDouble, box.DoubleSingle, box.Bold, box.Round, box.Hidden, box.Classic}
	colorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta", "White", "HiBlue", "HiRed", "HiGreen", "HiYellow", "HiCyan", "HiMagenta", "HiWhite"}

	for _, st := range styleCases {
		for _, c := range colorTypes {
			b := box.NewBox().Padding(2, 6).Style(st).Color(c)
			s, err := b.Render("Box \tCLI \tMaker \t📦", "Highly \t\tCustomizable \t\t\tTerminal\t Box\t Maker")
			if err != nil {
				panic(err)
			}
			fmt.Println(s)
		}
	}
}
