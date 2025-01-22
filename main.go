package main

import (
	"fmt"
	"reflect"
)

type ol int

var b ol

func main() {
	fmt.Println("Type of variable5:", reflect.TypeOf(b))
}
