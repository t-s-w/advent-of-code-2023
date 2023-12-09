package main

import (
	"fmt"
	"strings"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func Part01() (score int) {
	input := utils.GetInput(4)
	winners := calcWinnersPerScratchcard(input)
	for _, card := range winners {
		if card > 0 {
			score = score + intPow(2,card- 1)
		}
	}
	return score
}

func calcWinnersPerScratchcard(input []string) []int {
	output := make([]int, len(input))
	for i, line := range input {
		if line == "" {
			output = output [:i]
			return output
		}
		state := 0
		winning := make(map[string]bool)
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
					output[i] = output[i] + 1
				}
			}
		}
	}
	return output
}

func Part02() (result int) {
	input := utils.GetInput(4)
	winners := calcWinnersPerScratchcard(input)
	cards := make([]int, len(winners), len(winners))
	for i := 0; i < len(cards); i++ {
		cards[i] = cards[i] + 1
		for j := 1; j <= winners[i]; j++ {
			cards[i+j] = cards[i+j] + cards[i]
		}  
	}
	for _, v := range(cards) {
		result = result + v
	}
	return result
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
	fmt.Println(Part02())
}