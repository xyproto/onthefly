package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/onthefly"
)

// Set up the paths and handlers then start serving.
func main() {
	fmt.Println("onthefly ", onthefly.Version)

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	// Publish the generated page
	page := onthefly.ThreeTestPage()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, page.String())
	})

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 8080
	n.Run(":8080")
}
