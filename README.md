Browserspeak
============

[![Build Status](https://travis-ci.org/xyproto/browserspeak.svg?branch=master)](https://travis-ci.org/xyproto/browserspeak)
[![Build Status](https://drone.io/github.com/xyproto/browserspeak/status.png)](https://drone.io/github.com/xyproto/browserspeak/latest)

* Package for generating SVG (TinySVG) on the fly
* Can also be used for generating HTML and CSS, or templates for HTML and CSS
* The recommended use is mainly for generating SVG files
* Could be used to set up a diskless webserver that generates all the content
  (which could also be done with templates)

Online API Documentation
------------------------

[godoc.org](http://godoc.org/github.com/xyproto/browserspeak)

Generate content on the fly
---------------------------

<img src="https://raw.github.com/xyproto/browserspeak/master/img/inbrowser.png">

Example with [Negroni](https://github.com/codegangsta/negroni)
--------------------
~~~ go
package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/browserspeak"
)

// Generate a new SVG Page
func svgPage() *browserspeak.Page {
	page, svg := browserspeak.NewTinySVG(0, 0, 128, 64)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Hello SVG")
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")
	return page
}

// Generate a new browserspeak Page (HTML5 and CSS combined)
func indexPage(svgurl string) *browserspeak.Page {

	// Create a new HTML5 page, with CSS included
	page := browserspeak.NewHTML5Page("Demonstration")

	// Add some text
	page.AddContent(fmt.Sprintf("browserspeak %.1f", browserspeak.Version))

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("serif") // sans serif

	// Change the color scheme
	page.SetColor("black", "#d0d0d0")

	// Include the generated SVG image on the page
	body, err := page.GetTag("body")
	if err == nil {
		// CSS attributes for the body tag
		body.AddStyle("font-size", "2em")

		// Paragraph
		p := body.AddNewTag("p")

		// CSS style
		p.AddStyle("margin-top", "2em")

		// Image tag
		img := p.AddNewTag("img")

		// HTML attributes
		img.AddAttrib("src", svgurl)
		img.AddAttrib("alt", "Three circles")

		// CSS style
		img.AddStyle("width", "60%")
		img.AddStyle("border", "4px solid white")
	}

	return page
}

// Set up the paths and handlers then start serving.
func main() {
	fmt.Println("browserspeak ", browserspeak.Version)

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	// Publish the generated SVG as "/circles.svg"
	svgurl := "/circles.svg"
	mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		fmt.Fprint(w, svgPage().String())
	})

	// Generate a Page that includes the svg image
	page := indexPage(svgurl)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, "/", "/style.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 8080
	n.Run(":8080")
}
~~~

Example with [web.go](https://github.com/hoisie/web.go)
--------------------
~~~ go
package main

import (
	"fmt"

	"github.com/hoisie/web"
	"github.com/xyproto/browserspeak"
	"github.com/xyproto/webhandle"
)

// Generate a new SVG Page
func svgPage() *browserspeak.Page {
	page, svg := browserspeak.NewTinySVG(0, 0, 128, 64)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Hello SVG")
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")
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

// Generate a new browserspeak Page (HTML5 and CSS)
func indexPage(cssurl string) *browserspeak.Page {
	page := browserspeak.NewHTML5Page("Demonstration")

	// Link the page to the css file generated from this page
	page.LinkToCSS(cssurl)

	// Add some text
	page.AddContent(fmt.Sprintf("browserspeak %.1f", browserspeak.Version))

	// Change the margin (em is default)
	page.SetMargin(7)

	// Change the font family
	page.SetFontFamily("serif") // sans serif

	// Change the color scheme
	page.SetColor("#f02020", "#101010")

	// Include the generated SVG image on the page
	body, err := page.GetTag("body")
	if err == nil {
		// CSS attributes for the body tag
		body.AddStyle("font-size", "2em")

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

// Set up the paths and handlers then start serving.
func main() {
	fmt.Println("browserspeak ", browserspeak.Version)

	// Connect the url for the HTML and CSS with the HTML and CSS generated from indexPage
	webhandle.PublishPage("/", "/style.css", indexPage)

	// Connect /circles.svg with the generated handle
	web.Get("/circles.svg", svgHandlerGenerator())

	// Run the web server at port 8080
	web.Run(":8080")
}
~~~

Version, license and author
---------------------------

* Version: 0.8
* License: MIT
* Alexander Rødseth

