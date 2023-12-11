package two

import (
	"aoc_2023/utils"
	"fmt"
)

const INPUT_FILE = "two/input"

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

var MAX_ROUND = Game_round{MAX_RED, MAX_GREEN, MAX_BLUE}

type Game struct {
	ID int
	Blue int
	Green int
	Red int
}

type Game_round struct {
	Red int
	Green int
	Blue int
}

func sum_ids(data []string) int {
	for _, line := range data {
		_ = line
	}
	return 0
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	sum := sum_ids(data)
	fmt.Printf("Final Sum: %d\n", sum)
}