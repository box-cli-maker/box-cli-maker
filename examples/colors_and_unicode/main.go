package main

import (
	"fmt"
	"strings"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	styles := []box.BoxStyle{
		box.Single,
		box.Double,
		box.SingleDouble,
		box.DoubleSingle,
		box.Bold,
		box.Round,
		box.Hidden,
		box.Classic,
	}

	colors := []box.Color{
		box.Black, box.Blue, box.Red, box.Green, box.Yellow, box.Cyan, box.Magenta, box.White,
		box.BrightBlack, box.BrightBlue, box.BrightRed, box.BrightGreen, box.BrightYellow,
		box.BrightCyan, box.BrightMagenta, box.BrightWhite,
	}

	// Color and style combinations (border color only).
	for _, style := range styles {
		for _, c := range colors {
			b := box.NewBox().
				Padding(2, 5).
				Style(style).
				Color(c)

			out, err := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Style: %s, Color: %s\n%s\n\n", style, c, out)
		}
	}

	// Multiline + tabs with colored title/content.
	b := box.NewBox().
		Padding(2, 5).
		Style(box.Single).
		TitlePosition(box.Top).
		Color(box.Green).
		TitleColor(box.Cyan).
		ContentColor(box.Red)

	multi := "Make\n\tHighly\n\t\tCustomized\n\t\t\tTerminal\n\t\t\t\tBoxes"

	out, err := b.Render("Box CLI Maker", multi)
	if err != nil {
		panic(err)
	}
	fmt.Println("Multiline + tabs:")
	fmt.Println(out)

	// Unicode / emoji demo.
	titles := []string{
		"Box CLI Maker",
		"ボックスメーカー",
		"盒子製造商",
		"박스 메이커",
		"Créateur de boîte CLI",
		"Fabricante de cajas",
		"Qui fecit me arca CLI",
		"Κουτί CLI Maker",
	}
	lines := []string{
		"Make Highly Customized Terminal Boxes",
		"高度にカスタマイズされた端子ボックスを作成する",
		"製作高度定制的接線盒",
		"고도로 맞춤화 된 터미널 박스 만들기",
		"Créez des boîtes à bornes hautement personnalisées",
		"Haga cajas de terminales altamente personalizadas",
		"Fac multum Customized Terminal Pyxidas",
		"Δημιουργήστε πολύ προσαρμοσμένα τερματικά κουτιά",
	}

	for i := range titles {
		for _, style := range styles {
			b := box.NewBox().
				Padding(2, 5).
				Style(style)

			out, err := b.Render(titles[i], lines[i])
			if err != nil {
				panic(err)
			}

			fmt.Printf("Unicode style: %s\n%s\n\n", style, out)
		}
	}

	// Wrapping + emoji demo.
	content := strings.Repeat(" Box CLI Maker 盒子製 造商,📦 ", 10)

	bw := box.NewBox().
		Padding(2, 0).
		Style(box.Single).
		Color(box.Green).
		TitlePosition(box.Top).
		WrapContent(true).
		WrapLimit(40).
		ContentColor(box.Cyan).
		TitleColor(box.BrightRed)

	out, err = bw.Render("Box\tCLI\tMaker\t📦", content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrapped with tabs + emoji:")
	fmt.Println(out)
}
