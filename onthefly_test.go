package onthefly

import (
	"strings"
	"testing"
)

func TestSVG(t *testing.T) {
	SampleSVG1()
}

func TestSVG2(t *testing.T) {
	SampleSVG2()
}

func TestGen(t *testing.T) {
	page := NewHTML5Page("Hello Title")
	page.LinkToCSS("/test.css")
	page.AddContent("hello body")

	// Add a link to /test.svg
	body, err := page.GetTag("body")
	if err == nil {
		body.AddNewTag("br")
		//	a := body.AddNewTag("a")
		//	a.AddAttrib("href", "/test.svg")
		//	a.AddContent("See SVG")
	}

	p := *page
	p.prettyPrint()

	s := page.String()

	if strings.Count(s, "hello body") > 1 {
		t.Errorf("Error, text appears more than once!\n%s\n", s)
	}

}

func TestGetName(t *testing.T) {
	tag := NewTag("div")
	if tag.GetName() != "div" {
		t.Errorf("Expected tag name 'div', got '%s'", tag.GetName())
	}

	tag2 := NewTag("p")
	if tag2.GetName() != "p" {
		t.Errorf("Expected tag name 'p', got '%s'", tag2.GetName())
	}
}

func TestGetSetContent(t *testing.T) {
	tag := NewTag("p")

	if tag.GetContent() != "" {
		t.Errorf("Expected empty content, got '%s'", tag.GetContent())
	}

	tag.SetContent("Hello World")
	if tag.GetContent() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", tag.GetContent())
	}

	// Test SetContent replaces existing content
	tag.SetContent("New Content")
	if tag.GetContent() != "New Content" {
		t.Errorf("Expected 'New Content', got '%s'", tag.GetContent())
	}

	// Test that SetContent is different from AddContent
	tag.AddContent(" Added")
	if tag.GetContent() != "New Content Added" {
		t.Errorf("Expected 'New Content Added', got '%s'", tag.GetContent())
	}
}

func TestChildNavigation(t *testing.T) {
	parent := NewTag("div")
	child1 := NewTag("p")
	child2 := NewTag("span")
	child3 := NewTag("a")

	// Test empty parent
	if parent.GetFirstChild() != nil {
		t.Error("Expected nil first child for empty parent")
	}

	// Add children
	parent.AddChild(child1)
	parent.AddChild(child2)
	parent.AddChild(child3)

	first := parent.GetFirstChild()
	if first == nil {
		t.Error("Expected first child, got nil")
	}
	if first.GetName() != "p" {
		t.Errorf("Expected first child name 'p', got '%s'", first.GetName())
	}

	second := first.GetNextSibling()
	if second == nil {
		t.Error("Expected second child, got nil")
	}
	if second.GetName() != "span" {
		t.Errorf("Expected second child name 'span', got '%s'", second.GetName())
	}

	third := second.GetNextSibling()
	if third == nil {
		t.Error("Expected third child, got nil")
	}
	if third.GetName() != "a" {
		t.Errorf("Expected third child name 'a', got '%s'", third.GetName())
	}

	// Test last child has no next sibling
	if third.GetNextSibling() != nil {
		t.Error("Expected nil next sibling for last child")
	}
}

func TestClearChildren(t *testing.T) {
	parent := NewTag("div")
	child1 := NewTag("p")
	child2 := NewTag("span")

	parent.AddChild(child1)
	parent.AddChild(child2)

	if parent.CountChildren() != 2 {
		t.Errorf("Expected 2 children, got %d", parent.CountChildren())
	}

	parent.ClearChildren()

	// Verify that children are cleared
	if parent.CountChildren() != 0 {
		t.Errorf("Expected 0 children after clearing, got %d", parent.CountChildren())
	}

	if parent.GetFirstChild() != nil {
		t.Error("Expected nil first child after clearing")
	}
}

func TestAttributeManagement(t *testing.T) {
	tag := NewTag("input")

	// Test HasAttribute on empty tag
	if tag.HasAttribute("type") {
		t.Error("Expected HasAttribute to return false for non-existent attribute")
	}

	// Test GetAttribute on empty tag
	value, exists := tag.GetAttribute("type")
	if exists {
		t.Error("Expected GetAttribute to return false for non-existent attribute")
	}
	if value != "" {
		t.Errorf("Expected empty value for non-existent attribute, got '%s'", value)
	}

	tag.AddAttrib("type", "text")

	if !tag.HasAttribute("type") {
		t.Error("Expected HasAttribute to return true for existing attribute")
	}

	value, exists = tag.GetAttribute("type")
	if !exists {
		t.Error("Expected GetAttribute to return true for existing attribute")
	}
	if value != "text" {
		t.Errorf("Expected attribute value 'text', got '%s'", value)
	}

	tag.RemoveAttribute("type")
	if tag.HasAttribute("type") {
		t.Error("Expected HasAttribute to return false after removing attribute")
	}

	// Try removing a non-existent attribute
	tag.RemoveAttribute("nonexistent")
}

