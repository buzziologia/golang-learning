/* Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
*/

package exercicio_02

import (
	"fmt"
)

// declarando variaveis
var x int
var y string
var z bool

func exercicio_02() {
	fmt.Printf("%v\n, %v\n, %v\n", x, y, z)
	resposta := "Resposta Questao 2: os valores são chamados de zeros"
	fmt.Printf("\n%v", resposta)
}
