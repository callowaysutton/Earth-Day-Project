package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Now Listening on 80")
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	switch p {
	case "./":
		p = "./index.html"
	case "./class/":
		p = "./class/index.html"
	case "./econ/":
		p = "./econ/index.html"
	case "./sci/":
		p = "./sci/index.html"
	default:
		p = "./error/404.html"
	}
	http.ServeFile(w, r, p)
}
