// - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//- Utilizando slicing e append, crie uma slice y que contenha os valores:
//- [42, 43, 44, 48, 49, 50, 51]

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := x[:3]
	fmt.Println(y)

	y = append(y, x[6:]...)
	fmt.Println(y)

}