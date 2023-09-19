package main

import "fmt"

func bomDia() string {
	return "Bom dia!"
}

func soma(a, b int) int {
	resultado := a + b
	return resultado
}

func main() {

	fmt.Println(bomDia())
	fmt.Println(soma(10, 12))

}
