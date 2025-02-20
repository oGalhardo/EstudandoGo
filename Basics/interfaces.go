package main

// import "fmt"

// type pessoa struct {
// 	nome  string
// 	idade float64
// }

// type dentista struct {
// 	pessoa
// 	dentesExtraidos string
// 	salario         float64
// }

// type arquiteto struct {
// 	pessoa
// 	tipoDeConstrucao string
// }

// func (x dentista) manha() {
// 	fmt.Printf("MANHA ESTA CLARA, APROVEITE %v\n", x.nome)
// }

// func (x arquiteto) manha() {
// 	fmt.Printf("MANHA ESTA CLARA, APROVEITE a construcao %v\n", x.tipoDeConstrucao)
// }

// type people interface {
// 	manha()
// }

// func serHumano(g people) {
// 	g.manha()
// }

// func main() {
// 	srPredio := arquiteto{
// 		pessoa: pessoa{
// 			nome:  "JUNIOR",
// 			idade: 12,
// 		},
// 		tipoDeConstrucao: "BOLEANA",
// 	}
// 	srDente := dentista{
// 		pessoa: pessoa{
// 			nome:  "JUNIOR",
// 			idade: 12,
// 		},
// 		dentesExtraidos: "NENHUm",
// 		salario:         90000,
// 	}
// 	srDente.manha()
// 	srPredio.manha()
// }
