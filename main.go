package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user", NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	println("Go ORM Tutorial")
	InitialMigration()
	handleRequest()
	defer db.Close()
}
