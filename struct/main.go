package main

import (
	"fmt"
	"./data"
	"./dao"
)

func main() {
	fmt.Println("逃げ恥")

	user := data.NewUserData()
	showCollection(user)

	announce := data.NewAnnounceData()
	showCollection(announce)

	userDao := dao.NewUserDao()
	userDao.Find("user")
	userDao.FindById()

	announceDao := dao.NewAnnounceDao()
	announceDao.Find("announce")
}

func showCollection(data data.Data) {
	fmt.Println(data.CollectionName())
}
