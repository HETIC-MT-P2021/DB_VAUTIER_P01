package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var db *sql.DB

func InitializeDb(user string, password string, host string, name string, port int) {
	var err error
	dbUrl := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", user, password, port, name)
	db, err = sql.Open("mysql", dbUrl)

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    } else {
		fmt.Println("Connected to database")
	}
    // Close the connection after main function has finished
	defer db.Close()
}