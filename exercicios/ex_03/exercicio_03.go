/* Utilizando a solução do exercício anterior(exercicio_02):
   1. Em package-level scope, atribua os seguintes valores às variáveis:
       1. para "x" atribua 42
       2. para "y" atribua "James Bond"
       3. para "z" atribua true
   2. Na função main:
       1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
       2. Demonstre a variável "s".
*/

package exercicio_03

import (
	"fmt"
)

// declarando variaveis
var x int = 41
var y string = "James Bond"
var z bool = true

func exercicio_03() {
	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Println(s)
}
