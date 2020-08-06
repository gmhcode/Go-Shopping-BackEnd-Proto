package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Item - Object
type Item struct {
	UUID       string `json:"uuid" gorm:"primary_key"`
	Store      string `json:"store"`
	UserSentID string `json:"userSentId"`
	Name       string `json:"name"`
	ListID     string `json:"listID"`
}

//AllItems - returns all the items
func AllItems(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Could not connect to the database")
	}
	var items []Item

	db.Find(&items)
	json.NewEncoder(w).Encode(items)
	// fmt.Fprintf(w, "All Items Endpoint Hit")
}

//DeleteAllItems - Deletes All Items
func DeleteAllItems(w http.ResponseWriter, r *http.Request) {
	var items []Item

	db.Find(&items)
	print("User length ", len(items))

	for i, item := range items {
		fmt.Print(i, item.Name)
		db.Delete(item)
	}
}

//NewItem - Creates a new item
func NewItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	body, _ := ioutil.ReadAll(r.Body)
	//prints the body data
	fmt.Println(string(body))
	//converts user into json

	str, _ := json.Marshal(&item)
	//prints the user json
	fmt.Println(string(str))

	db.Where("UUID = ?", item.UUID).FirstOrCreate(&item)
	json.NewEncoder(w).Encode(&item)
}

//DeleteItem - Deletes item with the gived uuid
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var item Item

	db.Where("UUID = ?", id).Find(&item)
	db.Delete(&item)
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

//UpdateItem - updates the item in the body
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedItem Item
	var item Item
	//this turns updatedItem into the item in the r.body
	json.NewDecoder(r.Body).Decode(&updatedItem)

	db.Where("UUID = ?", updatedItem.UUID).Find(&item)

	item.Store = updatedItem.Store
	item.UserSentID = updatedItem.UserSentID
	item.Name = updatedItem.Name
	item.ListID = updatedItem.ListID

	db.Model(&item).Updates(updatedItem)

	json.NewEncoder(w).Encode(item)
	fmt.Fprintf(w, "Update User Endpoint Hit")

}

//GetItem - Responds with an item
func GetItem(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}

	vars := mux.Vars(r)
	uuid := vars["id"]

	var item Item

	db.Where("UUID = ?", uuid).Find(&item)
	json.NewEncoder(w).Encode(&item)

}
