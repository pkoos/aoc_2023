package two

import (
	"testing"
	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestGamesEqual(t *testing.T) {
	test_game1 := Game{
		ID: 14,
		Rounds: []Game_round {
			{1, 2, 3},
			{2, 3, 4},
		},
		Is_possible: true,
	}
	test_game2 := test_game1

	if !test_game1.Equals(test_game2) {
		t.Errorf("expected Equals from: %+v and: %+v\n", test_game1, test_game2)
	}
	test_game3 := test_game1
	test_game3.ID = 28
	if test_game1.Equals(test_game3) {
		t.Errorf("expected Not Equals from: %+v and:%+v\n", test_game1, test_game3)
	}
	test_game4 := test_game1
	test_game4.Is_possible = false
	if test_game1.Equals(test_game4) {
		t.Errorf("expected Not Equals from: %+v and:%+v\n", test_game4, test_game4)
	}
}

func TestProcessGameString(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expecteds := []Game {
		{},
		{},
		{},
		{},
		{},
	}

	for idx, line := range data {
		expected := expecteds[idx]
		actual := process_game_string(line)
		if !expected.Equals(actual) {
			t.Errorf("expected: %+v, actual: %+v\n", expected, actual)
		}
	}
}

func TestSumIds(t *testing.T) { 
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 0
	// expected := 8
	actual := sum_ids(data)
	if actual != expected {
		t.Errorf("expected: %d, actual :%d\n", expected, actual)
	}

	t.Skip("Not Yet Implemented") 
}