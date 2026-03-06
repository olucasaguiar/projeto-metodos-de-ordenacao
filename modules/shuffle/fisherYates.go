package shuffle

import "math/rand"

// FisherYates implementa o algoritmo de embaralhamento Fisher-Yates (também conhecido como Knuth Shuffle).
// Este algoritmo embaralha aleatoriamente os elementos de um slice de inteiros,
// garantindo que cada permutação possível tenha probabilidade igual de ocorrer.
//
// O algoritmo funciona iterando do último elemento até o segundo, selecionando aleatoriamente
// um índice anterior (incluindo o índice atual) e trocando os elementos nessas posições.
// Esta abordagem garante uma distribuição uniforme de todas as permutações possíveis.
//
// Parâmetros:
//   - array: um ponteiro para um slice de inteiros que será embaralhado in-place.
//
// Retorno:
//   - bool: retorna true se o embaralhamento foi bem-sucedido, ou se o array possui menos de 2 elementos.
//
// Complexidade:
//   - Tempo: O(n), onde n é o comprimento do array.
//
// Exemplo:
//
//	array := []int{1, 2, 3, 4}
//	FisherYates(&array) // embaralha o array aleatoriamente
//	// Exemplo de saída possível: [3, 1, 4, 2]
func FisherYatesStrategy(array *[]int) bool {
	if len(*array) < 2 {
		return true
	}

	for i := len(*array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}

	return true
}
