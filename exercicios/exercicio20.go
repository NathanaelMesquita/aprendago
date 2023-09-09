//- - Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {
	var x string
	fmt.Println("Digite qual esporte favoritor\n 1 - Futebol\n 2- Corrida\n 3- Natação")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Não foi possível ler a opção")
		return
	}
	switch x {
	case "1":
		fmt.Println("Seu esporte favorito é futebol!")

	case "2":
		fmt.Println("Seu esporte favorito é corrida!")

	case "3":
		fmt.Println("Seu esporte favorito é natação!")

	}
}
