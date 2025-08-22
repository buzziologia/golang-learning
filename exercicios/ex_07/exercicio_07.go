/*Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.*/

package main

import (
	"fmt"
)

var a int = 10
var b string = "10"
var c int = 100

func main() {

	// declaracoes falsas
	fmt.Println("False declarations")
	fmt.Println(a > c)
	fmt.Println(a != a)
	fmt.Println(a >= 100)

	fmt.Println("")

	// declaracoes verdadeiras
	fmt.Println("True declarations")
	fmt.Println(a <= a)
	fmt.Println(b == b)
	fmt.Println(c > a)
}
