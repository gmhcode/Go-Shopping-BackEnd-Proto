package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/user", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user", UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/lists", AllLists).Methods("GET")
	myRouter.HandleFunc("/list", NewList).Methods("POST")
	myRouter.HandleFunc("/list/{id}", DeleteList).Methods("DELETE")
	myRouter.HandleFunc("/list", UpdateList).Methods("PUT")
	myRouter.HandleFunc("/list/{id}", GetList).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	println("Go ORM Tutorial")
	InitialMigration()
	handleRequest()
	defer db.Close()
}
