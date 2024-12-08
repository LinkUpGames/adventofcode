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

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var pos []int
			character := crossword[i][j]

			// We found a character
			if character == word[0] {
				// Up
				pos = []int{i - 1, j}
				total += nextWord(crossword, pos, []int{-1, 0}, 1)
				// Down
				pos = []int{i + 1, j}
				total += nextWord(crossword, pos, []int{1, 0}, 1)
				// Left
				pos = []int{i, j + 1}
				total += nextWord(crossword, pos, []int{0, 1}, 1)
				// Right
				pos = []int{i, j - 1}
				total += nextWord(crossword, pos, []int{-1, -1}, 1)
				// Diagonal Left Up
				pos = []int{i - 1, j + 1}
				total += nextWord(crossword, pos, []int{-1, 1}, 1)
				// Diagonal Left Down
				pos = []int{i + 1, j + 1}
				total += nextWord(crossword, pos, []int{1, 1}, 1)
				// Diagonal Right Up
				pos = []int{i - 1, j - 1}
				total += nextWord(crossword, pos, []int{-1, -1}, 1)
				// Diagonal Right Down
				pos = []int{i + 1, j - 1}
				total += nextWord(crossword, pos, []int{1, -1}, 1)
			}
		}
	}

	return total
}

// check the next word in the sequence
func nextWord(crossword [][]byte, pos []int, amount []int, index int) int {
	i := pos[0]
	j := pos[1]

	// out of bounds
	if i < 0 || j < 0 || i >= len(crossword) || j >= len(crossword[0]) {
		return 0
	}

	// Bingo
	if index >= len(word) {
		return 1
	}

	character := crossword[i][j]

	// Check if the character differs at the position
	if character != word[index] {
		return 0
	}

	return nextWord(crossword, []int{i + amount[0], j + amount[1]}, amount, index+1)
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
