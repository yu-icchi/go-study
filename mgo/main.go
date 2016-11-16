package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"./mongo"
)

type People struct {
	Id string `bson:"_id"`
	FullName string `bson:"full_name"`
	PhoneNumberString string `bson:"phone"`
}

type Hoge struct {
	Id string `bson:"_id",json:"_id"`
	Name string `bson:"name",json:"name"`
	Age int `bson:"age",json:"age"`
	People []People
}

func (hoge *Hoge) GetCollectionName() string {
	return "hoge"
}

const (
	_ = iota
	A
	B
	C
)

func main() {

	fmt.Println(A, B, C)

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{"localhost:27017"},
		Database: "test",
		Timeout: 20 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	// 読み書きをプライマリにする指定(defaultでこうなっている)
	session.SetMode(mgo.Strong, true)

	db := session.DB("") // 空文字指定の場合はデフォルトDBになる仕様らしい

	hoge := Hoge{}
	if err = db.C(hoge.GetCollectionName()).Find(bson.M{"_id": "ok3"}).One(&hoge); err != nil {
		fmt.Println(err)
	}
	fmt.Println(hoge.Id, hoge.Age, hoge.Name, hoge.People)

	instance := mongo.GetInstance();
	instance.Collection("name")
}