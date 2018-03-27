package main

import (
	"fmt"

	"github.com/hoisie/web"
	"github.com/xyproto/onthefly"
	"github.com/xyproto/webhandle"
)

// Generate a new SVG Page
func svgPage() *onthefly.Page {
	page, svg := onthefly.NewTinySVG(0, 0, 128, 64)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Hello SVG")

	// x, y, radius, color
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")

	// x, y, font size, font family, text and color
	svg.Text(3, 60, 6, "Courier", "There will be cake", "#394851")

	return page
}

// Generator for a handle that returns the generated SVG content.
// Also sets the content type.
func svgHandlerGenerator() func(ctx *web.Context) string {
	page := svgPage()
	return func(ctx *web.Context) string {
		ctx.ContentType("image/svg+xml")
		return page.String()
	}
}

// Generate a new onthefly Page (HTML5 and CSS)
func indexPage(cssurl string) *onthefly.Page {
	page := onthefly.NewHTML5Page("Demonstration")

	// Link the page to the css file generated from this page
	page.LinkToCSS(cssurl)

	// Add some text
	page.AddContent(fmt.Sprintf("onthefly %.1f", onthefly.Version))

	// Change the margin (em is default)
	page.SetMargin(7)

	// Change the font family

	// Change the color scheme
	page.SetColor("#f02020", "#101010")

	// Include the generated SVG image on the page
	body, err := page.GetTag("body")
	if err == nil {
		// CSS attributes for the body tag
		body.AddStyle("font-size", "2em")
		body.AddStyle("font-family", "sans-serif")

		// Paragraph
		p := body.AddNewTag("p")

		// CSS style
		p.AddStyle("margin-top", "2em")

		// Image tag
		img := p.AddNewTag("img")

		// HTML attributes
		img.AddAttrib("src", "/circles.svg")
		img.AddAttrib("alt", "Three circles")

		// CSS style
		img.AddStyle("width", "60%")
	}

	return page
}

func main() {
	fmt.Println("onthefly ", onthefly.Version)

	// Connect the url for the HTML and CSS with the HTML and CSS generated from indexPage
	webhandle.PublishPage("/", "/style.css", indexPage)

	// Connect /circles.svg with the generated handle
	web.Get("/circles.svg", svgHandlerGenerator())

	// Listen for requests at port 3000
	web.Run(":3000")
}
