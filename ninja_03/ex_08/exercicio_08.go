/*Crie um programa que utilize a declaração switch, sem switch statement especificado.
 */

package main

import (
	"fmt"
)

func main() {

	nota_rec := 7.6

	switch {
	case (nota_rec >= 9.0):
		fmt.Println("Indice academico S")
	case (nota_rec >= 8.0) && (nota_rec < 9.0):
		fmt.Println("Indice academico A")
	case (nota_rec >= 7.0) && (nota_rec < 8.0):
		fmt.Println("Indice academico B")
	case (nota_rec >= 6.0) && (nota_rec < 7.0):
		fmt.Println("Indice academica C")
	case (nota_rec < 6.0):
		fmt.Println("Indice academica I")
	}
}
