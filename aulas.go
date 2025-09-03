/*arquivo destino para acompanhar o desenvolvimento das aulas, para acessar o conteúdo individual das aulas visualize o historico de modificacoes realizadas
 */

/*
 Declaracao switch é como se fosse um if com varios casos diferentes.
 ele avalia questoes e executa de cima pra baixo e todas que estiverem corretas tbm.
*/

package main

import (
	"fmt"
)

func main() {
	x := 10

	switch true {
	case x < 5:
		fmt.Println("x é menor que ")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")
	}

	if (x == 5) == true {
		fmt.Println("x é igual a 5")
	}

}
