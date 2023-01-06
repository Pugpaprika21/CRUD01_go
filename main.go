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
	t, _ := template.ParseFiles("template/basictemplating.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello GO</h1>")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		p := NewsAggPage{Title: r.Method, News: "Login Error"}
		t, _ := template.ParseFiles("template/loginErr.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", indexHandler)      // http://localhost:8080/
	http.HandleFunc("/agg/", newAggHandler) // http://localhost:8080/agg/
	http.HandleFunc("/login/", loginPage)   // http://localhost:8080/login/
	http.ListenAndServe(":8080", nil)
}
