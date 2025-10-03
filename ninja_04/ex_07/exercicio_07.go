/*Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.*/

package main

import (
	"fmt"
)

func main() {
	prontuario := [][]string{
		[]string{
			"Sovinha",
			"Stein",
			"Jogar valorant",
		},
		[]string{
			"Dakkels",
			"Kulique",
			"Jogar Lol",
		},
		[]string{
			"Pokas",
			"Possamai",
			"Pegar mulher casada",
		},
	}

	for valor := range prontuario {
		fmt.Println(valor)
	}
}
