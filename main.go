package browserspeak

import (
	"io/ioutil"
    "github.com/go-martini/martini"
	"github.com/martini-contrib/web"
	"github.com/xyproto/instapage"
)

// This is a test function
func testbuilder(cssurl string) *Page {
	page := NewHTML5Page("Hello")
	body, _ := page.SetMargin(3)

	h1 := body.AddNewTag("h1")
	h1.SetMargin(1)
	h1.AddContent("Browser")

	h1, err := page.root.GetTag("h1")
	if err == nil {
		h1.AddContent("Spe")
	}

	if err := page.LinkToCSS(cssurl); err == nil {
		h1.AddContent("ak")
	} else {
		h1.AddContent("akkkkkkkk")
	}

	page.SetColor("#202020", "#A0A0A0")
	page.SetFontFamily("sans serif")

	box, _ := page.addBox("box0", true)
	box.AddStyle("margin-top", "-2em")
	box.AddStyle("margin-bottom", "3em")

	image := body.AddImage("http://www.shoutmeloud.com/wp-content/uploads/2010/01/successful-Blogger.jpeg", "50%")
	image.AddStyle("margin-top", "2em")
	image.AddStyle("margin-left", "3em")

	return page
}

func errorlog() string {
	data, err := ioutil.ReadFile("error.log")
	if err != nil {
		return "No errors\n"
	}
	return "Errors:\n" + string(data) + "\n"
}

func hello(val string) string {
	return instapage.Message("root page", "hello: "+val)
}

func notFound(ctx *web.Context, message string) {
	ctx.NotFound("Page not found")
}

func exampleSVG() string {
	svg := newExampleSVG()
	return svg.String()
}

func main() {
	m := martini.Classic()
    m.Use(web.ContextWithCookieSecret(""))

	PublishPage(m, "/", "/main.css", testbuilder)

	m.Get("/error", errorlog)
	m.Get("/hello", hello)

	m.Get("/svg", exampleSVG)

	m.Get("/(.*)", notFound)

	m.Run() // "0.0.0.0:8080"
}
