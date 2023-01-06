package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func newAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Hahaha", News: "Some News"}
	t, err := template.ParseFiles("template/basictemplating.html")
	if err != nil {
		return
	}
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello GO</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)      // http://localhost:8080/
	http.HandleFunc("/agg/", newAggHandler) // http://localhost:8080/agg/
	http.ListenAndServe(":8080", nil)
}
