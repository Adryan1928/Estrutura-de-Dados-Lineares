package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type No struct {
	Coeficiente float64
	Expoente    int
	proximo     *No
}

type Termo struct {
	Coeficiente float64
	Expoente    int
}

type ListaOrdenada interface {
	ObterProximo(no *No) *No
	ObterValor(no *No) (coeficiente float64, expoente int, encontrado bool)
	AlterarNo(no *No, novoCoeficiente float64, novoExpoente int) bool
	Tamanho() int
	Existe(no *No) bool
	MostrarALL() []Termo
	Exibir() string
	FormatarTermo(no *No) string
	Buscar(expoente int) *No
	Inserir(coeficiente float64, expoente int) *No
	InserirNo(novo *No)
	Simplificar()
	RemoverNo(no *No) bool
	Excluir(no *No) bool
	Destrutor()
	Grau() int
	Avaliar(valor float64) float64
	CopiarTermos(origem *Lista, multiplicador int)
	Somar(outra *Lista) *Lista
	Subtrair(outra *Lista) *Lista
	Multiplicar(outra *Lista) *Lista
	LerPolinomio(linha string) error
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

func (l *Lista) ObterValor(no *No) (coeficiente float64, expoente int, encontrado bool) {
	if !l.Existe(no) {
		return 0, 0, false
	}
	return no.Coeficiente, no.Expoente, true
}

func (l *Lista) AlterarNo(no *No, novoCoeficiente float64, novoExpoente int) bool {
	if !l.RemoverNo(no) {
		return false
	}

	no.Expoente = novoExpoente
	no.Coeficiente = novoCoeficiente
	no.proximo = nil
	l.InserirNo(no)
	return true
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

func (l *Lista) MostrarALL() []Termo {
	termos := make([]Termo, 0, l.tamanho)
	for atualNo := l.inicio; atualNo != nil; atualNo = atualNo.proximo {
		termos = append(termos, Termo{
			Expoente:    atualNo.Expoente,
			Coeficiente: atualNo.Coeficiente,
		})
	}
	return termos
}

func (l *Lista) Exibir() string {
	if l.inicio == nil {
		return "0"
	}

	partes := make([]string, 0, l.tamanho)
	for atualNo := l.inicio; atualNo != nil; atualNo = atualNo.proximo {
		partes = append(partes, l.FormatarTermo(atualNo))
	}

	for esquerda, direita := 0, len(partes)-1; esquerda < direita; esquerda, direita = esquerda+1, direita-1 {
		partes[esquerda], partes[direita] = partes[direita], partes[esquerda]
	}

	polinomio := partes[0]
	for _, parte := range partes[1:] {
		if parte[0] == '-' {
			polinomio += " - " + parte[1:]
			continue
		}
		polinomio += " + " + parte
	}
	return polinomio
}

func (l *Lista) FormatarTermo(no *No) string {
	coeficiente, expoente, _ := l.ObterValor(no)

	if expoente == 0 {
		return strconv.FormatFloat(coeficiente, 'G', -1, 64)
	}

	variavel := "x"
	if expoente != 1 {
		variavel += "^" + strconv.Itoa(expoente)
	}

	modulo := coeficiente
	if modulo < 0 {
		modulo = -modulo
	}

	if modulo == 1 {
		if coeficiente < 0 {
			return "-" + variavel
		}
		return variavel
	}

	return strconv.FormatFloat(coeficiente, 'G', -1, 64) + variavel
}

func (l *Lista) Buscar(expoente int) *No {
	for atualNo := l.inicio; atualNo != nil && atualNo.Expoente <= expoente; atualNo = atualNo.proximo {
		if atualNo.Expoente == expoente {
			return atualNo
		}
	}
	return nil
}

func (l *Lista) Inserir(coeficiente float64, expoente int) *No {
	novo := &No{Coeficiente: coeficiente, Expoente: expoente}
	l.InserirNo(novo)
	return novo
}

func (l *Lista) InserirNo(novo *No) {
	if l.inicio == nil || novo.Expoente < l.inicio.Expoente {
		novo.proximo = l.inicio
		l.inicio = novo
		l.tamanho++
		return
	}

	atualNo := l.inicio
	for atualNo.proximo != nil && atualNo.proximo.Expoente <= novo.Expoente {
		atualNo = atualNo.proximo
	}
	novo.proximo = atualNo.proximo
	atualNo.proximo = novo
	l.tamanho++

	l.Simplificar()
}

func (l *Lista) Simplificar() {
	for atualNo := l.inicio; atualNo != nil && atualNo.proximo != nil; atualNo = atualNo.proximo {
		for proximoNo := atualNo.proximo; proximoNo != nil; proximoNo = proximoNo.proximo {
			if atualNo.Expoente == proximoNo.Expoente {
				atualNo.Coeficiente += proximoNo.Coeficiente
				l.RemoverNo(proximoNo)
			}
		}

		if atualNo.Coeficiente == 0 {
			l.RemoverNo(atualNo)
			continue
		}
	}
}

func (l *Lista) RemoverNo(no *No) bool {
	if no == nil || l.inicio == nil {
		return false
	}

	if l.inicio == no {
		l.inicio = no.proximo
		no.proximo = nil
		l.tamanho--
		return true
	}

	for atualNo := l.inicio; atualNo != nil && atualNo.proximo != nil; atualNo = atualNo.proximo {
		if atualNo.proximo == no {
			atualNo.proximo = no.proximo
			no.proximo = nil
			l.tamanho--
			return true
		}
	}
	return false
}

func (l *Lista) Excluir(no *No) bool {
	return l.RemoverNo(no)
}

func (l *Lista) Destrutor() {
	for l.inicio != nil {
		atual := l.inicio
		l.inicio = atual.proximo
		atual.proximo = nil
	}
	l.tamanho = 0
}

func (l *Lista) Grau() int {
	if l.inicio == nil {
		return -1
	}

	expoente := l.inicio.Expoente
	for atualNo := l.inicio; atualNo != nil; atualNo = atualNo.proximo {
		if atualNo.Expoente > expoente {
			expoente = atualNo.Expoente
		}
	}
	return expoente
}

func (l *Lista) Avaliar(valor float64) float64 {
	resultado := 0.0
	for atualNo := l.inicio; atualNo != nil; atualNo = atualNo.proximo {
		resultado += atualNo.Coeficiente * math.Pow(valor, float64(atualNo.Expoente))
	}
	return resultado
}

func (l *Lista) CopiarTermos(origem *Lista, multiplicador int) {
	if origem == nil {
		return
	}
	for atualNo := origem.inicio; atualNo != nil; atualNo = atualNo.proximo {
		l.Inserir(atualNo.Coeficiente*float64(multiplicador), atualNo.Expoente)
	}
}

func (l *Lista) Somar(outra *Lista) *Lista {
	resultado := NovaLista()
	resultado.CopiarTermos(l, 1)
	resultado.CopiarTermos(outra, 1)
	return resultado
}

func (l *Lista) Subtrair(outra *Lista) *Lista {
	resultado := NovaLista()
	resultado.CopiarTermos(l, 1)
	resultado.CopiarTermos(outra, -1)
	return resultado
}

func (l *Lista) Multiplicar(outra *Lista) *Lista {
	resultado := NovaLista()
	if outra == nil {
		return resultado
	}

	for termoEsquerdo := l.inicio; termoEsquerdo != nil; termoEsquerdo = termoEsquerdo.proximo {
		for termoDireito := outra.inicio; termoDireito != nil; termoDireito = termoDireito.proximo {
			resultado.Inserir(
				termoEsquerdo.Coeficiente*termoDireito.Coeficiente,
				termoEsquerdo.Expoente+termoDireito.Expoente,
			)
		}
	}
	return resultado
}

func (l *Lista) LerPolinomio(linha string) error {
	campos := strings.Fields(linha)
	if len(campos) == 0 {
		return fmt.Errorf("linha de polinômio vazia")
	}
	if len(campos)%2 != 0 {
		return fmt.Errorf("o polinômio deve ter pares de coeficiente e expoente: %q", linha)
	}

	for indice := 0; indice < len(campos); indice += 2 {
		coeficiente, erro := strconv.ParseFloat(campos[indice], 64)
		if erro != nil {
			return fmt.Errorf("coeficiente inválido %q: %w", campos[indice], erro)
		}
		expoente, erro := strconv.Atoi(campos[indice+1])
		if erro != nil {
			return fmt.Errorf("expoente inválido %q", campos[indice+1])
		}
		l.Inserir(coeficiente, expoente)
	}

	return nil
}