# Trapped Mouse

Implementação em Go do desafio de backtracking com pilha. A matriz é impressa no estado inicial e após cada movimento do rato. O programa tenta os vizinhos nesta ordem: direita, esquerda, baixo e cima.

## Executar

```bash
cd "Trapped Mouse"
go run .
```

No menu, escolha:

1. **Arquivo**: pressione Enter para ler o exemplo em `labirintos.txt` ou informe outro caminho. O arquivo pode conter vários labirintos; separe cada um por uma linha vazia.
2. **Terminal**: digite uma linha por vez e finalize com uma linha vazia. O programa deduz a quantidade de linhas e colunas.

Cada labirinto deve ser retangular, ter exatamente um `m` (rato) e um `e` (saída) e usar apenas `1`, `0`, `m` e `e`. As paredes externas são acrescentadas automaticamente.

Além do exemplo pequeno (`labirintos.txt`), há outros exemplos prontos para usar:

- Médio: `labirinto_medio_1.txt` e `labirinto_medio_2.txt`.
- Grande: `labirinto_grande_1.txt` e `labirinto_grande_2.txt`.

## Testes

```bash
go test .
```
