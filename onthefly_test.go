package onthefly

import (
	"strings"
	"testing"
)

func TestSVG(t *testing.T) {
	SampleSVG1()
}

func TestSVG2(t *testing.T) {
	SampleSVG2()
}

func TestGen(t *testing.T) {
	page := NewHTML5Page("Hello Title")
	page.LinkToCSS("/test.css")
	page.AddContent("hello body")

	// Add a link to /test.svg
	body, err := page.GetTag("body")
	if err == nil {
		body.AddNewTag("br")
		//	a := body.AddNewTag("a")
		//	a.AddAttrib("href", "/test.svg")
		//	a.AddContent("See SVG")
	}

	p := *page
	p.prettyPrint()

	s := page.String()

	if strings.Count(s, "hello body") > 1 {
		t.Errorf("Error, text appears more than once!\n%s\n", s)
	}

}
