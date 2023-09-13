package main

import "fmt"

func main() {
	a := (1 == 1)
	b := (2 != 1)
	c := (3 <= 4)
	d := (5 < 4)
	e := (6 >= 6)
	f := (7 > 9)

	fmt.Println(a, b, c, d, e, f)
}
