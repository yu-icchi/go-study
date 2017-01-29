package bar

import (
	"fmt"
)

type Bar struct {

}

func (b Bar) P() {
	fmt.Println("bar")
}
