package main

import (
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
	}
}

func main() {
	http.HandleFunc("/login/", LoginProcess)                                   // http://localhost:8080/login/
	http.HandleFunc("/formAddUser/", UserController.FormAddUser)               // http://localhost:8080/formAddUser/
	http.HandleFunc("/formAddUserProcess/", UserController.FormAddUserProcess) // http://localhost:8080/formAddUserProcess/
	http.HandleFunc("/show_users/", UserController.ShowUsers)                  // http://localhost:8080/show_users/
	http.HandleFunc("/show_user/", UserController.ShowUser)                    // http://localhost:8080/show_user/
	http.HandleFunc("/formUpdateUser/", UserController.FormUpdateUser)         // http://localhost:8080/formUpdateUser/
	http.HandleFunc("/updateUserProcess/", UserController.UpdateUserProcess)   // http://localhost:8080/updateUserProcess/
	http.ListenAndServe(":8080", nil)
}
