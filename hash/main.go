package main

import (
	"github.com/mitchellh/hashstructure"
	"fmt"
)

type ComplexStruct struct {
	Name     string
	Age      uint
	Metadata map[string]interface{}
}

func main() {
	v := ComplexStruct{
		Metadata: map[string]interface{}{
			"siblings": []string{"Bob", "John"},
			"location": "California",
			"car":      true,
		},
		Name: "mitchellh",
		Age:  64,
	}

	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x", hash)
	// 6691276962590150517
	// 6691276962590150517
	// 6691276962590150517
}