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

// AddRect a rectangle, given x and y position, width and height.
// No color is being set.
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

// AddText adds text. No color is being set
func (svg *Tag) AddText(x, y, fontSize int, fontFamily, text string) *Tag {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	fs := strconv.Itoa(fontSize)
	textTag := svg.AddNewTag("text")
	textTag.AddAttrib("x", sx)
	textTag.AddAttrib("y", sy)
	textTag.AddAttrib("font-family", fontFamily)
	textTag.AddAttrib("font-size", fs)
	textTag.AddContent(text)
	return textTag
}

// Box adds a rectangle, given x and y position, width, height and color
func (svg *Tag) Box(x, y, w, h int, color string) *Tag {
	rect := svg.AddRect(x, y, w, h)
	rect.Fill(color)
	return rect
}

// AddCircle adds a circle Add a circle, given a position (x, y) and a radius
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

// AddEllipse adds an ellipse with a given position (x,y) and radius (rx, ry).
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

// Line adds a line from (x1, y1) to (x2, y2) with a given stroke width and color
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

// Triangle adds a colored triangle
func (svg *Tag) Triangle(x1, y1, x2, y2, x3, y3 int, color string) *Tag {
	triangle := svg.AddNewTag("path")
	triangle.AddAttrib("d", fmt.Sprintf("M %d %d L %d %d L %d %d L %d %d", x1, y1, x2, y2, x3, y3, x1, y1))
	triangle.AddAttrib("fill", color)
	return triangle
}

// Poly4 adds a colored polygon with 4 points
func (svg *Tag) Poly4(x1, y1, x2, y2, x3, y3, x4, y4 int, color string) *Tag {
	poly4 := svg.AddNewTag("path")
	poly4.AddAttrib("d", fmt.Sprintf("M %d %d L %d %d L %d %d L %d %d L %d %d", x1, y1, x2, y2, x3, y3, x4, y4, x1, y1))
	poly4.AddAttrib("fill", color)
	return poly4
}

// Circle adds a circle, given x and y position, radius and color
func (svg *Tag) Circle(x, y, radius int, color string) *Tag {
	circle := svg.AddCircle(x, y, radius)
	circle.Fill(color)
	return circle
}

// Ellipse adds an ellipse, given x and y position, radiuses and color
func (svg *Tag) Ellipse(x, y, xr, yr int, color string) *Tag {
	ellipse := svg.AddEllipse(x, y, xr, yr)
	ellipse.Fill(color)
	return ellipse
}

// Fill selects the fill color that will be used when drawing
func (svg *Tag) Fill(color string) {
	svg.AddAttrib("fill", color)
}

// ColorString converts r, g and b (integers in the range 0..255)
// to a color string on the form "#nnnnnn".
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

// ColorStringAlpha converts integers r, g and b (the color) and also
// a given alpha (opacity) to a color-string on the form
// "rgba(255, 255, 255, 1.0)".
func ColorStringAlpha(r, g, b int, a float64) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, a)
}

// Pixel creates a rectangle that is 1 wide with the given color.
// Note that the size of the "pixel" depends on how large the viewBox is.
func (svg *Tag) Pixel(x, y, r, g, b int) *Tag {
	color := ColorString(r, g, b)
	rect := svg.Box(x, y, 1, 1, color)
	return rect
}

// AlphaDot creates a small circle that can be transparent.
// Takes a position (x, y) and a color (r, g, b, a).
func (svg *Tag) AlphaDot(x, y, r, g, b int, a float32) *Tag {
	color := fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, a)
	circle := svg.AddCircle(x, y, 1)
	circle.Fill(color)
	return circle
}

// Dot adds a small colored circle.
// Takes a position (x, y) and a color (r, g, b).
func (svg *Tag) Dot(x, y, r, g, b int) *Tag {
	color := ColorString(r, g, b)
	circle := svg.AddCircle(x, y, 1)
	circle.Fill(color)
	return circle
}

//
