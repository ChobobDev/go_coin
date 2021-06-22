package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ChobobDev/go_coin/blockchain"
)

const {
	port string = ":4000"
	templateDir string ="/templates"
}
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob())
	data := homeData{"버니 버니 바니 바니 고랭고랭", blockchain.GetBlockchain().AllBlocks()}

	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Litening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
