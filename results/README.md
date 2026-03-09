# Benchmarks de Algoritmos de Ordenação

O conteúdo deste diretório possui o registro das métricas de execução dos algoritmos de ordenação:

- Bubble Sort
- Selection Sort
- Insertion Sort
- Merge Sort
- Quick Sort

Os resultados desta execução podem ser encontrados no arquivo [benchmark.csv](./benchmark.csv).

## Especificações de Hardware

- CPU: `AMD Ryzen 5 3600X 6-Core Processor`
- Cores: `12`
- RAM: `7.7Gi`
- OS: `Linux DESKTOP-O2QNOLB 6.6.87.2-microsoft-standard-WSL2 #1 SMP PREEMPT_DYNAMIC Thu Jun  5 18:30:46 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux`

## Especificações de Software

- Versão Go: `go1.26.1`
- Implementações de Algoritmos de Ordenação:
  - Bubble Sort: [bubble.go](../app/sort/bubble.go)
  - Selection Sort: [selection.go](../app/sort/selection.go)
  - Insertion Sort: [insertion.go](../app/sort/insertion.go)
  - Merge Sort: [merge.go](../app/sort/merge.go)
  - Quick Sort: [quick.go](../app/sort/quick.go)
- Implementação do Benchmark: [benchmark.go](../app/benchmark.go)
