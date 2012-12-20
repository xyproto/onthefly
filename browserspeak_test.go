package browserspeak

import (
	"testing"
	"strings"
)

func TestSVG(t *testing.T) {
	svg := newExampleSVG()
	svg.String()
	//s := svg.GetXML(false)
	t.Log("hi")
	//t.Errorf("%s\n", s)
	//const in, out = 4, 2
	//if x := Sqrt(in); x != out {
	//	t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	//}
}

func TestSVG2(t *testing.T) {
	svg := newExampleSVG2()
	svg.String()
	//s := svg.String()
	//t.Errorf("%s\n", s)
}

func TestGen(t *testing.T) {
	page := NewHTML5Page("Hello Title")
	page.LinkCSS("/test.css")
	page.AddContent("hello body")

	// Add a link to /test.svg
	body, err := page.GetTag("body")
	if err == nil {
		body.AddNewTag("br")
	//	a := body.AddNewTag("a")
	//	a.AddAttr("href", "/test.svg")
	//	a.AddContent("See SVG")
	}

	if strings.Count(body.GetContent(), "hello body") > 1 {
		t.Errorf("Error, text appears more than once in body content!\n")
	}

	p := *page
	p.prettyPrint()

	s := page.String()

	if strings.Count(s, "hello body") > 1 {
		t.Errorf("Error, text appears more than once!\n%s\n", s)
	}

}
