package sort

import (
	"cmp"
)

// BubbleSort implementa o algoritmo de ordenação Bubble Sort.
//
// O Bubble Sort é um algoritmo simples que realiza ordenação comparando pares
// adjacentes de elementos e trocando-os se estiverem na ordem incorreta.
// O processo é repetido várias vezes até que a lista esteja ordenada.
//
// Complexidade de Tempo:
//   - Melhor caso: O(n) - quando o array já está ordenado
//   - Caso médio: O(n²)
//   - Pior caso: O(n²) - quando o array está em ordem inversa
//
// Parâmetros:
//   - array: slice de elementos do tipo T (genérico, deve ser ordenável)
//   - size: tamanho do array a ser ordenado
//
// Retorna: o slice ordenado em ordem crescente.
func BubbleSort[T cmp.Ordered](array []T, size int) []T {
	for i := 0; i < size-1; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return array
}
