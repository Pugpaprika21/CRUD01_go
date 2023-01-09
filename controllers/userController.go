package UserController

import (
	"database/sql"
	db "go_crud_2/database"
	"html"
	"html/template"
	_ "io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

type Users struct {
	USR_ID   int
	USR_NAME string
	USR_PASS string
}

//Create a global instance
//var tmplt *template.Template

func RowNumber(x, y int) int {
	return x + y
}

func DumpReq(r *http.Request, boolReq bool) ([]byte, error) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return reqDump, err
	}
	return reqDump, nil
}

func FormAddUser(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/form_add_user.html")
	t.Execute(w, nil)
}

func FormAddUserProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		usr_name := html.EscapeString(r.FormValue("USR_NAME"))
		usr_pass := html.EscapeString(r.FormValue("USR_PASS"))

		//reqDump, _ := DumpReq(r, true)
		// io.WriteString(w, string(reqDump))

		db, _ := sql.Open("mysql", db.Dsn())
		insertStmt, _ := db.Prepare("INSERT INTO USER_TB(USR_NAME, USR_PASS) VALUES(?, ?)")
		_, err := insertStmt.Exec(usr_name, usr_pass)

		if err != nil {
			panic(err.Error())
		} else {
			http.Redirect(w, r, "/formAddUser/ADD_USR?success="+strconv.Itoa(http.StatusOK), http.StatusSeeOther)
		}
		defer db.Close()
	}
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

		db, _ := sql.Open("mysql", db.Dsn())
		row := db.QueryRow("SELECT * FROM USER_TB WHERE USR_ID = ?", usr_id)

		u := new(Users)
		err := row.Scan(&u.USR_ID, &u.USR_NAME, &u.USR_PASS)

		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		t, _ := template.ParseFiles("template/show_user.html")
		t.Execute(w, u)
		defer db.Close()
	}
}
