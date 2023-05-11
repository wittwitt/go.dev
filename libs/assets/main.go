package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":3000", nil)
}
