package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *Pessoa) {
	fmt.Println("Atualizando dados...")
	(*p).idade = 30
}

func main() {
	pessoa1 := Pessoa{
		nome:      "Nathanael",
		sobrenome: "Mequita",
		idade:     29,
	}
	fmt.Println(pessoa1)

	mudeMe(&pessoa1)

	fmt.Println(pessoa1)

	fmt.Println("Novos dados!")

	mudeMe2(&pessoa1)
	fmt.Println(pessoa1)
}

func mudeMe2(p *Pessoa) {
	fmt.Println("Atualizando dados..")
	(*p).nome = "Mariana"
	(*p).sobrenome = "Costa"
	(*p).idade = 28
}
