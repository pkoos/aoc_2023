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

type Symbols []symbol_coordinates

type digit_coordinates struct {
	Digit string `json:"digit"`
	Row int `json:"row"`
	Col int `json:"col"`
}

type Digits []digit_coordinates

type part_number struct {
	Row int `json:"row"`
	Start int `json:"start"`
	End int `json:"end"`
	Digit string `json:"digit"`
}

func (p part_number) InRange(start int, end int) (in_range bool) {
	in_start := p.Start <= end && p.Start >= start
	in_end := p.End <= end && p.End >= start
	in_range = in_start || in_end
	return in_range
}

type Parts []part_number

func (parts Parts) Filter(row int, start int, end int) (filtered Parts) {
	for _, i := range parts {
		if i.InRange(start, end) && i.Row == row {
		// if i.Row == row && i.Start >= start || i.End <= end{
			filtered = append(filtered, i)
		}
	}
	return filtered
}

type gear_ratio struct {
	Symbol symbol_coordinates `json:"symbol"`
	Gears Parts `json:"gears"`
}

func (gear gear_ratio) IsEmpty() (empty bool) {
	empty_symbol := gear.Symbol == symbol_coordinates{}
	empty_gears := gear.Gears == nil

	empty = empty_symbol && empty_gears
	return empty
}

func (gear gear_ratio) CalcGearRatio() (product int) {
	gear_val_one, _ := strconv.Atoi(gear.Gears[0].Digit)
	gear_val_two, _ := strconv.Atoi(gear.Gears[1].Digit)
	product = gear_val_one * gear_val_two
	return product
}

type Gears []gear_ratio


func find_gear(symbol symbol_coordinates, parts Parts) (gear gear_ratio) {
		gear.Symbol = symbol
		r_min := symbol.Row - 1
		r_max := symbol.Row + 1
		c_min := symbol.Col - 1
		c_max := symbol.Col + 1
		// fmt.Printf("symbol: %+v, r_min: %d, r_max: %d, c_min: %d, c_max: %d\n", symbol, r_min, r_max, c_min, c_max)
		for i := r_min; i <= r_max; i++ {
			
			loop_gears := parts.Filter(i, c_min, c_max)
			// fmt.Printf("%d: loop_gears(%d): %+v\n",i , len(loop_gears), loop_gears)
			if len(loop_gears) > 0 {
				// for _, gear := range loop_gears {
				gear.Gears = append(gear.Gears, loop_gears...)
				// }
			}
		}
		if len(gear.Gears) != 2 {
			return gear_ratio{}
		}
	return gear
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
	symbols := find_symbols(data, GEAR_SYMBOLS)
	parts := find_part_numbers(data)
	for _, symbol := range symbols {
		if gear := find_gear(symbol, parts); !gear.IsEmpty() {
			sum += gear.CalcGearRatio()
		}
	}
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
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	sum := sum_part_numbers(data)
	fmt.Println("==== Day 3 - Part 1 ====")
	fmt.Printf("Answer: %d\n", sum)

	gear_ratios := sum_gear_ratios(data)
	fmt.Println("==== Day 3 - Part 2 ====")
	fmt.Printf("Answer: %d\n", gear_ratios)
}