package six

import (
	"aoc_2023/utils"
	"reflect"
	"testing"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestParseLine(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expecteds := [][]int{
		{7, 15, 30},
		{9, 40, 200},
	}
	for idx, line := range data {
		actual := parse_line(line)
		expected := expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestRacePossibilities(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	races := parse_races(data)
	expecteds := []int {4, 8, 9}
	for idx, race := range races {
		expected := expecteds[idx]
		actual := race.Possibilities()
		if expected != actual {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestParseRaces(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	races := parse_races(data)
	expecteds := Races{
		{7, 9}, {15, 40}, {30, 200},
	}
	for idx, actual := range races {
		expected := expecteds[idx]
		if !actual.Equals(expected) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestProductWinningPossibilities(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 4 * 8 * 9 // 288, r1 * r2 * r3
	actual := product_winning_possibilities(data)
	// t.Skip("Not yet implemented")
	if actual != expected {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
}