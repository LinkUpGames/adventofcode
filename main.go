package main

import (
	"fmt"
	"os"
	"strings"
)

var word string = "XMAS"

func main() {
	matrix := getMatrix("input.txt")

	printCrossword(matrix)
	total := checkCrossword(matrix)

	fmt.Printf("Total: %d\n", total)
}

func getMatrix(input string) [][]byte {
	raw, _ := os.ReadFile(input)
	data := string(raw)

	lines := strings.Split(data, "\n")

	rows := len(lines) - 1
	matrix := make([][]byte, rows)

	for i, line := range lines {
		if len(line) > 0 {
			matrix[i] = make([]byte, len(line))

			for j, c := range line {
				matrix[i][j] = byte(c)
			}
		}
	}

	return matrix
}

func checkCrossword(crossword [][]byte) int {
	total := 0
	rows := len(crossword)
	columns := len(crossword[0])

	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Up, Down, Left, Right
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonals
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if crossword[i][j] == word[0] { // Only start search if the first letter matches
				for _, dir := range directions {
					if nextWord(crossword, i, j, dir[0], dir[1], 1) {
						total++
					}
				}
			}
		}
	}

	return total
}

// check the next word in the sequence
func nextWord(crossword [][]byte, x, y, dx, dy, index int) bool {
	if index >= len(word) {
		return true // Found the word
	}

	nx, ny := x+dx, y+dy // Calculate next position
	if nx < 0 || nx >= len(crossword) || ny < 0 || ny >= len(crossword[0]) {
		return false // Out of bounds
	}

	if crossword[nx][ny] != word[index] {
		return false // Character doesn't match
	}

	return nextWord(crossword, nx, ny, dx, dy, index+1)
}

func printCrossword(crossword [][]byte) {
	rows := len(crossword)
	columns := len(crossword[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("[%c]", crossword[i][j])
		}
		fmt.Println()
	}
}
