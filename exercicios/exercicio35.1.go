package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçaoNasQuatro bool
}

type sedan struct {
	veiculo
	modelodeLuxo bool
}

func main() {

	minhacaminhonete := caminhonete{
		veiculo{
			portas: 4,
			cor:    "preta",
		},
		traçãoNasQuatro: true,
	}

	meusedan := sedan{
		veiculo{
			portas: 4,
			cor:    "vermelho",
		}, modeloLuxo: false,
	}

	fmt.Println("Caminhonete:")
	fmt.Println("Cor:", minhacaminhonete.veiculo.cor)
	fmt.Println("Quantidade de portas:", minhacaminhonete.veiculo.portas)
	fmt.Println("Possui 4x4?", minhacaminhonete.traçaoNasQuatro)
	fmt.Println("")

	fmt.Println("Sedan:")
	fmt.Println("Cor:", meusedan.veiculo.cor)
	fmt.Println("Quantidade de portas:", meusedan.veiculo.portas)
	fmt.Println("Modelo de luxo?", meusedan.modelodeLuxo)
	fmt.Println("")

	fmt.Println("Quantidade de portas:", meusedan.veiculo.portas, "\n")

	fmt.Println("Cor:", minhacaminhonete.veiculo.cor)

}
