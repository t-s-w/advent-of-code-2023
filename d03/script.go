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
}
