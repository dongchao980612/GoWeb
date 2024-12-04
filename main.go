package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintln(w, r.URL.Path)+
		fmt.Fprintln(w, r.Header)
		fmt.Fprintln(w, r.Header["Accept-Encoding"])
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})
	http.ListenAndServe(":8080", nil)
}
