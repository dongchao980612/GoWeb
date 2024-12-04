package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		id := query.Get("id")
		name := query.Get("name")

		log.Printf("id: %s, name: %s", id, name)

	})

	http.ListenAndServe(":8080", nil)
}
