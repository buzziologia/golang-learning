/*Escreva um programa que mostre um número em decimal, binário, e hexadecimal.*/

package main

import (
	"fmt"
)

var x int

func main() {
	x = 18
	fmt.Printf("%d\t%v\n", x, "d -> decimal")
	fmt.Printf("%b\t%v\n", x, "b -> binario")
	fmt.Printf("%#x\t%v\n", x, "#x -> hexadecimal")
}
