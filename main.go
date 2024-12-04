package main

import (
	"net/http"
)

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello page!"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome page!"))
}

func main() {

	hh := helloHandler{}
	ah := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // DefaultServeMux
	}
	http.Handle("/hello", &hh) // 注册路由, 路径为/
	http.Handle("/about", &ah) // 注册路由, 路径为/about

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home page!"))
	})

	// http.HandleFunc("/welcome", welcome) // mux.handle(pattern, HandlerFunc(handler))
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()

}
