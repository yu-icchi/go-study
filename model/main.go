package main

import (
	"./model"
	"fmt"
)

func main() {
	fmt.Println(model.User.Collection())
	fmt.Println(model.Family.Collection())

	model.User.FindByID("test_user")
	model.Family.FindByID("test_family")
	model.UserFile.Find()
	families, _ := model.Family.FindByIDs([]string{"1", "2"})
	for _, family := range families {
		fmt.Println(family.ID)
	}

	fmt.Println(model.Collections())
}
