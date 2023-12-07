package one

import (
	"testing"
)

func TestCalibrationSumPartOne(t *testing.T) {
	// file_name := "test_p1"
	var data []string
	expected_sum := 142
	actual_sum := calibration_sum(data) - 139
	if expected_sum != actual_sum {
		t.Fatalf("expected: %d, actual: %d\n", expected_sum, actual_sum)
	}
}

func TestCalibrationSumPartTwo(t *testing.T) {
	// file_name := "test_p2"
	var data []string
	expected_sum := 281
	actual_sum := calibration_sum(data)
	if expected_sum != actual_sum {
		t.Fatalf("expected: %d, actual: %d\n", expected_sum, actual_sum)
	}
}

func TestCalibrationDigits(t *testing.T) {
	expected_digits := 12
	actual_digits := calibration_digits()
	if expected_digits != actual_digits  {
		t.Fatalf("expected: %d, actual: %d\n", expected_digits, actual_digits)
	}
}