package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	// English, Japanese, Chinese(Traditional), Korean, French, Spanish, Latin, Greek
	titles := []string{"Box CLI Maker", "ボックスメーカー", "盒子製造商", "박스 메이커", "Créateur de boîte CLI", "Fabricante de cajas", "Qui fecit me arca CLI", "Κουτί CLI Maker"}
	lines := []string{"Make Highly Customized Terminal Boxes", "高度にカスタマイズされた端子ボックスを作成する", "製作高度定制的接線盒", "고도로 맞춤화 된 터미널 박스 만들기", "Créez des boîtes à bornes hautement personnalisées", "Haga cajas de terminales altamente personalizadas", "Fac multum Customized Terminal Pyxidas", "Δημιουργήστε πολύ προσαρμοσμένα τερματικά κουτιά"}

	styleCases := []box.BoxStyle{box.Single, box.Double, box.SingleDouble, box.DoubleSingle, box.Bold, box.Round, box.Hidden, box.Classic}

	for i := range titles {
		for _, st := range styleCases {
			b := box.NewBox().Padding(2, 5).Style(st)
			s, err := b.Render(titles[i], lines[i])
			if err != nil {
				panic(err)
			}
			fmt.Println(s)
		}
	}
}
