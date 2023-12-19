package seven

import (
	"aoc_2023/utils"
	"reflect"
	"testing"
)



const TEST_INPUT_FILE = "test_files/test_input"

func TestDetermineHandTypes(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_types  := HandTypes{
		{"One pair", 2},
		{"Three of a kind", 4},
		{"Two pair", 3},
		{"Two pair", 3},
		{"Three of a kind", 4},
	}
	hands := parse_hands(data)
	hands.DetermineHandTypes()
	for idx, hand := range hands {
		expected := expected_types[idx]
		actual := hand.Type

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestDetermineType(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_types  := HandTypes{
		{"One pair", 2},
		{"Three of a kind", 4},
		{"Two pair", 3},
		{"Two pair", 3},
		{"Three of a kind", 4},
	}
	for idx, line := range data {
		expected := expected_types[idx]
		actual := parse_hand(line)
		actual.DetermineType()
		if !reflect.DeepEqual(actual.Type, expected) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual.Type)
		}
	}
}

func TestParsehand(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expecteds := Hands{
		{"32T3K", 765, 0, HandType{}},
		{"T55J5", 684, 0, HandType{}},
		{"KK677", 28,  0, HandType{}},
		{"KTJJT", 220, 0, HandType{}},
		{"QQQJA", 483, 0, HandType{}},
	}
	for idx, line := range data {
		actual := parse_hand(line)
		expected := expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestParseHands(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	hands := parse_hands(data)
	expecteds := Hands{
		{"32T3K", 765, 0, HandType{}},
		{"T55J5", 684, 0, HandType{}},
		{"KK677", 28,  0, HandType{}},
		{"KTJJT", 220, 0, HandType{}},
		{"QQQJA", 483, 0, HandType{}},
	}
	for idx, actual := range hands {
		expected := expecteds[idx]
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestTotalWinnings(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	actual := total_winnings(data)
	expected := 765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5 // 6440
	t.Skip("Not yet implemented")
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}
	
}