package main

import (
	"fmt"
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
		fmt.Printf("Valores do nó: coeficiente=%d, expoente=%d\n", coeficiente, expoente)
	}

	if proximo := polinomio.ObterProximo(noUm); proximo != nil {
		coeficiente, expoente, _ := polinomio.ObterValor(proximo)
		fmt.Printf("Próximo de x¹: coeficiente=%d, expoente=%d\n", coeficiente, expoente)
	}

	noTres := polinomio.Buscar(3)
	fmt.Println("Busca pelo expoente 3 encontrou um nó?", noTres != nil)

	polinomio.AlterarNo(noUm, 4, -2)
	fmt.Println("Após alterar -2x para -2x⁴:", polinomio.MostrarALL())

	polinomio.Excluir(noTres)
	fmt.Println("Após excluir o termo de expoente 3:", polinomio.MostrarALL())

	polinomio.Destrutor()
	fmt.Println("Após o destrutor, tamanho:", polinomio.Tamanho())
}

func questao2() {
}