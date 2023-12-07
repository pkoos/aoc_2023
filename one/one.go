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
	first, last := number_digits(line)
	digits := first * 10 + last
	return digits
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
	data, err := utils.Slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}

	sum := calibration_sum(data)
	fmt.Printf("len(data): %d, sum: %d\n", len(data), sum)
}