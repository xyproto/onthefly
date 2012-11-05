package main

import (
	"strconv"
)

func (tag *Tag) SetMargin(em int) {
	value := strconv.Itoa(em) + "em"
	tag.AddStyle("margin", value)
}

// Rounded corners
func (tag *Tag) SetRounded(em string) {
	//value := strconv.Itoa(em) + "em"
	value := em
	tag.AddStyle("border-radius", value)
	tag.AddStyle("-webkit-border-radius", value)
	tag.AddStyle("-moz-border-radius", value)
}

func (tag *Tag) SetColor(fgColor, bgColor string) {
	tag.AddStyle("color", fgColor)
	tag.AddStyle("background-color", bgColor)
}

// Add a box
func (tag *Tag) AddBox(id string, rounded bool, em, text, fgColor, bgColor, leftPadding string) {
	div := tag.AddNewTag("div")
	div.AddAttr("id", id)
	div.AddContent(text)
	if rounded {
		div.SetRounded(em)
	}
	div.SetColor(fgColor, bgColor)
	div.AddStyle("padding-left", leftPadding)
}
