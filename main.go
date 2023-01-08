package main

import (
	"fmt"
	UserController "go_crud_2/controllers"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Login struct {
	Title  string
	Method string
}

type LoginSuccess struct {
	Username string
	Password string
	List     []string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello GO</h1>")
}

func LoginForm(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/login_form.html")
	t.Execute(w, nil)
}

func LoginProcess(w http.ResponseWriter, r *http.Request) {
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
		// err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)                        // http://localhost:8080/
	http.HandleFunc("/login_form/", LoginForm)                // http://localhost:8080/login_form/
	http.HandleFunc("/login/", LoginProcess)                  // http://localhost:8080/login/
	http.HandleFunc("/show_users/", UserController.ShowUsers) // http://localhost:8080/show_users/
	http.HandleFunc("/show_user/", UserController.ShowUser)   // http://localhost:8080/show_user/
	http.ListenAndServe(":8080", nil)
}
