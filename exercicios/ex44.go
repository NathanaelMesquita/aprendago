package main

import "fmt"

type Veiculo struct {
	Portas int
	Cor    string
}
type Caminhonete struct {
	Veiculo
	traçaoNasQuatro bool
}
type Sedan struct {
	Veiculo
	Luxo bool
}

func main() {

	Hillux := Caminhonete{
		Veiculo: Veiculo{
			Portas: 4,
			Cor:    "Preta",
		},
		traçaoNasQuatro: true,
	}
	Corolla := Sedan{
		Veiculo: Veiculo{
			Portas: 4,
			Cor:    "Prata",
		},
		Luxo: true,
	}

	fmt.Println(Corolla.Luxo)
	fmt.Println(Hillux.traçaoNasQuatro)
	fmt.Println(Corolla)
	fmt.Println(Hillux)

}
