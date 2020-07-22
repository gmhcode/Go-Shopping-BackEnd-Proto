package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

//InitialMigration sets up the database
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "shopping.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{}, &List{}, &Item{}, &ListMember{})
}

//DeleteAll - Deletes All
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	var lists []List
	db.Find(&lists)
	for i, list := range lists {
		fmt.Print(i, list.Title)
		db.Delete(list)
	}
	var users []User
	db.Find(&users)
	for i, user := range users {
		fmt.Print(i, user.Name)
		db.Delete(user)
	}
	var items []Item
	db.Find(&items)
	for i, item := range items {
		fmt.Print(i, item.Name)
		db.Delete(item)
	}
	var listMembers []ListMember
	db.Find(&listMembers)
	for i, listMember := range listMembers {
		fmt.Print(i, listMember)
		db.Delete(listMember)
	}
}
