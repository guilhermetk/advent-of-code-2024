package day8

import (
	"fmt"

	"guilhermetiscoski.com/advent-of-code-2024/helper"
)

type Pos struct {
	X int
	Y int
}

func Task1() int {
	input := helper.ReadFileByCharacters("./day8/day8.input")
	antennas := map[string][]Pos{}
	antinodes := map[Pos]int{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			if input[y][x] != "." {
				char := input[y][x]

				_, exists := antennas[char]
				if !exists {
					antennas[char] = []Pos{}
				}

				antennas[char] = append(antennas[char], Pos{x, y})
			}
		}
	}

	for key := range antennas {
		values := antennas[key]
		for i := 0; i < len(values)-1; i++ {
			for j := i + 1; j < len(values); j++ {
				antennaA := values[i]
				antennaB := values[j]

				xDist := antennaA.X - antennaB.X
				yDist := antennaA.Y - antennaB.Y

				antinode1 := Pos{antennaA.X + xDist, antennaA.Y + yDist}
				antinode2 := Pos{antennaB.X - xDist, antennaB.Y - yDist}

				if isSafe(antinode1, input) {
					antinodes[antinode1] = 1
				}

				if isSafe(antinode2, input) {
					antinodes[antinode2] = 1
				}

			}
		}
	}

	return len(antinodes)
}

func Task2() int {
	input := helper.ReadFileByCharacters("./day8/day8.input")
	antennas := map[string][]Pos{}
	antinodes := map[Pos]int{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input); x++ {
			if input[y][x] != "." {
				char := input[y][x]

				_, exists := antennas[char]
				if !exists {
					antennas[char] = []Pos{}
				}

				antennas[char] = append(antennas[char], Pos{x, y})
			}
		}
	}

	for key := range antennas {
		values := antennas[key]
		for i := 0; i < len(values)-1; i++ {
			for j := i + 1; j < len(values); j++ {
				antennaA := values[i]
				antennaB := values[j]

				antinodes[antennaA] = 1
				antinodes[antennaB] = 1

				xDist := antennaA.X - antennaB.X
				yDist := antennaA.Y - antennaB.Y

				antinode1 := Pos{antennaA.X + xDist, antennaA.Y + yDist}
				for isSafe(antinode1, input) {
					antinodes[antinode1] = 1
					antinode1 = Pos{antinode1.X + xDist, antinode1.Y + yDist}
				}

				antinode2 := Pos{antennaB.X - xDist, antennaB.Y - yDist}
				for isSafe(antinode2, input) {
					antinodes[antinode2] = 1
					antinode2 = Pos{antinode2.X - xDist, antinode2.Y - yDist}
				}

			}
		}
	}

	return len(antinodes)
}

func isSafe(nextPos Pos, input [][]string) bool {
	if nextPos.Y < 0 || nextPos.X < 0 || nextPos.Y >= len(input) || nextPos.X >= len(input[0]) {
		return false
	}

	return true
}

func print2DArray(array [][]string) {
	for _, row := range array {
		fmt.Println(row)
	}
}
