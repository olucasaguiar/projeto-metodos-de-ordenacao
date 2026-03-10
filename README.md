# Análise crítica sobre o tempo de execução em algoritmos de ordenação

Este projeto foi desenvolvido dentro do programa de Pós-Graduação em Computação Aplicada (Mestrado Profissional) para a disciplina de Algoritmos e Programação no primeiro semestre de 2026.

No [README.md](./README.md) principal, será apresentado o enunciado do projeto disponibilizado na disciplina e um breve detalhamento dos arquivos existentes no projeto para auxiliar na compreensão e navegação dos arquivos.  

O presente projeto tem fins didáticos e compartilhamento de conhecimento e, portanto, não possui valor científico ou acadêmico. O projeto se trata de uma prática dentro da disciplina com foco em desenvolver no aluno aspectos da escrita acadêmica, conhecimento do rigor metódologico em produção científica e treinar a análise crítica de resultados experimentais.

## Introdução

O projeto propõe atingir ao objetivo proposto pelo seguinte enunciado:

> Faça um programa (em qualquer linguagem de programação) para:
>
> - Entradas de vetores de tamanhos 1.000, 10.000 e 100.000
>   - aleatórias
>   - ordenadas em forma crescente
>   - ordenadas em forma decrescente
> - Utilize os métodos: bolha (melhorado ou original), seleção, inserção, mergesort e quicksort para ordená-los.
> - Marque os tempos comparativamente entre todos os métodos em todos os vetores
> - Gere um relatório com:
>   - método: esclareça, por exemplo, tipo de equipamento utilizado, massa de dados, algoritmos utilizados, linguagem de programação, etc.
>   - gráficos com as comparações de tempos medidos entre todos os métodos.
>   - análise crítica sobre a eficiência dos algoritmos.
>   - análise crítica sobre a análise assintótica X os tempos obtidos.

## Estrutura do Projeto

O projeto foi escrito utilizando linguagem Go e organizado em um módulo principal. A organização dos pacotes não segue a estrutura padrão `src` para deixar o projeto com aspecto simples. Assim, os pacotes da aplicação foram distribuídos no diretório `app`, os testes de unidade em `tests` e os resultados da execução em `results`.

O ponto de entrada da aplicação é o arquivo [benchmark.go](./app/benchmark.go), responsável por tratar os argumentos passados para a aplicação e executar o fluxo de testes. Para saber mais sobre os parâmetros da aplicação, rode pacote principal com o `go` e observe a saída no terminal.

```bash
go run "app/benchmark.go"
```

Os algoritmos de ordenação se encontram no diretório `./app/sort/`, e cada algoritmo tem seu próprio arquivo com sua implementação documentada nos padrões de documentação da linguagem.

```text
/app
|---| sort
    |---| bubble.go     # Implementação do algoritmo Bubble Sort
        | insertion.go  # Implementação do algoritmo Insertion Sort
        | selection.go  # Implementação do algoritmo Selection Sort
        | merge.go      # Implementação do algoritmo Merge Sort
        | quick.go      # Implementação do algoritmo Quick Sort
```

No projeto, foram utilizados algoritmos de apoio para as necessidades: geração de vetores em tempo de execução ([generator.go](./app/utils/generator.go)), embaralhamento de valores em vetor ([fisherYates.go](./app/shuffle/fisherYates.go)), manipulação de arquivos ([file.go](./app/utils/file.go)) e relatório ([report.go](./app/utils/report.go)).

## Análise e Resultados

O artigo está separado nos seguintes documentos:

1. [Resumo](./docs/SUMMARY.md)
2. [Metodologia](./docs/METODOLOGY.md)
3. [Resultados](./docs/RESULTS.md)
4. [Conclusão](./docs/CONCLUSION.md)

## Outros documentos

1. [Relatório compilado](./docs/_relatorio.pdf)
2. [Apresentação em Slides](./docs/_apresentacao.pdf)
3. [Python Notebook utilizado nas análises](./docs/_analises.ipynb)
