//-- Usando uma literal composta:
//- Crie um array que suporte 5 valores to tipo int
//- Atribua valores aos seus índices
//- Utilize range e demonstre os valores do array.
//- Utilizando format printing, demonstre o tipo do array.
//- Solução: https://play.golang.org/p/tpWDzzlO2l

package main

import "fmt"

func main() {
	array := [5]int{12, 23, 34, 45, 56}

	for i, value := range array {
		fmt.Println(i, value)
	}
	fmt.Printf("%T\n", array)

}
