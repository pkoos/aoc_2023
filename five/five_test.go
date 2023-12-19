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
		Next: 0,
		Destinations: []int{11, 5},
		Sources: []int{3, 8},
		Ranges: []int{4, 5},
	}
	
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

func TestLowestLocationRange(t *testing.T) {
	data, _ := utils.String_slice_file("input")
	actual := find_lowest_location_seedrange(data)
	expected := 47909639
	if actual != expected {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
}


func TestFindLowestLocationSeedrange(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual := find_lowest_location_seedrange(data)
	expected := 46
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
}

func TestFindLowestLocation(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual := find_lowest_location(data)
	expected := 35
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}


}