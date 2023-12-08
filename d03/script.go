package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func Part1() {
	output := 0
	lines := utils.GetInput(3)
	for i, line := range lines {
		var prevLine string
		if i > 0 {
			prevLine = lines[i-1]
		}
		var nextLine string
		if i < len(line)-1 {
			nextLine = lines[i+1]
		}
		output = output + GetPartNumbers(line, prevLine, nextLine)
	}
	fmt.Println(output)
}

type CharSet map[string]bool

var digitSet = CharSet{
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
	"0": true,
}

func (c CharSet) has(str string) bool {
	return c[str]
}

type Schematic []string

func (schematic Schematic) completeNumber(x, y int) (int, bool) {
	line := schematic[y]
	if !digitSet.has(string(line[x])) {
		return 0, false
	}
	start := x
	end := x + 1
	for start > 0 && digitSet.has(string(line[start-1])) {
		start = start - 1
	}
	for end < len(line) && digitSet.has(string(line[end])) {
		end = end + 1
	}
	result, err := strconv.Atoi(string(line[start:end]))
	if err != nil {
		fmt.Println(err)
	}
	return result, true
}

func (schematic Schematic) appendAdjacentNumber(a []int, x, y int) []int {
	if x < 0 || x > len(schematic[0]) || y < 0 || y > len(schematic) {
		return a
	}
	if num, ok := schematic.completeNumber(x, y); ok {
		a = append(a, num)
	}
	return a
}

func Part2() (output int) {
	lines := utils.GetInput(3)
	schematic := Schematic(lines)
	for y, line := range schematic {
		for x, b := range line {
			if string(b) == "*" {
				adjacents := make([]int, 0, 6)
				adjacents = schematic.appendAdjacentNumber(adjacents, x-1, y)
				adjacents = schematic.appendAdjacentNumber(adjacents, x+1, y)
				if y > 0 {
					if digitSet.has(string(schematic[y-1][x])) {
						adjacents = schematic.appendAdjacentNumber(adjacents, x, y-1)
					} else {
						adjacents = schematic.appendAdjacentNumber(adjacents, x-1, y-1)
						adjacents = schematic.appendAdjacentNumber(adjacents, x+1, y-1)
					}
				}
				if y < len(schematic) {
					if digitSet.has(string(schematic[y+1][x])) {
						adjacents = schematic.appendAdjacentNumber(adjacents, x, y+1)
					} else {
						adjacents = schematic.appendAdjacentNumber(adjacents, x-1, y+1)
						adjacents = schematic.appendAdjacentNumber(adjacents, x+1, y+1)
					}
				}
				if len(adjacents) == 2 {
					output = output + adjacents[0]*adjacents[1]
				}
			}
		}
	}
	return output
}

func findSymbol(line string, start, end int) bool {
	symbolRegex := regexp.MustCompile("[^\\.0-9]")
	if start < 0 {
		start = 0
	}
	if end > len(line) {
		end = len(line)
	}
	return symbolRegex.FindStringIndex(line[start:end]) != nil
}

func GetPartNumbers(line, prevLine, nextLine string) (value int) {
	numberRegex := regexp.MustCompile("[0-9]+")
	numIndices := numberRegex.FindAllStringIndex(line, -1)
	for _, startEnd := range numIndices {
		num, err := strconv.Atoi(line[startEnd[0]:startEnd[1]])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		if findSymbol(line, startEnd[0]-1, startEnd[1]+1) ||
			(prevLine != "" && findSymbol(prevLine, startEnd[0]-1, startEnd[1]+1)) ||
			(nextLine != "" && findSymbol(nextLine, startEnd[0]-1, startEnd[1]+1)) {
			value = value + num
			continue
		}
	}
	return value
}

func main() {
	Part1()
	fmt.Print(Part2())
}
