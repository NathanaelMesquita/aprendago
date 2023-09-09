// - Crie um programa que demonstre o funcionamento da declaração if.
// - Solução: https://play.golang.org/p/OGPgTJZbiy
package main

import "fmt"

func main() {

	var x int

	fmt.Println("Digite um número")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Erro ao ler o número", err)
		return
	}
	if x%2 == 0 {
		fmt.Println("Número é par")
	} else if x%2 != 0 {
		fmt.Println("Número é ímpar")
	} else {
		fmt.Println("O número é zero")
	}
}
