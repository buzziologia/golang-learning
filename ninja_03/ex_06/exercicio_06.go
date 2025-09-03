/*Crie um programa que demonstre o funcionamento da declaração if.
 */

package main

import (
	"fmt"
)

var fav_color string = "azul"
var color string

func main() {
	color = "roxo"
	if fav_color == color {
		fmt.Println("minha cor favorita é azul")
	} else {
		fmt.Printf("minha cor favorita não é %s", color)
	}
}
