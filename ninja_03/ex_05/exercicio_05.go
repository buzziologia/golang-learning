/*Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
 */

package main

import (
	"fmt"
)

var resto int

func main() {

	fmt.Println("Resto das divisão por 4 dos numeros de 10 a 100")
	for x := 10; x <= 100; x++ {
		resto = x % 4
		fmt.Printf("%d -> resto = %d\n", x, resto)
	}
}
