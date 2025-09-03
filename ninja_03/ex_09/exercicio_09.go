/*Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".*/

package main

import (
	"fmt"
)

func main() {
	esporte_favorito := "volei"

	switch {
	case esporte_favorito == "volei":
		fmt.Println("esporte favorito é volei")
	case esporte_favorito == "basquete":
		fmt.Println("esporte favorito é basquete")
	case esporte_favorito == "valorant":
		fmt.Println("espero que esteja tudo bem contigo :), esporte favorito é valorant")
	}

}
