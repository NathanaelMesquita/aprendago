package main

import "fmt"

func main() {

	Mep := map[string][]string{
		"Nathanael": []string{
			"comer",
			"dormir",
			"jogar",
		},
		"Mariana": []string{
			"comer",
			"dormir",
			"malhar",
		},
	}

	Mep["Alaska"] = []string{
		"Comer",
		"Dormir",
		"Encher o saco",
	}
	for i, x := range Mep {
		fmt.Println(i)
		for j, y := range x {
			fmt.Println(j, "-", y)
		}
	}

	delete(Mep, "Alaska")

	for i, x := range Mep {
		fmt.Println(i)
		for j, y := range x {
			fmt.Println(j, "-", y)
		}
	}
}
