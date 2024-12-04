package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(w, r, filepath.Join("www", r.URL.Path))
	})

	http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", http.FileServer(http.Dir("www")))
}
