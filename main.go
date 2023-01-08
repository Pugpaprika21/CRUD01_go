package main

import (
	"database/sql"
	"fmt"
	db "go_crud_2/database"
	"html/template"
	"io"
	"log"
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

type Users struct {
	USR_ID   int
	USR_NAME string
	USR_PASS string
}

//Create a global instance
//var tmplt *template.Template

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

func RowNumber(x, y int) int {
	return x + y
}

func ShowUsers(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("mysql", db.Dsn())
	rows, err := db.Query("SELECT * FROM USER_TB")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	u := Users{}
	showUsers := []Users{}

	for rows.Next() {
		var usrId int
		var usrName string
		var usrPass string
		err := rows.Scan(&usrId, &usrName, &usrPass)
		if err != nil {
			log.Println(err)
			http.Error(w, "there was an error", http.StatusInternalServerError)
			return
		}
		u.USR_ID = usrId
		u.USR_NAME = usrName
		u.USR_PASS = usrPass
		showUsers = append(showUsers, u)
	}

	row := template.FuncMap{"RowNumber": RowNumber}
	t := template.Must(template.New("show_users.html").Funcs(row).ParseFiles("template/show_users.html"))
	t.Execute(w, showUsers)
	defer db.Close()
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		usr_id := r.Form.Get("USR_ID")
		io.WriteString(w, "USR_ID: "+usr_id)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)         // http://localhost:8080/
	http.HandleFunc("/login_form/", LoginForm) // http://localhost:8080/login_form/
	http.HandleFunc("/login/", LoginProcess)   // http://localhost:8080/login/
	http.HandleFunc("/show_users/", ShowUsers) // http://localhost:8080/show_users/
	http.HandleFunc("/show_user/", ShowUser)   // http://localhost:8080/show_user/
	http.ListenAndServe(":8080", nil)
}
