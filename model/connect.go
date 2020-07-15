package model

import (
	"database/sql"
	"fmt"
	"log"
)

// global connection to mysql db
var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:OmFericitsiVesel2!1@tcp(localhost:3306)/todos")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}