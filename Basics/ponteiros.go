package main

import "fmt"

func main() {
	x := 12
	y := &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*y)

}
