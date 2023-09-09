// -- Usando uma literal composta:
// - Crie um array que suporte 5 valores to tipo int
// - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.
// - - Usando uma literal composta:
// - Crie uma slice de tipo int
// - Atribua 10 valores a ela
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range slice {

		fmt.Println(value)
	}
}
