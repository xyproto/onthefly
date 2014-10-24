package onthefly

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	Version     = 0.8
	noAttribute = "NIL"
)

type Tag struct {
	name        string
	style       map[string]string
	content     string
	lastContent string
	xmlContent  string
	attrs       map[string]string
	nextSibling *Tag // siblings
	firstChild  *Tag // first child
}

type Page struct {
	title  string
	root   *Tag
	cursor *Tag
}

// Create a new XML/HTML page, with a root tag
// If rootTagName contains "<" or ">", it can be used for preceding declarations,
// like <!DOCTYPE html> or <?xml version=\"1.0\"?>.
// Returns a pointer to a Page.
func NewPage(title, rootTagName string) *Page {
	var page Page
	page.title = title
	rootTag := NewTag(rootTagName)
	page.root = rootTag
	return &page
}

// Create a new tag based on a given name.
// name is what will appear right after the "<" in the tag output
func NewTag(name string) *Tag {
	var tag Tag
	tag.name = name
	tag.style = make(map[string]string)
	tag.attrs = make(map[string]string)
	tag.nextSibling = nil
	tag.firstChild = nil
	tag.content = ""
	tag.lastContent = ""
	return &tag
}

// Add a new tag to another tag. This will place it one step
// lower in the hierarchy of tags. You can for example add
// a body tag to an html tag.
func (tag *Tag) AddNewTag(name string) *Tag {
	child := NewTag(name)
	tag.AddChild(child)
	return child
}

// Add a tag to another tag
func (tag *Tag) AddTag(child *Tag) {
	tag.AddChild(child)
}

// Add CSS style to a tag, for instance
// "background-color" and "red"
func (tag *Tag) AddStyle(styleName, styleValue string) {
	tag.style[styleName] = styleValue
}

// Add an attribute to a tag, for instance
// "size" and "20"
func (tag *Tag) AddAttrib(attrName, attrValue string) {
	tag.attrs[attrName] = attrValue
}

// Add an attribute without a value
func (tag *Tag) AddSingularAttrib(attrName string) {
	tag.attrs[attrName] = noAttribute
}

// Generate the CSS text for a given tag
// The generated string can be used directly in a CSS file
func (tag *Tag) GetCSS() string {
	if len(tag.style) == 0 {
		return ""
	}

	ret := ""

	// If there is an id="name" defined, use that id instead of the tag name

	if value, found := tag.attrs["id"]; found {
		ret = "#" + value
	} else if value, found := tag.attrs["class"]; found {
		ret = "." + value
	} else {
		ret = tag.name
	}
	ret += " {\n"

	// Attributes may appear in any order
	for key, value := range tag.style {
		ret += "  " + key + ": " + value + ";\n"
	}

	ret += "}\n\n"
	return ret
}

// Get a string that represents all the attribute keys and values
// of a tag. This can be used when generating HTML, for example.
func (tag *Tag) GetAttrString() string {
	ret := ""
	for key, value := range tag.attrs {
		if value == noAttribute {
			ret += key + " "
		} else {
			ret += key + "=\"" + value + "\"" + " "
		}
	}
	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}
	return ret
}

// Get spaces for indenting based on a given level
func getSpaces(level int) string {
	spacing := ""
	for i := 1; i < level; i++ {
		spacing += "  "
	}
	return spacing
}

// Generate a string for a tag, non-recursively
// indent is if the output should be indented or nto
// level is how many levels deep the output should be indented.
func (tag *Tag) getFlatXML(indent bool, level int) string {
	newLine := ""
	if indent {
		newLine = "\n"
	}
	// For the root tag
	if (len(tag.name) > 0) && (tag.name[0] == '<') {
		return tag.name + newLine + tag.content + tag.xmlContent + tag.lastContent
	}
	// For indenting
	spacing := ""
	if indent {
		spacing = getSpaces(level)
	}
	// Generate the XML based on the tag
	attrs := tag.GetAttrString()
	ret := spacing + "<" + tag.name
	if len(attrs) > 0 {
		ret += " " + attrs
	}
	if (len(tag.content) == 0) && (len(tag.xmlContent) == 0) && (len(tag.lastContent) == 0) {
		ret += " />"
	} else {
		if len(tag.xmlContent) > 0 {
			if tag.xmlContent[0] != ' ' {
				ret += ">" + newLine + spacing + tag.xmlContent + newLine + spacing + "</" + tag.name + ">"
			} else {
				ret += ">" + newLine + tag.xmlContent + spacing + "</" + tag.name + ">"
			}
		} else {
			ret += ">" + tag.content + tag.lastContent + "</" + tag.name + ">"
		}
	}
	return ret
}

// Get all the children for a given tag
// Returns a slice of pointers to tags
func (tag *Tag) GetChildren() []*Tag {
	var children []*Tag
	current := tag.firstChild
	for current != nil {
		children = append(children, current)
		current = current.nextSibling
	}
	return children
}

// Add a tag as a child to another tag
func (tag *Tag) AddChild(child *Tag) {
	if tag.firstChild == nil {
		tag.firstChild = child
		return
	}
	lastChild := tag.LastChild()
	child.nextSibling = nil
	lastChild.nextSibling = child
}

// Add content to a tag. This is what will appear
// between two tags, for example: <tag>content</tag>
// If the tag contains child tags, they will be added after this
func (tag *Tag) AddContent(content string) {
	tag.content += content
}

