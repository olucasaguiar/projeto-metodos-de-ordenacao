package shuffle

import (
	"cmp"
	"math/rand"
)

// FisherYates implementa o algoritmo de embaralhamento Fisher-Yates (também conhecido como Knuth Shuffle).
//
// O algoritmo funciona iterando do último elemento até o segundo, selecionando aleatoriamente
// um índice anterior (incluindo o índice atual) e trocando os elementos nessas posições.
// Esta abordagem garante uma distribuição uniforme de todas as permutações possíveis.
//
// Complexidade de Tempo:
//   - Qualquer caso: O(n)
//
// Parâmetros:
//   - array: slice de elementos do tipo T (genérico, deve ser ordenável)
//
// Retorno: retorna true se o embaralhamento foi bem-sucedido, ou se o array possui menos de 2 elementos.
func FisherYatesStrategy[T cmp.Ordered](array *[]T) bool {
	if len(*array) < 2 {
		return true
	}

	for i := len(*array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}

	return true
}
