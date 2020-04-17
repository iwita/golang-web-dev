package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("starting-files/public"))
	http.HandleFunc("/", home)
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)

}
