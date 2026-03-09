package sort

import (
	"cmp"
)

// QuickSort implementa o algoritmo de ordenação Quick Sort.
//
// O Quick Sort é um algoritmo de divisão e conquista que seleciona um elemento
// como pivô e particiona o array em torno dele, colocando elementos menores à
// esquerda e maiores à direita. O processo é repetido recursivamente em cada
// partição. É considerado um dos algoritmos mais eficientes na prática.
//
// Complexidade de Tempo:
//   - Melhor caso: O(n log n) - quando o pivô divide o array equilibradamente
//   - Caso médio: O(n log n)
//   - Pior caso: O(n²) - quando o pivô é sempre o menor ou maior elemento
//
// Parâmetros:
//   - array: slice de elementos do tipo T (genérico, deve ser ordenável)
//   - size: tamanho do array a ser ordenado
//
// Retorna: o slice ordenado em ordem crescente.
func QuickSort[T cmp.Ordered](array []T, size int) []T {
	if size <= 1 {
		return array
	}
	quickSortHelper(&array, 0, size-1)
	return array
}

func quickSortHelper[T cmp.Ordered](array *[]T, low, high int) {
	if low < high {
		pi := partition(array, low, high)
		quickSortHelper(array, low, pi-1)
		quickSortHelper(array, pi+1, high)
	}
}

func partition[T cmp.Ordered](array *[]T, low, high int) int {
	pivot := (*array)[high]
	i := low - 1

	for j := low; j < high; j++ {
		if (*array)[j] < pivot {
			i++
			(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
		}
	}
	(*array)[i+1], (*array)[high] = (*array)[high], (*array)[i+1]
	return i + 1
}
