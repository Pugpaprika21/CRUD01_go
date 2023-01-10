package main

import (
	UserController "go_crud_2/controllers"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", IndexHandler)                                         // http://localhost:8080/login/
	http.HandleFunc("/formAddUser/", UserController.FormAddUser)               // http://localhost:8080/formAddUser/
	http.HandleFunc("/formAddUserProcess/", UserController.FormAddUserProcess) // http://localhost:8080/formAddUserProcess/
	http.HandleFunc("/show_users/", UserController.ShowUsers)                  // http://localhost:8080/show_users/
	http.HandleFunc("/show_user/", UserController.ShowUser)                    // http://localhost:8080/show_user/
	http.HandleFunc("/formUpdateUser/", UserController.FormUpdateUser)         // http://localhost:8080/formUpdateUser/
	http.HandleFunc("/updateUserProcess/", UserController.UpdateUserProcess)   // http://localhost:8080/updateUserProcess/
	http.ListenAndServe(":8080", nil)
}
