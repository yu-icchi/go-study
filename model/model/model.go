package model

import (
	"./admin"
	"./user"
	"./family"
	"reflect"
)

var (
	M = newModel()
)

type Model struct {
	User user.User
	UserFile user.File
	Family family.Family
	AdminRole admin.Role
}

func (m Model) Collections() []string {
	v := reflect.Indirect(reflect.ValueOf(m))
	t := v.Type()

	cols := []string{}
	for i := 0; i < t.NumField(); i++ {
		cols = append(cols, t.Field(i).Name)
	}
	return cols
}

func newModel() Model {
	return Model{
		User: user.User{},
		UserFile: user.File{},
		Family: family.Family{},
		AdminRole: admin.Role{},
	}
}
