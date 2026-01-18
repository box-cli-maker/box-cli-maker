package main

import (
	"fmt"

	box "github.com/Delta456/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 0).
		Style(box.Single).
		Color(box.Green).
		ContentAlign(box.Center). // box.Left by default, change to box.Right or box.Center to align content
		WrapContent(true)         // Enable content wrapping, incase terminal width is small

	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed dignissim, arcu nec interdum faucibus, elit ante luctus erat, vitae malesuada lacus justo non risus.\n" +
		"Integer vestibulum, justo vitae cursus tristique, ligula nibh lacinia arcu, ac dictum ipsum mi nec erat. Vestibulum vel arcu ac justo finibus pulvinar. Suspendisse potenti.\n" +
		"Phasellus viverra sem vel est volutpat, sed bibendum erat malesuada. Nunc posuere, mauris in efficitur vehicula, nunc nisl malesuada lacus, id aliquam orci justo ut ante. Curabitur malesuada, justo nec rhoncus ultricies, ipsum nunc lacinia odio, sed tristique tortor mi vitae mi.\n" +
		"Ut volutpat, eros ut eleifend facilisis, nisi nisl facilisis risus, a accumsan tortor tellus non leo. Mauris in ipsum libero. Vestibulum scelerisque feugiat orci, id volutpat odio eleifend a. Mauris ornare dolor eu arcu faucibus, in aliquet risus laoreet.\n" +
		"Donec maximus, elit non facilisis interdum, leo ligula dignissim neque, ut ultrices sapien arcu id nulla. Suspendisse consequat orci sit amet risus ullamcorper, et ultrices quam efficitur."

	s, err := b.Render("Lorem Ipsum", content)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
