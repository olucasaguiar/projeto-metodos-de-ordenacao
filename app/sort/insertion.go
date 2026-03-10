package sort

import (
	"cmp"
)

// InsertionSort implementa o algoritmo de ordenação Insertion Sort.
//
// O Insertion Sort é um algoritmo simples que constrói a sequência ordenada
// um item de cada vez. Ele itera sobre um array, e para cada elemento, encontra
// sua posição correta na sequência já ordenada e o insere lá.
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
func InsertionSort[T cmp.Ordered](array []T, size int) []T {
	for i := 1; i < size; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}
