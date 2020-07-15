package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	// EXAMPLE: 
	// curl -H 'Content-Type: application/json' --request POST -d '{"name":"Kratos", "todo":"Kill the Zeus"}' localhost:3000/create
	mux.HandleFunc("/", create())
	return mux
}