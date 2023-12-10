package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func Part01(input []Race) (result int) {
	result = 1
	for _, race := range input {
		result = result * race.WaysToWin()
	}
	return result
}

func Part02(input []string) (result int) {
	digitRegex := regexp.MustCompile("[0-9]")
	timesText := strings.Join(digitRegex.FindAllString(input[0],-1),"")
	distsText := strings.Join(digitRegex.FindAllString(input[1],-1), "")
	time, err1 := strconv.Atoi(timesText)
	dists, err2 := strconv.Atoi(distsText)
	if err1 == nil && err2 == nil {
		race := Race{Time: time, Distance: dists}
		return race.WaysToWin()
	}
	return result
}

type Race struct {
	Time int
	Distance int
}

func (race *Race) WaysToWin() int {
	discriminator := math.Pow((math.Pow(float64(race.Time), 2) - float64(race.Distance) * 4),0.5) / 2
	centre := float64(race.Time) / 2
	return (int(math.Ceil(centre + discriminator)) - int(math.Floor(centre - discriminator)) - 1)
}

func main() {
	input := utils.GetInput(6)
	numberRegex := regexp.MustCompile("[0-9]+")
	timesText := numberRegex.FindAllString(input[0],-1)
	distsText := numberRegex.FindAllString(input[1],-1)
	races := make([]Race, len(timesText),len(timesText))
	for i := 0; i < len(timesText); i++ {
		t, o1 := strconv.Atoi(timesText[i])
		d, o2 := strconv.Atoi(distsText[i])
		if o1 == nil && o2 == nil {
			races[i] = Race{Time:t, Distance:d}
		}
	}
	fmt.Println(Part01(races))
	fmt.Println(Part02(input))
}