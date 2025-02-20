package main

// import "fmt"

// // var x [5]int

// // func ex1() {
// // 	x[0] = 1
// // 	x[2] = 2
// // 	fmt.Println(len(x))
// // }

// // func slices() {

// // 	slices := []int{1, 2, 3, 4, 5}
// // 	fmt.Println(slices)
// // 	slices = append(slices, 12)
// // 	fmt.Println(slices)

// // 	slice := []string{"ILove", "Topster"}
// // 	for _, valor := range slice {
// // 		if valor == "ILove" {
// // 			fmt.Println("CONGRATULATIONS")
// // 		} else {
// // 			slice = append(slice, "21")
// // 		}
// // 	}
// // 	fmt.Println(slice)

// // }

// // func slide2() {
// // 	sabores := []string{"mussarela", "strognoff", "calabresa"}
// // 	// fatia := sabores[0:2]
// // 	// fmt.Println(fatia[0:len(sabores)])
// // 	for i := 0; i < len(sabores); i++ {
// // 		fmt.Println(sabores[i])
// // 	}
// // 	//Deletar uma fatia
// // 	sabores = append(sabores[:0], sabores[1:]...)
// // 	fmt.Println(sabores)
// // }

// func appen2d() {
// 	slice := []int{12, 33, 44}
// 	slice2 := []int{322, 22}
// 	outraSlice := append(slice, slice2...) //Pegando contéudo de slice 2
// 	fmt.Println(outraSlice)

// 	outro := make([]int, 5, 6)
// 	outro[0], outro[1], outro[2] = 3, 4, 5
// 	outro = append(outro, 3)
// 	fmt.Println(cap(outro))

// 	outro = append(outro, 4)

// 	fmt.Println(cap(outro))
// 	outro = append(outro, 22)

// 	fmt.Println(outro)
// }

// func main() {
// 	appen2d()
// }
