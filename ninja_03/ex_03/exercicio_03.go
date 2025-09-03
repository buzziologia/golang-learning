/*Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

func main() {
	my_birth := 2000
	atual_year := 2025

	for my_birth <= atual_year {
		fmt.Println(my_birth)
		my_birth++
	}
}
