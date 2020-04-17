package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/", home)
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.ListenAndServe(":8080", nil)
}
