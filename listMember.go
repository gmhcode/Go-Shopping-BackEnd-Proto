package main

import (
	"encoding/json"
	"fmt"
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

// GetListMembersWithLists - Gets all the listMembers for the provided lists
func GetListMembersWithLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var lists []List

	json.NewDecoder(r.Body).Decode(&lists)

	var listMembers = make([]ListMember, 0)

	for _, list := range lists {
		var listMemberArray []ListMember
		db.Where("listID = ?", list.UUID).Find(&listMemberArray)
		listMembers = append(listMembers, listMemberArray...)
	}
	json.NewEncoder(w).Encode(listMembers)
}

//NewListMember - Creates and saves new ListMember (adds a member to a list)
func NewListMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var listMember ListMember

	json.NewDecoder(r.Body).Decode(&listMember)

	//prints the body data
	fmt.Println("printing body")
	fmt.Println(listMember.UserID)
	// listMember.UUID = listMember.UserID + listMember.ListID
	listMember = CreateNewListMember(listMember.UserID, listMember.ListID)
	//prints the user json
	// db.Where("UUID = ?", listMember.UUID).FirstOrCreate(&listMember)

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

//DeleteAllLMLocal - locally deletes all list members
func DeleteAllLMLocal(list List) {
	var listMembers []ListMember

	db.Where("ListID = ?", list.UUID).Find(&listMembers)

	for i, listMember := range listMembers {
		fmt.Print("Deleting: ", i, list.Title)
		db.Delete(listMember)
	}
	// db.Delete(list)
}

//CreateNewListMember - creates a new ListMember
func CreateNewListMember(uID string, lID string) ListMember {
	var listMember = ListMember{UserID: uID, ListID: lID, UUID: uID + lID}

	db.Where("UUID = ?", listMember.UUID).FirstOrCreate(listMember)
	str, _ := json.Marshal(listMember)
	//prints the user json
	fmt.Println("listMember Created", string(str))
	return listMember
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
