package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

func Task1() int {
	inputs := helper.ReadFileDay5("./day5/day5.input")
	rules := parseInputs(inputs[0])
	middleNumbersSum := 0

	for _, update := range inputs[1] {
		correctlyOrdered := true

		pages := strings.Split(update, ",")

		reversePages := slices.Clone(pages)
		slices.Reverse(reversePages)

		for i := 0; i < len(pages); i++ {
			curPage := pages[i]
			curPageRules, exists := rules[curPage]

			revCurPage := reversePages[i]
			revCurPageRules, exists := rules[revCurPage]

			if !exists {
				continue
			}

			for _, nextPage := range pages[i+1:] {
				if findPageInRules(nextPage, curPageRules) {
					continue
				}

				nextPageRules := rules[nextPage]
				if findPageInRules(curPage, nextPageRules) {
					correctlyOrdered = false
				}
			}

			for _, nextPage := range reversePages[i+1:] {
				if findPageInRules(nextPage, revCurPageRules) {
					correctlyOrdered = false
				}
			}
		}
		if correctlyOrdered {
			middleIndex := int64(math.Floor(float64((len(pages) - 1) / 2)))
			middleItem, _ := strconv.Atoi(pages[middleIndex])
			middleNumbersSum += middleItem
		}
	}

	return middleNumbersSum
}

func Task2() int {
	inputs := helper.ReadFileDay5("./day5/day5.input")
	rules := parseInputs(inputs[0])
	middleNumbersSum := 0

	for _, update := range inputs[1] {
		correctlyOrdered := true

		pages := strings.Split(update, ",")

		reversePages := slices.Clone(pages)
		slices.Reverse(reversePages)

		for i := 0; i < len(pages); i++ {
			curPage := pages[i]
			curPageRules, exists := rules[curPage]

			revCurPage := reversePages[i]
			revCurPageRules, exists := rules[revCurPage]

			if !exists {
				continue
			}

			for _, nextPage := range pages[i+1:] {
				if findPageInRules(nextPage, curPageRules) {
					continue
				}

				nextPageRules := rules[nextPage]
				if findPageInRules(curPage, nextPageRules) {
					correctlyOrdered = false
				}
			}

			for _, nextPage := range reversePages[i+1:] {
				if findPageInRules(nextPage, revCurPageRules) {
					correctlyOrdered = false
				}
			}
		}
		if !correctlyOrdered {
			for i := 0; i < len(pages)-1; i++ {
				for j := i + 1; j < len(pages); j++ {
					curPage := strings.Clone(pages[i])
					nextPage := strings.Clone(pages[j])
					nextPageRules, exists := rules[nextPage]

					if !exists {
						continue
					}

					if findPageInRules(curPage, nextPageRules) {
						pages[i] = nextPage
						pages[j] = curPage

					}

				}
			}

			middleIndex := int64(math.Floor(float64((len(pages) - 1) / 2)))
			middleItem, _ := strconv.Atoi(pages[middleIndex])
			middleNumbersSum += middleItem
		}
	}

	return middleNumbersSum
}

func parseInputs(input []string) map[string][]string {
	rules := map[string][]string{}

	for _, rule := range input {
		split := strings.Split(rule, "|")
		_, exists := rules[split[0]]

		if !exists {
			rules[split[0]] = []string{}
		}
		rules[split[0]] = append(rules[split[0]], split[1])
	}

	return rules
}

func findPageInRules(page string, rules []string) bool {
	return slices.Index(rules, page) >= 0
}
