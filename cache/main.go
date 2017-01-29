package main

import (
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
)

type Hoge struct {
	ID string
	Age int
}

func main() {

	hoge := Hoge{"test", 24}

	c := cache.New(5*time.Minute, 30*time.Second)
	c.Set("foo", &hoge, cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		f := foo.(*Hoge)
		fmt.Println(f.ID)
		fmt.Println(f.Age)
	}

}
