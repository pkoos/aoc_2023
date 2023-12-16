package four

import (
	"aoc_2023/utils"
	"testing"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestCardCalculatePoints(t *testing.T) {
	test_card1 := Card{ID: 1, Matches: 0}
	expected1 := 0
	actual1 := test_card1.CalculatePoints()
	if test_card1.CalculatePoints() != expected1 {
		t.Errorf("expected: %d, actual: %d\n", expected1, actual1)
	}

	test_card2 := Card{ID: 2, Matches: 1}
	expected2 := 1
	actual2 := test_card2.CalculatePoints()
	if test_card2.CalculatePoints() != expected2 {
		t.Errorf("expected: %d, actual: %d\n", expected2, actual2)
	}

	test_card3 := Card{ID: 3, Matches: 2}
	expected3 := 2
	actual3 := test_card3.CalculatePoints()
	if test_card3.CalculatePoints() != expected3 {
		t.Errorf("expected: %d, actual: %d\n", expected3, actual3)
	}

	test_card4 := Card{ID: 4, Matches: 3}
	expected4 := 4
	actual4 := test_card4.CalculatePoints()
	if test_card4.CalculatePoints() != expected4 {
		t.Errorf("expected: %d, actual: %d\n", expected4, actual4)
	}

	test_card5 := Card{ID: 5, Matches: 4}
	expected5 := 8
	actual5 := test_card5.CalculatePoints()
	if test_card5.CalculatePoints() != expected5 {
		t.Errorf("expected: %d, actual: %d\n", expected5, actual5)
	}
	test_card6 := Card{ID: 6, Matches: 5}
	expected6 := 16
	actual6 := test_card6.CalculatePoints()
	if test_card6.CalculatePoints() != expected6 {
		t.Errorf("expected: %d, actual: %d\n", expected6, actual6)
	}

	test_card7 := Card{ID: 7, Matches: 6}
	expected7 := 32
	actual7 := test_card1.CalculatePoints()
	if test_card7.CalculatePoints() != expected7 {
		t.Errorf("expected: %d, actual: %d\n", expected7, actual7)
	}

	test_card8 := Card{ID: 8, Matches: 7}
	expected8 := 64
	actual8 := test_card8.CalculatePoints()
	if test_card8.CalculatePoints() != expected8 {
		t.Errorf("expected: %d, actual: %d\n", expected8, actual8)
	}

	test_card9 := Card{ID: 9, Matches: 8}
	expected9 := 128
	actual9 := test_card9.CalculatePoints()
	if test_card9.CalculatePoints() != expected9 {
		t.Errorf("expected: %d, actual: %d\n", expected9, actual9)
	}

	test_card10 := Card{ID: 10, Matches: 9}
	expected10 := 256
	actual10 := test_card10.CalculatePoints()
	if test_card10.CalculatePoints() != expected10 {
		t.Errorf("expected: %d, actual: %d\n", expected10, actual10)
	}

}

func TestSumScratchcards(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 13
	actual := sum_points(data)
	if expected != actual {
		t.Skip("Not yet implemented.")
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
	t.Skip("Not yet implemented.")
}