package foo

import (
	"fmt"
)

type Foo struct {}

func (f Foo) P() {
	fmt.Println("foo")
}
