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
		actual := parse_line(line, true)
		expected := expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 1 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}

	expecteds = [][]int{
		{ 71530 },
		{ 940200 },
	}
	for idx, line := range data {
		actual := parse_line(line, false)
		expected := expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestRacePossibilities(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	races := parse_races(data, true)
	expecteds := []int {4, 8, 9}
	for idx, race := range races {
		expected := expecteds[idx]
		actual := race.Possibilities()
		if expected != actual {
			t.Errorf("Part 1 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}

	races = parse_races(data, false)
	expecteds = []int {71503}
	for idx, race := range races {
		expected := expecteds[idx]
		actual := race.Possibilities()
		if expected != actual {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestParseRaces(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	races_p1 := parse_races(data, true)
	expecteds_p1 := Races{
		{7, 9}, {15, 40}, {30, 200},
	}
	for idx, actual := range races_p1 {
		expected := expecteds_p1[idx]
		if !actual.Equals(expected) {
			t.Errorf("Part 1 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
	races_p2 := parse_races(data, false)
	expecteds_p2 := Races{
		{71530, 940200},
	}
	for idx, actual := range races_p2 {
		expected := expecteds_p2[idx]
		if !actual.Equals(expected) {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestProductWinningPossibilities(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_p1 := 4 * 8 * 9 // 288, r1 * r2 * r3
	actual_p1 := product_winning_possibilities(data, true)
	// t.Skip("Not yet implemented")
	if actual_p1 != expected_p1 {
		t.Errorf("Part 1 - expected: %d, actual: %d\n", expected_p1, actual_p1)
	}
	expected_p2 := 71503
	actual_p2 := product_winning_possibilities(data, false)
	if actual_p2 != expected_p2 {
		t.Errorf("Part 2 - expected: %d, actual: %d\n", expected_p2, actual_p2)
	}
}