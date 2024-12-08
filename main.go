package main

import (
	"fmt"
	"os"
	"strings"
)

var word string = "XMAS"

func main() {
	matrix := getMatrix("default.txt")

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
				total += nextWord(crossword, pos, 1)
				// Down
				pos = []int{i + 1, j}
				total += nextWord(crossword, pos, 1)
				// Left
				pos = []int{i, j + 1}
				total += nextWord(crossword, pos, 1)
				// Right
				pos = []int{i, j - 1}
				total += nextWord(crossword, pos, 1)
			}
		}
	}

	return total
}

func nextWord(crossword [][]byte, coords []int, index int) int {
	i := coords[0]
	j := coords[1]

	// Correct boundary check
	if i < 0 || i >= len(crossword) || j < 0 || j >= len(crossword[0]) {
		return 0
	}

	// Correct character comparison and bounds check
	if index >= len(word) {
		return 0 // Word has been found
	}
	if crossword[i][j] != word[index] {
		return 0 // Character does not match
	}

	if index == len(word)-1 {
		return 1 // Word found!
	}

	return nextWord(crossword, []int{i - 1, j}, index+1) + // Up
		nextWord(crossword, []int{i + 1, j}, index+1) + // Down
		nextWord(crossword, []int{i, j + 1}, index+1) + // Right
		nextWord(crossword, []int{i, j - 1}, index+1) + // Left
		nextWord(crossword, []int{i - 1, j + 1}, index+1) + // Diagonal Right Up
		nextWord(crossword, []int{i + 1, j + 1}, index+1) + // Diagonal Right down
		nextWord(crossword, []int{i - 1, j - 1}, index+1) + // Diagonal left up
		nextWord(crossword, []int{i + 1, j - 1}, index+1) // Diagonal left down
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
