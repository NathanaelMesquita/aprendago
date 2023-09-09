//- Crie um loop utilizando a sintaxe: for condition {}
//- Utilize-o para demonstrar os anos desde que você nasceu.
//- Solução: https://play.golang.org/p/qnFjiDJzLor

package main

import "fmt"

var y int

func main() {

	fmt.Println("Quer contar até que ano? ")
	fmt.Scanln(y)

	for x := 1992; x <= y; x++ {
		fmt.Println(x)
	}

}
