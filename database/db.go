package database

import "fmt"

func Dsn() string {
	username := "root"
	password := ""
	hostname := "127.0.0.1:3306"
	dbname := "example_db"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
