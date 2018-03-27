package onthefly

// Only convenience functions and functions for backward compatibility here

func NewTinySVG(x, y, w, h int) (*Page, *Tag) {
	return NewTinySVG2(&Pos{x, y}, &Size{w, h})
}

// AddRect a rectangle, given x and y position, width and height.
// No color is being set.
func (svg *Tag) AddRect(x, y, w, h int) *Tag {
	return svg.Rect2(&Pos{x, y}, &Size{w, h}, nil)
}

// AddText adds text. No color is being set
func (svg *Tag) AddText(x, y, fontSize int, fontFamily, text string) *Tag {
	return svg.Text2(&Pos{x, y}, &Font{fontFamily, fontSize}, text, nil)
}

// Box adds a rectangle, given x and y position, width, height and color
func (svg *Tag) Box(x, y, w, h int, color string) *Tag {
	return svg.Rect2(&Pos{x, y}, &Size{w, h}, ColorByName(color))
}

// AddCircle adds a circle Add a circle, given a position (x, y) and a radius.
// No color is being set.
func (svg *Tag) AddCircle(x, y, radius int) *Tag {
	return svg.Circle2(&Pos{x, y}, radius, nil)
}

// AddEllipse adds an ellipse with a given position (x,y) and radius (rx, ry).
// No color is being set.
func (svg *Tag) AddEllipse(x, y, rx, ry int) *Tag {
	return svg.Ellipse2(&Pos{x, y}, &Radius{rx, ry}, nil)
}

// Line adds a line from (x1, y1) to (x2, y2) with a given stroke width and color
func (svg *Tag) Line(x1, y1, x2, y2, thickness int, color string) *Tag {
	return svg.Line2(&Pos{x1, y1}, &Pos{x2, y2}, thickness, ColorByName(color))
}

// Triangle adds a colored triangle
func (svg *Tag) Triangle(x1, y1, x2, y2, x3, y3 int, color string) *Tag {
	return svg.Triangle2(&Pos{x1, y1}, &Pos{x2, y2}, &Pos{x3, y3}, ColorByName(color))
}

// Poly4 adds a colored polygon with 4 points
func (svg *Tag) Poly4(x1, y1, x2, y2, x3, y3, x4, y4 int, color string) *Tag {
	return svg.Poly2(&Pos{x1, y1}, &Pos{x2, y2}, &Pos{x3, y3}, &Pos{x4, y4}, ColorByName(color))
}

// Circle adds a circle, given x and y position, radius and color
func (svg *Tag) Circle(x, y, radius int, color string) *Tag {
	return svg.Circle2(&Pos{x, y}, radius, ColorByName(color))
}

// Ellipse adds an ellipse, given x and y position, radiuses and color
func (svg *Tag) Ellipse(x, y, xr, yr int, color string) *Tag {
	return svg.Ellipse2(&Pos{x, y}, &Radius{xr, yr}, ColorByName(color))
}

// Fill selects the fill color that will be used when drawing
func (svg *Tag) Fill(color string) {
	svg.AddAttrib("fill", color)
}

// ColorString converts r, g and b (integers in the range 0..255)
// to a color string on the form "#nnnnnn".
func ColorString(r, g, b int) string {
	return RGB(r, g, b).String()
}

// ColorStringAlpha converts integers r, g and b (the color) and also
// a given alpha (opacity) to a color-string on the form
// "rgba(255, 255, 255, 1.0)".
func ColorStringAlpha(r, g, b int, a float64) string {
	return RGBA(r, g, b, a).String()
}

// Pixel creates a rectangle that is 1 wide with the given color.
// Note that the size of the "pixel" depends on how large the viewBox is.
func (svg *Tag) Pixel(x, y, r, g, b int) *Tag {
	return svg.Rect2(&Pos{x, y}, &Size{1, 1}, RGB(r, g, b))
}

// AlphaDot creates a small circle that can be transparent.
// Takes a position (x, y) and a color (r, g, b, a).
func (svg *Tag) AlphaDot(x, y, r, g, b int, a float32) *Tag {
	return svg.Circle2(&Pos{x, y}, 1, RGBA(r, g, b, float64(a)))
}

// Dot adds a small colored circle.
// Takes a position (x, y) and a color (r, g, b).
func (svg *Tag) Dot(x, y, r, g, b int) *Tag {
	return svg.Circle2(&Pos{x, y}, 1, RGB(r, g, b))
}
