package day6

import (
	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

type Pos struct {
	X int
	Y int
}

var directions = []Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func Task1() int {
	input := helper.ReadFileByCharacters("./day6/day6.input")
	curPos := getStartPos(input)

	curDir := 0
	nextStep := directions[curDir]
	stepsCount := 1
	visitedPos := map[Pos]bool{}

	for {
		visitedPos[curPos] = true
		nextPos := Pos{curPos.X + nextStep.X, curPos.Y + nextStep.Y}
		if isSafe(nextPos, input) {
			nextChar := input[nextPos.Y][nextPos.X]
			if nextChar == "#" {
				if curDir == 3 {
					curDir = 0
				} else {
					curDir++
				}
				nextStep = directions[curDir]
			} else {
				curPos = nextPos
				_, exists := visitedPos[nextPos]
				if !exists {
					stepsCount++
				}
			}
		} else {
			break
		}

	}

	return stepsCount
}

func getStartPos(input [][]string) Pos {
	var curPos Pos
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == "^" {
				curPos = Pos{j, i}
			}
		}
	}
	return curPos
}

func Task2() int {
	input := helper.ReadFileByCharacters("./day6/day6.input")
	curPos := getStartPos(input)

	curDir := 0
	nextStep := directions[curDir]
	visitedPos := map[Pos]bool{}

	for {
		visitedPos[curPos] = true
		nextPos := Pos{curPos.X + nextStep.X, curPos.Y + nextStep.Y}
		if isSafe(nextPos, input) {
			nextChar := input[nextPos.Y][nextPos.X]
			if nextChar == "#" {
				if curDir == 3 {
					curDir = 0
				} else {
					curDir++
				}
				nextStep = directions[curDir]
			} else {
				curPos = nextPos
			}
		} else {
			break
		}
	}

	loopsCount := 0
	doneCount := 0

	for pos := range visitedPos {
		if pos.Y != 6 && pos.X != 4 {
			visitedObs := map[Pos]int{}
			curDir = 0
			nextStep = directions[curDir]
			curPos = getStartPos(input)

			inputCopy := DeepClone(input)
			inputCopy[pos.Y][pos.X] = "0"

			for {
				nextPos := Pos{curPos.X + nextStep.X, curPos.Y + nextStep.Y}
				if isSafe(nextPos, inputCopy) {
					nextChar := inputCopy[nextPos.Y][nextPos.X]
					if nextChar == "#" || nextChar == "0" {
						value, exists := visitedObs[nextPos]

						if !exists {
							visitedObs[nextPos] = 1
						} else {
							visitedObs[nextPos]++
						}

						if value > 4 {
							loopsCount++
							break
						}

						if curDir == 3 {
							curDir = 0
						} else {
							curDir++
						}
						nextStep = directions[curDir]
					} else {
						curPos = nextPos
					}
				} else {
					break
				}
			}
		}
		doneCount++
	}

	return loopsCount
}

func checkForLoop(obs []Pos) bool {
	if len(obs) < 8 {
		return false
	}

	return comparePos(obs[0], obs[4]) &&
		comparePos(obs[1], obs[5]) &&
		comparePos(obs[2], obs[6]) &&
		comparePos(obs[3], obs[7])
}

func comparePos(a, b Pos) bool {
	return a.X == b.X && a.Y == b.Y
}

func isSafe(nextPos Pos, input [][]string) bool {
	if nextPos.Y < 0 || nextPos.X < 0 || nextPos.Y >= len(input) || nextPos.X >= len(input[0]) {
		return false
	}

	return true
}

func DeepClone(input [][]string) [][]string {
	clone := make([][]string, len(input))
	for i := range input {
		clone[i] = append([]string{}, input[i]...) // Create a new slice for each inner slice
	}
	return clone
}
