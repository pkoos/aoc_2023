package five

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"aoc_2023/utils"
)

type Farm struct {
	Seeds []int
}

const INPUT = "five/input"

func parse_seeds(line string) (seeds []int) {
	var seed_str string
	part_of_digit := false
	for idx, char := range line {
		if unicode.IsDigit(char) {
			seed_str += string(char)
			part_of_digit = true
			if idx + 1 == len(line) {
				val, _ := strconv.Atoi(seed_str)
				seeds = append(seeds, val)
			}
		} else {
			if part_of_digit {
				val, _ := strconv.Atoi(seed_str)
				seeds = append(seeds, val)
				part_of_digit = false
				seed_str = ""
			}
		}
	}
	return seeds
}

func parse_file(data [] string) (farm Farm) {
	for idx, line := range data {
		fmt.Printf("%d: ", idx)
		if strings.Contains(line, "seeds:") {
			farm.Seeds = parse_seeds(line)
		}
	}
	return farm
}

func find_lowest_location(data []string) (location int) {
	farm := parse_file(data)
	_= farm
	return location
}

func Run() {
	data, err := utils.String_slice_file("five/input")
	if err != nil {
		panic(err)
	}

	min_loc_part_one := find_lowest_location(data)
	fmt.Println("==== Day 5 - Part 1 ====")
	fmt.Printf("Answer: %d\n", min_loc_part_one)
	
	fmt.Println("==== Day 5 - Part 2 ====")
}