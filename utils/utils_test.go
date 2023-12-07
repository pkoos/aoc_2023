package utils

import (
	"testing"
)

const TEST_FILE1_PATH = "test_file1"
const TEST_FILE2_PATH = "test_file2"

func TestTheTest(t *testing.T) {
	expected:= "one"
	actual := "one"
	if expected != actual {
		t.Fatalf("expected: '%s', actual: '%s'\n", expected, actual)
	}
}

func TestSlice_file(t *testing.T) {
	expected_test_data1_len := 4

	test_data1, err1 := Slice_file(TEST_FILE1_PATH)
	if err1 != nil {
		t.Fatalf("expected: nil, actual: %s\n", err1)
	}
	if test_data1 == nil {
		t.Fatalf("expected non-nil, actual: %v\n", test_data1)
	}
	
	test_data1_len := len(test_data1)

	if expected_test_data1_len != test_data1_len {
		t.Fatalf("expected: %d, actual: %d\n", expected_test_data1_len, test_data1_len)
	}

	expected_test_data2_len := 7

	test_data2, err2 := Slice_file(TEST_FILE2_PATH)
	if err2 != nil {
		t.Fatalf("expected: nil, actual: %s\n", err2)
	}
	if test_data2 == nil {
		t.Fatalf("expected non-nil, actual: %v\n", test_data2)
	}
	test_data2_len := len(test_data2)

	if expected_test_data2_len != test_data2_len {
		t.Fatalf("expected: %d, actual: %d\n", expected_test_data1_len, test_data1_len)
	}

	test_data_invalid_path, err_path := Slice_file("invalid_file")
	if err_path == nil {
		t.Fatalf("expected: %s, actual: nil\n", err_path)
	}
	if test_data_invalid_path != nil {
		t.Fatalf("expected: nil, actual: %v\n", test_data_invalid_path)
	}
}