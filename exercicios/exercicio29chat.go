package main

import "fmt"

func main() {
	x := [][]string{
		[]string{"Nathanael", "Mesquita", "Jogar"},
		[]string{"Mariana", "Costa", "Viajar"},
		[]string{"Alaska", "Dramática", "Fazer drama"},
	}

	for _, pessoa := range x {
		fmt.Println("Nome:", pessoa[0])
		fmt.Println("Sobrenome:", pessoa[1])
		fmt.Println("Hobby favorito:", pessoa[2])
		fmt.Println()
	}
}
