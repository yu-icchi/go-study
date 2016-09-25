package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id string `bson:"_id"`
	Name string	`bson:"name"`
	Phone string `bson:"phone"`
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"0001", "Ale", "+55 53 8116 9639"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}