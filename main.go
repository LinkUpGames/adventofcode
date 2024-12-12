package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, sequences := parseFile("default.txt")

	printRules(rules)
	for _, sequence := range sequences {
		checkSequence(sequence, rules)
	}
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

func checkSequence(sequence string, rules map[int][]int) int {
	seen := make(map[int]bool)
	nums := strings.Split(sequence, ",")

	for _, strnum := range nums {
		num, _ := strconv.Atoi(strnum)

		// Seen this one
		seen[num] = true

		beforenums, _ := rules[num]

		for _, num := range beforenums {
			_, exists := seen[num]
			if exists {
				fmt.Printf("Hello: %d\n", exists)
			}
		}
	}

	return 0
}

func rulesForSequence(sequence string, rules map[int][]int) map[int][]int {
	seen := make(map[int]bool)

	nums := strings.Split(sequence, ",")

	// Go through the sequence and get the numbers that actually apply
	for _, strnum := range nums {
		num, _ := strconv.Atoi(strnum)
		seen[num] = true
	}

	return rules
}

func removeFromArray[T comparable](element any, array []T) []T {
	index := 0
	for i, value := range array {
		if value == element {
			index = i
			break
		}
	}

	array[index] = array[len(array)-1]

	return array[:len(array)-1]
}
