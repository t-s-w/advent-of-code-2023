package main

import (
	"fmt"
	"regexp"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

type Junction struct {
	left  string
	right string
}

func Part01(input []string) int {
	directions := input[0]
	locationRegex := regexp.MustCompile("[A-Z]+")
	locationMap := make(map[string]Junction)
	for _, line := range input[2:] {
		r := locationRegex.FindAllString(line, -1)
		if len(r) < 3 {
			continue
		}
		locationMap[r[0]] = Junction{left: r[1], right: r[2]}
	}
	location := "AAA"
	current_steps := 0
	for true {
		for i, r := range directions {
			if string(r) == "L" {
				location = locationMap[location].left
			} else if string(r) == "R" {
				location = locationMap[location].right
			}
			if location == "ZZZ" {
				return current_steps + i + 1
			}
		}
		current_steps += len(directions)
	}
	return 0
}

func Part02(input []string) int {
	return 0
}

func main() {
	input := utils.GetInput(8)
	fmt.Println(Part01(input))
	fmt.Println(Part02(input))
}
