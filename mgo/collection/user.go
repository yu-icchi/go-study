package collection

import "time"

type User struct {
	ID string `bson:"_id"`
	Time int64 `bson:"time"`
}

func (user *User) GetCollectionName() string {
	return "User"
}

func (user *User) GetTime() time.Time {
	return time.Unix(user.Time, 0)
}
