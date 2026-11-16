package main

import (
	"fmt"
	"time"
)

const defaultMazeFile = "labirintos.txt"

func main() {
	var q string
	fmt.Println("=== Trapped Mouse ===")
	fmt.Println("1 - Ler labirinto(s) de um arquivo")
	fmt.Println("2 - Digitar um labirinto no terminal")
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&q)

	var mazes []*Maze
	var err error
	switch q {
	case "1":
		fmt.Printf("Caminho do arquivo (Enter para %s): ", defaultMazeFile)
		var path string
		fmt.Scanln(&path)
		if path == "" {
			path = defaultMazeFile
		}
		mazes, err = readMazesFromFile(path)
	case "2":
		var maze *Maze
		maze, err = readMazeFromTerminal()
		if err == nil {
			mazes = []*Maze{maze}
		}
	default:
		fmt.Println("Opção inválida.")
		return
	}

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	for index, maze := range mazes {
		if len(mazes) > 1 {
			fmt.Printf("\n=== Labirinto %d ===\n", index+1)
		}

		escaped := maze.ExitMaze(func(step int, current *Maze) {
			if step == 0 {
				fmt.Println("\nEstado inicial:")
			} else {
				fmt.Printf("\nMovimento %d:\n", step)
			}
			fmt.Print(current)
			time.Sleep(200 * time.Millisecond)
		})

		if escaped {
			fmt.Println("O rato encontrou a saída!")
		} else {
			fmt.Println("Não existe caminho até a saída.")
		}
	}
}
