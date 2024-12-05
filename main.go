package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "application/x-www-form-urlencoded")
		r.ParseForm()
		fmt.Fprintln(w, "Form", r.Form)
		fmt.Fprintln(w, "PostForm", r.PostForm)
		fmt.Fprintln(w, "multipart/form-data")
		fmt.Fprintln(w, "MultipartForm", r.MultipartForm)
		r.ParseMultipartForm(1024)
		fmt.Fprintln(w, r.MultipartForm)

		fmt.Fprintln(w, "application/x-www-form-urlencoded")
		fmt.Fprintln(w, r.FormValue("username"))
		fmt.Fprintln(w, r.PostFormValue("username"))
	})

	http.ListenAndServe(":8080", nil)
}
