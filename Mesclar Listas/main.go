package main

import (
	"fmt"
);

// No primeiro exemplo lida com entradas de listas não ordenadas, e consegue ordenalas com uma complexidade de o(n elevado a 2). No segundo exemplo, lida com entradas de listas ordenadas, e consegue mesclar as listas ordenadas com complexidade de o(n).

func main() {
	lista1 := NovaLista()
	lista1.Inserir(2)
	lista1.Inserir(3)
	lista1.Inserir(5)

	lista2 := NovaLista()
	lista2.Inserir(1)
	lista2.Inserir(4)
	lista2.Inserir(6)
	lista2.Inserir(101)

	mergedList := NovaLista()
	mergedList.MesclarListas(lista1, lista2)

	fmt.Println("Lista mesclada:")
	for no := mergedList.inicio; no != nil; no = no.proximo {
		fmt.Printf("%d ", no.item)
	}
	fmt.Println()


	fmt.Println("Lista com intercalar")
	no3 := Intercalar(lista1.inicio, lista2.inicio)

	for no3 != nil {
		fmt.Printf("%d ", no3.item)
		no3 = no3.proximo
	}
	fmt.Println()
}