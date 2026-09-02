package main

import (
	"fmt"
	"strings"
)

type Stack interface {
	Push(value rune)
	Pop() (rune, bool)
	Peek() (rune, bool)
	IsEmpty() bool
	Len() int
}

type Pilha struct {
	items []rune
}

func (p *Pilha) Push(value rune) {
	p.items = append(p.items, value)
}

func (p *Pilha) Pop() (rune, bool) {
	if p.IsEmpty() {
		return 0, false
	}

	lastIndex := len(p.items) - 1
	value := p.items[lastIndex]
	p.items = p.items[:lastIndex]
	return value, true
}

func (p *Pilha) Peek() (rune, bool) {
	if p.IsEmpty() {
		return 0, false
	}

	return p.items[len(p.items)-1], true
}

func (p *Pilha) IsEmpty() bool {
	return len(p.items) == 0
}

func (p *Pilha) Len() int {
	return len(p.items)
}

var pares = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func estaBalanceada(expressao string) (bool, map[rune]rune) {
	var pilha Stack = &Pilha{}

	
	items_faltando := map[rune]rune{
		'(': 0,
		')': 0,
		'[': 0,
		']': 0,
		'{': 0,
		'}': 0,
	}

	for _, caractere := range expressao {
		switch caractere {
			case '(', '[', '{':
				pilha.Push(caractere)
			case ')', ']', '}':
				abertura, ok := pilha.Peek()
				if (abertura == 0 || !ok) {
					items_faltando[caractere]++
					continue
				} else if pares[caractere] == abertura {
					pilha.Pop()
					continue
				}

				items_faltando[caractere]++
		}
	}

	empty := pilha.IsEmpty()

	for !pilha.IsEmpty() {
		abertura, _ := pilha.Pop()
		items_faltando[abertura]++
	}

	for key, value := range items_faltando {
		if (items_faltando[pares[key]] > 0) {
			less := items_faltando[pares[key]] > value
			if less {
				items_faltando[pares[key]] -= value
				items_faltando[key] = 0
			} else {
				items_faltando[key] -= items_faltando[pares[key]]
				items_faltando[pares[key]] = 0
			}
		} else if (value > 0) {
			items_faltando[key] = value
		}
	}

	return empty, items_faltando
}

func main() {
	fmt.Println("Digite a expressão matemática: ")
	var expressao string
	fmt.Scanln(&expressao)

	if len(expressao) == 0 {
		// ((a + b] * [c - d])
		// a + (b, {[a * b], (a + b)}, a + {b - [c * d}
		expressao = "a + (b, {[a * b], (a + b)}, a + {b - [c * d}"
		fmt.Println("Expressão padrão utilizada: ", expressao)
	}

	expressao = strings.TrimSpace(expressao)
	empty, items_faltando := estaBalanceada(expressao)
	if empty {
		fmt.Println("A expressão está balanceada.")
		return
	}


	for key, value := range items_faltando {
		if value > 0 {
			for key_p, pair := range pares {
				if key_p == key {
					fmt.Printf("Faltam %d itens de %c\n", value, pair)
				} else if pair == key {
					fmt.Printf("Faltam %d itens de %c\n", value, key_p)
				}
			}
		}
	}

	fmt.Println("A expressão não está balanceada.")
}
