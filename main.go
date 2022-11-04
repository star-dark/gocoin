package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/star-dark/gocoin/blockchain"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlock()}
	templates.ExecuteTemplate(rw, "home", data)
}
func add(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "add", nil)
}
func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partial/*.gohtml"))
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
	fmt.Printf("Listening on http://localhost%s\n", port)
}
