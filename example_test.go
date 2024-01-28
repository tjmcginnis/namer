package namer_test

import (
	"fmt"

	"github.com/tjmcginnis/namer"
)

func ExampleNamer_Name() {
	n := namer.New()
	fmt.Println(n.Name())
}
