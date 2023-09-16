package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice completo: ", x)

	y := x[:3]
	fmt.Println("Primeiro ao terceiro: ", y)

	z := x[5:]
	fmt.Println("Quinto ao último: ", z)

	w := x[2:8]
	fmt.Println("Segundo ao sétimo: ", w)

	q := x[3 : len(x)-1]
	fmt.Println("Terceiro ao penúltimo com len: ", q)

}
