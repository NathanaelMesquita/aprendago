//
//- Crie um programa que:
// Atribua um valor int a uma variável
//- Demonstre este valor em decimal, binário e hexadecimal
//- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
//- Demonstre esta outra variável em decimal, binário e hexadecimal
package main

import "fmt"

func main() {
	x:=10
	fmt.Printf("%b\n,%d\n,#%x\n", x, x ,x)
	fmt.Println(x)
	
	y:=x <<1
	fmt.Printf("%b\n,%d\n,#%x\n", y, y , y)
	fmt.Println(y)

	
}
