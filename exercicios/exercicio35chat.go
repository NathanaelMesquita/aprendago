package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	minhaCaminhonete := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "preta",
		},
		traçãoNasQuatro: true,
	}

	meuSedan := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "vermelho",
		},
		modeloLuxo: false,
	}

	fmt.Println("Caminhonete:")
	fmt.Println("Cor:", minhaCaminhonete.veiculo.cor)
	fmt.Println("Quantidade de portas:", minhaCaminhonete.veiculo.portas)
	fmt.Println("Possui tração nas quatro rodas?", minhaCaminhonete.traçãoNasQuatro)
	fmt.Println()

	fmt.Println("Sedan:")
	fmt.Println("Cor:", meuSedan.veiculo.cor)
	fmt.Println("Quantidade de portas:", meuSedan.veiculo.portas)
	fmt.Println("Modelo de luxo?", meuSedan.modeloLuxo)
	fmt.Println()

	fmt.Println("Quantidade de portas do Sedan:", meuSedan.veiculo.portas)
	fmt.Println()

	fmt.Println("Cor da Caminhonete:", minhaCaminhonete.veiculo.cor)
}
