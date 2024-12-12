package day7

import (
	"strconv"
	"strings"

	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

func Task1() int {
	input := helper.ReadFile("./day7/day7.input")
	sum := 0

	for _, line := range input {
		split := strings.Split(line, ":")
		result, _ := strconv.Atoi(split[0])
		values := getIntValues(split[1])

		operations := []rune{'*', '+', '|'}
		combinations := generateCombinations(operations, len(values)-1)

		for _, combination := range combinations {
			symbols := []rune(combination)
			symbolCount := 0
			guess := calculate(values[0], values[1], symbols[symbolCount])
			symbolCount++

			for i := 2; i < len(values); i++ {
				guess = calculate(guess, values[i], symbols[symbolCount])
				symbolCount++
			}

			if guess == result {
				sum += result
				break
			}

		}

	}

	return sum
}

func Task2() int {

	return 0
}

func getIntValues(str string) []int {
	valuesStr := strings.Split(str[1:], " ")
	valuesInt := []int{}

	for _, value := range valuesStr {
		conv, _ := strconv.Atoi(value)
		valuesInt = append(valuesInt, conv)
	}

	return valuesInt
}

func generateCombinations(chars []rune, length int) []string {
	if length <= 0 {
		return []string{}
	}

	var results []string
	var backtrack func(path []rune)

	backtrack = func(path []rune) {
		if len(path) == length {
			results = append(results, string(path))
			return
		}

		for _, char := range chars {
			backtrack(append(path, char))
		}
	}

	backtrack([]rune{})
	return results
}

func calculate(a, b int, symbol rune) int {
	if symbol == '*' {
		return a * b
	}

	if symbol == '|' {
		strJoin := strconv.Itoa(a) + strconv.Itoa(b)
		conv, _ := strconv.Atoi(strJoin)
		return conv
	}

	return a + b
}
