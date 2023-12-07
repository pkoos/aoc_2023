package one

import (
	"fmt"
	"os"
)

const INPUT_FILE = "one/input"
const TEST_P1_FILE = "one/test_p1"
const TEST_P2_FILE = "one/test_p2"

func TestOneExport() string {
	return "Exporting the 'one' package works!"
}

func Calibration_sum() int {
	return 281
}

func calibration_digits() int {
	return 12
}

func Run() {
	raw_data, err := os.ReadFile(TEST_P1_FILE)
	if err != nil {
		panic(err)
	}
	data := string(raw_data)
	fmt.Printf("len(data): %d\n", len(data))
}