package three

import (
	"encoding/json"
	"os"
	"testing"

	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestFindGear(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	symbols := find_symbols(data, GEAR_SYMBOLS)
	parts := find_part_numbers(data)

	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_found_gears"); err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	var expected_gears Gears
	if err = json.Unmarshal(test_data, &expected_gears); err != nil {
		t.Fatalf("Error:%s\n", err)
	}

	for idx, symbol := range symbols {
		expected := expected_gears[idx]
		actual := find_gear(symbol, parts)
		if expected.Symbol != actual.Symbol {
			t.Errorf("Symbol expected: %+v, actual: %+v\n", expected.Symbol, actual.Symbol)
		}		
		for idx, actual_part := range actual.Gears {
			expected_part := expected.Gears[idx]
			if actual_part != expected_part {
				t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected_part, actual_part)
			}
		}
	}
}
func TestGearAdjacents(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_adjacents := find_adjacent_digits(data, find_symbols(data, GEAR_SYMBOLS))

	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_gear_adjacents"); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	var expected_adjacents []digit_coordinates
	if err = json.Unmarshal(test_data, &expected_adjacents); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, actual := range actual_adjacents {
		if expected := expected_adjacents[idx]; actual != expected {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}


func TestGearSymbols(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_symbols := find_symbols(data, GEAR_SYMBOLS)

	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_gear_symbols"); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	var expected_symbols []symbol_coordinates
	if err = json.Unmarshal(test_data, &expected_symbols); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, actual := range actual_symbols {
		if expected := expected_symbols[idx]; actual != expected {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}

}

func TestSumGearRatios(t *testing.T) {
	
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 467835

	if actual := sum_gear_ratios(data); expected != actual {
		t.Errorf("expected :%d, actual :%d\n", expected, actual)
	}
}

func TestFindSpecificPart(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	parts := find_part_numbers(data)
	adjacent := find_adjacent_digits(data, find_symbols(data, SYMBOLS))
	
	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_found_parts");err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	
	var expecteds []part_number
	if err = json.Unmarshal(test_data, &expecteds);err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, adj := range adjacent {
		expected := expecteds[idx]
		
		if actual := find_specific_partnum(adj, parts); actual != expected {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestPartsUnique(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	adjacent := find_adjacent_digits(data, find_symbols(data, SYMBOLS))
	actual_unique_parts := parts_unique(adjacent, find_part_numbers(data))
	
	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_unique_parts"); err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	
	var expected_unique_parts []part_number
	if err = json.Unmarshal(test_data, &expected_unique_parts);err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	for idx, actual := range actual_unique_parts {
		if expected := expected_unique_parts[idx]; actual != expected {
			t.Errorf("%d: expected %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestFindPartNumbers(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_parts := find_part_numbers(data)

	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_parts"); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	var expected_parts []part_number
	if err = json.Unmarshal(test_data, &expected_parts); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, actual := range actual_parts {
		if expected := expected_parts[idx]; actual != expected {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestFindAdjacentDigits(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_adjacent_digits := find_adjacent_digits(data, find_symbols(data, SYMBOLS))

	var test_data []byte
	var err error
	if test_data, err = os.ReadFile("test_files/test_adjacent_digits"); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	var expected_adjacent_digits []digit_coordinates
	if err = json.Unmarshal(test_data, &expected_adjacent_digits); err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	for idx, actual := range actual_adjacent_digits {
		if expected := expected_adjacent_digits[idx]; expected != actual {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestFindSymbols(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual_symbols := find_symbols(data, SYMBOLS)

	var test_symbols_data []byte
	var err error
	
	if test_symbols_data, err = os.ReadFile("test_files/test_symbols"); err != nil {
		panic(err)
	}
	var expected_symbols []symbol_coordinates
	
	if err = json.Unmarshal(test_symbols_data, &expected_symbols); err != nil {
		panic(err)
	}

	for idx, actual := range actual_symbols {
		if expected := expected_symbols[idx]; expected != actual {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestSumPartNumbers(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 4361

	if actual := sum_part_numbers(data); expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
}
