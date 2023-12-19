package five

import (
	"aoc_2023/utils"
	// "fmt"
	"testing"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestParseMapData(t *testing.T) {
	// t.Skip("Not yet implemented")
	data, _ := utils.String_slice_file("test_files/test_map_input")
	actual := parse_map_data(data)
	expected := Farm_map{
		ID: 7,
		Name: "humidity-to-location",
		Destination: 0,
		Mapping: map[int]int{
			11:3, 12:4, 13:5, 14:6,
			5:8, 6:9, 7:10, 8:11, 9:12}}
	
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
	actual := parse_file(data)
	_ = actual
	// fmt.Printf("actual: %+v\n", actual)
	// t.Skip("Not yet implemented")

}