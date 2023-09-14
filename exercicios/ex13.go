package main

import "fmt"

func main() {

	for x := 65; x < 91; x++ {
		fmt.Println(x)
		for y := 1; y < 4; y++ {
			fmt.Printf("\t%#U\n", x)

		}
	}
}
