package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello dog")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Main page")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Achilleas")
}
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
