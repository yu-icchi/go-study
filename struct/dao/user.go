package dao

import (
	"fmt"
	"../data"
)

type UserDao struct {
	BaseDao
}

func NewUserDao() UserDao {
	u := UserDao{}
	u.data = data.NewUserData()
	return u
}

func (u *UserDao) FindById() {
	fmt.Println("UserDao#FindById", u.data.CollectionName(), u.data)
}
