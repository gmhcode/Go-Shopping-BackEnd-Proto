package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//List struct
type List struct {
	UUID         string `json:"uuid"`
	Title        string `json:"title"`
	ListMasterID string `json:"listMasterID"`
}

//AllLists Returns all the users
func AllLists(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		panic("Could not connect to the database")
	}
	//Create an empty array of users
	var lists []List

	//Finds all users
	db.Find(&lists)
	json.NewEncoder(w).Encode(lists)
	fmt.Fprintf(w, "All lists Endpoint Hit")
}

//DeleteAllLists - Deletes All Lists
func DeleteAllLists(w http.ResponseWriter, r *http.Request) {
	var lists []List

	db.Find(&lists)
	print("User length ", len(lists))

	for i, list := range lists {
		fmt.Print(i, list.Title)
		db.Delete(list)
	}
}

//NewList Creates New List
func NewList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var list List
	json.NewDecoder(r.Body).Decode(&list)

	body, _ := ioutil.ReadAll(r.Body)
	//prints the body data
	fmt.Println(string(body))

	//converts user into json
	str, _ := json.Marshal(list)
	//prints the user json
	fmt.Println(string(str))

	db.Where("UUID = ?", list.UUID).FirstOrCreate(&list)
	CreateNewListMember(list.ListMasterID, list.UUID)
	json.NewEncoder(w).Encode(list)
}

//DeleteList - deletes list from ID
func DeleteList(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var list List

	db.Where("UUID = ?", id).Find(&list)
	db.Delete(&list)
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

//UpdateList - Updates List
func UpdateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedList List
	var list List

	json.NewDecoder(r.Body).Decode(&updatedList)

	db.Where("UUID = ?", updatedList.UUID).Find(&list)

	list.UUID = updatedList.UUID
	list.Title = updatedList.Title
	list.ListMasterID = updatedList.ListMasterID

	db.Save(&list)
	json.NewEncoder(w).Encode(list)
	fmt.Fprintf(w, "Update User Endpoint Hit")

}

//GetList - responds with a list
func GetList(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}

	vars := mux.Vars(r)
	uuid := vars["id"]

	var list List

	db.Where("UUID = ?", uuid).Find(&list)
	json.NewEncoder(w).Encode(&list)
}
