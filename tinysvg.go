package browserspeak

import (
	"strconv"
	"fmt"
)

func NewTinySVG(x, y, w, h int) (*Page, *Tag) {
	page := NewPage("", "<?xml version=\"1.0\"?>")
	svg := page.root.AddNewTag("svg")
	svg.AddAttr("xmlns", "http://www.w3.org/2000/svg")
	svg.AddAttr("version", "1.2")
	svg.AddAttr("baseProfile", "tiny")
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sw := strconv.Itoa(w)
	sh := strconv.Itoa(h)
	svg.AddAttr("viewBox", sx + " " + sy + " " + sw + " " + sh)
	return page, svg
}

func (svg *Tag) AddRect(x, y, w, h int) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sw := strconv.Itoa(w)
	sh := strconv.Itoa(h)
	rect := svg.AddNewTag("rect")
	rect.AddAttr("x", sx)
	rect.AddAttr("y", sy)
	rect.AddAttr("width", sw)
	rect.AddAttr("height", sh)
	return rect
}

func (svg *Tag) Circle(x, y, radius int, color string) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sradius := strconv.Itoa(radius)
	circle := svg.AddNewTag("circle")
	circle.AddAttr("cx", sx)
	circle.AddAttr("cy", sy)
	circle.AddAttr("r", sradius)
	circle.Fill(color)
	return circle
}

func (svg *Tag) AddCircle(x, y, radius int) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sradius := strconv.Itoa(radius)
	circle := svg.AddNewTag("circle")
	circle.AddAttr("cx", sx)
	circle.AddAttr("cy", sy)
	circle.AddAttr("r", sradius)
	return circle
}

func (rect *Tag) Fill(color string) {
	rect.AddAttr("fill", color)
}

// Converts r, g and b which are integers in the range from 0..255
// to a color-string on the form "#nnnnnn".
func ColorString(r, g, b int) string {
	rs := strconv.FormatInt(int64(r), 16)
	gs := strconv.FormatInt(int64(g), 16)
	bs := strconv.FormatInt(int64(b), 16)
	if len(rs) == 1 {
		rs = "0" + rs
	}
	if len(gs) == 1 {
		gs = "0" + gs
	}
	if len(bs) == 1 {
		bs = "0" + bs
	}
	return "#" + rs + gs + bs
}

// Creates a rectangle that is 1 wide with the given color
// Note that the size of the "pixel" depends on how large the viewBox is
func (svg *Tag) Pixel(x, y, r, g, b int) *Tag {
	color := ColorString(r, g, b)
	rect := svg.AddRect(x, y, 1, 1)
	rect.Fill(color)
	return rect
}

func (svg *Tag) AlphaDot(cx, cy, r, g, b int, a float32) *Tag {
	color := fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, a)
	circle := svg.AddCircle(cx, cy, 1)
	circle.Fill(color)
	return circle
}

func (svg *Tag) Dot(cx, cy, r, g, b int) *Tag {
	color := ColorString(r, g, b)
	circle := svg.AddCircle(cx, cy, 1)
	circle.Fill(color)
	return circle
}

func NewExampleSVG() *Page {
	page, svg := NewTinySVG(0, 0, 30, 30)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Example SVG file")
	rect := svg.AddRect(10, 10, 10, 10)
	rect.Fill("green")
	svg.Pixel(10, 10, 255, 0, 0)
	svg.AlphaDot(5, 5, 0, 0, 255, 0.5)
	return page
}

func NewExampleSVG2() *Page {
	w := 160
	h := 90
	stepx := 8
	stepy := 8
	page, svg := NewTinySVG(0, 0, w, h)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Example SVG file")
	increase := 0
	decrease := 0
	for y := stepy; y < h; y += stepy {
		for x := stepx; x < w; x += stepx {
			increase = int((float32(x) / float32(w)) * 255.0)
			decrease = 255 - increase
			svg.Dot(x, y, 255, decrease, increase)
		}
	}
	return page
}
