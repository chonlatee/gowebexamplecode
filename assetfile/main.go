package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("get current dir fail", err)
	}

	fs := http.FileServer(http.Dir(filepath.Join(cwd, "assetfile", "assets")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
