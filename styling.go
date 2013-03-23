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

func AddGoogleFonts(page *Page, googleFonts []string) {
	for _, fontname := range googleFonts {
		page.LinkToGoogleFont(fontname)
	}
}

func AddScriptToHeader(page *Page, js string) error {
	// Check if there's anything to add
	if js == "" {
		// Nope
		return nil
	}
	// Add a script tag
	head, err := page.GetTag("head")
	if err == nil {
		script := head.AddNewTag("script")
		script.AddAttr("type", "text/javascript")
		script.AddContent(js)
	}
	return err
}

func AddBodyStyle(page *Page, bgimageurl string, stretchBackground bool) {
	body, _ := page.SetMargin(1)
	body.SansSerif()
	if stretchBackground {
		body.AddStyle("background", "url('"+bgimageurl+"') no-repeat center center fixed")
	} else {
		body.AddStyle("background", "url('"+bgimageurl+"')")
	}
}
