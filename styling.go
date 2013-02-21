package browserspeak

// Various "hardcoded" stylistic choices

// Boxes for content
func (tag *Tag) RoundedBox() {
	tag.AddStyle("border", "solid 1px #b4b4b4")
	tag.AddStyle("border-radius", "10px")
	tag.AddStyle("box-shadow", "1px 1px 3px rgba(0,0,0, .5)")
}

// Set the tag font to some sort of sans-serif
func (tag *Tag) SansSerif() {
	tag.AddStyle("font-family", "Verdana, Geneva, sans-serif")
}

// Set the tag font to the given font or just some sort of sans-serif
func (tag *Tag) CustomSansSerif(custom string) {
	tag.AddStyle("font-family", custom+", Verdana, Geneva, sans-serif")
}
