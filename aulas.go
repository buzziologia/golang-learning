/*arquivo destino para acompanhar o desenvolvimento das aulas, para acessar o conteúdo individual das aulas visualize o historico de modificacoes realizadas
 */

package main

import (
	"fmt"
)

func main() {
	y := 24
	x := y << 2
	z := y >> 2

	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
}
