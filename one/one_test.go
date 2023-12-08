package one

import (
	"testing"

	"aoc_2023/utils"
)

func TestCalibrationSumP1(t *testing.T) { t.Skip("Not yet implemented.") }

func TestCalibrationDigitsP1(t *testing.T) { t.Skip("Not yet implemented.") }

func TestCalibrationDigitsP2(t *testing.T) { t.Skip("Not yet implemented.") }

func TestFindDigits(t *testing.T) {
	data, _ := utils.String_slice_file("test_input")
	expecteds, _ := utils.Int_slice_file("test_last")

	if len(data) != len(expecteds) {
		t.Fatalf("test is broken, test files len(data): %d, len(expecteds): %d\n", 
			len(data), len(expecteds))
	}

	for idx, line := range data {
		value := find_last(line)
		expected := expecteds[idx]

		if value != expected {
			t.Errorf("line: %s, expected: %d, actual: %d\n", line, expected, value)
		}
	}
}

func TestFindFirst(t *testing.T) {
	data, _ := utils.String_slice_file("test_input")
	expecteds, _ := utils.Int_slice_file("test_first")
	if len(data) != len(expecteds) {
		t.Fatalf("test is broken, test files len(data): %d, len(expecteds): %d\n", 
			len(data), len(expecteds))
	}

	for idx, line := range data {
		value := find_first(line)
		expected := expecteds[idx]
		if value != expected {
			t.Errorf("line: %s, expected: %d, actual: %d\n", line, expected, value)
		}
	}
}

func TestFirstNumber(t *testing.T) {
	data, _ := utils.String_slice_file("test_input")
	expected_idxs, _ := utils.Int_slice_file("test_first_number_index")
	expected_vals, _ := utils.Int_slice_file("test_first_number")
	
	if len(data) != len(expected_idxs) || len(data) != len(expected_vals) {
		t.Fatalf("test is broken, test files len(data): %d, len(idxs): %d, len(vals): %d\n", 
			len(data), len(expected_idxs), len(expected_vals))
	}

	for idx, line := range data {
		first := first_number(line)
		expected_idx := expected_idxs[idx]
		expected_val := expected_vals[idx]

		if first.Index != expected_idx || first.Value != expected_val {
			t.Errorf("line: %s, index (e:%d, a:%d) value: (e: %d, a:%d)\n", 
				line, expected_idx, first.Index, expected_val, first.Value)
		}
	}
}

func TestFirstString(t *testing.T) {
	// t.Skip("Not yet implemented.")
	data, _ := utils.String_slice_file("test_input")
	expected_idxs, _ := utils.Int_slice_file("test_first_string_index")
	expected_vals, _ := utils.Int_slice_file("test_first_string")

	if len(data) != len(expected_idxs) || len(data) != len(expected_vals) {
		t.Fatalf("test is broken, test files len(data): %d, len(idxs): %d, len(vals): %d\n", 
			len(data), len(expected_idxs), len(expected_vals))
	}

	for idx, line := range data {
		first := first_string(line)
		expected_idx := expected_idxs[idx]
		expected_val := expected_vals[idx]

		if first.Index != expected_idx || first.Value != expected_val {
			t.Errorf("line: %s, index (e:%d, a:%d) value: (e: %d, a:%d)\n", 
				line, expected_idx, first.Index, expected_val, first.Value)
		}
	}
}
func TestFindLast(t *testing.T) {
	data, _ := utils.String_slice_file("test_input")
	expecteds, _ := utils.Int_slice_file("test_last")	
	if len(data) != len(expecteds) {
		t.Fatalf("test is broken, test files len(data): %d, len(expecteds): %d\n", 
			len(data), len(expecteds))
	}

	for idx, line := range data {
		value := find_last(line)
		expected := expecteds[idx]
		if value != expected {
			t.Errorf("line: %s, expected: %d, actual: %d\n", line, expected, value)
		}
	}
}

func TestLastNumber(t *testing.T) {
	data, _ := utils.String_slice_file("test_input")
	expected_idxs, _ := utils.Int_slice_file("test_last_number_index")
	expected_vals, _ := utils.Int_slice_file("test_last_number")
	if len(data) != len(expected_idxs) || len(data) != len(expected_vals) {
		t.Fatalf("test is broken, test files len(data): %d, len(idxs): %d, len(vals): %d\n", 
			len(data), len(expected_idxs), len(expected_vals))
	}

	for idx, line := range data {
		last := last_number(line)
		
		expected_idx := expected_idxs[idx]
		expected_val := expected_vals[idx]
		if last.Index != expected_idx || last.Value != expected_val {
			t.Errorf("line: %s, index (e:%d, a:%d) value: (e: %d, a:%d)\n", 
				line, expected_idx, last.Index, expected_val, last.Value)
		}
	}
}

func TestLastString(t *testing.T) {
	data, _ := utils.String_slice_file("test_input")
	expected_idxs, _ := utils.Int_slice_file("test_last_string_index")
	expected_vals, _ := utils.Int_slice_file("test_last_string")

	if len(data) != len(expected_idxs) || len(data) != len(expected_vals) {
		t.Fatalf("test is broken, test files len(data): %d, len(idxs): %d, len(vals): %d\n", 
			len(data), len(expected_idxs), len(expected_vals))
	}

	for idx, line := range data {
		last := last_string(line)
		expected_idx := expected_idxs[idx]
		expected_val := expected_vals[idx]

		if last.Index != expected_idx || last.Value != expected_val {
			t.Errorf("%d: line: %s, index(e:%d, a:%d), value: (e:%d, a:%d)\n", 
				idx, line, expected_idx, last.Index, expected_val, last.Value)
		}
	}
}
