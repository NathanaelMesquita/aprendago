// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//- Nome
//- Sobrenome
//- Sabores favoritos de sorvete
//- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
//- Solução: https://play.golang.org/p/Pyp7vmTJfY

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{"Nathanael", "Mesquita", []string{"chocolate", "morango", "maracujá"}}
	fmt.Println(pessoa1)

	pessoa2 := pessoa{"Mariana", "Mesquita", []string{"nutella", "ninho", "cookies"}}
	fmt.Println(pessoa2)

	fmt.Println("")

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}

}
