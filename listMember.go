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
	ListID string `json:"listID" gorm:"column:listID"`
	UserID string `json:"userID" gorm:"column:userID"`
	UUID   string `json:"uuid" gorm:"primary_key"`
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
	fmt.Println(listMember.UUID)
	db.Where("UUID = ?", listMember.UUID).FirstOrCreate(&listMember)
	// fmt.Println(string(str))
	json.NewEncoder(w).Encode(listMember)
}

//DeleteAllListMembers - Deletes All ListMembers
func DeleteAllListMembers(w http.ResponseWriter, r *http.Request) {
	var listMembers []ListMember

	db.Find(&listMembers)
	print("listmember length ", len(listMembers))

	for i, listMember := range listMembers {
		fmt.Print(i, listMember)
		db.Delete(listMember)
	}
}

//CreateNewListMember - creates a new ListMember
func CreateNewListMember(uID string, lID string) {

	var listMember = ListMember{UserID: uID, ListID: lID, UUID: uID + lID}

	db.Where("UUID = ?", listMember.UUID).FirstOrCreate(listMember)

	str, _ := json.Marshal(listMember)
	//prints the user json
	fmt.Println("listMember Created", string(str))
}

//DeleteListMember - Deletes ListMember
func DeleteListMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var listMember ListMember
	db.Where("UUID = ?", id).Find(&listMember)
	db.Delete(listMember)
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

//UpdateListMember - Updates List Member
func UpdateListMember(w http.ResponseWriter, r *http.Request) {

}

// //GetListMembersWithList - Gets All list members for specified list
// func GetListMembersWithList(list List) []ListMember {
// 	listMembers := make([]ListMember, 0)

// 	db.Where("")
// }

//GetListMember - Gets a specific list member

func GetListMember(w http.ResponseWriter, r *http.Request) {

}
