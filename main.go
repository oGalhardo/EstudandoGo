package main

import "fmt"

// func main() {
// 	idade := map[string]int{}
// 	idade["Junior"] = 29
// 	idade["Iara"] = 212
// 	fmt.Println(idade["Junior"])
// 	createMap()
// }

// func createMap() {
// 	anoNasc := map[string]int{
// 		"Junio": 1290,
// 		"SETe":  12,
// 	}
// 	fmt.Println(anoNasc)
// 	anoNasc["SCE"] = 212
// 	fmt.Println(anoNasc)
// }

// STRUCKS
type Carro struct {
	Modelo  string
	Tamanho int
}

type Transformer struct {
	Carro
	Nome string
}

func main() {
	C := Carro{}
	C.Modelo = "4x5"
	cars := []Carro{}
	cars = append(cars, C)
	// fmt.Println(cars)
	concessionarai := map[string][]Carro{}
	concessionarai["Modelo"] = cars
	// fmt.Println(concessionarai)
	robo := Transformer{
		Carro: Carro{Modelo: "4x5", Tamanho: 12},
		Nome:  "Optimus Prime",
	}
	fmt.Println(robo)
}
