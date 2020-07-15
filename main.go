package main

import (
	"log"
	"fmt"
	"net/http"
	"./controller"
	"./model"
	_ "github.com/go-sql-driver/mysql" // driver
)

// unexpected directory layout problem 
// Solution -> export GOPATH=$HOME/go/src/hello/todo_api

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}