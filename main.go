package main

import (
	"io/ioutil"
	"fmt"

	"github.com/hoisie/web"
	"github.com/xyproto/browserspeak"
)

// TODO Add functions for building templates as well

// This is a test function
func testbuilder(cssurl string) *browserspeak.Page {
	page := browserspeak.NewHTML5Page("Hello")
	body, _ := page.SetMargin(3)

	h1 := body.AddNewTag("h1")
	h1.SetMargin(1)
	h1.AddContent("Browser")

	h1, err := page.root.GetTag("h1")
	if err == nil {
		h1.AddContent("Spe")
	}

	head, err := page.root.GetTag("head")
	if err == nil {
		link := head.AddNewTag("link")
		link.AddAttr("rel", "stylesheet")
		link.AddAttr("href", cssurl)
		link.AddAttr("type", "text/css")
		h1.AddContent("ak")
	} else {
		h1.AddContent("akkkkkkkk")
	}

	page.SetColor("#202020", "#A0A0A0")
	page.SetFont("sans serif")

	box, _ := page.AddBox("box0", true)
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
	return browserspeak.Message("root page", "hello: "+val)
}

func notFound(ctx *web.Context, val string) {
	ctx.NotFound("Page not found")
}

func hi() string {
	return "hi\n"
}

func onTheFlyRepo(ctx *web.Context) {
	fmt.Println("on the fly")
	ctx.SetHeader("Content-Type", "text/html; charset=utf-8", true)
	ctx.WriteString("hello")
	//ctx.ContentType("tgz")
	// writeTarGz(ctx.ResponseWriter)
}

func main() {
	Publish("/", "/main.css", testbuilder)

	web.Get("/error", errorlog)

	web.Get("/hi", hi)

	web.Get("/hello/(.*)", hello)

	web.Get("/archlinux/my/os/x86_64/my.db", onTheFlyRepo)

	web.Get("/(.*)", notFound)
	web.Run("0.0.0.0:8080")
}
