package UserController

import (
	"database/sql"
	db "go_crud_2/database"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Users struct {
	USR_ID   int
	USR_NAME string
	USR_PASS string
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
