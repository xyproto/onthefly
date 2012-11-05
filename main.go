package main

import (
	"io/ioutil"

	"github.com/hoisie/web"
)

// TODO Change this to buildtemplate instead
func buildpage(cssurl string) *Page {
	page := NewHTML5Page("Hello")
	body, _ := page.SetMargin("body", 3)

	h1 := body.AddNewTag("h1")
	h1.SetMargin(1)
	h1.AddContent("hello")

	h1, err := page.root.GetTag("h1")
	if err == nil {
		h1.AddContent("OSTEBOLLEEEEEEyyyy")
	}

	head, err := page.root.GetTag("head")
	if err == nil {
		link := head.AddNewTag("link")
		link.AddAttr("rel", "stylesheet")
		link.AddAttr("href", cssurl)
		link.AddAttr("type", "text/css")
		h1.AddContent("FOUND")
	} else {
		h1.AddContent("NOT FOUND")
	}

	page.SetColor("#202020", "#A0A0A0")
	page.SetFont("sans serif")

	page.AddBox("box0", true)

	return page
}

// func rootpage(w http.ResponseWriter, r *http.Request) {
// r.URL.Path[1:]
//fmt.Fprintf(w, message("root page", "hello: " + r.URL.RawQuery + " - " + r.URL.Host + " - " + html.EscapeString(r.URL.Path[1:])))
//}

func errorlog() string {
	data, err := ioutil.ReadFile("error.log")
	if err != nil {
		return "No errors\n"
	}
	return "Errors:\n" + string(data) + "\n"
}

func hello(val string) string {
	return message("root page", "hello: "+val)
}

func notFound(ctx *web.Context, message string) {
	ctx.NotFound("Page not found")
}

func main() {

	cssurl := "/main.css"
	page := buildpage(cssurl)

	web.Get("/error", errorlog)

	web.Get("/", HTML(page))
	web.Get(cssurl, CSS(page))

	web.Get("/hello", hello)
	web.Get("/(.*)", notFound)
	web.Run("0.0.0.0:8080")

}
