package main

import (
	"fmt"
	"strings"

	box "github.com/Delta456/box-cli-maker/v2/v3"
)

func main() {
	bx := box.NewBox().Width(2).Height(3).Color("Red").TitlePosition(box.Top).Style(box.Single).TitleColor("#427ef5").ContentColor("#f5b342")
	b, _ := bx.Render("Box CLI Maker", "	Make Highly Customizable Boxes for CLI in Go				")
	fmt.Println(b)

	//Broken, need to fix
	// ContentColor is not working
	// bx1 := box.NewBox().Width(2).Height(5).Color("Cyan").Style(box.Single).TitleColor("#427ef5")
	// b1, _ := bx1.Render("", `
	// __________                 _________  .____     .___     _____            __
	// \______   \  ____ ___  ___ \_   ___ \ |    |    |   |   /     \  _____   |  | __  ____ _______
	//  |    |  _/ /  _ \\  \/  / /    \  \/ |    |    |   |  /  \ /  \ \__  \  |  |/ /_/ __ \\_  __ \
	//  |    |   \(  <_> )>    <  \     \____|    |___ |   | /    Y    \ / __ \_|    < \  ___/ |  | \/
	//  |______  / \____//__/\_ \  \______  /|_______ \|___| \____|__  /(____  /|__|_ \ \___  >|__|
	// 		\/              \/         \/         \/              \/      \/      \/     \/ `)
	// fmt.Println(b1)

	wrap := box.NewBox().Width(2).Height(5).Color("Cyan").TitlePosition(box.Top).Style(box.Single).TitleColor("#427ef5").AllowWrapping(true)
	b1, _ := wrap.Render("Wrapping Works", strings.Repeat("Box CLI Maker 盒子製 造商,📦 ", 160))
	fmt.Println(b1)

}
