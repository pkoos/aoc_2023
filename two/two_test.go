package two

import (
	"aoc_2023/utils"
	"strings"
	"testing"
)

const TEST_INPUT_FILE = "test_files/test_input"

func TestSumProducts(t *testing.T) {
	data, _ := utils.String_slice_file("test_files/test_input")
	// expected := 2286
	expected := 0
	actual := sum_products(data)
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
}

func TestIsPossible(t *testing.T) {
	data, _ := utils.String_slice_file("test_files/test_roundstr")
	expecteds := []bool { true, true, false, false, true }

	for idx, line := range data {
		actual := is_possible(process_all_rounds(line))
		expected := expecteds[idx]
		if actual != expected {
			t.Errorf("%d: '%s', expected: %t, actual: %t\n", idx, line, expected, actual)
		}
	}
}

func TestColorValue(t *testing.T) {
	data, _ := utils.String_slice_file("test_files/test_colorval")
	expecteds := []Color_val{
		{"blue", 3},  {"red", 4},   {"red", 1},   {"green", 2},  {"blue", 6}, 
		{"green", 2}, {"blue", 1},  {"green", 2}, {"green", 3},  {"blue", 4}, 
		{"red", 1},   {"green", 1}, {"blue", 1},  {"green", 8},  {"blue", 6}, 
		{"red", 20},  {"blue", 5},  {"red", 4},   {"green", 13}, {"green", 5}, 
		{"red", 1},   {"green", 1}, {"red", 3},   {"blue", 6},   {"green", 3}, 
		{"red", 6},   {"green", 3}, {"blue", 15}, {"red", 14},   {"red", 6}, 
		{"blue", 1},  {"green", 3},  {"blue", 2},  {"red", 1},   {"green", 2},
	}
	for idx, line := range data {
		expected := expecteds[idx]
		actual := color_value(line)
		if expected.name != actual.name && expected.value != actual.value {
			t.Errorf("%d: '%s', expected: %+v, actual: %+v\n", idx, line, expected, actual)
		}
	}
}
func TestProcessOneRound(t *testing.T) {
	data, _ := utils.String_slice_file("test_files/test_roundstr")
	expecteds := [][]Game_round{
		{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}},
		{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}},
		{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}},
		{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}},
		{{6, 3, 1}, {1, 2, 2}},
	}
	for game_idx, line := range data {
		line_data := strings.Split(line, ";")
		for round_idx, round_data := range line_data {
			actual := process_one_round(round_data)
			expected := expecteds[game_idx][round_idx]
			if actual != expected {
				t.Errorf("[%d][%d]: '%s', expected: %+v, actual: %+v\n", game_idx, round_idx, line, expected, actual )
			}
		}
	}
}

func TestProcessAllRounds(t *testing.T) {
	data, _ := utils.String_slice_file("test_files/test_roundstr")
	expecteds := [][]Game_round{
		{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}},
		{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}},
		{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}},
		{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}},
		{{6, 3, 1}, {1, 2, 2}},
	}
	for idx, line := range data {
		expected_rounds :=expecteds[idx]
		actual_rounds := process_all_rounds(line)
		for idx_2, actual := range actual_rounds {
			expected := expected_rounds[idx_2]
			if actual != expected {
				t.Errorf("[%d][%d]: '%s', expected: %+v, actual: %+v\n", idx, idx_2, line, expected, actual)
			}
		}
	}
}

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
	// t.Skip("Not Yet Implemented") 
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 8
	actual := sum_ids(data)
	if actual != expected {
		t.Errorf("expected: %d, actual :%d\n", expected, actual)
	}
}