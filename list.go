package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

//List struct
type List struct {
	gorm.Model
	uuid         string `gorm:"unique;not null"`
	title        string
	listMasterID string
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
