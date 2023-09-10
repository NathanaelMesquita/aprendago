package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	x := make(map[string]pessoa)

	x["Mesquita"] = pessoa{"Nathanael", "Mesquita", []string{"chocolate", "morango", "maracujá"}}
	x["Costa"] = pessoa{"Mariana", "Costa", []string{"nutella", "ninho", "cookies"}}

	for sobrenome, p := range x {
		fmt.Println("Sobrenome:", sobrenome)
		fmt.Println("Nome:", p.nome)
		fmt.Println("Sabores:")
		for _, sabor := range p.sabores {
			fmt.Println("\t", sabor)
		}
		fmt.Println()
	}
}
