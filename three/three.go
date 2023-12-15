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

type gear_ratio struct {
	symbol symbol_coordinates
	gears []digit_coordinates
}

func find_symbols(data []string, symbols string) (found []symbol_coordinates) {
	fmt.Printf("symbols: '%s'\n", symbols)
	for r_idx, line := range data {
		for c_idx, val := range line {
			s_idx := strings.IndexAny(string(val), symbols)
			// is_symbol := strings.Contains(SYMBOLS, string(val))
			if s_idx != -1 {
				found = append(found, symbol_coordinates{string(val), r_idx, c_idx})
			}
			// if is_symbol {
			// 	found = append(found, symbol_coordinates{string(val), r_idx, c_idx})
			// }
			
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

// func find_gear_ratios(data []string, symbols []symbol_coordinates)(ratios []gear_ratio) {
// 	var working_gear gear_ratio
// 	for _, symbol := range symbols {
// 		r_min := symbol.Row -1
// 		r_max := symbol.Row +1
// 		c_min := symbol.Col -1
// 		c_max := symbol.Col +1
// 		for i := r_min; i <= r_max; i++ {
// 			var gears []gear_ratio
// 			for j := c_min; j <= c_max; j++ {
// 				working_gear = gear_ratio{}
// 				if i != symbol.Row || j != symbol.Col {
// 					if unicode.IsDigit(rune(data[i][j])) {
// 						working_gear.symbol = symbol
// 						working_gear.gears = append(working_gear.gears, )
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return ratios
// }

// func sum_gear_ratios(data [] string) (sum int) {
// 	symbols := find_symbols(data, GEAR_SYMBOLS)
// 	fmt.Printf("found symbols: %+v\n", symbols)

// 	adjacent_digits := find_adjacent_digits(data, symbols)
// 	fmt.Printf("found adjacent digits: %+v\n", adjacent_digits)

// 	parts := find_part_numbers(data)
// 	return sum
// }

func sum_part_numbers(data []string) (sum int) {

	symbols := find_symbols(data, SYMBOLS)

	adjacent_digits := find_adjacent_digits(data, symbols)
	parts := find_part_numbers(data)
	var found_parts_unique []part_number
	for _, adj := range adjacent_digits {
		found_part := find_specific_partnum(adj, parts)
		if !slices.Contains(found_parts_unique, found_part) {
			found_parts_unique = append(found_parts_unique, found_part)
		}
	}

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