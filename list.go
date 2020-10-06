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
	UUID         string `json:"uuid" gorm:"primary_key"`
	Title        string `json:"title" gorm:"column:title"`
	ListMasterID string `json:"listMasterID" gorm:"column:listMasterID"`
}

//ListAndItemsAndListMembers struct
type ListAndItemsAndListMembers struct {
	Lists       []List
	Items       []Item
	ListMembers []ListMember
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
	// fmt.Fprintf(w, "All lists Endpoint Hit")
}

//DeleteAllLists - Deletes All Lists
func DeleteAllLists(w http.ResponseWriter, r *http.Request) {
	var lists []List

	db.Find(&lists)
	print("User length ", len(lists))

	for i, list := range lists {
		fmt.Print(i, list.Title)
		DeleteAllLMLocal(list)
		db.Delete(list)
	}
	println("Delete All Lists Hit")
}

//NewList Creates New List
func NewList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		panic("Could not connect to the database")
	}
	var list List
	json.NewDecoder(r.Body).Decode(&list)

	body, _ := ioutil.ReadAll(r.Body)
	//prints the body data
	fmt.Println(string(body))

	//converts user into json
	str, _ := json.Marshal(list)

	//prints the user json
	fmt.Println("printing JSON", string(str))

	db.Where("UUID = ?", list.UUID).FirstOrCreate(&list)
	CreateNewListMember(list.ListMasterID, list.UUID)
	json.NewEncoder(w).Encode(list)
}

//GetListsAndItemsAndLMsWith - Gets all the lists and items and listMembers for a specified UserID..and listID associated with the user ID
func GetListsAndItemsAndLMsWith(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("userID")
	var listMembersForUser []ListMember

	db.Where("userID = ?", userID).Find(&listMembersForUser)

	lists := make([]List, 0)

	for _, listMember := range listMembersForUser {
		var list List
		db.Where("UUID = ?", listMember.ListID).Find(&list)
		lists = append(lists, list)
	}

	items := make([]Item, 0)
	listMembers := make([]ListMember, 0)

	for _, list := range lists {
		var itemArray []Item
		var listMemberArray []ListMember
		db.Where("listID = ?", list.UUID).Find(&itemArray)
		db.Where("listID = ?", list.UUID).Find(&listMemberArray)
		items = append(items, itemArray...)
		listMembers = append(listMembers, listMemberArray...)
	}
	var listAndItems = ListAndItemsAndListMembers{lists, items, listMembers}
	println("GetListsAndItemsAndLMsWith Hit")
	json.NewEncoder(w).Encode(listAndItems)
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
	DeleteAllLMLocal(list)
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
	db.Model(&list).Updates(updatedList)

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

	var lists []List
	fmt.Println("Get List Hit")
	db.Where("ListMasterID = ?", uuid).Find(&lists)
	json.NewEncoder(w).Encode(&lists)
}
