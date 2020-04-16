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

	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
