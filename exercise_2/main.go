package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "home")
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "me")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
