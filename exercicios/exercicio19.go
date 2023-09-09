//- Crie um programa que utilize a declaração switch, sem switch statement especificado.
//- Solução: https://play.golang.org/p/TyIGp4Hi8B

package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite qual opção")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Não foi possível ler a opção")
		return
	}
	switch x {
	case 1:
		fmt.Println("Você escolheu a primeira opção")

	case 2:
		fmt.Println("Você escolheu a segunda opção")

	}
}
