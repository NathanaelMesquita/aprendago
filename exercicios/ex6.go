//- Utilizando a solução do exercício anterior:

//	2. Agora faça tambem:
//1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", ////atribua o valor de "x" a "y"
//		2. Demonstre o valor de "y"
//		3. Demonstre o tipo de "y"
//- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
//- Solução: https://play.golang.org/p/X7qm3aWSa6

package main

import "fmt"

func main() {
	x := 215
	fmt.Printf("%d, %b, %#x", x, x, x)
	fmt.Println("")
}
