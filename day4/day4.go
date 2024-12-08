package day4

import (
	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

type Point struct {
	X int
	Y int
}

type Word struct {
	Start  Point
	Finish Point
}

func Task1() int {
	var directions = []Point{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	var found = map[Word]int{}
	input := helper.ReadFileByCharacters("./day4/day4.input")
	count := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			for _, dir := range directions {
				curPos := Point{x + (3 * dir.X), y + (3 * dir.Y)}
				if checkLimit(curPos, input) {
					word := ""
					startChar := input[y][x]
					startPos := Point{x + (0 * dir.X), y + (0 * dir.Y)}
					finishPos := Point{x + (3 * dir.X), y + (3 * dir.Y)}
					if startChar == "X" || startChar == "S" {
						for w := 0; w < 4; w++ {
							word += input[y+(w*dir.Y)][x+(w*dir.X)]
						}
						_, exists := found[Word{startPos, finishPos}]
						if (word == "XMAS" || word == "SAMX") && !exists {
							count++
							found[Word{startPos, finishPos}] = 1
							found[Word{finishPos, startPos}] = 1
						}
					}
				}
			}
		}
	}

	return count
}

func Task2() int {
	var directions = []Point{
		{-1, -1}, {1, -1},
		{-1, 1}, {1, 1},
	}

	input := helper.ReadFileByCharacters("./day4/day4.input")
	count := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			safeEdges := 0
			for _, dir := range directions {
				limitPos := Point{x + (1 * dir.X), y + (1 * dir.Y)}
				if checkLimit(limitPos, input) {
					safeEdges++
				}
			}
			if safeEdges == 4 {
				word1 := ""
				word2 := ""
				startChar := input[y][x]
				if startChar == "A" {
					word1 += input[y-1][x-1] + startChar + input[y+1][x+1]
					word2 += input[y+1][x-1] + startChar + input[y-1][x+1]
					if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
						count++
					}
				}
			}
		}
	}

	return count
}

func checkLimit(cur Point, input [][]string) bool {
	if cur.Y < 0 || cur.Y > len(input)-1 || cur.X < 0 || cur.X > len(input[0])-1 {
		return false
	}

	return true
}
