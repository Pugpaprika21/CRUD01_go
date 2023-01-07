package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type NewsAggPage struct {
	Title string
	News  string
}

type Login struct {
	Title  string
	Method string
}

type LoginSuccess struct {
	Username string
	Password string
	List     []string
}

type Users struct {
	USR_ID   int32
	USR_NAME string
	USR_PASS string
}

//Create a global instance
//var tmplt *template.Template

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "example_db"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
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
	t, _ := template.ParseFiles("template/login_form.html")
	t.Execute(w, nil)
}

func loginProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := Login{Title: "Login Error", Method: r.Method}
		t, _ := template.ParseFiles("template/login_error.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()

		makeList := []string{
			"Technology Explained",
			"Programming",
			"Linux",
			"Android",
			"iOS",
			"Many More................"}

		p := LoginSuccess{Username: r.FormValue("username"), Password: r.FormValue("password"), List: makeList}
		t, _ := template.ParseFiles("template/login_success.html")
		t.Execute(w, p)

		// fmt.Println("username:", r.Form["username"])
		// fmt.Println("password:", r.Form["password"])
	}
}

func getUsers() {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		fmt.Print("connected fail")
	} else {
		fmt.Print("connected")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM USER_TB")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var users Users
		err := rows.Scan(&users.USR_ID, &users.USR_NAME, &users.USR_PASS)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(users.USR_NAME)
	}
}

func main() {
	//getUsers()
	http.HandleFunc("/", indexHandler)         // http://localhost:8080/
	http.HandleFunc("/agg/", newAggHandler)    // http://localhost:8080/agg/
	http.HandleFunc("/login_form/", loginForm) // http://localhost:8080/login_form/
	http.HandleFunc("/login/", loginProcess)   // http://localhost:8080/login/
	http.ListenAndServe(":8080", nil)
}
