package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func calibrate(str string) int {
	numberStrings := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	indices := make(map[string][2]int)
	var output struct {
		tensIndex int
		tens      int
		onesIndex int
		ones      int
	}
	output.tensIndex = len(str)
	output.onesIndex = -1
	for key, _ := range numberStrings {
		pattern := regexp.MustCompile(key)
		matches := pattern.FindAllStringIndex(str, -1)
		if len(matches) > 0 {
			indices[key] = [2]int{matches[0][0], matches[len(matches)-1][1]}
		}
	}
	for key, value := range indices {
		if value[0] < output.tensIndex {
			output.tensIndex = value[0]
			output.tens = numberStrings[key]
		}
		if value[1] > output.onesIndex {
			output.onesIndex = value[1]
			output.ones = numberStrings[key]
		}
	}
	return output.tens*10 + output.ones
}

func main() {
	data := utils.Get("https://adventofcode.com/2023/day/1/input")
	lines := strings.Split(string(data), "\n")
	output := 0
	for _, str := range lines {
		output += calibrate(str)
	}
	fmt.Print(output)
}
