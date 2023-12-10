package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/t-s-w/advent-of-code-2023/utils"
)

func Part01(input []Race) (result int) {
	result = 1
	for _, race := range input {
		discriminator := math.Pow((math.Pow(float64(race.Time), 2) - float64(race.Distance) * 4),0.5) / 2
		centre := float64(race.Time) / 2
		result = result * (int(math.Ceil(centre + discriminator)) - int(math.Floor(centre - discriminator)) - 1)
	}
	return result
}

func Part02(input []string) (result int) {
	return result
}

type Race struct {
	Time int
	Distance int
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