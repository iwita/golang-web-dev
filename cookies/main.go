package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

//var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
// }

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home")
}

func counter(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err != nil {
		log.Println(err)
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: "1",
		})
	}
	val, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln("Wrong cookie value")
	}
	val++

	// TODO
	// Verify why fmt.Fprintf() does not allow SetCookie to write back
	// to response writer
	c.Value = strconv.Itoa(val)
	http.SetCookie(w, c)
	io.WriteString(w, c.Value)

}

func main() {
	http.HandleFunc("/", counter)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
