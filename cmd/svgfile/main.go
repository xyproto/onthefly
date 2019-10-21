package main

import (
	"github.com/xyproto/tinysvg"
)

// Generate a new SVG file

func main() {
	document, svg := tinysvg.NewTinySVG(256, 256)
	svg.Describe("Diagram")

	rr := svg.AddRoundedRect(30, 10, 5, 5, 20, 20)
	rr.Fill("red")

	document.SaveSVG("output.svg")
}
