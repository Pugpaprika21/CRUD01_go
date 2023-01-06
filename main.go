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

type Login struct {
	Title  string
	Method string
}

func newAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Hahaha", News: "Some News"}
	t, _ := template.ParseFiles("template/basictemplating.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello GO</h1>")
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	message := "Welcome login Page"
	t, _ := template.ParseFiles("template/login_form.html")
	t.Execute(w, message)
}

func loginProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := Login{Title: "Login Error", Method: r.Method}
		t, _ := template.ParseFiles("template/login_error.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", indexHandler)         // http://localhost:8080/
	http.HandleFunc("/agg/", newAggHandler)    // http://localhost:8080/agg/
	http.HandleFunc("/login_form/", loginForm) // http://localhost:8080/login_form/
	http.HandleFunc("/login/", loginProcess)   // http://localhost:8080/login/
	http.ListenAndServe(":8080", nil)
}