// Add content that will be added within a tag, but after
// the other content and any child tags
func (tag *Tag) AddLastContent(content string) {
	tag.lastContent += content
}

// Count how many children a tag has
// Returns an integer
func (tag *Tag) CountChildren() int {
	child := tag.firstChild
	if child == nil {
		return 0
	}
	count := 1
	if child.nextSibling == nil {
		return count
	}
	child = child.nextSibling
	for child != nil {
		count++
		child = child.nextSibling
	}
	return count
}

// Count the number of siblings a tag has
func (tag *Tag) CountSiblings() int {
	sib := tag.nextSibling
	if sib == nil {
		return 0
	}
	count := 1
	if sib.nextSibling == nil {
		return count
	}
	sib = sib.nextSibling
	for sib != nil {
		count++
		sib = sib.nextSibling
	}
	return count
}

// Find the last child of the children of a tag
func (tag *Tag) LastChild() *Tag {
	child := tag.firstChild
	for child.nextSibling != nil {
		child = child.nextSibling
	}
	return child
}

// Given the name of a tag, finds the first tag that matches
func (page *Page) GetTag(name string) (*Tag, error) {
	return page.root.GetTag(name)
}

// Find a tag by name, returns an error if not found
// Returns the first tag that matches
func (tag *Tag) GetTag(name string) (*Tag, error) {
	if strings.Index(tag.name, name) == 0 {
		return tag, nil
	}
	couldNotFindError := errors.New("Could not find tag: " + name)
	if tag.CountChildren() == 0 {
		// No children. Not found so far
		return nil, couldNotFindError
	}

	var child *Tag = tag.firstChild
	for child != nil {
		found, err := child.GetTag(name)
		if err == nil {
			return found, err
		}
		child = child.nextSibling
	}

	return nil, couldNotFindError
}

// Generate XML for a tag, recursively
// indent is if the output should be indented or not
// level is the indentation level
// Returns the generated XML as a string
func getXMLRecursively(cursor *Tag, indent bool, level int) string {

	newLine := ""
	if indent {
		newLine = "\n"
	}

	if cursor.CountChildren() == 0 {
		return cursor.getFlatXML(indent, level) + newLine
	}

	content := ""
	xmlContent := ""

	level++

	child := cursor.firstChild
	for child != nil {
		xmlContent = getXMLRecursively(child, indent, level)
		if len(xmlContent) > 0 {
			content += xmlContent
		}
		child = child.nextSibling
	}

	level--

	cursor.xmlContent = cursor.content + content + cursor.lastContent

	ret := cursor.getFlatXML(indent, level)
	if level > 0 {
		ret += newLine
	}
	return ret
}

// Generate CSS for a tag, recursively
// Returns the generated CSS as a string
// The output can go directly in a CSS file
func getCSSRecursively(cursor *Tag) string {
	if cursor.CountChildren() == 0 {
		return cursor.GetCSS()
	}

	style := ""
	cssContent := ""

	child := cursor.firstChild
	for child != nil {
		cssContent = getCSSRecursively(child)
		if len(cssContent) > 0 {
			style += cssContent
		}
		child = child.nextSibling
	}

	return cursor.GetCSS() + style
}

// Generate XML for a page
// The output can go directly in an XML file
func (page *Page) GetXML(indent bool) string {
	return getXMLRecursively(page.root, indent, 0)
}

// Generate CSS for a page
// The output can go directly in a CSS file
func (page *Page) GetCSS() string {
	return getCSSRecursively(page.root)
}

// Generate HTML for a page
func (page *Page) GetHTML() string {
	return page.GetXML(true)
}

// Show various information for a page, used for debugging
func (page *Page) prettyPrint() {
	root := *page.root
	fmt.Println("Page title:", page.title)
	fmt.Println("Page root tag name:", root.name)
	fmt.Println("Root tag children count:", root.CountChildren())
	fmt.Printf("HTML:\n%s\n", page.GetXML(true))
	fmt.Printf("CSS:\n%s\n", page.GetCSS())
}

// Add content to the body tag
// Returns the body tag and nil if successful
// Returns and an error if no body tag is found, else nil
func (page *Page) AddContent(content string) (*Tag, error) {
	body, err := page.root.GetTag("body")
	if err == nil {
		body.AddContent(content)
	}
	return body, err
}

// Get the string for the page
func (page *Page) String() string {
	return page.GetHTML()
}

// Publish the linked HTML and CSS for a Page. Specify if the contents should be generated every time or not.
func (page *Page) Publish(mux *http.ServeMux, htmlurl, cssurl string, generate bool) {
	page.LinkToCSS(cssurl)
	if generate {
		// Serve HTML that is generated for each call
		mux.HandleFunc(htmlurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			fmt.Fprint(w, page.GetHTML())
		})
		// Serve CSS that is generated for each call
		mux.HandleFunc(cssurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/css")
			fmt.Fprint(w, page.GetCSS())
		})
	} else {
		// Cached
		html := page.GetHTML()
		css := page.GetCSS()
		// Serve HTML
		mux.HandleFunc(htmlurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			fmt.Fprint(w, html)
		})
		// Serve CSS
		mux.HandleFunc(cssurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/css")
			fmt.Fprint(w, css)
		})
	}
}
