package main

import (
	goml "goml/pkg"
)

func main() {
	p := goml.Text(goml.Tags{}, "hi")
	h1 := goml.Header(goml.Tags{}, goml.Text(goml.Tags{}, "header is here!!"))

	button := goml.Button(goml.Tags{}, "hi")

	options := goml.RenderParams{}
	render := goml.Render([]goml.Element{h1, p, button}, options)
	render.Write()
}
