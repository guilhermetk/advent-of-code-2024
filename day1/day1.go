package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

func Task2() int64 {
	day1Input := helper.ReadFile("./day1/day1.input")

	list1 := []int{}
	list2 := []int{}

	for i := 0; i < len(day1Input); i++ {
		values := strings.Split(day1Input[i], "   ")

		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])

		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	return calculateSimilarity(list1, list2)
}

func Task1() int64 {
	day1Input := helper.ReadFile("./day1/day1.input")

	list1 := []int{}
	list2 := []int{}

	for i := 0; i < len(day1Input); i++ {
		values := strings.Split(day1Input[i], "   ")

		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])

		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	return calculateDistance(list1, list2)

}

func calculateDistance(list1, list2 []int) int64 {
	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0.0

	for i := 0; i < len(list1); i++ {
		sum += math.Abs(float64(list1[i]) - float64(list2[i]))
	}

	return int64(sum)
}

func calculateSimilarity(list1, list2 []int) int64 {
	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		count := 0
		for j := 0; j < len(list2) && list2[j] <= list1[i]; j++ {
			if list1[i] == list2[j] {
				count += 1
			}

		}
		sum += count * list1[i]
	}

	return int64(sum)
}
