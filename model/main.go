package main

import (
	db "./model"
	"fmt"
)

func main() {
	fmt.Println(db.M.User.Collection())
	fmt.Println(db.M.Family.Collection())

	db.M.User.FindByID("test_user")
	db.M.Family.FindByID("test_family")
	db.M.UserFile.Find()

	fmt.Println(db.M.Collections())
}
