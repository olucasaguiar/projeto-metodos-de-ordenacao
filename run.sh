#!/bin/bash

set -e

rm -rf "results" && mkdir -p "results"
cat > "results/README.md" <<EOF
# Benchmarks de Algoritmos de Ordenação

O conteúdo deste diretório possui o registro das métricas de execução dos algoritmos de ordenação:

- Bubble Sort
- Selection Sort
- Insertion Sort
- Merge Sort
- Quick Sort

Os resultados desta execução podem ser encontrados no arquivo [benchmark.csv](./benchmark.csv).

## Especificações de Hardware

- CPU: \`$(lscpu | grep "Model name" | awk -F: '{print $2}' | xargs)\`
- Cores: \`$(nproc)\`
- RAM: \`$(free -h | grep "Mem" | awk '{print $2}')\`
- OS: \`$(uname -a)\`

## Especificações de Software

- Versão Go: \`$(go version | awk '{print $3}')\`
- Implementações de Algoritmos de Ordenação:
  - Bubble Sort: [bubble.go](../app/sort/bubble.go)
  - Selection Sort: [selection.go](../app/sort/selection.go)
  - Insertion Sort: [insertion.go](../app/sort/insertion.go)
  - Merge Sort: [merge.go](../app/sort/merge.go)
  - Quick Sort: [quick.go](../app/sort/quick.go)
- Implementação do Benchmark: [benchmark.go](../app/benchmark.go)
EOF

echo "Running benchmarks..."
BENCHMARK_SCRIPT="app/benchmark.go"
BENCHMARKS=("bubble" "selection" "insertion" "merge" "quick")

for algorithm in "${BENCHMARKS[@]}"; do
    go run "$BENCHMARK_SCRIPT" "$algorithm"
done

echo "Benchmarks completed. Results are stored in the 'results' directory."
