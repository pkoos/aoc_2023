package two

import (
	"testing"
	"aoc_2023/utils"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestGameIdInt(t *testing.T) {
	data, _ := utils.String_slice_file("test_files/test_gamestr")
	expecteds, _ := utils.Int_slice_file("test_files/test_ids")
	for idx, line := range data {
		actual := game_id_int(line)
		expected := expecteds[idx]
		if actual != expected {
			t.Errorf("%d: '%s' expected: %d, actual: %d\n", idx, line, expected, actual)
		}
	}
}
func TestSplitGameAndRounds(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_games, _ := utils.String_slice_file("test_files/test_gamestr")
	expected_rounds, _ := utils.String_slice_file("test_files/test_roundstr")
	for idx, line := range data {
		actual_game, actual_round := split_game_and_rounds(line)
		expected_game := expected_games[idx]
		expected_round := expected_rounds[idx]
		if actual_game != expected_game && actual_round != expected_round {
			t.Errorf("%d: '%s', expected: ('%s', '%s'), actual: ('%s', '%s')\n", 
				idx, line, expected_game, expected_round, actual_game, actual_round)
		}
	}
}

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
	t.Skip("Not fully implemented")
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expecteds := []Game {
		{ ID: 1},
		{ ID: 2},
		{ ID: 3},
		{ ID: 4},
		{ ID: 5},
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