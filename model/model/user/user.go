package user

import "fmt"

type User struct {
	ID string `bson:"_id"`
}

func (u *User) Collection() string {
	return "User"
}

func (u *User) FindByID(id string) {
	fmt.Println(u.Collection(), id)
}
