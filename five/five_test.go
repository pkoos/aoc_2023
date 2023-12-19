package five

import (
	"aoc_2023/utils"
	"testing"
)

const TEST_INPUT_FILE = "test_files/test_input"


func TestTraverseMapping(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	farm := parse_file(data)
	expected_locations := []int{82, 43, 86, 35}
	for idx, expected := range expected_locations {
		seed := farm.Seeds[idx]
		actual := farm.TraverseMapping(seed).Location
		if expected != actual {
			t.Errorf("%d: seed: %d, expected: %d, actual: %d\n", idx, seed, expected, actual)
		}
	}
}

func TestParseMapData(t *testing.T) {
	// t.Skip("Not yet implemented")
	data, _ := utils.String_slice_file("test_files/test_map_input")
	actual := parse_map_data(data)
	expected := Farm_map{
		ID: 7,
		Name: "humidity-to-location",
		Destination: 0,
		Mapping: map[int]int{
			3:11, 4:12, 5:13, 6:14,
			8:5, 9:6, 10:7, 11:8, 12:9}}
	
	if !actual.Equals(expected) {
		t.Errorf("expected: %+v, actual: %+v\n", expected, actual)
	}
}

func TestParseSeeds(t *testing.T) {
	actual := parse_seeds("seeds: 79 14 55 13")
	var expected = []int{ 79, 14, 55, 13}
	for idx, act := range actual {
		exp := expected[idx]
		if exp != act {
			t.Errorf("%d: expected: %d, actual: %d\n", idx, exp, act)
		}
	}
}

func TestParseFile(t *testing.T) {
	t.Skip("Not yet implemented")
}

func TestFindLowestLocation(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual := find_lowest_location(data)
	expected := 35
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}


}