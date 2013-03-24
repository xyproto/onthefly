package browserspeak

import (
	"strconv"
)

// Create an empty page only containing the given tag
// Returns both the page and the tag
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

func HTMLPageRedirect(url string) string {
	return "<html><head><script type=\"text/javascript\">window.location.href = \"" + url + "\";</script></head></html>"
}

func LoginForm() string {
	// TODO: Use a CowboyTag instead and create a more general form generator
	return "<form id=\"loginForm\" action=\"/login\" method=\"POST\"><div style=\"margin: 1em;\"><label for=\"username\" style=\"display: inline-block; float: left; clear: left; width: 150px; text-align: right; margin-right: 2em;\">Username:</label><input style=\"display:inline-block; float:left;\" id=\"username\"><br /><label for=\"password\" style=\"display: inline-block; float: left; clear: left; width: 150px; text-align: right; margin-right: 2em;\">Password:</label><input style=\"display:inline-block; float:left;\" id=\"password\" type=\"password\" name=\"password\"></div><p><button style=\"margin-left: 300px; margin-top: 1em;\" id=\"loginButton\">Login</button></p></form>"
}

func RegisterForm() string {
	// TODO: Use a CowboyTag instead and create a more general form generator
	return "<form id=\"registerForm\" action=\"/register\" method=\"POST\"><div style=\"margin: 1em;\"><label for=\"username\" style=\"display: inline-block; float: left; clear: left; width: 150px; text-align: right; margin-right: 2em;\">Username:</label><input style=\"display:inline-block; float:left;\" id=\"username\"><br /><label for=\"password1\" style=\"display: inline-block; float: left; clear: left; width: 150px; text-align: right; margin-right: 2em;\">Password:</label><input style=\"display:inline-block; float:left;\" id=\"password1\" type=\"password\" name=\"password1\"><br /><label for=\"password2\" style=\"display: inline-block; float: left; clear: left; width: 150px; text-align: right; margin-right: 2em;\">Confirm password:</label><input style=\"display:inline-block; float:left;\" id=\"password2\" type=\"password\" name=\"password2\"><br /><label for=\"email\" style=\"display: inline-block; float: left; clear: left; width: 150px; text-align: right; margin-right: 2em;\">Email:</label><input name=\"email\" style=\"display:inline-block; float:left;\" id=\"email\"></div><p><button style=\"margin-left: 300px; margin-top: 1em;\" id=\"registerButton\">Register</button></p></form>"
}

func messageComposer(title, msg, javascript string) string {
	// TODO: Use a CowboyTag instead
	return "<!DOCTYPE html><html><head><title>" + title + "</title></head><body style=\"margin:4em; font-family:courier; color:gray; background-color:light gray;\"><h2>" + title + "</h2><hr style=\"margin-top:-1em; margin-bottom: 2em; margin-right: 20%; text-align: left; border: 1px dotted #b0b0b0; height:1px;\"><div style=\"margin-left: 2em;\">" + msg + "<br /><br /><button style=\"margin-top: 2em; margin-left: 20em;\" onclick=\"" + javascript + "\">OK</button></div></body></html>"
}

// Displays a title, a message and an OK button that goes to the previous page
func MessageOKback(title, msg string) string {
	return messageComposer(title, msg, "history.go(-1);")
}

// Displays a title, a message and an OK button that goes to the given url
func MessageOKurl(title, msg, url string) string {
	return messageComposer(title, msg, "location.href='"+url+"';")
}
