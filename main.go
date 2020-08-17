package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/query", GetUsersWith).Methods("GET").Queries("listID", "{listID}")
	myRouter.HandleFunc("/user/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/user", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/users", DeleteAllUsers).Methods("DELETE")

	myRouter.HandleFunc("/lists", AllLists).Methods("GET")
	myRouter.HandleFunc("/lists/query", GetListsAndItemsAndLMsWith).Methods("GET").Queries("userID", "{userID}")
	myRouter.HandleFunc("/list", NewList).Methods("POST")
	myRouter.HandleFunc("/list/{id}", DeleteList).Methods("DELETE")
	myRouter.HandleFunc("/list", UpdateList).Methods("PUT")
	myRouter.HandleFunc("/list/{id}", GetList).Methods("GET")
	myRouter.HandleFunc("/lists", DeleteAllLists).Methods("DELETE")

	myRouter.HandleFunc("/items", AllItems).Methods("GET")
	myRouter.HandleFunc("/items/query", GetItemsWith).Methods("GET").Queries("userID", "{userID}", "listID", "{listID}")
	myRouter.HandleFunc("/item", NewItem).Methods("POST")
	myRouter.HandleFunc("/item/{id}", DeleteItem).Methods("DELETE")
	myRouter.HandleFunc("/item", UpdateItem).Methods("PUT")
	myRouter.HandleFunc("/item/{id}", GetItem).Methods("GET")
	myRouter.HandleFunc("/items", DeleteAllItems).Methods("DELETE")

	myRouter.HandleFunc("/listMembers", AllListMembers).Methods("GET")
	myRouter.HandleFunc("/listMember", NewListMember).Methods("POST")
	myRouter.HandleFunc("/listMember/{id}", DeleteListMember).Methods("DELETE")
	myRouter.HandleFunc("/listMember", UpdateListMember).Methods("PUT")
	myRouter.HandleFunc("/listMember/{id}", GetListMember).Methods("GET")
	myRouter.HandleFunc("/listMembers", DeleteAllListMembers).Methods("DELETE")

	myRouter.HandleFunc("/deleteAll", DeleteAll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	println("Go ORM Tutorial")
	InitialMigration()
	handleRequest()

	defer db.Close()
}
