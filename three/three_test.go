package three

import (
	"fmt"
	"testing"

	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestSumParts(t *testing.T) {
	
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	for idx, line := range data {
		fmt.Printf("%d: %s\n", idx, line)
	}
	t.Skip("Not yet implemented")
}