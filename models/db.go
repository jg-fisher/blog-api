package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// package level DB context
var db *sql.DB

func InitDB(connString string) {
	var err error
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
