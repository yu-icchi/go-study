package data

type User struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Age int `bson:"age"`
	Date int64 `bson:"time"`
}

func NewUser() *User {
	return &User{"", "", 0, 0}
}

func (u *User) CollectionName() string {
	return "User"
}
