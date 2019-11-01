On The Fly
==========

[![Build Status](https://travis-ci.org/xyproto/onthefly.svg?branch=master)](https://travis-ci.org/xyproto/onthefly)
[![GoDoc](https://godoc.org/github.com/xyproto/onthefly?status.svg)](http://godoc.org/github.com/xyproto/onthefly)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/xyproto/onthefly/master/LICENSE)
[![Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat)](http://goreportcard.com/report/xyproto/onthefly)


* Package for generating HTML and CSS together, on the fly.
* Can also be used for generating HTML, XML or CSS (or templates).
* HTML and CSS can be generated together, but be presented as two seperate (but linked) files.
* Could be used to set up a diskless webserver that generates all the content on the fly
  (something similar could also be achieved by using templates that are not stored on disk).

New/experimental features:
* Generating WebGL graphics with Three.JS on the fly.
* Generating AngularJS applications on the fly.

The SVG related code has been refactored out, optimized and moved to the [tinysvg](https://github.com/xyproto/tinysvg) package.

Online API Documentation
------------------------

[godoc.org](http://godoc.org/github.com/xyproto/onthefly)

Generate HTML, CSS and SVG in one go
------------------------------------
<img src="https://raw.github.com/xyproto/onthefly/master/img/onthefly.png">

Create hardware accelerated 3D-graphics with [three.js](http://threejs.org/)
-----------------------------------------------------
<img src="https://raw.github.com/xyproto/onthefly/master/img/threejs.png">

Experiment with [AngularJS](https://angularjs.org/)
-----------------------
<img src="https://raw.github.com/xyproto/onthefly/master/img/angular.png">

Example for [Negroni](https://github.com/urfave/negroni)
--------------------

~~~go
package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/xyproto/onthefly"
	"github.com/xyproto/tinysvg"
)

// Generate a new SVG image
func svgImage() []byte {
	document, svg := tinysvg.NewTinySVG(128, 64)
	svg.Describe("Hello SVG")

	// x, y, radius, color
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")

	// x, y, font size, font family, text and color
	svg.Text(3, 60, 6, "Courier", "There will be cake", "#394851")

	return document.Bytes()
}

// Generate a new onthefly Page (HTML5 and CSS combined)
func indexPage(svgurl string) *onthefly.Page {

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page("Demonstration")

	// Add some text
	page.AddContent(fmt.Sprintf("onthefly %.1f", onthefly.Version))

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("serif") // or: sans-serif

	// Change the color scheme
	page.SetColor("black", "#d0d0d0")

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
	fmt.Println("onthefly ", onthefly.Version)

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	// Publish the generated SVG as "/circles.svg"
	svgurl := "/circles.svg"
	mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		w.Write(svgImage())
	})

	// Generate a Page that includes the svg image
	page := indexPage(svgurl)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, "/", "/style.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":3000")
}
~~~

Example for `net/http`
--------------------

~~~go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xyproto/onthefly"
	"github.com/xyproto/tinysvg"
)

// Generate a new SVG image
func svgImage() []byte {
	document, svg := tinysvg.NewTinySVG(128, 64)
	svg.Describe("Hello SVG")

	// x, y, radius, color
	svg.Circle(30, 10, 5, "red")
	svg.Circle(110, 30, 2, "green")
	svg.Circle(80, 40, 7, "blue")

	// x, y, font size, font family, text and color
	svg.Text(3, 60, 6, "Courier", "There will be cake", "#394851")

	return document.Bytes()
}

// Generate a new onthefly Page (HTML5 and CSS combined)
func indexPage(svgurl string) *onthefly.Page {

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page("Demonstration")

	// Add some text
	page.AddContent(fmt.Sprintf("onthefly %.1f", onthefly.Version))

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("serif") // or: sans-serif

	// Change the color scheme
	page.SetColor("black", "#d0d0d0")

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
	fmt.Println("onthefly ", onthefly.Version)

	// Create a mux
	mux := http.NewServeMux()

	// Publish the generated SVG as "/circles.svg"
	svgurl := "/circles.svg"
	mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		w.Write(svgImage())
	})

	// Generate a Page that includes the svg image
	page := indexPage(svgurl)

	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, "/", "/style.css", false)

	// Configure the HTTP server and permissionHandler struct
	s := &http.Server{
		Addr:           ":3000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening for requests on port 3000")

	// Start listening
	log.Fatal(s.ListenAndServe())
}
~~~

Additional screenshots
----------------------

Screenshots from version 0.8:

<img src="https://raw.github.com/xyproto/onthefly/master/img/svg_dark.png">

<img src="https://raw.github.com/xyproto/onthefly/master/img/svg_light.png">

TODO
----

* Create a version 2 that is more focused on performance and has more consistent API function names.

Version, license and author
---------------------------

* Version: 1.2.2
* Alexander F Rødseth &lt;xyproto@archlinux.org&gt;
* License: MIT

