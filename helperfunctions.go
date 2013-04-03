package browserspeak

import (
	"strconv"
)

// Create an empty page only containing the given tag
// Returns both the page and the tag
// A lone tag, without a proper page
// TODO: Find a better name
func CowboyTag(tagname string) (*Page, *Tag) {
	page := NewPage("blank", tagname)
	tag, _ := page.GetTag(tagname)
	return page, tag
}

func TagString(tagname string) string {
	page := NewPage("blank", tagname)
	return page.String()
}

func SetPixelPosition(tag *Tag, xpx, ypx int) {
	tag.AddStyle("position", "absolute")
	xpxs := strconv.Itoa(xpx) + "px"
	ypxs := strconv.Itoa(ypx) + "px"
	tag.AddStyle("top", xpxs)
	tag.AddStyle("left", ypxs)
}

func SetRelativePosition(tag *Tag, x, y string) {
	tag.AddStyle("position", "relative")
	tag.AddStyle("top", x)
	tag.AddStyle("left", y)
}

func SetWidthAndSide(tag *Tag, width string, leftside bool) {
	side := "right"
	if leftside {
		side = "left"
	}
	tag.AddStyle("float", side)
	tag.AddStyle("width", width)
}
