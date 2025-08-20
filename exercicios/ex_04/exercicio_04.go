/*Crie um tipo. O tipo subjacente deve ser int.
Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x
*/

package exercicio_04

import (
	"fmt"
)

type my_var int

var x my_var

func exericio_04() {
	fmt.Println(x)
	x = 42
	fmt.Printf("%T\n", x)
	fmt.Printf("%v", x)
}

func main() {
	exercicio_01()
}
