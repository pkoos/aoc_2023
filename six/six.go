package six

import (
	"aoc_2023/utils"
	"fmt"
)

type Race struct {
	Time int `json:"time"`
	Distance int `json:"distance"`
}

func (r Race) Possibilities() (result int) {
	return result
}

type Races []Race

func parse_races(data []string) (races Races) {
	return races
}

func product_winning_possibilities(data []string) (result int) {
	races := parse_races(data)
	for _, race := range races {
		result += race.Possibilities()
	}
	return result
}

func Run() {
	data, err := utils.String_slice_file("six/input")
	if err != nil {
		panic(err)
	}
	fmt.Println("==== Day 6 - Part 1 ====")
	part_one := product_winning_possibilities(data)
	fmt.Printf("Answer: %d\n", part_one)
	fmt.Println("==== Day 6 - Part 2 ====")
}