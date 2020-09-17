package models

import (
	"database/sql"
	"fmt"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitializeDb(user string, password string, host string, name string, port int) {
	var err error
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	db, err = sql.Open("mysql", dbUrl)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to database")
	}

	// defer db.Close()
}

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}