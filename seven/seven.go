package seven

import (
	"aoc_2023/utils"
	"fmt"
)

const INPUT_FILE = "seven/input"

func total_winnings(data []string) (result int) {
	return result
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	part_one := total_winnings(data)
	fmt.Println("==== Day 7 - Part 1 ====")
	fmt.Printf("Answer: %d\n", part_one)

	fmt.Println("==== Day 7 - Part 2 ====")
}