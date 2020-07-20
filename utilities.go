package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//InitialMigration sets up the database
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "shopping.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{}, &List{}, &Item{})
}
