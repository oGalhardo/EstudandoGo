package main

import "fmt"

// type ol int

// var b ol

// func ex1() {
// x := 42
// y := "James Bond"
// z := true
// fmt.Println(x, y, z)
// fmt.Println(x)
// fmt.Println(z)
// fmt.Println(y)
// }

// func ex2() {
// 	fmt.Println(x, y, z)
// 	fmt.Println("Esses valores se chamam:valores zero")
// }

// var x int = 4
// var y string = " James Bond"
// var z bool

// func ex3() {
// 	s := fmt.Sprint(x, y, z)
// 	fmt.Println(s)
// }

type top int

var x top

func ex4() {
	fmt.Printf("valor de x=%v e seu tipo %T", x, x)
	x = 42
	fmt.Println(x)
}

func main() {
	ex4()
}
