package family

import "fmt"

type Family struct {
	ID string `bson:"_id"`
}

func (f Family) Collection() string {
	return "Family"
}

func (f Family) FindByID(id string) {
	fmt.Println(f.Collection(), id)
}
