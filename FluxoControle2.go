package main

import "fmt"

// func ex1() {
// 	x := 10
// 	if x < 10 {
// 		fmt.Println(x)
// 	}
// 	if !(x < 10) {
// 		x++
// 		fmt.Println(x)
// 	}
// 	fmt.Println(x)
// 	if i := 10; i <= 10 {
// 		fmt.Println(i)
// 	}
// }

func extensionsIf() {
	if x := 5; x > 10 {
		fmt.Println("Is maior que 10")
	} else if x < 1 {
		fmt.Println("X < 10")
	} else {
		fmt.Println("500")

	}
}

func main() {
	extensionsIf()
}
