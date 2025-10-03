/*Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
Do quinto ao último item do slice (incluindo o último item!)
Do segundo ao sétimo item do slice (incluindo o sétimo item!)
Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item*/

package main

import (
	"fmt"
)

func main() {
	my_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	new_slice := my_slice[:3]
	fmt.Println(new_slice)

	new_slice = my_slice[4:]
	fmt.Println(new_slice)

	new_slice = my_slice[1:8]
	fmt.Println(new_slice)

	new_slice = my_slice[3 : len(my_slice)-1]
	fmt.Println(new_slice)

}
