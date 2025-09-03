/*Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

func main() {
	for my_birth := 2000; my_birth <= 2025; my_birth++ {
		fmt.Println(my_birth)
	}
}
