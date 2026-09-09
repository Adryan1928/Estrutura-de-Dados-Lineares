package main

import (
	"fmt"
)

type No struct {
	item int
	proximo     *No
}

type Termo struct {
	item int
}

type ListaOrdenada interface {
	ObterProximo(no *No) *No
	Tamanho() int
	Existe(no *No) bool
	Inserir(item int) *No
	InserirNo(novo *No)
	OrdenarLista()
	MesclarListas(lista1, lista2 *Lista)
}

type Lista struct {
	inicio  *No
	tamanho int
}

var _ ListaOrdenada = (*Lista)(nil)

func NovaLista() *Lista {
	return &Lista{}
}

func (l *Lista) ObterProximo(no *No) *No {
	if !l.Existe(no) {
		return nil
	}
	return no.proximo
}

func (l *Lista) Tamanho() int {
	return l.tamanho
}

func (l *Lista) Existe(no *No) bool {
	for atualNo := l.inicio; atualNo != nil; atualNo = atualNo.proximo {
		if atualNo == no {
			return true
		}
	}
	return false
}

func (l *Lista) Inserir(item int) *No {
	if (l.Tamanho() >= 50) {
		fmt.Println("A lista atingiu o tamanho máximo de 50 elementos. Não é possível inserir mais elementos.")
		return nil
	}

	if (item > 100 || item < -100) {
		fmt.Println("O item deve estar entre -100 e 100.")
		return nil
	}
	novo := &No{item: item}
	l.InserirNo(novo)
	return novo
}

func (l *Lista) InserirNo(novo *No) {
	if l.inicio == nil {
		l.inicio = novo
		l.tamanho++
		return
	}

	for atualNo := l.inicio; atualNo != nil; atualNo = atualNo.proximo {
		if atualNo.proximo == nil {
			atualNo.proximo = novo
			l.tamanho++
			return
		}
	}
}

func (l *Lista) MesclarListas(lista1, lista2 *Lista) {
	if lista1 == nil || lista2 == nil {
		return
	}

	no1 := lista1.inicio
	no2 := lista2.inicio

	for no1 != nil {
		l.Inserir(no1.item)
		no1 = no1.proximo
	}

	for no2 != nil {
		l.Inserir(no2.item)
		no2 = no2.proximo
	}

	l.OrdenarLista()
}

func (l *Lista) OrdenarLista() {
	if l.inicio == nil || l.inicio.proximo == nil {
		return
	}

	for atual := l.inicio; atual != nil; atual = atual.proximo {
		for proximo := atual.proximo; proximo != nil; proximo = proximo.proximo {
			if atual.item > proximo.item {
				atual.item, proximo.item = proximo.item, atual.item
			}
		}
	}
}

func Intercalar(esquerda *No, direita *No) *No {
    var inicio *No
    var fim *No

    for esquerda != nil && direita != nil {
        var menor *No

        if esquerda.item <= direita.item {
            menor = esquerda
            esquerda = esquerda.proximo
        } else {
            menor = direita
            direita = direita.proximo
        }

        if inicio == nil {
            inicio = menor
            fim = menor
        } else {
            fim.proximo = menor
            fim = menor
        }
    }

    if esquerda != nil {
        fim.proximo = esquerda
    }

    if direita != nil {
        fim.proximo = direita
    }

    return inicio
}