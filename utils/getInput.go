package utils

import (
	"fmt"
	"strings"
)

func GetInput(day int) []string {
	body := Get(fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day))
	return strings.Split(string(body), "\n")
}
