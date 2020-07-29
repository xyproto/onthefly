package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/xyproto/onthefly"
)

const (
	angularVersion = "1.6.9"
)

// Generate a new onthefly Page (HTML5, Angular and CSS combined)
func indexPage() *onthefly.Page {

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewAngularPage("Demonstration", angularVersion)

	// Rely on the body tag being present
	body, _ := page.GetTag("body")

	// Add a title paragraph
	title := body.AddNewTag("p")
	// Use id attributes to style similar tags separately
	title.AddAttrib("id", "title")
	title.AddContent(fmt.Sprintf("onthefly %.1f and AngularJS %s", onthefly.Version, angularVersion))
	title.AddStyle("font-size", "2em")
	title.AddStyle("font-family", "sans-serif")
	title.AddStyle("font-style", "italic")

	// Add a paragraph for the angular related tags
	angularp := body.AddNewTag("p")
	angularp.AddAttrib("id", "angular")
	angularp.AddStyle("margin-top", "2em")

	// Label for the input box
	label := angularp.AddNewTag("label")
	inputID := "input1"
	label.AddAttrib("for", inputID)
	label.AddContent("Enter text:")
	label.AddStyle("margin-right", "3em")

	// Angular input
	input := angularp.AddNewTag("input")
	input.AddAttrib("id", inputID)
	input.AddAttrib("type", "text")
	dataBindingName := "sometext"
	input.AddAttrib("ng-model", dataBindingName)

	// Angular output
	h1 := angularp.AddNewTag("h1")
	h1.AddAttrib("ng-show", dataBindingName)
	h1.AddContent("HI {{ " + dataBindingName + " | uppercase }}")
	h1.AddStyle("color", "blue")
	h1.AddStyle("margin", "2em")
	h1.AddStyle("font-size", "3em")
	h1.AddStyle("font-family", "courier")

	// Set the margin (em is default)
	page.SetMargin(4)

	// Set the font family
	page.SetFontFamily("serif")

	// Set the color scheme (fg, bg)
	page.SetColor("black", "#e0e0e0")

	return page
}

// Set up the paths and handlers then start serving.
func main() {
	fmt.Println("onthefly", onthefly.Version)

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	// Publish the generated Page in a way that connects the HTML and CSS. Cached.
	indexPage().Publish(mux, "/", "/style.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":3000")
}
