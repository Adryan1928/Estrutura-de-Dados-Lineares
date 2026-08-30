package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	q := 1
	for q != 0 {
		fmt.Println("Escolhar a questão desejada: (1 o 2):")
		fmt.Scanln(&q)

		switch q {
		case 1:
			questao1()
		case 2:
			questao2()
		}
	}
}

func questao1() {
	var polinomio ListaOrdenada = NovaLista()

	polinomio.Inserir(4, 3)          // 4x³
	noUm := polinomio.Inserir(-2, 1) // -2x
	polinomio.Inserir(7, 5)          // 7x⁵
	polinomio.Inserir(9, 0)          // 9

	fmt.Println("Termos ordenados:", polinomio.MostrarALL())
	fmt.Println("Tamanho:", polinomio.Tamanho())
	fmt.Println("O nó de expoente 1 existe?", polinomio.Existe(noUm))

	if coeficiente, expoente, ok := polinomio.ObterValor(noUm); ok {
		fmt.Printf("Valores do nó: coeficiente=%f, expoente=%d\n", coeficiente, expoente)
	}

	if proximo := polinomio.ObterProximo(noUm); proximo != nil {
		coeficiente, expoente, _ := polinomio.ObterValor(proximo)
		fmt.Printf("Próximo de x¹: coeficiente=%f, expoente=%d\n", coeficiente, expoente)
	}

	noTres := polinomio.Buscar(3)
	fmt.Println("Busca pelo expoente 3 encontrou um nó?", noTres != nil)

	polinomio.AlterarNo(noUm, -2, 4)
	fmt.Println("Após alterar -2x para -2x⁴:", polinomio.MostrarALL())

	fmt.Println("Grau do polinômio:", polinomio.Grau())
	valor := 2.0
	resultado := polinomio.Avaliar(valor)
	fmt.Printf("Avaliação do polinômio em x=%.2f: %.2f\n", valor, resultado)

	fmt.Println("Polinômio formatado:", polinomio.Exibir())

	outra := NovaLista()
	outra.Inserir(3, 4) // 3x²
	outra.Inserir(5, 0) // 5x

	soma := polinomio.Somar(outra)
	fmt.Println("Soma dos polinômios:", soma.Exibir())

	subtracao := polinomio.Subtrair(outra)
	fmt.Println("Subtração dos polinômios:", subtracao.Exibir())

	multiplicacao := polinomio.Multiplicar(outra)
	fmt.Println("Multiplicação dos polinômios:", multiplicacao.Exibir())

	polinomio.Excluir(noTres)
	fmt.Println("Após excluir o termo de expoente 3:", polinomio.MostrarALL())

	polinomio.Destrutor()
	fmt.Println("Após o destrutor, tamanho:", polinomio.Tamanho())
}

func questao2() {
	caminho, erro := localizarArquivoDeOperacoes()
	if erro != nil {
		fmt.Println("Erro:", erro)
		return
	}

	resultados, erro := executarOperacoes(caminho)
	if erro != nil {
		fmt.Println("Erro ao processar o arquivo:", erro)
		return
	}

	fmt.Printf("Operações lidas de %s:\n", caminho)
	for _, resultado := range resultados {
		fmt.Println(resultado)
	}
}

func localizarArquivoDeOperacoes() (string, error) {
	for _, caminho := range []string{"Polinomio/polinomios.txt", "polinomios.txt"} {
		if _, erro := os.Stat(caminho); erro == nil {
			return caminho, nil
		}
	}
	return "", fmt.Errorf("arquivo polinomios.txt não encontrado")
}

func executarOperacoes(caminho string) ([]string, error) {
	arquivo, erro := os.Open(caminho)
	if erro != nil {
		return nil, erro
	}
	defer arquivo.Close()

	linhas := make([]string, 0)
	leitor := bufio.NewScanner(arquivo)
	for leitor.Scan() {
		linha := strings.TrimSpace(leitor.Text())
		if linha != "" && !strings.HasPrefix(linha, "#") {
			linhas = append(linhas, linha)
		}
	}
	if erro := leitor.Err(); erro != nil {
		return nil, erro
	}

	return processarLinhas(linhas)
}

func processarLinhas(linhas []string) ([]string, error) {
	resultados := make([]string, 0)
	indice := 0

	for indice < len(linhas) {
		operacao := strings.ToLower(linhas[indice])
		indice++

		switch operacao {
			case "+", "-", "*":
				primeiro, erro := proximoPolinomio(linhas, &indice)
				if erro != nil {
					return nil, erro
				}
				segundo, erro := proximoPolinomio(linhas, &indice)
				if erro != nil {
					return nil, erro
				}

				var resultado *Lista
				switch operacao {
					case "+":
						resultado = primeiro.Somar(segundo)
					case "-":
						resultado = primeiro.Subtrair(segundo)
					case "*":
						resultado = primeiro.Multiplicar(segundo)
				}
				resultados = append(resultados, fmt.Sprintf("(%s) %s (%s) = %s", primeiro.Exibir(), operacao, segundo.Exibir(), resultado.Exibir()))

			case "g":
				polinomio, erro := proximoPolinomio(linhas, &indice)
				if erro != nil {
					return nil, erro
				}
				resultados = append(resultados, fmt.Sprintf("Grau de %s: %d", polinomio.Exibir(), polinomio.Grau()))

			case "t":
				polinomio, erro := proximoPolinomio(linhas, &indice)
				if erro != nil {
					return nil, erro
				}
				resultados = append(resultados, fmt.Sprintf("Quantidade de termos de %s: %d", polinomio.Exibir(), polinomio.Tamanho()))

			case "p":
				polinomio, erro := proximoPolinomio(linhas, &indice)
				if erro != nil {
					return nil, erro
				}
				resultados = append(resultados, fmt.Sprintf("Polinômio: %s", polinomio.Exibir()))

			case "a":
				valorTexto, erro := proximaLinha(linhas, &indice, "o valor de avaliação")
				if erro != nil {
					return nil, erro
				}
				valor, erro := strconv.ParseFloat(valorTexto, 64)
				if erro != nil {
					return nil, fmt.Errorf("valor de avaliação inválido %q", valorTexto)
				}
				polinomio, erro := proximoPolinomio(linhas, &indice)
				if erro != nil {
					return nil, erro
				}
				resultados = append(resultados, fmt.Sprintf("p(x) = %s; p(%.2f) = %.2f", polinomio.Exibir(), valor, polinomio.Avaliar(valor)))

			default:
				return nil, fmt.Errorf("operação desconhecida %q", operacao)
			}
	}

	return resultados, nil
}

func proximoPolinomio(linhas []string, indice *int) (*Lista, error) {
	linha, erro := proximaLinha(linhas, indice, "um polinômio")
	if erro != nil {
		return nil, erro
	}
	polinomio := NovaLista()
	erro = polinomio.LerPolinomio(linha)
	if erro != nil {
		return nil, erro
	}
	return polinomio, nil
}

func proximaLinha(linhas []string, indice *int, descricao string) (string, error) {
	if *indice >= len(linhas) {
		return "", fmt.Errorf("faltou informar %s", descricao)
	}
	linha := linhas[*indice]
	*indice = *indice + 1
	return linha, nil
}
