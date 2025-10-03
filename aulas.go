/*arquivo destino para acompanhar o desenvolvimento das aulas, para acessar o conteúdo individual das aulas visualize o historico de modificacoes realizadas
 */

/*
 Declaracao switch é como se fosse um if com varios casos diferentes.
 ele avalia questoes e executa de cima pra baixo e todas que estiverem corretas tbm.
*/

// agrupamento de dados
// array, slices

// slices é criado com um número fixo de elementos, so posso adicionar um novo elemento com um append ou len(slice) + 1

// como fatiar um slice

package main

import (
	"fmt"
)

func main() {
	sabores := []string{"calabresa", "margherita", "abacaxi", "mozzarela"}
	fatia := sabores[2 : len(sabores)-1]

	fmt.Println(fatia)
}
