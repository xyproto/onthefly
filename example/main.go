package main

import (
	"fmt"

	"github.com/hoisie/web"
	"github.com/xyproto/browserspeak"
)

// Generate a new Page (HTML with CSS)
func helloPage(cssurl string) *browserspeak.Page {
	page := browserspeak.NewHTML5Page("Hello Title")

	// Link the page to the css file generated from this page
	page.LinkToCSS(cssurl)

	// Add some text
	page.AddContent("hello body")

	// Change the margin
	page.SetMargin(3)

	// Change the font family
	page.SetFontFamily("sans serif")

	// Change the color scheme
	page.SetColor("#202020", "#a0a0a0")

	// Add a link to /test.svg
	body, err := page.GetTag("body")
	if err == nil {
		a := body.AddNewTag("a")
		a.AddAttr("href", "/test.svg")
		a.AddContent("See SVG")
	}

	return page
}

// Generate a new Page (SVG)
func svgPage() *browserspeak.Page {
	page, svg := browserspeak.NewTinySVG(0, 0, 128, 64)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Hello SVG")
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")
	return page
}

// Get the string from the SVG Page
func svgGenerator() string {
	page := svgPage()
	return page.String()
}

func main() {
	fmt.Println("BrowserSpeak Version:", browserspeak.Version)

	// Connect the url for the HTML and CSS with the HTML and CSS generated from helloPage
	browserspeak.PublishPage("/", "/hello.css", helloPage)

	// Connect /test.svg with svgGenerator
	web.Get("/test.svg", svgGenerator)

	// Run the web server at port 8080
	web.Run("0.0.0.0:8080")
}
