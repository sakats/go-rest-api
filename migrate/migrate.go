package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbconn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbconn)
	dbconn.AutoMigrate(&model.User{}, &model.Task{})
}
