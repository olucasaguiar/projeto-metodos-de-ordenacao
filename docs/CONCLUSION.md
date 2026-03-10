# Conclusão

Os resultados obtidos confirmam que a análise assintótica é uma ferramenta fundamental para compreender a escalabilidade dos algoritmos, mas também revelam que fatores práticos — como a estrutura inicial dos dados e constantes ocultas — influenciam diretamente o desempenho real. Algoritmos quadráticos, como *Bubble Sort*, *Selection Sort* e *Insertion Sort*, mostraram-se inviáveis para grandes volumes de dados, embora o *Insertion Sort* tenha se destacado em cenários favoráveis, aproximando-se de comportamento linear.

Por outro lado, os algoritmos \(O(n log n)\), *Merge Sort* e *Quick Sort*, mantiveram tempos significativamente menores e confirmaram sua superioridade teórica. *Merge Sort* demonstrou robustez e estabilidade em todos os cenários, enquanto *Quick Sort* foi altamente eficiente em entradas aleatórias, mas vulnerável em casos extremos, evidenciando a importância da estratégia de escolha do pivô.

Em síntese, conclui-se que a teoria assintótica fornece uma base sólida para prever o comportamento dos algoritmos, mas a escolha prática deve considerar também o contexto de aplicação. Para grandes conjuntos de dados, *Merge Sort* e *Quick Sort* são os mais indicados, enquanto *Insertion Sort* pode ser útil em vetores pequenos ou quase ordenados. Já *Bubble Sort* e *Selection Sort* devem ser restritos a fins didáticos, dada sua baixa eficiência prática.
