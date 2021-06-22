package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from Home !!!!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Litening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
