package routing

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "views/layouts"
var bootstrap *template.Template
var err error

// parsing
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gotpl")
	if err != nil {
		panic(err)
	}
	return files
}

// Source for handlers/middlewares
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "INDEX PAGE!") // send data to client side
}

func createUser(w http.ResponseWriter, r *http.Request) {
	bootstrap, err = template.ParseFiles(layoutFiles()...)
	if err != nil {
		panic(err)
	}

	bootstrap.ExecuteTemplate(w, "bootstrap", nil)
}
