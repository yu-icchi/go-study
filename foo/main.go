package main

import (
	"./bar"
	"./foo"
	"./buz"
)

func main() {
	buz.Call(foo.Foo{})
	buz.Call(bar.Bar{})
}
