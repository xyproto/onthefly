package onthefly

import (
	"fmt"
	"log"
)

// Create a HTML5 page that links with Three.JS and sets up a scene
func NewThreeJS(titleText string) (*Page, *Tag) {
	page := NewHTML5Page(titleText)
	page.SetMargin(0)
	page.LinkToJSInBody("http://threejs.org/build/three.min.js")

	// Script for setting canvas css. The canvas tag is added by three.js.
	/*canvasStyling := `var sheet = document.createElement('style');
	                      sheet.innerHTML = "canvas {width: 100%; height: 100%; }";
	                      document.body.appendChild(sheet);`
		page.AddScriptToHead(canvasStyling)*/

	page.AddStyle("body { margin: 0; }; canvas { width: 100%; height: 100% }")

	// The script tag to be used for adding additional javascript code
	script, _ := page.AddScriptToBody("var scene = new THREE.Scene();")

	return page, script
}

// Add a camera with default settings
// todo: create an AddCustomCamera function
func (three *Tag) AddCamera() {
	three.AddContent("var camera = new THREE.PerspectiveCamera(75, window.innerWidth/window.innerHeigth, 0.1, 1000);")
}

// Add a WebGL renderer with default settings
// todo: create an AddCustomRenderer function
func (three *Tag) AddRenderer() {
	three.AddContent("var renderer = new THREE.WebGLRenderer();")
	three.AddContent("renderer.setSize(window.innerWidth, window.innerHeight);")
	three.AddContent("document.body.appendChild(renderer.domElement);")
}

type Mesh struct {
	id string
	js string
}

func (three *Tag) AddToScene(mesh *Mesh) {
	three.AddContent(mesh.js)
	three.AddContent("scene.add(" + mesh.id + ");")
}

func NewMesh(id, geometryid, materialid string) *Mesh {
	js := "var " + id + " = new THREE.Mesh(" + geometryid + ", " + materialid + ");"
	return &Mesh{id, js}
}

func (three *Tag) CameraPos(axis string, value int) {
	if (axis != "x") && (axis != "y") && (axis != "z") {
		log.Fatalln("camera axis must be x, y or z")
	}
	three.AddContent(fmt.Sprintf("camera.position.%s = %d;", axis, value))
}

// Add a test cube to the scene
// todo: create functions for adding geometry, material and creating meshes
func (three *Tag) AddTestCube() {
	three.AddContent("var geometry = new THREE.BoxGeometry(1, 1, 1);")
	three.AddContent("var material = new THREE.MeshBasicMaterial({color: 0x00ff00});")
	cube := NewMesh("cube", "geometry", "material")
	three.AddToScene(cube)
}

type RenderFunc struct {
	head, mid, tail string
}

func NewRenderFunction() *RenderFunc {
	head := "var render = function() { requestAnimationFrame(render);"
	tail := "renderer.render(scene, camera); };"
	return &RenderFunc{head, "", tail}
}

func (r *RenderFunc) Add(s string) {
	r.mid += s
}

func (three *Tag) AddRenderFunction(r *RenderFunc, call bool) {
	three.AddContent(r.head + r.mid + r.tail)
	if call {
		three.AddContent("render();")
	}
}

// Add a test scene to a scene
func ThreeTestPage() *Page {
	p, t := NewThreeJS("My first Three.js app")
	t.AddCamera()
	t.AddRenderer()
	t.AddTestCube()
	t.CameraPos("z", 5)

	r := NewRenderFunction()
	r.Add("cube.rotation.x += 0.1;")
	r.Add("cube.rotation.y += 0.1;")

	t.AddRenderFunction(r, true)
	return p
}
