package onthefly

//
// Support for TinySVG 1.2
//
// The functions are suffixed with "2" here, to keep the function names in tinysvg.go,
// to avoid breaking backward compatibility.
//
// All functions will be reorganized in a future release of onthefly, but then under a new package name.
//

import (
	"fmt"
	"strconv"
)

const (
	TRANSPARENT = 0.0
	OPAQUE      = 1.0
)

type Pos struct {
	x int
	y int
}

type Size struct {
	w int
	h int
}

type Radius struct {
	x int
	y int
}

type Color struct {
	r int     // red, 0..255
	g int     // green, 0..255
	b int     // blue, 0..255
	a float64 // alpha, 0.0..1.0
	n string  // name (optional, will override the above values)
}

type Font struct {
	family string
	size   int
}

// NewTinySVG2 creates new TinySVG 1.2 image. Pos and Size defines the viewbox
func NewTinySVG2(p *Pos, s *Size) (*Page, *Tag) {
	// No page title is needed when building an SVG tag tree
	page := NewPage("", `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.2 Tiny//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-tiny.dtd">`)

	// Add the root tag
	svg := page.root.AddNewTag("svg")
	svg.AddAttrib("xmlns", "http://www.w3.org/2000/svg")
	svg.AddAttrib("version", "1.2")
	svg.AddAttrib("baseProfile", "tiny")
	svg.AddAttrib("viewBox", fmt.Sprintf("%d %d %d %d", p.x, p.y, s.w, s.h))
	return page, svg
}

// Rect a rectangle, given x and y position, width and height.
// No color is being set.
func (svg *Tag) Rect2(p *Pos, s *Size, c *Color) *Tag {
	rect := svg.AddNewTag("rect")
	rect.AddAttrib("x", strconv.Itoa(p.x))
	rect.AddAttrib("y", strconv.Itoa(p.y))
	rect.AddAttrib("width", strconv.Itoa(s.w))
	rect.AddAttrib("height", strconv.Itoa(s.h))
	rect.Fill2(c)
	return rect
}

// Text adds text. No color is being set
func (svg *Tag) Text2(p *Pos, f *Font, message string, c *Color) *Tag {
	text := svg.AddNewTag("text")
	text.AddAttrib("x", strconv.Itoa(p.x))
	text.AddAttrib("y", strconv.Itoa(p.y))
	text.AddAttrib("font-family", f.family)
	text.AddAttrib("font-size", strconv.Itoa(f.size))
	text.Fill2(c)
	text.AddContent(message)
	return text
}

// Circle adds a circle, given a position, radius and color
func (svg *Tag) Circle2(p *Pos, radius int, c *Color) *Tag {
	circle := svg.AddNewTag("circle")
	circle.AddAttrib("cx", strconv.Itoa(p.x))
	circle.AddAttrib("cy", strconv.Itoa(p.y))
	circle.AddAttrib("r", strconv.Itoa(radius))
	circle.Fill2(c)
	return circle
}

// Ellipse adds an ellipse with a given position (x,y) and radius (rx, ry).
func (svg *Tag) Ellipse2(p *Pos, r *Radius, c *Color) *Tag {
	ellipse := svg.AddNewTag("ellipse")
	ellipse.AddAttrib("cx", strconv.Itoa(p.x))
	ellipse.AddAttrib("cy", strconv.Itoa(p.y))
	ellipse.AddAttrib("rx", strconv.Itoa(r.x))
	ellipse.AddAttrib("ry", strconv.Itoa(r.y))
	ellipse.Fill2(c)
	return ellipse
}

// Line adds a line from (x1, y1) to (x2, y2) with a given stroke width and color
func (svg *Tag) Line2(p1, p2 *Pos, thickness int, c *Color) *Tag {
	line := svg.AddNewTag("line")
	line.AddAttrib("x1", strconv.Itoa(p1.x))
	line.AddAttrib("y1", strconv.Itoa(p1.y))
	line.AddAttrib("x2", strconv.Itoa(p2.x))
	line.AddAttrib("y2", strconv.Itoa(p2.y))
	line.AddAttrib("stroke-width", strconv.Itoa(thickness))
	line.Stroke2(c)
	return line
}

// Triangle adds a colored triangle
func (svg *Tag) Triangle2(p1, p2, p3 *Pos, c *Color) *Tag {
	triangle := svg.AddNewTag("path")
	triangle.AddAttrib("d", fmt.Sprintf("M %d %d L %d %d L %d %d L %d %d", p1.x, p1.y, p2.x, p2.y, p3.x, p3.y, p1.x, p1.y))
	triangle.Fill2(c)
	return triangle
}

// Poly2 adds a colored polygon with 4 points
func (svg *Tag) Poly2(p1, p2, p3, p4 *Pos, c *Color) *Tag {
	poly4 := svg.AddNewTag("path")
	poly4.AddAttrib("d", fmt.Sprintf("M %d %d L %d %d L %d %d L %d %d L %d %d", p1.x, p1.y, p2.x, p2.y, p3.x, p3.y, p4.x, p4.y, p1.x, p1.y))
	poly4.Fill2(c)
	return poly4
}

// Fill selects the fill color that will be used when drawing
func (svg *Tag) Fill2(c *Color) {
	// If no color name is given and the color is transparent, don't set a fill color
	if (c == nil) || (len(c.n) == 0 && c.a == TRANSPARENT) {
		return
	}
	svg.AddAttrib("fill", c.String())
}

// Stroke selects the stroke color that will be used when drawing
func (svg *Tag) Stroke2(c *Color) {
	// If no color name is given and the color is transparent, don't set a stroke color
	if (c == nil) || (len(c.n) == 0 && c.a == TRANSPARENT) {
		return
	}
	svg.AddAttrib("stroke", c.String())
}

// RGBString converts r, g and b (integers in the range 0..255)
// to a color string on the form "#nnnnnn".
func RGBString(r, g, b int) string {
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

// RGBAString converts integers r, g and b (the color) and also
// a given alpha (opacity) to a color-string on the form
// "rgba(255, 255, 255, 1.0)".
func RGBAString(r, g, b int, a float64) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, a)
}

// RGBA creates a new Color with the given red, green and blue values.
// The colors are in the range 0..255
func RGB(r, g, b int) *Color {
	return &Color{r, g, b, OPAQUE, ""}
}

// RGBA creates a new Color with the given red, green, blue and alpha values.
// Alpha is between 0 and 1, the rest are 0..255.
// For the alpha value, 0 is transparent and 1 is opaque.
func RGBA(r, g, b int, a float64) *Color {
	return &Color{r, g, b, a, ""}
}

// ColorByName creates a new Color with a given name, like "blue"
func ColorByName(name string) *Color {
	return &Color{n: name}
}

// String returns the color as an RGB (#1234FF) string
// or as an RGBA (rgba(0, 1, 2 ,3)) string.
func (c *Color) String() string {
	// Return an empty string if nil
	if c == nil {
		return ""
	}
	// Return the name, if specified
	if len(c.n) != 0 {
		return c.n
	}
	// Return a regular RGB string if alpha is 1.0
	if c.a == OPAQUE {
		// Generate a rgb string
		return RGBString(c.r, c.g, c.b)
	}
	// Generate a rgba string if alpha is < 1.0
	return RGBAString(c.r, c.g, c.b, c.a)
}
