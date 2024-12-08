package day3

import (
	"regexp"
	"strconv"
	"strings"

	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

func Task1() int {
	day3Input := helper.ReadFile("./day3/day3.input")
	input := strings.Join(day3Input, "")

	extractFunctionRegex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	matchedFunctions := extractFunctionRegex.FindAllString(input, -1)

	return parseFunctions(matchedFunctions)
}

func Task2() int {
	day3Input := helper.ReadFile("./day3/day3.input")
	input := strings.Join(day3Input, "")

	clearDontRangesRegex, _ := regexp.Compile(`don't\(\).*?do\(\)`)
	matchedDontIntervals := clearDontRangesRegex.FindAllString(input, -1)

	for _, match := range matchedDontIntervals {
		input = strings.Replace(input, match, "", 1)
	}

	cleanTrailingDontRegex, _ := regexp.Compile(`don't\(\)`)
	matchedTrailingDont := cleanTrailingDontRegex.FindStringIndex(input)

	input = input[:matchedTrailingDont[0]-1]

	extractFunctionRegex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	matchedFunctions := extractFunctionRegex.FindAllString(input, -1)

	return parseFunctions(matchedFunctions)
}

func parseFunctions(results []string) int {
	sum := 0

	for _, result := range results {
		str1 := strings.ReplaceAll(result, `mul(`, "")
		str2 := strings.ReplaceAll(str1, `)`, "")
		nbrs := strings.Split(str2, ",")
		nbr1, _ := strconv.Atoi(nbrs[0])
		nbr2, _ := strconv.Atoi(nbrs[1])

		mult := nbr1 * nbr2
		sum += mult
	}
	return sum
}
