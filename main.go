package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, _ := parseFile("default.txt")

	printRules(rules)
}

func parseFile(input string) (map[int][]int, []string) {
	// [key] = [1,2,3] == key depends on 1,2,3
	rules := make(map[int][]int)

	data, _ := os.ReadFile(input)
	sections := strings.Split(string(data), "\n\n")

	// Rules
	rulesText := strings.Split(sections[0], "\n")
	for _, rule := range rulesText {
		nums := strings.Split(rule, "|")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		_, exists := rules[b]
		if !exists {
			rules[b] = make([]int, 0)
		}
		rules[b] = append(rules[b], a)
	}

	// Sequence
	sequences := strings.Split(sections[1], "\n")

	return rules, sequences
}

func printRules(rules map[int][]int) {
	for k, v := range rules {
		fmt.Printf("[%d]: [", k)

		for _, num := range v {
			fmt.Printf("%d,", num)
		}
		fmt.Println("]")
	}
}

func checkSequence(sequence string) int {
	return 0
}
