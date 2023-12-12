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

type symbol_coordinates struct {
	Symbol string
	Row int
	Col int
}

type digit_coordinates struct {
	Digit string
	Row int
	Col int
}

type part_number struct {
	Row int
	Start int
	End int
	Digit string
}

func find_symbols(data []string) (found []symbol_coordinates) {
	for r_idx, line := range data {
		for c_idx, val := range line {
			is_symbol := strings.Contains(SYMBOLS, string(val))
			if is_symbol {
				found = append(found, symbol_coordinates{string(val), r_idx, c_idx})
			}
			
			// fmt.Printf("[%d][%d]: '%s' symbol? %t\n",
			// 	r_idx, c_idx, string(val), is_symbol)
		}
	}
	return found
}

func find_adjacent_digits(data []string, symbols []symbol_coordinates)(adjacent_digits []digit_coordinates) {
	for _, symbol := range symbols {
		// fmt.Printf("%+v, row: %d, col: %d\n", symbol, symbol.Row, symbol.Col)
		r_min := symbol.Row - 1
		r_max := symbol.Row + 1
		c_min := symbol.Col - 1
		c_max := symbol.Col + 1
		for i := r_min; i <= r_max; i++ {
			for j := c_min; j <= c_max; j++ {
				if i != symbol.Row || j != symbol.Col {
					if unicode.IsDigit(rune(data[i][j])) {
						// fmt.Printf("[%d][%d]: %s, digit? %t\n", i, j, string(data[i][j]), unicode.IsDigit(rune(data[i][j])))
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
				// fmt.Printf("[%d][%d]: '%s'\n", r_idx, c_idx, string(character))
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
					// fmt.Printf("line[%d][%d:%d]: %s\n", ridx, sidx, cidx -1, the_digit)
					parts = append(parts, part_number{ridx, sidx, cidx -1, the_digit})
					is_start = false
				}
			}
		}
		if is_start {
			the_digit := line[sidx:]
			// fmt.Printf("right before next row: line[%d][%d:%d]: %s\n", ridx, sidx, len(line) -1, the_digit)
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
				// fmt.Printf("%d: digit: %+v\n", idx, digit)
				// fmt.Printf("%d: found_part: %+v\n", idx, part)
				return part
			}
		}
	}
	return found_part
}

func sum_part_numbers(data []string) (sum int) {
	// var found_symbols []symbol_coordinates

	symbols := find_symbols(data)
	// fmt.Printf("found symbols: %+v\n", symbols)

	adjacent_digits := find_adjacent_digits(data, symbols)
	// fmt.Printf("adjacent digits: %+v\n", adjacent_digits)

	// all_digits := find_all_digits(data)
	// fmt.Printf("all digits: %+v\n", all_digits)

	parts := find_part_numbers(data)
	// fmt.Printf("parts: %+v\n", parts)

	// var found_parts []part_number
	var found_parts_unique []part_number
	for _, adj := range adjacent_digits {
		found_part := find_specific_partnum(adj, parts)
		if !slices.Contains(found_parts_unique, found_part) {
			found_parts_unique = append(found_parts_unique, found_part)
		}
		// found_parts = append(found_parts, found_part)
	}

	for _, part := range found_parts_unique {
		value, _ := strconv.Atoi(part.Digit)
		sum += value
	}
	// fmt.Printf("found parts: %+v\n", found_parts)
	// fmt.Printf("\nunique found parts: %+v\n", found_parts_unique)
	
	// fmt.Printf("========== THE SUM ==========\n========== %d ==========", sum)
	return sum
}

func Run() {
	fmt.Print("Day 3 - Advent of Code 2023")
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	sum := sum_part_numbers(data)
	fmt.Printf("Final Day 2, Part 1: %d\n", sum)
}