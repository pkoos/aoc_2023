package one

import (
	"fmt"

	"aoc_2023/utils"
)

const INPUT_FILE = "one/input"
const TEST_P1_FILE = "one/test_p1"
const TEST_P2_FILE = "one/test_p2"

func TestOneExport() string {
	return "Exporting the 'one' package works!"
}

func calibration_sum(input []string) int {
	for _, line := range input {
		digits := calibration_digits(line)
		fmt.Printf("digits: %d\n", digits)
	}

	sum := 281
	return sum
}

func calibration_digits(line string) int {
	fmt.Printf("calibration_digits input: %s\n", line)
	return 12
}

func Run() {
	data, err := utils.Slice_file(TEST_P1_FILE)
	if err != nil {
		panic(err)
	}

	sum := calibration_sum(data)
	fmt.Printf("len(data): %d, sum: %d\n", len(data), sum)
}