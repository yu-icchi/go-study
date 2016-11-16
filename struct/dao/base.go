package dao

import (
	"fmt"
	"../data"
)

type BaseDao struct {
	data data.Data
}

func (b *BaseDao) Find(id interface{}) {
	fmt.Println("BaseDao#Find", id, b.data.CollectionName(), b.data)
}

func (b *BaseDao) Remove() {
	fmt.Println("BaseDao#Remove")
}

func (b *BaseDao) Insert() {
	fmt.Println("BaseDao#Insert")
}

func (b *BaseDao) Update() {
	fmt.Println("BaseDao#Update")
}
