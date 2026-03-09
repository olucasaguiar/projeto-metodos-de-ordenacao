# Benchmarks de Algoritmos de Ordenação

O conteúdo deste diretório possui o registro das métricas de execução dos algoritmos de ordenação:

- Bubble Sort
- Selection Sort
- Insertion Sort
- Merge Sort
- Quick Sort

Os resultados desta execução podem ser encontrados no arquivo [benchmark.csv](./benchmark.csv).

## Especificações de Hardware

- CPU: `AMD EPYC 7763 64-Core Processor`
- Cores: `2`
- RAM: `7.8Gi`
- OS: `Linux codespaces-363dba 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 GNU/Linux`

## Especificações de Software

- Versão Go: `go1.25.6`
- Implementações de Algoritmos de Ordenação:
  - Bubble Sort: [bubble.go](../app/sort/bubble.go)
  - Selection Sort: [selection.go](../app/sort/selection.go)
  - Insertion Sort: [insertion.go](../app/sort/insertion.go)
  - Merge Sort: [merge.go](../app/sort/merge.go)
  - Quick Sort: [quick.go](../app/sort/quick.go)
- Implementação do Benchmark: [benchmark.go](../app/benchmark.go)
