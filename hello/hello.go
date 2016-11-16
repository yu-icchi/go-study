package main

import (
	"fmt"
	"reflect"
)

type Test1 struct {
	name string `json:"name"`
}

type Test2 struct {
	name string `json:"test"`
}

var typeRegistry = make(map[string]reflect.Type)

func init() {
	typeRegistry["Test1"] = reflect.TypeOf(Test1{})
	typeRegistry["Test2"] = reflect.TypeOf(Test2{})
}

func makeInstance(name string) interface{} {
	v := reflect.New(typeRegistry[name]).Elem()
	return v.Interface()
}

func main() {
	str := "Test2"
	fmt.Println(str)
	v := makeInstance(str)
	fmt.Print(reflect.TypeOf(v))
}
