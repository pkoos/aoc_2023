package three

import (
	"encoding/json"
	_ "fmt"
	"os"
	"testing"

	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

// func TestSumGearRatios(t *testing.T) {
// 	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
// 	actual := sum_gear_ratios(data)
// 	_=actual
// 	// t.Skip("Not yet implemented")
// }

func TestFindSpecificPart(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	parts := find_part_numbers(data)
	adjacent := find_adjacent_digits(data, find_symbols(data, SYMBOLS))
	test_data, err := os.ReadFile("test_files/test_found_parts")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	var expecteds []part_number
	err = json.Unmarshal(test_data, &expecteds)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, adj := range adjacent {
		actual := find_specific_partnum(adj, parts)
		expected := expecteds[idx]
		if actual != expected {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestPartsUnique(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	adjacent := find_adjacent_digits(data, find_symbols(data, SYMBOLS))
	parts := find_part_numbers(data)
	
	unique_parts := parts_unique(adjacent, parts)
	_ = unique_parts
	// fmt.Printf("len(adjacent): %d, len(parts): %d, len(unique_parts): %d\n", len(adjacent), len(parts), len(unique_parts))
	// fmt.Printf("unique_parts: %+v\n", unique_parts)
	t.Skip("Not Implemented Yet.")
}

func TestFindPartNumbers(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_parts := find_part_numbers(data)

	test_data, err := os.ReadFile("test_files/test_parts")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	var expected_parts []part_number
	err = json.Unmarshal(test_data, &expected_parts)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, actual := range actual_parts {
		expected := expected_parts[idx]
		if actual != expected {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestFindAdjacentDigits(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_adjacent_digits := find_adjacent_digits(data, find_symbols(data, SYMBOLS))

	test_data, err := os.ReadFile("test_files/test_adjacent_digits")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	var expected_adjacent_digits []digit_coordinates
	err = json.Unmarshal(test_data, &expected_adjacent_digits)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, actual := range actual_adjacent_digits {
		expected := expected_adjacent_digits[idx]
		if expected != actual {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestFindSymbols(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_symbols := find_symbols(data, SYMBOLS)

	test_symbols_data, err := os.ReadFile("test_files/test_symbols")
	if err != nil {
		panic(err)
	}
	var expected_symbols []symbol_coordinates
	err = json.Unmarshal(test_symbols_data, &expected_symbols)
	if err != nil {
		panic(err)
	}

	for idx, actual := range actual_symbols {
		expected := expected_symbols[idx]
		if expected != actual {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestSumPartNumbers(t *testing.T) {
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