/*Escreva expressões utilizando os operadores*/

package main

import (
	"fmt"
)

var a int = 10
var b string = "10"
var c int = 100
var d string = "10"
var e int = 10

func main() {

	// declaracoes falsas
	fmt.Println("False declarations")
	fmt.Println(a > c)
	fmt.Println(a != e)
	fmt.Println(a >= 100)

	fmt.Println("")

	// declaracoes verdadeiras
	fmt.Println("True declarations")
	fmt.Println(a <= e)
	fmt.Println(b == d)
	fmt.Println(c > a)
}
