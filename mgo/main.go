package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type People struct {
	Id string `bson:"_id"`
	Name string	`bson:"name"`
	Phone string `bson:"phone"`
}

type Hoge struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Age int `bson:"age"`
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	db := session.DB("test")

	result := bson.M{}
	query := bson.M{}
	key := "_id"
	value := "ok3"
	query[key] = value
	err = db.C("hoge").Find(query).One(&result)
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

	//hoge := []bson.M{}
	//err = db.C("hoge").Find(bson.M{}).All(&hoge)
	//if err != nil {
	//	panic(err)
	//}
	//jh, _ := json.Marshal(hoge)
	//fmt.Println(string(jh))
}