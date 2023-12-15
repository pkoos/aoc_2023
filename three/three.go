package three

import (
	"aoc_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const INPUT_FILE = "three/input"
const SYMBOLS = "*/-@$=%+#&"
const GEAR_SYMBOLS = "*"

type symbol_coordinates struct {
	Symbol string `json:"symbol"`
	Row int `json:"row"`
	Col int `json:"col"`
}

type digit_coordinates struct {
	Digit string `json:"digit"`
	Row int `json:"row"`
	Col int `json:"col"`
}

type part_number struct {
	Row int `json:"row"`
	Start int `json:"start"`
	End int `json:"end"`
	Digit string `json:"digit"`
}

type gear_ratio struct {
	symbol symbol_coordinates
	gears []digit_coordinates
}

func find_symbols(data []string, symbols string) (found []symbol_coordinates) {
	for r_idx, line := range data {
		for c_idx, val := range line {
			s_idx := strings.IndexAny(string(val), symbols)
			if s_idx != -1 {
				found = append(found, symbol_coordinates{string(val), r_idx, c_idx})
			}
		}
	}
	return found
}

func find_adjacent_digits(data []string, symbols []symbol_coordinates)(adjacent_digits []digit_coordinates) {
	for _, symbol := range symbols {
		r_min := symbol.Row - 1
		r_max := symbol.Row + 1
		c_min := symbol.Col - 1
		c_max := symbol.Col + 1
		for i := r_min; i <= r_max; i++ {
			for j := c_min; j <= c_max; j++ {
				if i != symbol.Row || j != symbol.Col {
					if unicode.IsDigit(rune(data[i][j])) {
						adjacent_digits = append(adjacent_digits, digit_coordinates{string(data[i][j]), i, j})
					}
					
					
				}
			}
		}
	}
	return adjacent_digits
}

func find_all_digits(data []string)(all_digits []digit_coordinates) {
	for r_idx, line := range data {
		for c_idx, character := range line {
			if unicode.IsDigit(character) {
				all_digits = append(all_digits, digit_coordinates{string(character), r_idx, c_idx})
			}
		}
	}
	return all_digits
}

func find_part_numbers(data []string)(parts []part_number) {
	sidx := 0
	is_start := false
	for ridx, line := range data {
		for cidx, char := range line {
			if unicode.IsDigit(char) {
				if !is_start {
					is_start = true
					sidx = cidx
				}
			} else {
				if is_start {
					the_digit := line[sidx:cidx]
					parts = append(parts, part_number{ridx, sidx, cidx -1, the_digit})
					is_start = false
				}
			}
		}
		if is_start {
			the_digit := line[sidx:]
			parts = append(parts, part_number{ridx, sidx, len(line) -1, the_digit})
			is_start = false
		}
	}
	
	return parts
}

func find_specific_partnum(digit digit_coordinates, parts []part_number) (found_part part_number) {
	for _, part := range parts {
		if digit.Row == part.Row {
			if digit.Col <= part.End && digit.Col >= part.Start {
				return part
			}
		}
	}
	return found_part
}

func sum_gear_ratios(data [] string) (sum int) {
	return sum
}

func parts_unique(adjacent []digit_coordinates, parts []part_number) (unique_parts []part_number) {
	for _, adj := range adjacent {
		found_part := find_specific_partnum(adj, parts)
		if !slices.Contains(unique_parts, found_part) {
			unique_parts = append(unique_parts, found_part)	
		}
	}

	return unique_parts
}

func sum_part_numbers(data []string) (sum int) {

	symbols := find_symbols(data, SYMBOLS)

	adjacent_digits := find_adjacent_digits(data, symbols)
	parts := find_part_numbers(data)
	found_parts_unique := parts_unique(adjacent_digits, parts)

	for _, part := range found_parts_unique {
		value, _ := strconv.Atoi(part.Digit)
		sum += value
	}

	return sum
}

func Run() {
	fmt.Print("Day 3 - Advent of Code 2023")
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	sum := sum_part_numbers(data)
	fmt.Printf("Day 3, Part 1: %d\n", sum)
	// gear_ratios := sum_gear_ratios(data)
	// fmt.Printf("Day 3, Part 2: %d\n", gear_ratios)
}