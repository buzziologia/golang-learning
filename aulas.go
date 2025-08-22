package main

import (
	"fmt"
)

func main() {
	a := "e"
	b := "é"
	c := "終"

	fmt.Printf("%v\tap%v\t%v", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("\n%v\t%v\t%v", d, e, f)

}
