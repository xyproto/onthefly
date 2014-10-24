package onthefly

import (
	"fmt"
	"log"
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

// Add a camera with default settings
// todo: create an AddCustomCamera function
func (three *Tag) AddCamera() {
	three.AddContent("var camera = new THREE.PerspectiveCamera(75, window.innerWidth/window.innerHeight, 0.1, 1000);")
}

// Add a WebGL renderer with default settings
func (three *Tag) AddRenderer() {
	three.AddContent("var renderer = new THREE.WebGLRenderer();")
	three.AddContent("renderer.setSize(window.innerWidth, window.innerHeight);")
	three.AddContent("document.body.appendChild(renderer.domElement);")
}

type Mesh struct {
	Id string // name of the mesh variable
	Js string // javascript code for creating the mesh
}

func (three *Tag) AddToScene(mesh *Mesh) {
	three.AddContent(mesh.Js)
	three.AddContent("scene.add(" + mesh.Id + ");")
}

func NewMesh(id string, geometry *Geometry, material *Material) *Mesh {
	js := geometry.Js + material.Js
	js += "var " + id + " = new THREE.Mesh(" + geometry.Id + ", " + material.Id + ");"
	return &Mesh{id, js}
}

func (three *Tag) CameraPos(axis string, value int) {
	if (axis != "x") && (axis != "y") && (axis != "z") {
		log.Fatalln("camera axis must be x, y or z")
	}
	three.AddContent(fmt.Sprintf("camera.position.%s = %d;", axis, value))
}

type Material struct {
	Id string // name of the material variable
	Js string // javascript code for creating the material
}

// Very simple type of material
func NewMaterial(id, color string) *Material {
	js := "var " + id + " = new THREE.MeshBasicMaterial({color: " + color + "});"
	return &Material{id, js}
}

type Geometry struct {
	Id string // name of the geometry variable
	Js string // javascript code for creating the geometry
}

func NewBoxGeometry(id string, w, h, d int) *Geometry {
	js := fmt.Sprintf("var %s = new THREE.BoxGeometry(%d, %d, %d);", id, w, h, d)
	return &Geometry{id, js}
}

// Add a test cube to the scene
// todo: create functions for adding geometry, material and creating meshes
func (three *Tag) AddTestCube(id, color string) *Mesh {
	material := NewMaterial(id+"_material", color)
	geometry := NewBoxGeometry(id+"_geometry", 1, 1, 1)
	cube := NewMesh(id, geometry, material)
	three.AddToScene(cube)
	return cube
}

type RenderFunc struct {
	head, mid, tail string
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
