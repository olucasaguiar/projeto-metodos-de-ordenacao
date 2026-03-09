package sort

import (
	"cmp"
)

// SelectionSort implementa o algoritmo de ordenação Selection Sort.
//
// O Selection Sort é um algoritmo simples que funciona dividindo o array em
// duas partes: a parte ordenada (à esquerda) e a parte não ordenada (à direita).
// A cada iteração, encontra o menor elemento da parte não ordenada e o coloca
// no final da parte ordenada.
//
// Complexidade de Tempo:
//   - Melhor caso: O(n²)
//   - Caso médio: O(n²)
//   - Pior caso: O(n²)
//
// Parâmetros:
//   - array: slice de elementos do tipo T (genérico, deve ser ordenável)
//   - size: tamanho do array a ser ordenado
//
// Retorna: o slice ordenado em ordem crescente.
func SelectionSort[T cmp.Ordered](array []T, size int) []T {
	for i := 0; i < size-1; i++ {
		minIdx := i
		for j := i + 1; j < size; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			array[i], array[minIdx] = array[minIdx], array[i]
		}
	}
	return array
}
