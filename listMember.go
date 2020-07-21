package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//ListMember - Connects the lists to the users, so we can know which user is a member of their respective lists
type ListMember struct {
	ListID string `json:"listID"`
	UserID string `json:"userID"`
	ID     string `json:"id"`
}

//AllListMembers - Returns all list members
func AllListMembers(w http.ResponseWriter, r *http.Request) {
	var listMembers []ListMember

	db.Find(&listMembers)
	json.NewEncoder(w).Encode(listMembers)
	// fmt.Fprintf(w, "All listMembers Endpoint Hit")
}

//NewListMember - Creates and saves new ListMember
func NewListMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var listMember ListMember

	json.NewDecoder(r.Body).Decode(&listMember)

	body, _ := ioutil.ReadAll(r.Body)
	//prints the body data
	fmt.Println(string(body))

	// str, _ := json.Marshal(listMember)
	//prints the user json
	fmt.Println(listMember.ID)
	db.Where("ID = ?", listMember.ID).FirstOrCreate(&listMember)
	// fmt.Println(string(str))
	json.NewEncoder(w).Encode(listMember)
}

//CreateNewListMember - creates a new ListMember
func CreateNewListMember(uID string, lID string) {

	var listMember = ListMember{UserID: uID, ListID: lID, ID: uID + lID}

	db.Where("ID = ?", listMember.ID).FirstOrCreate(listMember)

	str, _ := json.Marshal(listMember)
	//prints the user json
	fmt.Println("listMember Created", string(str))
}

//DeleteListMember - Deletes ListMember
func DeleteListMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var listMember ListMember
	db.Where("ID = ?", id).Find(&listMember)
	db.Delete(listMember)
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

//UpdateListMember - Updates List Member
func UpdateListMember(w http.ResponseWriter, r *http.Request) {

}

//GetListMember - Gets a specific list member
func GetListMember(w http.ResponseWriter, r *http.Request) {

}
