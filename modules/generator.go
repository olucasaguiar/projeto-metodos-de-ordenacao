package modules

func VetorOrdenadoCrescente(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range tamanho {
		vetor[i] = i + 1
	}
	return vetor
}

func VetorOrdenadoDecrescente(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range tamanho {
		vetor[i] = tamanho - i
	}
	return vetor
}

func VetorAleatorio(tamanho int) []int {
	vetor := make([]int, tamanho)
	for i := range tamanho {
		vetor[i] = i + 1
	}
	success := FisherYates(&vetor)
	if !success {
		return nil
	}
	return vetor
}
