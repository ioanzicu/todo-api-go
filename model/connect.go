package model

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:OmFericitsiVesel2!1@/todos(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	return db
}