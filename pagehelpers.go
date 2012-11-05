package main

import (
	"github.com/hoisie/web"
)

func NewHTML5Page(titleText string) *Page {
	page := NewPage(titleText, "<!DOCTYPE html>")
	html := page.root.AddNewTag("html")
	head := html.AddNewTag("head")
	title := head.AddNewTag("title")
	title.AddContent(titleText)
	html.AddNewTag("body")
	return page
}

// Get a function that returns a string that is the html for this page
func HTML(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		return page.GetHTML()
	}
}

// Get a function that returns a string that is the css for this page
func CSS(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		ctx.ContentType("css")
		return page.GetCSS()
	}
}

func (page *Page) SetMargin(tagname string, em int) (*Tag, error) {
	tag, err := page.root.GetTag(tagname)
	if err == nil {
		tag.SetMargin(em)
	}
	return tag, err
}

// Set one of the css styles of the <body>
func (page *Page) bodyAttr(key, value string) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddStyle(key, value)
	}
	return tag, err
}

func (page *Page) SetColor(fgColor string, bgColor string) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddStyle("color", fgColor)
		tag.AddStyle("background-color", bgColor)
	}
	return tag, err
}

func (page *Page) SetFont(fontFamily string) (*Tag, error) {
	return page.bodyAttr("font-family", fontFamily)
}

// Add a box
func (page *Page) AddBox(id string, rounded bool) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddBox(id, rounded, "0.9em", "HEOI HEOI", "green", "black", "3em")
	}
	return tag, err
}
