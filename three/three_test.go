package three

import (
	"testing"

	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestSumPartNumbers(t *testing.T) {
	// t.Skip("Not yet implemented.")
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	// data, _ := utils.String_slice_file(INPUT_FILE)
	// expected := 0
	expected := 4361
	actual := sum_part_numbers((data))
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
	t.Skip("Not yet implemented")
}