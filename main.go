package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	println("Go ORM Tutorial")
	InitialMigration()
	handleRequest()
	defer db.Close()
}
