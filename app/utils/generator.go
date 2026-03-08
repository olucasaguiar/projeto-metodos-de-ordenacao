package utils

import "github.com/olucasaguiar/projeto-metodos-de-ordenacao/app/shuffle"

func BuildAscendingSortedArray(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range tamanho {
		vetor[i] = i + 1
	}
	return vetor
}

func BuildDescendingSortedArray(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range tamanho {
		vetor[i] = tamanho - i
	}
	return vetor
}

func BuildRandomArray(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range tamanho {
		vetor[i] = i + 1
	}
	success := shuffle.FisherYatesStrategy(&vetor)
	if !success {
		return nil
	}
	return vetor
}
