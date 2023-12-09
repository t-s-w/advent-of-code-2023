package main

import (
	"fmt"
	"strconv"
	"strings"

	maaps "github.com/t-s-w/advent-of-code-2023/d05/maps"
	"github.com/t-s-w/advent-of-code-2023/utils"
)

type GameData struct {
	Seeds []int
	SeedSoilMap maaps.ResourceMap
	SoilFertMap maaps.ResourceMap
	FertWaterMap maaps.ResourceMap
	WaterLightMap maaps.ResourceMap
	LightTempMap maaps.ResourceMap
	TempHumidMap	maaps.ResourceMap
	HumidLocMap maaps.ResourceMap
}

func (g *GameData) SeedLocMap(seed int) (loc int){ 
	loc = g.SeedSoilMap.Map(seed)
	loc = g.SoilFertMap.Map(loc)
	loc = g.FertWaterMap.Map(loc)
	loc = g.WaterLightMap.Map(loc)
	loc = g.LightTempMap.Map(loc)
	loc = g.TempHumidMap.Map(loc)
	loc = g.HumidLocMap.Map(loc)
	return loc
}

func Part01(game *GameData) int {
	g := *game
	output := int(^uint(0) >> 1)
	for _, seed := range g.Seeds {
		loc := g.SeedLocMap(seed)
		if loc < output {
			output = loc
		}
	}
	return output
}

func Part02(game *GameData) (result int) {
	g := *game
	result = int(^uint(0) >> 1) 
	for j := 0; j < len(game.Seeds) / 2; j++ {
		start := game.Seeds[j*2]
		l := game.Seeds[j*2 + 1]
		for seed := start; seed < start + l; seed++ {
			loc := g.SeedLocMap(seed)
			if loc < result {
				result = loc
			}
		}
	}
	return result
}



func readData(input []string) GameData {
	var data GameData
	state := 0
	var currentMap *maaps.ResourceMap
	for _, line := range input {
		switch line {
		case "":
			continue
		case "seed-to-soil map:":
			state = 1
			continue
		case "soil-to-fertilizer map:":
			state = 2
			continue
		case "fertilizer-to-water map:":
			state = 3
			continue
		case "water-to-light map:":
			state = 4
			continue
		case "light-to-temperature map:":
			state = 5
			continue
		case "temperature-to-humidity map:":
			state = 6
			continue
		case "humidity-to-location map:":
			state = 7
			continue
		}
		switch state {
		case 0:
			if strings.Contains(line, "seeds: ") {
				seeds := strings.Split(line, "seeds: ")[1]
				for _, seed := range strings.Split(seeds, " ") {
					n, o := strconv.Atoi(seed)
					if o == nil {
						data.Seeds = append(data.Seeds, n)
					}
				}
			}
		case 1:
			currentMap = &data.SeedSoilMap
		case 2:
			currentMap = &data.SoilFertMap
		case 3:
			currentMap = &data.FertWaterMap
		case 4:
			currentMap = &data.WaterLightMap
		case 5:
			currentMap = &data.LightTempMap
		case 6:
			currentMap = &data.TempHumidMap
		case 7:
			currentMap = &data.HumidLocMap

		}
		if state > 0 {
			nums := strings.Split(line, " ")
			i := 0
			var d maaps.MapData
			for _, num := range nums {
				n, o := strconv.Atoi(num)
				if o == nil {
					d[i] = n
					i++
				}
			}
			currentMap.Data = append(currentMap.Data,d)
		}
	}
	return data
}

func main() {
	input := utils.GetInput(5)
	gamedata := readData(input)
	fmt.Println(Part01(&gamedata))
	fmt.Println(Part02(&gamedata))
}