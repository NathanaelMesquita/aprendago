package main

import "fmt"

func main() {
	x := 1992
	y := 2023

	for {
		if x > y {
			break
		}
		fmt.Println(x)
		x++
	}
}
