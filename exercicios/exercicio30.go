// - Crie um map com key tipo string e value tipo []string.
//- Key deve conter nomes no formato sobrenome_nome
//- Value deve conter os hobbies favoritos da pessoa
//- Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {

	pessoa := map[string][]string{

		"nathanael_mesquita": []string{"Jogar", "Transar", "Dormir"},
		"mariana_mesquita":   []string{"Comer", "Transar", "Dormir"},
		"alaska_mesquita":    []string{"Comer", "encher o saco", "Dormir"},
	}

	for key, hobbies := range pessoa {
		fmt.Println("Pessoa: ", key)
		for _, hobby := range hobbies {
			fmt.Printf("%v\n", hobby)
		}
		fmt.Println("")
	}

}
