package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//User Struct
type User struct {
	gorm.Model
	UUID       string 
	listID     string
	store      string
	userSentID string
	name       string
}

//AllUsers Returns all the users
func AllUsers(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		panic("Could not connect to the database")
	}
	//Create an empty array of users
	var users []User

	//Finds all users
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
	fmt.Fprintf(w, "All Users Endpoint Hit")
}