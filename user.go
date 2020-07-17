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

//NewUser Creates a new user
func NewUser(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}
	w.Header().Set("Content-Type", "application/json")
	var user User
	
	//decodes the user from the body and turns it into data
	json.NewDecoder(r.Body).Decode(&user)

	db.Where("UUID = ?", user.UUID).FirstOrCreate(&user)
	json.NewEncoder(w).Encode(user)
}
