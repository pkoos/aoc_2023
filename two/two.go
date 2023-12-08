package two

import (
	"aoc_2023/utils"
	"fmt"
)

const INPUT_FILE = "two/input"

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

type Game struct {
	ID int
	Blue int
	Green int
	Red int
}

func a_sum_function_of_some_type(data []string) int {
	return 0
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	sum := a_sum_function_of_some_type(data)
	fmt.Printf("Final Sum: %d\n", sum)
}