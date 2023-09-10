// - // - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

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

	for _, p := range x {
		fmt.Println(p.nome)
		fmt.Println(p.sobrenome)
		fmt.Println("Sabores:")
		for _, sabor := range p.sabores {
			fmt.Println("\t", sabor)

		}
	}

}
