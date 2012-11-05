package browserspeak

import (
	"errors"
	"fmt"
	"strings"
)

type Tag struct {
	name        string
	style       map[string]string
	content     string
	htmlContent string
	attrs       map[string]string
	nextSibling *Tag // siblings
	firstChild  *Tag // first child
}

type Page struct {
	title  string
	root   *Tag
	cursor *Tag
}

func NewPage(title, rootTagName string) *Page {
	var page Page
	page.title = title
	rootTag := NewTag(rootTagName)
	page.root = rootTag
	return &page
}

func NewTag(name string) *Tag {
	var tag Tag
	tag.name = name
	tag.style = make(map[string]string)
	tag.attrs = make(map[string]string)
	tag.nextSibling = nil
	tag.firstChild = nil
	return &tag
}

func (tag *Tag) AddNewTag(name string) *Tag {
	var child *Tag = NewTag(name)
	tag.AddChild(child)
	return child
}

func (tag Tag) AddStyle(key, value string) {
	tag.style[key] = value
}

func (tag Tag) AddAttr(key, value string) {
	tag.attrs[key] = value
}

// Generate CSS that can go in a CSS file, for a given tag
func (tag *Tag) GetCSS() string {
	if len(tag.style) == 0 {
		return ""
	}

	ret := ""

	// If there is an id="name" defined, use that id instead of the tag name
	value, found := tag.attrs["id"]
	if found {
		ret = "#" + value
	} else {
		ret = tag.name
	}
	ret += " {\n"

	// NB! The attributes may appear in any order!
	for key, value := range tag.style {
		ret += "  " + key + ": " + value + ";\n"
	}
	ret += "}\n\n"
	return ret
}

func (tag *Tag) GetAttrString() string {
	ret := ""
	for key, value := range tag.attrs {
		ret += key + "=\"" + value + "\"" + " "
	}
	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}
	return ret
}

// Get spaces for indenting based on a given level
func GetSpaces(level int) string {
	spacing := ""
	for i := 1; i < level; i++ {
		spacing += "  "
	}
	return spacing
}

func (tag *Tag) GetFlatHTML(level int) string {
	// For the root tag
	if (len(tag.name) > 0) && (tag.name[0] == '<') {
		return tag.name + "\n" + tag.content + tag.htmlContent
	}
	// For indenting
	spacing := GetSpaces(level)
	// Generate the HTML based on the tag
	attrs := tag.GetAttrString()
	ret := spacing + "<" + tag.name
	if len(attrs) > 0 {
		ret += " " + attrs
	}
	if (len(tag.content) == 0) && (len(tag.htmlContent) == 0) {
		ret += " />"
	} else {
		if len(tag.htmlContent) > 0 {
			ret += ">" + "\n" + tag.content + tag.htmlContent + spacing + "</" + tag.name + ">"
		} else {
			ret += ">" + tag.content + "</" + tag.name + ">"
			// Indented content
			//ret += ">" + "\n" + GetSpaces(level + 1) + tag.content + "\n" + spacing + "</" + tag.name + ">"
		}
	}
	return ret
}

func (tag *Tag) GetChildren() []*Tag {
	var children []*Tag
	current := tag.firstChild
	for current != nil {
		children = append(children, current)
		current = current.nextSibling
	}
	return children
}

func (tag *Tag) AddChild(child *Tag) {
	if tag.firstChild == nil {
		tag.firstChild = child
		return
	}
	lastChild := tag.LastChild()
	child.nextSibling = nil
	lastChild.nextSibling = child
}

func (tag *Tag) AddContent(content string) {
	tag.content += content
}

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

func (tag *Tag) LastChild() *Tag {
	child := tag.firstChild
	for child.nextSibling != nil {
		child = child.nextSibling
	}
	return child
}

// Find a tag by name, returns an error if not found
func (cursor *Tag) GetTag(name string) (*Tag, error) {
	if strings.Index(cursor.name, name) == 0 {
		return cursor, nil
	}
	couldNotFindError := errors.New("Could not find tag: " + name)
	if cursor.CountChildren() == 0 {
		// No children. Not found so far
		return nil, couldNotFindError
	}

	var child *Tag = cursor.firstChild
	for child != nil {
		found, err := child.GetTag(name)
		if err == nil {
			return found, err
		}
		child = child.nextSibling
	}

	return nil, couldNotFindError
}

func GetHTMLRecursively(cursor *Tag, level int) string {

	if cursor.CountChildren() == 0 {
		return cursor.GetFlatHTML(level) + "\n"
	}

	content := ""
	htmlContent := ""

	level++

	child := cursor.firstChild
	for child != nil {
		htmlContent = GetHTMLRecursively(child, level)
		if len(htmlContent) > 0 {
			content += htmlContent
		}
		child = child.nextSibling
	}

	level--

	cursor.htmlContent = cursor.content + content

	ret := cursor.GetFlatHTML(level)
	if level > 0 {
		ret += "\n"
	}
	return ret
}

func GetCSSRecursively(cursor *Tag) string {
	if cursor.CountChildren() == 0 {
		return cursor.GetCSS()
	}

	style := ""
	cssContent := ""

	child := cursor.firstChild
	for child != nil {
		cssContent = GetCSSRecursively(child)
		if len(cssContent) > 0 {
			style += cssContent
		}
		child = child.nextSibling
	}

	return cursor.GetCSS() + style
}

func (page Page) GetHTML() string {
	return GetHTMLRecursively(page.root, 0)
}

func (page Page) GetCSS() string {
	return GetCSSRecursively(page.root)
}

func (page Page) PrettyPrint() {
	fmt.Println("Page title:", page.title)
	fmt.Println("Page root tag name:", page.root.name)
	rootPointer := page.root
	fmt.Println("Root tag children count:", rootPointer.CountChildren())
	fmt.Printf("HTML:\n%s\n", page.GetHTML())
	fmt.Printf("CSS:\n%s\n", page.GetCSS())
}

