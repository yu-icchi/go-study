package dao

import (
	"../data"
)

type AnnounceDao struct {
	BaseDao
}

func NewAnnounceDao() AnnounceDao {
	a := AnnounceDao{}
	a.data = data.NewAnnounce()
	return a
}
