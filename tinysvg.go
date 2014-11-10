package onthefly

import (
	"fmt"
	"strconv"
)

func NewTinySVG(x, y, w, h int) (*Page, *Tag) {
	page := NewPage("", `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Tiny//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-tiny.dtd">`)
	svg := page.root.AddNewTag("svg")
	svg.AddAttrib("xmlns", "http://www.w3.org/2000/svg")
	svg.AddAttrib("version", "1.1")
	svg.AddAttrib("baseProfile", "tiny")
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sw := strconv.Itoa(w)
	sh := strconv.Itoa(h)
	svg.AddAttrib("viewBox", sx+" "+sy+" "+sw+" "+sh)
	return page, svg
}

func (svg *Tag) AddRect(x, y, w, h int) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sw := strconv.Itoa(w)
	sh := strconv.Itoa(h)
	rect := svg.AddNewTag("rect")
	rect.AddAttrib("x", sx)
	rect.AddAttrib("y", sy)
	rect.AddAttrib("width", sw)
	rect.AddAttrib("height", sh)
	return rect
}

// Add a box/rectangle, given x and y position, width, height and color
func (svg *Tag) Box(x, y, w, h int, color string) *Tag {
	rect := svg.AddRect(x, y, w, h)
	rect.Fill(color)
	return rect
}

func (svg *Tag) AddCircle(x, y, radius int) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sradius := strconv.Itoa(radius)
	circle := svg.AddNewTag("circle")
	circle.AddAttrib("cx", sx)
	circle.AddAttrib("cy", sy)
	circle.AddAttrib("r", sradius)
	return circle
}

func (svg *Tag) AddEllipse(x, y, rx, ry int) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	srx := strconv.Itoa(rx)
	sry := strconv.Itoa(ry)
	ellipse := svg.AddNewTag("ellipse")
	ellipse.AddAttrib("cx", sx)
	ellipse.AddAttrib("cy", sy)
	ellipse.AddAttrib("rx", srx)
	ellipse.AddAttrib("ry", sry)
	return ellipse
}

// Add a line from (x1, y1) to (x2, y2) with a given stroke width and color
func (svg *Tag) Line(x1, y1, x2, y2, thickness int, color string) *Tag {
	sx1 := strconv.Itoa(x1)
	sy1 := strconv.Itoa(y1)
	sx2 := strconv.Itoa(x2)
	sy2 := strconv.Itoa(y2)
	sw := strconv.Itoa(thickness)
	line := svg.AddNewTag("line")
	line.AddAttrib("x1", sx1)
	line.AddAttrib("y1", sy1)
	line.AddAttrib("x2", sx2)
	line.AddAttrib("y2", sy2)
	line.AddAttrib("stroke-width", sw)
	line.AddAttrib("stroke", color)
	return line
}

func (svg *Tag) Triangle(x1, y1, x2, y2, x3, y3 int, color string) *Tag {
	triangle := svg.AddNewTag("path")
	triangle.AddAttrib("d", fmt.Sprintf("M %d %d L %d %d L %d %d L %d %d", x1, y1, x2, y2, x3, y3, x1, y1))
	triangle.AddAttrib("fill", color)
	return triangle
}

func (svg *Tag) Poly4(x1, y1, x2, y2, x3, y3, x4, y4 int, color string) *Tag {
	poly4 := svg.AddNewTag("path")
	poly4.AddAttrib("d", fmt.Sprintf("M %d %d L %d %d L %d %d L %d %d L %d %d", x1, y1, x2, y2, x3, y3, x4, y4, x1, y1))
	poly4.AddAttrib("fill", color)
	return poly4
}

// Add a circle, given x and y position, radius and color
func (svg *Tag) Circle(x, y, radius int, color string) *Tag {
	circle := svg.AddCircle(x, y, radius)
	circle.Fill(color)
	return circle
}

// Add an ellipse, given x and y position, x radius, y radius and color
func (svg *Tag) Ellipse(x, y, xr, yr int, color string) *Tag {
	ellipse := svg.AddEllipse(x, y, xr, yr)
	ellipse.Fill(color)
	return ellipse
}

func (rect *Tag) Fill(color string) {
	rect.AddAttrib("fill", color)
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

// Converts r, g and b which are integers in the range from 0..255
// and a which is the alpha color in the range from 0.0 to 1.0
// to a color-string on the form "rgba(255, 255, 255, 1.0)".
func ColorStringAlpha(r, g, b int, a float64) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, a)
}

// Creates a rectangle that is 1 wide with the given color
// Note that the size of the "pixel" depends on how large the viewBox is
func (svg *Tag) Pixel(x, y, r, g, b int) *Tag {
	color := ColorString(r, g, b)
	rect := svg.Box(x, y, 1, 1, color)
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

func SampleSVG1() *Page {
	page, svg := NewTinySVG(0, 0, 30, 30)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Sample SVG file 1")
	rect := svg.AddRect(10, 10, 10, 10)
	rect.Fill("green")
	svg.Pixel(10, 10, 255, 0, 0)
	svg.AlphaDot(5, 5, 0, 0, 255, 0.5)
	return page
}

func SampleSVG2() *Page {
	w := 160
	h := 90
	stepx := 8
	stepy := 8
	page, svg := NewTinySVG(0, 0, w, h)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Sample SVG file 2")
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
