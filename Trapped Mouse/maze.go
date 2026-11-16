package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	exitMarker  = 'e'
	entryMarker = 'm'
	visited     = '.'
	passage     = '0'
	wall        = '1'
)

type Cell struct {
	x int
	y int
}

func (c Cell) Equals(other Cell) bool {
	return c.x == other.x && c.y == other.y
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	s.items = s.items[:last]
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

type Maze struct {
	currentCell Cell
	exitCell    Cell
	entryCell   Cell

	mazeStack Stack[Cell]
	maze      [][]rune
}

func NewMaze(lines []string) (*Maze, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("o labirinto não possui linhas")
	}

	width := len([]rune(strings.TrimSpace(lines[0])))
	if width == 0 {
		return nil, fmt.Errorf("a primeira linha está vazia")
	}

	var rows Stack[[]rune]
	entryCount, exitCount := 0, 0
	for lineNumber, line := range lines {
		runes := []rune(strings.TrimSpace(line))
		if len(runes) != width {
			return nil, fmt.Errorf("linha %d possui %d colunas; eram esperadas %d", lineNumber+1, len(runes), width)
		}

		for column, value := range runes {
			switch value {
				case wall, passage:
				case entryMarker:
					entryCount++
				case exitMarker:
					exitCount++
				default:
					return nil, fmt.Errorf("caractere inválido %q na linha %d, coluna %d", value, lineNumber+1, column+1)
			}
		}

		rowWithWalls := make([]rune, 0, width+2)
		rowWithWalls = append(rowWithWalls, wall)
		rowWithWalls = append(rowWithWalls, runes...)
		rowWithWalls = append(rowWithWalls, wall)
		rows.Push(rowWithWalls)
	}

	if entryCount != 1 || exitCount != 1 {
		return nil, fmt.Errorf("o labirinto deve possuir exatamente um 'm' e um 'e' (encontrados: m=%d, e=%d)", entryCount, exitCount)
	}

	maze := make([][]rune, 0, len(lines)+2)
	border := make([]rune, width+2)
	for i := range border {
		border[i] = wall
	}
	maze = append(maze, border)

	innerRows := make([][]rune, 0, len(lines))
	for !rows.IsEmpty() {
		row, _ := rows.Pop()
		innerRows = append(innerRows, row)
	}
	for i := len(innerRows) - 1; i >= 0; i-- {
		maze = append(maze, innerRows[i])
	}
	maze = append(maze, append([]rune(nil), border...))

	m := &Maze{maze: maze}
	for x, row := range maze {
		for y, value := range row {
			switch value {
			case entryMarker:
				m.entryCell = Cell{x: x, y: y}
			case exitMarker:
				m.exitCell = Cell{x: x, y: y}
			}
		}
	}
	m.currentCell = m.entryCell
	return m, nil
}

func (m *Maze) String() string {
	var builder strings.Builder
	for x, row := range m.maze {
		for y, value := range row {
			if m.currentCell.Equals(Cell{x: x, y: y}) {
				builder.WriteRune(entryMarker)
				continue
			}
			builder.WriteRune(value)
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (m *Maze) popNextUnvisitedCell() (Cell, bool) {
	for !m.mazeStack.IsEmpty() {
		cell, _ := m.mazeStack.Pop()
		if m.hasUnvisitedPassage(cell) {
			return cell, true
		}
	}
	return Cell{}, false
}

func (m *Maze) hasUnvisitedPassage(cell Cell) bool {
	value := m.maze[cell.x][cell.y]
	return value == passage || value == exitMarker
}

func (m *Maze) pushUnvisitedNeighbors(cell Cell) {
	neighbors := []Cell{
		{x: cell.x - 1, y: cell.y},
		{x: cell.x + 1, y: cell.y},
		{x: cell.x, y: cell.y - 1},
		{x: cell.x, y: cell.y + 1},
	}
	for _, neighbor := range neighbors {
		if m.hasUnvisitedPassage(neighbor) {
			m.mazeStack.Push(neighbor)
		}
	}
}

func (m *Maze) ExitMaze(show func(step int, maze *Maze)) bool {
	step := 0
	show(step, m)

	for !m.currentCell.Equals(m.exitCell) {
		m.maze[m.currentCell.x][m.currentCell.y] = visited
		m.pushUnvisitedNeighbors(m.currentCell)

		next, ok := m.popNextUnvisitedCell()
		if !ok {
			return false
		}

		m.currentCell = next
		step++
		show(step, m)
	}

	return true
}

func readMazesFromFile(path string) ([]*Maze, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readMazes(file)
}

func readMazes(reader io.Reader) ([]*Maze, error) {
	scanner := bufio.NewScanner(reader)
	var mazes []*Maze
	var lines []string
	lineNumber := 0

	createMaze := func() error {
		if len(lines) == 0 {
			return nil
		}
		maze, err := NewMaze(lines)
		if err != nil {
			return err
		}
		mazes = append(mazes, maze)
		lines = nil
		return nil
	}

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if err := createMaze(); err != nil {
				return nil, fmt.Errorf("labirinto terminado na linha %d: %w", lineNumber, err)
			}
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if err := createMaze(); err != nil {
		return nil, err
	}
	if len(mazes) == 0 {
		return nil, fmt.Errorf("nenhum labirinto foi informado")
	}
	return mazes, nil
}

func readMazeFromTerminal() (*Maze, error) {
	fmt.Println("Digite as linhas do labirinto (use somente 1, 0, m e e).")
	fmt.Println("Uma linha vazia finaliza a entrada.")

	var lines []string
	for {
		fmt.Printf("> ")
		var line string
		fmt.Scanln(&line)
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return NewMaze(lines)
}
