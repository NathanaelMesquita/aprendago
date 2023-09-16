//- Crie e use um struct anônimo.
//- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

package main

import "fmt"

func main() {
	x := 1250
	fmt.Println(x)

	y := &x
	fmt.Println("Esse é o endereço de x", y)
	fmt.Println("Esse é o valor de x, usando ponteiro,", *y)
}
