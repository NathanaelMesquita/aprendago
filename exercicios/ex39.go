package main

import "fmt"

func main() {
x :=10

if x%2 == 0{
	fmt.Println("o número é par!")
}else if x%2 == 1{
	fmt.Println("o número é impar!")
}else{
	fmt.Println("o número é zero")
}
}
