package day2

import (
	"math"
	"strconv"
	"strings"

	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

func Task1() int {
	reports := helper.ReadFile("./day2/day2.input")

	safeLevelsCount := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		safeLevelsCount += isReportSafe(levels)
	}
	return safeLevelsCount
}

func Task2() int {
	reports := helper.ReadFile("./day2/day2.input")

	safeReportsCount := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		isOriginalReportSafe := isReportSafe(levels)

		if isOriginalReportSafe > 0 {
			safeReportsCount += 1
		} else {
			for i := 0; i < len(levels); i++ {
				partialReport := []string{}
				for j := 0; j < len(levels); j++ {
					if i != j {
						partialReport = append(partialReport, levels[j])
					}
				}
				partialLevelSafe := isReportSafe(partialReport)
				if partialLevelSafe > 0 {
					safeReportsCount += 1
					break
				}
			}
		}
	}
	return safeReportsCount
}

func isReportSafe(levels []string) int {
	lvlComparisons := len(levels) - 1

	safeLevel := 1
	direction := 0

	for i := 0; i < lvlComparisons; i++ {
		a, _ := strconv.Atoi(levels[i])
		b, _ := strconv.Atoi(levels[i+1])

		diff := math.Abs(float64(b) - float64(a))

		if b > a {
			direction += 1
		}

		if a > b {
			direction -= 1
		}

		if diff < 1.0 || diff > 3.0 {
			safeLevel = 0
		}

	}

	if direction != lvlComparisons && direction != lvlComparisons*-1 {
		safeLevel = 0
	}

	return safeLevel
}
