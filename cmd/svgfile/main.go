package main

import (
	"github.com/xyproto/onthefly"
)

// Generate a new SVG file

func main() {
	page, svg := onthefly.NewTinySVG(0, 0, 256, 256)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Diagram")

	rr := svg.AddRoundedRect(30, 10, 5, 5, 20, 20)
	rr.Fill2(onthefly.ColorByName("red"))
	page.SaveSVG("output.svg")
}
