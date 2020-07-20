package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//User Struct
type User struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
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

	body, _ := ioutil.ReadAll(r.Body)
	//prints the body data
	fmt.Println(string(body))

	//converts user into json
	str, _ := json.Marshal(&user)
	//prints the user json
	fmt.Println(string(str))

	db.Where("UUID = ?", user.UUID).FirstOrCreate(&user)
	json.NewEncoder(w).Encode(user)
}

//GetUser - responds with a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var user User

	db.Where("UUID = ?", id).Find(&user)
	json.NewEncoder(w).Encode(&user)
}

//DeleteUser - Deletes user with given ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Error in NewUser")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var user User

	db.Where("UUID = ?", id).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

//UpdateUser - Updates user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var userUpdates User
	var user User

	json.NewDecoder(r.Body).Decode(&userUpdates)

	db.Where("UUID = ?", userUpdates.UUID).Find(&user)

	user.Name = userUpdates.Name
	user.Email = userUpdates.Email

	db.Save(&user)
	json.NewEncoder(w).Encode(user)
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