func TestCloneTag(t *testing.T) {
	original := NewTag("div")
	original.AddAttrib("id", "test")
	original.AddAttrib("class", "container")
	original.AddStyle("color", "red")
	original.AddStyle("background", "blue")
	original.SetContent("Original Content")

	// Add child to original
	child := NewTag("p")
	child.SetContent("Child content")
	original.AddChild(child)

	// Clone the tag
	clone := original.CloneTag()

	// Test basic properties
	if clone.GetName() != "div" {
		t.Errorf("Expected clone name 'div', got '%s'", clone.GetName())
	}

	if clone.GetContent() != "Original Content" {
		t.Errorf("Expected clone content 'Original Content', got '%s'", clone.GetContent())
	}

	// Test attributes are copied
	if !clone.HasAttribute("id") {
		t.Error("Expected clone to have 'id' attribute")
	}

	value, _ := clone.GetAttribute("id")
	if value != "test" {
		t.Errorf("Expected clone id 'test', got '%s'", value)
	}

	if !clone.HasAttribute("class") {
		t.Error("Expected clone to have 'class' attribute")
	}

	// Test that modifying the clone does not affect the original
	clone.SetContent("Modified Content")
	if original.GetContent() != "Original Content" {
		t.Error("Modifying clone affected original content")
	}

	clone.AddAttrib("id", "modified")
	if originalID, _ := original.GetAttribute("id"); originalID != "test" {
		t.Error("Modifying clone affected original attributes")
	}
}

func TestFindChildByName(t *testing.T) {
	parent := NewTag("div")
	child1 := NewTag("p")
	child2 := NewTag("span")
	child3 := NewTag("a")
	grandchild := NewTag("strong")

	child2.AddChild(grandchild)
	parent.AddChild(child1)
	parent.AddChild(child2)
	parent.AddChild(child3)

	// Test finding direct child
	found := parent.FindChildByName("span")
	if found == nil {
		t.Error("Expected to find 'span' child")
	}
	if found.GetName() != "span" {
		t.Errorf("Expected found tag name 'span', got '%s'", found.GetName())
	}

	// Test finding grandchild
	found = parent.FindChildByName("strong")
	if found == nil {
		t.Error("Expected to find 'strong' grandchild")
	}
	if found.GetName() != "strong" {
		t.Errorf("Expected found tag name 'strong', got '%s'", found.GetName())
	}

	// Test not finding non-existent tag
	found = parent.FindChildByName("img")
	if found != nil {
		t.Error("Expected not to find 'img' tag")
	}
}

func TestFindChildByAttribute(t *testing.T) {
	parent := NewTag("div")
	child1 := NewTag("p")
	child1.AddAttrib("id", "paragraph")
	child2 := NewTag("span")
	child2.AddAttrib("class", "highlight")
	child3 := NewTag("a")
	child3.AddAttrib("href", "http://example.com")
	grandchild := NewTag("strong")
	grandchild.AddAttrib("id", "emphasis")

	child2.AddChild(grandchild)
	parent.AddChild(child1)
	parent.AddChild(child2)
	parent.AddChild(child3)

	// Test finding by id
	found := parent.FindChildByAttribute("id", "paragraph")
	if found == nil {
		t.Error("Expected to find tag with id 'paragraph'")
	}
	if found.GetName() != "p" {
		t.Errorf("Expected found tag name 'p', got '%s'", found.GetName())
	}

	// Test finding by class
	found = parent.FindChildByAttribute("class", "highlight")
	if found == nil {
		t.Error("Expected to find tag with class 'highlight'")
	}
	if found.GetName() != "span" {
		t.Errorf("Expected found tag name 'span', got '%s'", found.GetName())
	}

	// Test finding grandchild by attribute
	found = parent.FindChildByAttribute("id", "emphasis")
	if found == nil {
		t.Error("Expected to find grandchild with id 'emphasis'")
	}
	if found.GetName() != "strong" {
		t.Errorf("Expected found tag name 'strong', got '%s'", found.GetName())
	}

	// Test not finding non-existent attribute
	found = parent.FindChildByAttribute("id", "nonexistent")
	if found != nil {
		t.Error("Expected not to find tag with non-existent id")
	}

	// Test not finding non-existent attribute name
	found = parent.FindChildByAttribute("nonexistent", "value")
	if found != nil {
		t.Error("Expected not to find tag with non-existent attribute name")
	}
}

func BenchmarkGetName(b *testing.B) {
	tag := NewTag("div")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tag.GetName()
	}
}

func BenchmarkCloneTag(b *testing.B) {
	tag := NewTag("div")
	tag.AddAttrib("id", "test")
	tag.AddAttrib("class", "container")
	tag.AddStyle("color", "red")
	tag.SetContent("Test content")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tag.CloneTag()
	}
}

func BenchmarkFindChildByName(b *testing.B) {
	parent := NewTag("div")
	for i := 0; i < 100; i++ {
		child := NewTag("span")
		if i == 50 {
			child = NewTag("target")
		}
		parent.AddChild(child)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parent.FindChildByName("target")
	}
}
