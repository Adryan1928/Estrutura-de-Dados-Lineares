package main

type No struct {
	Coeficiente int
	Expoente int
	proximo *No
}

type Termo struct {
	Coeficiente int
	Expoente int
}

type ListaOrdenada interface {
	ObterProximo(no *No) *No
	ObterValor(no *No) (coeficiente int, expoente int, encontrado bool)
	AlterarNo(no *No, novoCoeficiente int, novoExpoente int) bool
	Tamanho() int
	Existe(no *No) bool
	MostrarALL() []Termo
	Buscar(expoente int) *No
	Inserir(coeficiente int, expoente int) *No
	Excluir(no *No) bool
	Destrutor()
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

func (l *Lista) ObterValor(no *No) (coeficiente int, expoente int, encontrado bool) {
	if !l.Existe(no) {
		return 0, 0, false
	}
	return no.Coeficiente, no.Expoente, true
}

func (l *Lista) AlterarNo(no *No, novoCoeficiente int, novoExpoente int) bool {
	if !l.removerNo(no) {
		return false
	}

	no.Expoente = novoExpoente
	no.Coeficiente = novoCoeficiente
	no.proximo = nil
	l.inserirNo(no)
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

func (l *Lista) Buscar(expoente int) *No {
	for atualNo := l.inicio; atualNo != nil && atualNo.Expoente <= expoente; atualNo = atualNo.proximo {
		if atualNo.Expoente == expoente {
			return atualNo
		}
	}
	return nil
}

func (l *Lista) Inserir(coeficiente int, expoente int) *No {
	novo := &No{Coeficiente: coeficiente, Expoente: expoente}
	l.inserirNo(novo)
	return novo
}

func (l *Lista) Excluir(no *No) bool {
	return l.removerNo(no)
}

func (l *Lista) Destrutor() {
	for l.inicio != nil {
		atual := l.inicio
		l.inicio = atual.proximo
		atual.proximo = nil
	}
	l.tamanho = 0
}

// TODO: Juntar os expoentes iguais
func (l *Lista) inserirNo(novo *No) {
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

	l.simplificar()
}

func (l *Lista) removerNo(no *No) bool {
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

func (l *Lista) simplificar() {
	for atualNo := l.inicio; atualNo != nil && atualNo.proximo != nil; atualNo = atualNo.proximo {
		for proximoNo := atualNo.proximo; proximoNo != nil; proximoNo = proximoNo.proximo {
			if atualNo.Expoente == proximoNo.Expoente {
				atualNo.Coeficiente += proximoNo.Coeficiente
				l.removerNo(proximoNo)
			}
		}
	}
}
