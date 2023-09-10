// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
//- Solução: https://play.golang.org/p/3fcvHlt8Lm

package main

import "fmt"

func main() {

	pessoa := map[string][]string{

		"nathanael_mesquita": []string{"Jogar", "Transar", "Dormir"},
		"mariana_mesquita":   []string{"Comer", "Transar", "Dormir"},
		"alaska_mesquita":    []string{"Comer", "encher o saco", "Dormir"},
	}

	pessoa["adicionando_teste"] = []string{"correr", "pular ", "nadar"}

	for key, hobbies := range pessoa {
		fmt.Println("Pessoa: ", key)
		for _, hobby := range hobbies {
			fmt.Printf("%v\n", hobby)
		}
		fmt.Println("")
	}

}
