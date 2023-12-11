package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func main() {
	input := utils.GetInput(7)
	fmt.Print(Part01(input))
	fmt.Print(Part02(input))
}

type Hand struct {
	Cards [5]int
	Bid int
}

func (h *Hand) Type() int {
	counts := make(map[int]int)
	for _, card := range h.Cards {
		counts[card] = counts[card] + 1
	}
	types := make(map[int]int)
	for _, count := range counts {
		types[count] += 1
	}
	if types[5] > 0 {
		return 7
	} else if types[4] > 0{
		return 6
	} else if types[3] > 0 && types[2] > 0 {
		return 5
	} else if types[3] > 0 {
		return 4
	} else if types[2] >= 2 {
		return 3
	} else if types[2] > 0 {
		return 2
	} else {
		return 1
	}
}

var CardRank = map[string]int {
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type SetOfHands []Hand

func (s SetOfHands) Len() int { return len(s)}
func (s SetOfHands) Swap(i,j int) {s[i], s[j] = s[j], s[i]}
func (s SetOfHands) Less(i, j int) bool {
	if s[i].Type() != s[j].Type() {
		return s[i].Type() < s[j].Type()
	} else if s[i].Cards[0] != s[j].Cards[0] {
		return s[i].Cards[0] < s[j].Cards[0]
	} else if s[i].Cards[1] != s[j].Cards[1] {
		return s[i].Cards[1] < s[j].Cards[1]
	} else if s[i].Cards[2] != s[j].Cards[2] {
		return s[i].Cards[2] < s[j].Cards[2]
	} else if s[i].Cards[3] != s[j].Cards[3] {
		return s[i].Cards[3] < s[j].Cards[3]
	} else if s[i].Cards[4] != s[j].Cards[4] {
		return s[i].Cards[4] < s[j].Cards[4]
	}
	return false
} 

func Part01 (input []string) (result int) {
	var set = make(SetOfHands,len(input)-1)
	for i, line := range input {
		split := strings.Split(line, " ")
		if len(split) < 2 {
			continue
		}
		for j, card := range split[0] {
			set[i].Cards[j] = CardRank[string(card)]
		}
		bid, _ := strconv.Atoi(split[1])
		set[i].Bid = bid
	}
	sort.Sort(set)
	for i, hand := range set{
		result = result + hand.Bid * (i+1)
	} 
	return result
}

func Part02 (input []string) (result int) {
return result 
}