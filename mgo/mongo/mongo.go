package mongo

import (
	"fmt"
)

type mongodb struct {
}

func (mongo *mongodb) Collection(name string) {
	fmt.Println("col:", name)
}

var instance *mongodb = new(mongodb)

func GetInstance() *mongodb {
	if instance == nil {
		instance = new(mongodb)
		fmt.Println("Create MongoDB Instance")
	} else {
		fmt.Println("Not Create MongoDB Instance")
	}
	return instance
}
