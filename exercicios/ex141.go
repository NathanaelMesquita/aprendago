package main

import "fmt"

func main() {
	var x int

	fmt.Println("Digite um número:")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("Erro ao ler o número", err)
		return
	}

	if x%2 == 0 {
		fmt.Println("É par!")
	} else if x%2 == 1 {
		fmt.Println("É ímpar!")
	} else {
		fmt.Println("É zero!")
	}
}

