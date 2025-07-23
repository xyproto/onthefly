// Package main demonstrates Angular.JS usage with onthefly
package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/xyproto/onthefly"
)

// Generate a new onthefly Page (HTML5, AngularJS and CSS combined)
func indexPage() *onthefly.Page {

	// Create a new HTML5 page with AngularJS included
	page := onthefly.NewAngularPage("onthefly with AngularJS")

	// Add ng-app with controller name to html tag
	html, _ := page.GetTag("html")
	html.AddAttrib("ng-app", "demoApp")
	html.AddAttrib("ng-controller", "DemoController")

	// Rely on the body tag being present
	body, _ := page.GetTag("body")

	// Add a title paragraph
	title := body.AddNewTag("h1")
	title.AddAttrib("id", "title")
	title.AddContent(fmt.Sprintf("onthefly %.1f with AngularJS", onthefly.Version))
	title.AddStyle("font-size", "2.5em")
	title.AddStyle("font-family", "Arial, sans-serif")
	title.AddStyle("color", "#333")
	title.AddStyle("text-align", "center")
	title.AddStyle("margin-bottom", "2em")

	// Create a container div
	container := body.AddNewTag("div")
	container.AddAttrib("class", "container")
	container.AddStyle("max-width", "800px")
	container.AddStyle("margin", "0 auto")
	container.AddStyle("padding", "2em")

	// Add input section
	inputSection := container.AddNewTag("div")
	inputSection.AddAttrib("class", "input-section")
	inputSection.AddStyle("margin-bottom", "2em")

	// Label for the input box
	label := inputSection.AddNewTag("label")
	label.AddAttrib("for", "textInput")
	label.AddContent("Enter your message:")
	label.AddStyle("display", "block")
	label.AddStyle("margin-bottom", "0.5em")
	label.AddStyle("font-weight", "bold")
	label.AddStyle("color", "#555")

	// AngularJS input with improved styling
	input := inputSection.AddNewTag("input")
	input.AddAttrib("id", "textInput")
	input.AddAttrib("type", "text")
	input.AddAttrib("ng-model", "userMessage")
	input.AddAttrib("placeholder", "Type something here...")
	input.AddStyle("width", "100%")
	input.AddStyle("padding", "12px")
	input.AddStyle("font-size", "16px")
	input.AddStyle("border", "2px solid #ddd")
	input.AddStyle("border-radius", "8px")
	input.AddStyle("box-sizing", "border-box")
	input.AddStyle("transition", "border-color 0.3s")

	// Add CSS for input focus
	page.AddStyle("input:focus { border-color: #4CAF50; outline: none; }")

	// Output section
	outputSection := container.AddNewTag("div")
	outputSection.AddAttrib("class", "output-section")
	outputSection.AddAttrib("ng-show", "userMessage && userMessage.length > 0")

	// Display message in different formats
	formats := []struct {
		title string
		expr  string
		style map[string]string
	}{
		{"Uppercase:", "{{ userMessage | uppercase }}", map[string]string{"color": "#2196F3", "font-weight": "bold"}},
		{"Lowercase:", "{{ userMessage | lowercase }}", map[string]string{"color": "#FF9800", "font-style": "italic"}},
		{"Character count:", "{{ userMessage.length }} characters", map[string]string{"color": "#4CAF50", "font-size": "14px"}},
		{"Reversed:", "{{ userMessage.split('').reverse().join('') }}", map[string]string{"color": "#9C27B0", "font-family": "monospace"}},
	}

	for _, format := range formats {
		formatDiv := outputSection.AddNewTag("div")
		formatDiv.AddStyle("margin", "1em 0")
		formatDiv.AddStyle("padding", "1em")
		formatDiv.AddStyle("background-color", "#f9f9f9")
		formatDiv.AddStyle("border-radius", "6px")
		formatDiv.AddStyle("border", "1px solid #ddd")

		titleSpan := formatDiv.AddNewTag("strong")
		titleSpan.AddContent(format.title + " ")
		titleSpan.AddStyle("color", "#333")

		valueSpan := formatDiv.AddNewTag("span")
		valueSpan.AddContent(format.expr)
		for key, value := range format.style {
			valueSpan.AddStyle(key, value)
		}
	}

	// Add AngularJS controller script
	script := body.AddNewTag("script")
	script.AddContent(`
angular.module('demoApp', [])
.controller('DemoController', function($scope) {
    $scope.userMessage = '';

    // Add a method to reverse string (since AngularJS doesn't have built-in reverse filter)
    $scope.reverseString = function(str) {
        return str ? str.split('').reverse().join('') : '';
    };
});`)

	// Set overall page styling
	page.SetMargin(0)
	page.SetFontFamily("Arial, sans-serif")
	page.SetColor("#333", "#ffffff")

	// Add responsive CSS
	page.AddStyle(`
@media (max-width: 600px) {
    .container { padding: 1em; }
    h1 { font-size: 1.8em !important; }
}`)

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
