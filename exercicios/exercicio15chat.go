package main

import (
	"fmt"
)

func main() {
	var y int // Declare a variable to store the user input

	fmt.Print("Quer contar até que ano? ")
	_, err := fmt.Scanln(&y) // Read user input and store it in 'y'
	if err != nil {
		fmt.Println("Erro ao ler o ano:", err)
		return
	}

	for x := 1992; x <= y; x++ {
		fmt.Println(x)
	}
}
