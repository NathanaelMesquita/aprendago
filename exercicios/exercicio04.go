package main

import "fmt"

type treinogo int

var x treinogo

func main() {
    fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)	
}
