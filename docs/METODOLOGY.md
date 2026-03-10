# Metodologia

A metodologia adotada neste estudo foi estruturada em seis seções principais: ambiente de simulação, linguagem utilizada, implementação dos algoritmos, preparação dos dados, cenários de teste e modo de execução. Essa organização buscou garantir rigor experimental, reprodutibilidade e clareza na análise comparativa dos resultados.

## 1. Ambiente de Simulação  

Os experimentos foram realizados em ambiente controlado, utilizando um computador pessoal com CPU **AMD Ryzen 5 3600X (6-Core, 12 threads)**, memória RAM de **7,7 GiB** e sistema operacional **Linux (WSL2)**. Esse ambiente foi selecionado por se tratar de um equipamento pessoal comum e acessível para replicabilidade do experimento. Apesar de se tratar de um computador pessoal, a configuração ainda oferece estabilidade e capacidade de processamento adequadas para a execução dos algoritmos para os volumes de dados do experimento.

## 2. Linguagem Utilizada  

A linguagem escolhida foi **Go (versão 1.26.1)**, devido à sua eficiência na manipulação de memória, suporte nativo a concorrência e precisão nas medições de tempo. A clareza sintática e a facilidade de documentação também foram fatores determinantes para sua adoção neste projeto.

## 3. Implementação dos Algoritmos  

Foram implementados cinco algoritmos clássicos de ordenação: **Bubble Sort, Selection Sort, Insertion Sort, Merge Sort e Quick Sort**. Cada algoritmo foi desenvolvido em módulos independentes, seguindo boas práticas de programação e documentação da linguagem Go. Antes da execução dos testes, todas as implementações foram validadas quanto à correção funcional, assegurando que os resultados refletissem apenas diferenças de desempenho e não erros de implementação.

## 4. Preparação dos Dados  

Os vetores utilizados nos experimentos foram gerados por meio de laços sequenciais, garantindo elementos únicos e sem repetição. Para o cenário embaralhado, foi aplicado o algoritmo **Fisher-Yates**, que assegura distribuição uniforme das permutações, garantindo que todos os elementos tenham probabilidade equivalente de serem reposicionados. Essa preparação buscou simular condições iniciais distintas de ordenação, permitindo avaliar o impacto da estrutura dos dados sobre o desempenho dos algoritmos.

## 5. Cenários de Teste  

Foram definidos três cenários distintos de entrada, cada um com vetores de tamanhos **1.000, 10.000 e 100.000 elementos**:  

- Vetores ordenados em ordem crescente.  
- Vetores ordenados em ordem decrescente.  
- Vetores embaralhados aleatoriamente.  

Essa diversidade de cenários permitiu avaliar o comportamento dos algoritmos em condições favoráveis (dados já ordenados), desfavoráveis (dados em ordem inversa) e neutras (dados aleatórios).

## 6. Modo de Execução  

Cada combinação de algoritmo e cenário foi executada **cinco vezes consecutivas**, com o objetivo de reduzir interferências externas do ambiente de execução. Os tempos de execução foram medidos em nanossegundos, utilizando funções nativas da linguagem Go. Para análise, foram calculadas as métricas de **melhor caso, pior caso, média e mediana**. Os resultados foram organizados em tabelas e gráficos comparativos, servindo de base para a análise crítica da eficiência prática dos algoritmos frente às suas complexidades assintóticas teóricas.
