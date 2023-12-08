package one

import (
	"fmt"
	"strconv"
	"strings"

	"aoc_2023/utils"
)

const INPUT_FILE = "one/input"
const TEST_P1_FILE = "one/test_p1"
const TEST_P2_FILE = "one/test_p2"

var numbers = []number_data{
	{1, "one"}, {2, "two"}, {3, "three"}, 
	{4, "four"}, {5, "five"}, {6, "six"},
	{7, "seven"}, {8, "eight"}, {9, "nine"},
}

type number_data struct {
	Value int
	Name string
}

type number_index struct {
	Value int
	Index int
}

func TestOneExport() string {
	return "Exporting the 'one' package works!"
}

func calibration_sum(input []string) int {
	var sum int
	for _, line := range input {
		digits := calibration_digits(line)
		sum += digits
	}
	return sum
}

func calibration_digits(line string) int {
	first, last := find_digits(line)
	digits := first * 10 + last
	return digits
}

func find_first(line string) (int) {
	num := first_number(line)
	str := first_string(line)

	if num.Index == -1 || str.Index == -1 {
		// one of the pairs of values doesn't actually have something in it
		if num.Index == -1 {
			return str.Value
		}
		if str.Index == -1 {
			return num.Value
		}
	}

	if num.Index < str.Index {
		return num.Value
	} else {
		return str.Value
	}
}

func find_last(line string) (int) {
	num := last_number(line)
	str := last_string(line)
	
	if num.Index == -1 || str.Index == -1 {
		// one of the pairs of values doesn't actually have something in it
		if num.Index == -1 {
			return str.Value
		}
		if str.Index == -1 {
			return num.Value
		}
	}
	
	if num.Index > str.Index { 
		return num.Value 
	} else { 
		return str.Value 
	}
}

func first_number(line string) number_index {
	index := strings.IndexAny(line, "123456789")
	if index == -1 { // a number wasn't found
		return number_index{-1, -1}
	}
	value, _ := strconv.Atoi(string(line[index]))
	
	return number_index{value, index}
}

func last_number(line string) number_index {
	index := strings.LastIndexAny(line, "123456789")
	if index == -1 {
		return number_index{-1, -1}
	}
	value, _ := strconv.Atoi(string(line[index]))
	
	return number_index{value, index}
}

func first_string(line string) number_index {
	first := number_index{Index: len(line)}
	var index int
	for _, num := range numbers {
		index = strings.Index(line, num.Name)
		// fmt.Printf("first_string line: %s, string: %s, index: %d\n", line, num.Name, index)
		if index != -1 {
			if index < first.Index {
				// fmt.Printf("line: %s, first index: %d, value: %d, new index: %d value: %d\n", line, first.Index, first.Value, index, num.Value)
				first.Index = index
				first.Value = num.Value
			}
		}
	}

	if first.Index == len(line) {
		return number_index {-1, -1}
	}

	return first
}

func last_string(line string) number_index {
	last := number_index{Index:0}
	var index int
	for _, num := range numbers {
		index = strings.LastIndex(line, num.Name)
		if index != -1 {
			if index > last.Index {
				last.Index = index
				last.Value = num.Value
			}
		}
	}

	if last.Index == 0 {
		return number_index{-1, -1}
	}

	return last
}

func find_digits(line string) (int, int) {
	return find_first(line), find_last(line)
}

func number_digits(line string) (int, int) {
	first := strings.IndexAny(line, "123456789")
	last := strings.LastIndexAny(line, "123456789")
	if last == -1 { last = first }
	first_val, err_first := strconv.Atoi(string(line[first]))
	last_val, err_last := strconv.Atoi(string(line[last]))
	if err_first != nil || err_last != nil {
		return -1, -1
	}

	return first_val, last_val
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}

	sum := calibration_sum(data)
	fmt.Printf("len(data): %d, sum: %d\n", len(data), sum)
}