package onthefly

// This is a test function
func SamplePage(cssurl string) *Page {
	page := NewHTML5Page("Hello")
	body, _ := page.SetMargin(3)

	h1 := body.AddNewTag("h1")
	h1.SetMargin(1)
	h1.AddContent("Browser")

	h1, err := page.root.GetTag("h1")
	if err == nil {
		h1.AddContent("Spe")
	}

	if err := page.LinkToCSS(cssurl); err == nil {
		h1.AddContent("ak")
	} else {
		h1.AddContent("akkkkkkkk")
	}

	page.SetColor("#202020", "#A0A0A0")
	page.SetFontFamily("sans serif")

	box, _ := page.addBox("box0", true)
	box.AddStyle("margin-top", "-2em")
	box.AddStyle("margin-bottom", "3em")

	image := body.AddImage("http://www.shoutmeloud.com/wp-content/uploads/2010/01/successful-Blogger.jpeg", "50%")
	image.AddStyle("margin-top", "2em")
	image.AddStyle("margin-left", "3em")

	return page
}
