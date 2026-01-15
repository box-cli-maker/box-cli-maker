package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	styleCases := []box.BoxStyle{box.Single, box.Double, box.SingleDouble, box.DoubleSingle, box.Bold, box.Round, box.Hidden, box.Classic}
	colorTypes := []box.Color{
		box.Black, box.Blue, box.Red, box.Green, box.Yellow, box.Cyan, box.Magenta, box.White,
		box.HiBlue, box.HiRed, box.HiGreen, box.HiYellow, box.HiCyan, box.HiMagenta, box.HiWhite,
	}

	for _, st := range styleCases {
		for _, c := range colorTypes {
			b := box.NewBox().Padding(2, 6).
				Style(st).
				Color(c).
				ContentColor(box.Cyan).
				TitleColor("#d73a4a")

			s, err := b.Render("Box CLI Maker 📦", "Highly Customizable Terminal\tBox Maker")
			if err != nil {
				panic(err)
			}
			fmt.Println(s)
		}
	}
}
