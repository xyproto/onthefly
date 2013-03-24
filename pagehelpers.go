package browserspeak

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/drbawb/mustache"
	"github.com/xyproto/web"
)

var globalStringCache map[string]string

// Create a blank HTML5 page
func NewHTML5Page(titleText string) *Page {
	page := NewPage(titleText, "<!DOCTYPE html>")
	html := page.root.AddNewTag("html")
	head := html.AddNewTag("head")
	title := head.AddNewTag("title")
	title.AddContent(titleText)
	html.AddNewTag("body")
	return page
}

// Create a web.go compatible function that returns a string that is the HTML for this page
func GenerateHTML(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		return page.GetXML(true)
	}
}

// Create a web.go compatible function that returns a string that is the HTML for this page
func GenerateHTMLwithTemplate(page *Page, values map[string]string) func(*web.Context) string {
	return func(ctx *web.Context) string {
		return mustache.Render(page.GetXML(true), values)
	}
}

// Create a web.go compatible function that returns a string that is the CSS for this page
func GenerateCSS(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		ctx.ContentType("css")
		return page.GetCSS()
	}
}

// Create a web.go compatible function that returns a string that is the XML for this page
func GenerateXML(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		ctx.ContentType("xml")
		return page.GetXML(false)
	}
}

// Set the margins of the body
func (page *Page) SetMargin(em int) (*Tag, error) {
	value := strconv.Itoa(em) + "em"
	return page.bodyAttr("margin", value)
}

// Set one of the CSS styles of the body
func (page *Page) bodyAttr(key, value string) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddStyle(key, value)
	}
	return tag, err
}

// Set the foreground and background color of the body
func (page *Page) SetColor(fgColor string, bgColor string) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddStyle("color", fgColor)
		tag.AddStyle("background-color", bgColor)
	}
	return tag, err
}

// Set the font family
func (page *Page) SetFontFamily(fontFamily string) (*Tag, error) {
	return page.bodyAttr("font-family", fontFamily)
}

// Add a box, for testing
func (page *Page) addBox(id string, rounded bool) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		return tag.AddBox(id, rounded, "0.9em", "Speaks browser so you don't have to", "white", "black", "3em"), nil
	}
	return tag, err
}

// Creates a HTML page with a title that displays a message
// Can be used for debugging
func Message(title, msg string) string {
	return "<!DOCTYPE html><html><head><title>" + title + "</title></head><body style=\"margin:4em; font-family:courier; color:gray; background-color:light gray;\"><h2>" + title + "</h2><hr style=\"margin-top:-1em; margin-bottom: 2em; margin-right: 20%; text-align: left; border: 1px dotted #b0b0b0; height:1px;\"><div style=\"margin-left: 2em;\">" + msg + "</div></body></html>"
}

// Link a page up with a CSS file
// Takes the url to a CSS file as a string
// The given page must have a "head" tag for this to work
// Returns an error if no "head" tag is found, or nil
func (page *Page) LinkToCSS(cssurl string) error {
	head, err := page.GetTag("head")
	if err == nil {
		link := head.AddNewTag("link")
		link.AddAttr("rel", "stylesheet")
		link.AddAttr("href", cssurl)
		link.AddAttr("type", "text/css")
	}
	return err
}

// Link a page up with a JS file
// Takes the url to a JS file as a string
// The given page must have a "head" tag for this to work
// Returns an error if no "head" tag is found, or nil
func (page *Page) LinkToJS(jsurl string) error {
	head, err := page.GetTag("head")
	if err == nil {
		src := head.AddNewTag("script")
		src.AddAttr("src", jsurl)
		src.AddAttr("type", "text/javascript")
		src.AddContent(" ")
	}
	return err
}

// Link a page up with a Favicon file
// Takes the url to a favicon file as a string
// The given page must have a "head" tag for this to work
// Returns an error if no "head" tag is found, or nil
func (page *Page) LinkToFavicon(favurl string) error {
	head, err := page.root.GetTag("head")
	if err == nil {
		link := head.AddNewTag("link")
		link.AddAttr("rel", "shortcut icon")
		link.AddAttr("href", favurl)
	}
	return err
}

// Takes a charset, for example UTF-8, and creates a <meta> tag in <head>
func (page *Page) MetaCharset(charset string) error {
	// Add a meta tag
	head, err := page.GetTag("head")
	if err == nil {
		meta := head.AddNewTag("meta")
		meta.AddAttr("http-equiv", "Content-Type")
		meta.AddAttr("content", "text/html; charset="+charset)
	}
	return err
}

// Link to Google Fonts
func (page *Page) LinkToGoogleFont(name string) error {
	url := "http://fonts.googleapis.com/css?family="
	// Replace space with +, if needed
	if strings.Contains(name, " ") {
		url += strings.Replace(name, " ", "+", -1)
	} else {
		url += name
	}
	// Link to the CSS for the given font name
	return page.LinkToCSS(url)
}

// Creates a page based on the contents of "error.log". Useful for showing compile errors while creating an application.
func Errorlog() string {
	data, err := ioutil.ReadFile("error.log")
	if err != nil {
		return Message("Good", "No errors")
	}
	errors := strings.Replace(string(data), "\n", "</br>", -1)
	return Message("Errors", errors)
}

// Handles pages that are not found
func NotFound(ctx *web.Context, val string) string {
	ctx.NotFound(Message("No", "Page not found"))
	return ""
}

// Takes a filename and returns a function that can handle the request
func File(filename string) func(ctx *web.Context) string {
	var extension string
	if strings.Contains(filename, ".") {
		extension = filepath.Ext(filename)
	}
	return func(ctx *web.Context) string {
		if extension != "" {
			ctx.ContentType(extension)
		}
		imagebytes, _ := ioutil.ReadFile(filename)
		buf := bytes.NewBuffer(imagebytes)
		return buf.String()
	}
}

// Takes an url and a filename and offers that file at the given url
func PublishFile(url, filename string) {
	web.Get(url, File(filename))
}

// Takes a filename and offers that file at the root url
func PublishRootFile(filename string) {
	web.Get("/"+filename, File(filename))
}

// Expose the HTML and CSS generated by a page building function to the two given urls
func PublishPage(htmlurl, cssurl string, buildfunction (func(string) *Page)) {
	page := buildfunction(cssurl)
	web.Get(htmlurl, GenerateHTML(page))
	web.Get(cssurl, GenerateCSS(page))
}

func AddHeader(page *Page, js string) {
	page.MetaCharset("UTF-8")
	AddScriptToHeader(page, js)
}

// Wrap a SimpleContextHandle so that the output is cached (with an id)
// Do not cache functions with side-effects! (that sets the mimetype for instance)
// The safest thing for now is to only cache images.
func CacheWrapper(id string, f SimpleContextHandle) SimpleContextHandle {
	return func(ctx *web.Context) string {
		if _, ok := globalStringCache[id]; !ok {
			globalStringCache[id] = f(ctx)
		}
		return globalStringCache[id]
	}
}

func Publish(url, filename string, cache bool) {
	if cache {
		web.Get(url, CacheWrapper(url, File(filename)))
	} else {
		web.Get(url, File(filename))
	}
}


