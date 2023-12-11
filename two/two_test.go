package two

import (
	"testing"
	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestSumIds(t *testing.T) { 
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 0
	// expected := 8
	actual := sum_ids(data)
	if actual != expected {
		t.Errorf("expected: %d, actual :%d\n", expected, actual)
	}

	t.Skip("Not Yet Implemented") 
}