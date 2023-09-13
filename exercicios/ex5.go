//- Utilizando a solução do exercício anterior:

//	2. Agora faça tambem:
//1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", ////atribua o valor de "x" a "y"
//		2. Demonstre o valor de "y"
//		3. Demonstre o tipo de "y"
//- Solução: https://play.golang.org/p/uq81T_fasP

package main

import "fmt"

type meme int

var x meme
var y int

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Println(x)

	y := int(x)
	fmt.Println(y)

}
