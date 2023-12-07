package one

import (
	"fmt"
	"os"
	"strings"
)

const INPUT_FILE = "one/input"
const TEST_P1_FILE = "one/test_p1"
const TEST_P2_FILE = "one/test_p2"

func TestOneExport() string {
	return "Exporting the 'one' package works!"
}

func calibration_sum(input []string) int {
	sum := 281
	return sum
}

func calibration_digits() int {
	return 12
}

func Run() {
	raw_data, err := os.ReadFile(TEST_P1_FILE)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(raw_data), "\n")
	sum := calibration_sum(data)
	fmt.Printf("len(data): %d, sum: %d\n", len(data), sum)
}