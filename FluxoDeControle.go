package main

import "fmt"

// func aula1() {
// 	for x := 0; x < 10; x++ {
// 		fmt.Println(x)
// 	}
// }

func ex1() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 6; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break

	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func main() {
	ex1()
}
