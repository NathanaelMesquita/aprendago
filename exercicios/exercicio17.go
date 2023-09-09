//- Crie um programa que demonstre o funcionamento da declaração if.
//- Solução: https://play.golang.org/p/OGPgTJZbiy

package main

import "fmt"

func main() {

	x := 9

	if x%2 == 0 {
		fmt.Println("Número é par")
	} else {
		fmt.Println("Número é impar")
	}

}
