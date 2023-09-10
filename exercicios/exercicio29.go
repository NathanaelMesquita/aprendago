// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//- "Nome", "Sobrenome", "Hobby favorito"
//- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {

	x := [][]string{[]string{
		"Nathanael",
		"Mesquita",
		"Jogar",
	},
		[]string{
			"Mariana",
			"Costa",
			"Viajar",
		},
		[]string{
			"Alaska",
			"Dramática",
			"Fazer drama",
		},
	}

	for _, v := range x {
		fmt.Println(v[0])

		for _, j := range v {
			fmt.Println("\t", j)
		}
	}

}
