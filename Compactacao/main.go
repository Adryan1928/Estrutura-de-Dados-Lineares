package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	nomeEntrada        = "entrada.txt"
	nomeCompactado     = "compactado.bin"
	nomeDescompactacao = "descomptacao.txt"
)

var baseParaBits = map[byte]byte{
	'A': 0b00,
	'C': 0b01,
	'G': 0b10,
	'T': 0b11,
}

var bitsParaBase = [4]byte{'A', 'C', 'G', 'T'}

func main() {
	diretorio, err := localizarDiretorio()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	entradaPath := filepath.Join(diretorio, nomeEntrada)
	compactadoPath := filepath.Join(diretorio, nomeCompactado)
	descompactacaoPath := filepath.Join(diretorio, nomeDescompactacao)

	dna, err := lerDNA(entradaPath)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	compactado, err := compactar(dna)
	if err != nil {
		fmt.Println("Erro ao compactar:", err)
		return
	}
	if err := os.WriteFile(compactadoPath, compactado, 0o644); err != nil {
		fmt.Println("Erro ao salvar o arquivo compactado:", err)
		return
	}

	compactadoLido, err := os.ReadFile(compactadoPath)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo compactado:", err)
		return
	}
	descompactado, err := descompactar(compactadoLido)
	if err != nil {
		fmt.Println("Erro ao descompactar:", err)
		return
	}
	if err := os.WriteFile(descompactacaoPath, []byte(descompactado), 0o644); err != nil {
		fmt.Println("Erro ao salvar a descompactação:", err)
		return
	}

	fmt.Printf("%d bases compactadas em %d bytes.\n", len(dna), len(compactado))
	fmt.Println("Arquivo compactado:", compactadoPath)
	fmt.Println("Arquivo descompactado:", descompactacaoPath)
}

func localizarDiretorio() (string, error) {
	for _, diretorio := range []string{".", "Compactacao"} {
		if info, err := os.Stat(filepath.Join(diretorio, nomeEntrada)); err == nil && !info.IsDir() {
			return diretorio, nil
		}
	}
	return "", fmt.Errorf("%s não encontrado", nomeEntrada)
}

func lerDNA(caminho string) (string, error) {
	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		return "", err
	}

	var dna strings.Builder
	for posicao, caractere := range string(conteudo) {
		if unicode.IsSpace(caractere) {
			continue
		}

		base := byte(unicode.ToUpper(caractere))
		if _, existe := baseParaBits[base]; !existe {
			return "", fmt.Errorf("base inválida na posição %d: %q (use somente A, C, G ou T)", posicao+1, caractere)
		}
		dna.WriteByte(base)
	}

	if dna.Len() == 0 {
		return "", fmt.Errorf("a sequência de DNA está vazia")
	}
	return dna.String(), nil
}

func compactar(dna string) ([]byte, error) {
	if len(dna) == 0 {
		return nil, fmt.Errorf("a sequência de DNA está vazia")
	}

	quantidadeBytesDNA := (len(dna) + 3) / 4
	resultado := make([]byte, quantidadeBytesDNA+1)

	for indice, base := range []byte(dna) {
		bits, existe := baseParaBits[base]
		if !existe {
			return nil, fmt.Errorf("base inválida na posição %d: %q", indice+1, base)
		}

		indiceByte := indice / 4
		deslocamento := uint(6 - (indice%4)*2)
		resultado[indiceByte] |= bits << deslocamento
	}

	quantidadeNoUltimoByte := (len(dna)-1)%4 + 1
	resultado[len(resultado)-1] = byte(quantidadeNoUltimoByte - 1)
	return resultado, nil
}

func descompactar(compactado []byte) (string, error) {
	if len(compactado) < 2 {
		return "", fmt.Errorf("arquivo compactado inválido: faltam dados ou metadado final")
	}

	metadado := compactado[len(compactado)-1]
	if metadado&0b11111100 != 0 {
		return "", fmt.Errorf("arquivo compactado inválido: metadado final desconhecido")
	}

	dados := compactado[:len(compactado)-1]
	quantidadeNoUltimoByte := int(metadado&0b11) + 1
	var dna strings.Builder
	dna.Grow((len(dados)-1)*4 + quantidadeNoUltimoByte)

	for indiceByte, byteCompactado := range dados {
		quantidadeBases := 4
		if indiceByte == len(dados)-1 {
			quantidadeBases = quantidadeNoUltimoByte
		}

		for indiceBase := 0; indiceBase < quantidadeBases; indiceBase++ {
			deslocamento := uint(6 - indiceBase*2)
			bits := (byteCompactado >> deslocamento) & 0b11
			dna.WriteByte(bitsParaBase[bits])
		}
	}

	return dna.String(), nil
}
