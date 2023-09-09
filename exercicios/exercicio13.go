// - Põe na tela: todos os números de 1 a 10000.
// - Solução: https://play.golang.org/p/MkdZiDW8SQ
package main

import "fmt"

func main() {

	for x := 65; x <= 90; x++ {
		for y := 1; y <= 3; y++ {
			fmt.Println("\t%#U\n", x)
		}
	}

}
