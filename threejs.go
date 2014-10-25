package onthefly

import (
	"fmt"
)

// For generating IDs
var (
	geometryCounter = 0
	materialCounter = 0
)

// Unique prefixes when generating IDs
const (
	geometryPrefix = "g"
	materialPrefix = "ma"
)

type (
	// For Three.JS elements, like a mesh or material
	Element struct {
		ID string // name of the variable
		JS string // javascript code for creating the element
	}
	// The Three.JS render function, where heand and tail are standard
	RenderFunc struct {
		head, mid, tail string
	}
	// Different types of elements. Mainly used for passing types safely around.
	Geometry Element
	Material Element
)

// Create a HTML5 page that links with Three.JS and sets up a scene
func NewThreeJS(titleText string) (*Page, *Tag) {
	page := NewHTML5Page(titleText)

	// Style the page for showing a fullscreen canvas
	page.FullCanvas()

	// Link to Three.JS
	page.LinkToJSInBody("http://threejs.org/build/three.min.js")

	// Add a scene
	script, _ := page.AddScriptToBody("var scene = new THREE.Scene();")

	// Return the sript tag that can be used for adding additional javascript/Three.JS code
	return page, script
}

// Add a WebGL renderer with default settings
func (three *Tag) AddRenderer() {
	three.AddContent("var renderer = new THREE.WebGLRenderer();")
	three.AddContent("renderer.setSize(window.innerWidth, window.innerHeight);")
	three.AddContent("document.body.appendChild(renderer.domElement);")
}

func (three *Tag) AddToScene(mesh *Mesh) {
	three.AddContent(mesh.JS)
	three.AddContent("scene.add(" + mesh.ID + ");")
}

// Very simple type of material
func NewMaterial(color string) *Material {
	id := fmt.Sprintf("%s%d", materialPrefix, materialCounter)
	materialCounter++
	js := "var " + id + " = new THREE.MeshBasicMaterial({color: " + color + "});"
	return &Material{id, js}
}

func NewNormalMaterial() *Material {
	id := fmt.Sprintf("%s%d", materialPrefix, materialCounter)
	materialCounter++
	js := "var " + id + " = new THREE.MeshNormalMaterial();"
	return &Material{id, js}
}

func NewBoxGeometry(w, h, d int) *Geometry {
	id := fmt.Sprintf("%s%d", geometryPrefix, geometryCounter)
	geometryCounter++
	js := fmt.Sprintf("var %s = new THREE.BoxGeometry(%d, %d, %d);", id, w, h, d)
	return &Geometry{id, js}
}

// Add a test cube to the scene
// todo: create functions for adding geometry, material and creating meshes
func (three *Tag) AddTestCube() *Mesh {
	//material := NewMaterial(color)
	material := NewNormalMaterial()
	geometry := NewBoxGeometry(1, 1, 1)
	cube := NewMesh(geometry, material, true)
	three.AddToScene(cube)
	return cube
}

func NewRenderFunction() *RenderFunc {
	head := "var render = function() { requestAnimationFrame(render);"
	tail := "renderer.render(scene, camera); };"
	return &RenderFunc{head, "", tail}
}

func (r *RenderFunc) AddJS(s string) {
	r.mid += s
}

func (three *Tag) AddRenderFunction(r *RenderFunc, call bool) {
	three.AddContent(r.head + r.mid + r.tail)
	if call {
		three.AddContent("render();")
	}
}
