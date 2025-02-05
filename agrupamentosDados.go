package main

import "fmt"

// var x [5]int

// func ex1() {
// 	x[0] = 1
// 	x[2] = 2
// 	fmt.Println(len(x))
// }

func slices() {

	slices := []int{1, 2, 3, 4, 5}
	fmt.Println(slices)
	slices = append(slices, 12)
	fmt.Println(slices)

	slice := []string{"ILove", "Topster"}
	for _, valor := range slice {
		if valor == "ILove" {
			fmt.Println("CONGRATULATIONS")
		} else {
			slice = append(slice, "21")
		}
	}
	fmt.Println(slice)

}

func main() {
	slices()
}
