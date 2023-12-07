package one

import (
	"testing"
	"aoc_2023/utils"
)

func TestCalibrationSumPartOne(t *testing.T) {
	file_name := "test_p1"
	data, _ := utils.Slice_file(file_name)
	expected_sum := 142
	actual_sum := calibration_sum(data)
	if expected_sum != actual_sum {
		t.Fatalf("expected: %d, actual: %d\n", expected_sum, actual_sum)
	}
}

func TestCalibrationSumPartTwo(t *testing.T) {
	t.Skip("Not yet Implemented")
	// file_name := "test_p2"
	var data []string
	expected_sum := 281
	actual_sum := calibration_sum(data)
	if expected_sum != actual_sum {
		t.Fatalf("expected: %d, actual: %d\n", expected_sum, actual_sum)
	}
}

func TestCalibrationDigitsP1(t *testing.T) {
	line1 := "1abc2"
	expected1 := 12
	actual1 := calibration_digits(line1)
	if expected1 != actual1 {
		t.Errorf("expected: %d, actual: %d\n", expected1, actual1)
	}

	line2 := "pqr3stu8vwx"
	expected2 := 38
	actual2 := calibration_digits(line2)
	if expected2 != actual2 {
		t.Errorf("expected: %d, actual: %d\n", expected2, actual2)
	}

	line3 := "a1b2c3d4e5f"
	expected3 := 15
	actual3 := calibration_digits(line3)
	if expected3 != actual3 {
		t.Errorf("expected: %d, actual: %d\n", expected3, actual3)
	}

	line4 := "treb7uchet"
	expected4 := 77
	actual4 := calibration_digits(line4)
	if expected4 != actual4 {
		t.Errorf("expected: %d, actual: %d\n", expected4, actual4)
	}
}

func TestCalibrationDigitsP2(t *testing.T) {
	t.Skip("Not yet implemented")
	line5 := "two1nine"
	expected5 := 29
	actual5 := calibration_digits(line5)
	if expected5 != actual5 {
		t.Errorf("expected: %d, actual: %d\n", expected5, actual5)
	}

	line6 := "eightwothree"
	expected6 := 83
	actual6 := calibration_digits(line6)
	if expected6 != actual6 {
		t.Errorf("expected: %d, actual: %d\n", expected6, actual6)
	}

	line7 := "abcone2threexyz"
	expected7 := 13
	actual7 := calibration_digits(line7)
	if expected7 != actual7 {
		t.Errorf("expected: %d, actual: %d\n", expected7, actual7)
	}

	line8 := "xtwone3four"
	expected8 := 24
	actual8 := calibration_digits(line8)
	if expected8 != actual8 {
		t.Errorf("expected: %d, actual: %d\n", expected8, actual8)
	}

	line9 := "4nineeightseven2"
	expected9 := 42
	actual9 := calibration_digits(line9)
	if expected9 != actual9 {
		t.Errorf("expected: %d, actual: %d\n", expected9, actual9)
	}

	line10 := "zoneight234"
	expected10 := 52
	actual10 := calibration_digits(line10)
	if expected10 != actual10 {
		t.Errorf("expected: %d, actual: %d\n", expected10, actual10)
	}

	line11 := "7pqrstsixteen"
	expected11 := 76
	actual11 := calibration_digits(line11)
	if expected11 != actual11 {
		t.Errorf("expected: %d, actual: %d\n", expected11, actual11)
	}
}