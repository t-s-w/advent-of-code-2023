package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

type reading struct {
	red   int
	green int
	blue  int
}

type game struct {
	id       int
	readings []reading
}

var numbersRegex = regexp.MustCompile("[0-9]+")

func strToReading(str string) (r reading) {
	cubes := strings.Split(str, ", ")
	for _, cube := range cubes {
		count, err := strconv.Atoi(numbersRegex.FindString(cube))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if strings.Contains(cube, "red") {
			r.red = count
		} else if strings.Contains(cube, "green") {
			r.green = count
		} else if strings.Contains(cube, "blue") {
			r.blue = count
		}
	}
	return r
}

func readGame(line string) (game game) {
	arr := strings.Split(line, ": ")
	id, err := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(arr[0]))
	if err != nil {
		fmt.Print(err)
		return game
	}
	game.id = id
	readingStrs := strings.Split(arr[1], "; ")
	game.readings = make([]reading, len(readingStrs), len(readingStrs))
	for i, readingStr := range readingStrs {
		game.readings[i] = strToReading(readingStr)
	}
	return game
}

func (game *game) Check(criterion reading) bool {
	for _, reading := range game.readings {
		if reading.red > criterion.red || reading.green > criterion.green || reading.blue > criterion.blue {
			return false
		}
	}
	return true
}

func part1() {
	input := utils.Get("https://adventofcode.com/2023/day/2/input")
	lines := strings.Split(string(input), "\n")
	criterion := reading{
		red:   12,
		green: 13,
		blue:  14,
	}
	output := 0
	for _, line := range lines {
		game := readGame(line)
		if game.Check(criterion) {
			output = output + game.id
		}
	}
	fmt.Println(output)
}

func part2() {
	input := utils.Get("https://adventofcode.com/2023/day/2/input")
	lines := strings.Split(string(input), "\n")
	output := 0
	for _, line := range lines {
		game := readGame(line)
		minReading := reading{}
		for _, reading := range game.readings {
			if minReading.red < reading.red {
				minReading.red = reading.red
			}
			if minReading.green < reading.green {
				minReading.green = reading.green
			}
			if minReading.blue < reading.blue {
				minReading.blue = reading.blue
			}
		}
		output = output + minReading.red*minReading.green*minReading.blue
	}
	fmt.Println(output)
}

func main() {
	part2()
}
