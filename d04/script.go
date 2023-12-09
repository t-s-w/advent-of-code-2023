package main

import (
	"fmt"
	"strings"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func Part01() (score int) {
	input := utils.GetInput(4)
	for _, line := range input {
		winning := make(map[string]bool)
		winners := 0
		state := 0
		a := strings.Split(line, " ")
		for _, b := range a {
			switch state {
			case 0:
				if strings.Contains(b, ":") {
					state = 1
				}
			case 1:
				if b == "|" {
					state = 2
					continue
				}
				if b != "" {
					winning[b] = true
				}
				
			case 2:
				if b != "" && winning[b] {
					winners = winners + 1
				}
			}
		}
		if winners > 0 {
			score = score + intPow(2,winners-1)
		}
	}
	return score
}

func intPow(x,n int) int {
	if n == 0 {
		return 1
	}
	bin := 1
	for (bin*2) <= n {
		bin = bin * 2
	}
	result := x
	n = n - bin
	for bin > 1 {
		bin = bin / 2
		result = result * result
		if n >= bin {
			n = n - bin
			result = result * x
		}
		
	}
	return result
}

func main() {
	fmt.Println(Part01())
}