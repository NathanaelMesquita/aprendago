//- -tente acessar todos os itens de uma slice *sem* utilizar range

package main

import "fmt"

func main() {
	slice := []string{"1", "2", "3", "4", "5", "6"}

	fmt.Println(slice)

	fatia := (slice[:])

	fmt.Println(fatia)

}
