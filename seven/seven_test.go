package seven

import (
	"aoc_2023/utils"
	"reflect"
	"testing"
)



const TEST_INPUT_FILE = "test_files/test_input"

func TestCalculateWinnings(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 6440
	actual := parse_hands(data).CalculateWinnings(true)
	if expected != actual {
		t.Errorf("Part 1 - expected: %d, actual: %d\n", expected, actual)
	}

	expected = 5905
	actual = parse_hands(data).CalculateWinnings(false)
	if expected != actual {
		t.Errorf("Part 2 - expected: %d, actual: %d\n", expected, actual)
	}
}

func TestRankHands(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	part_one := true

	expecteds := Hands{
		{"32T3K", 765, 1, HandType{"One pair", 6}},
		{"KTJJT", 220, 2, HandType{"Two pair", 5}},
		{"KK677", 28,  3, HandType{"Two pair", 5}},
		{"T55J5", 684, 4, HandType{"Three of a kind", 4}},
		{"QQQJA", 483, 5, HandType{"Three of a kind", 4}},
	}
	hands := parse_hands(data)
	hands.DetermineHandTypes(part_one)
	hands.SortHands(part_one)
	hands.RankHands()
	for idx, actual := range hands {
		expected :=expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 1 - %d: expected: %+v, actual: %+v", idx, expected, actual)
		}
	}

	expecteds = Hands{
		{"32T3K", 765, 1, HandType{"One pair", 6}},
		{"KK677", 28,  2, HandType{"Two pair", 5}},
		{"T55J5", 684, 3, HandType{"Four of a kind", 2}},
		{"QQQJA", 483, 4, HandType{"Four of a kind", 2}},
		{"KTJJT", 220, 5, HandType{"Four of a kind", 2}},
		
	}
	hands = parse_hands(data)
	hands.DetermineHandTypes(!part_one)
	hands.SortHands(!part_one)
	hands.RankHands()
	for idx, actual := range hands {
		expected :=expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v", idx, expected, actual)
		}
	}
}

func TestSortHands(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	part_one := true

	expecteds := Hands{
		{"32T3K", 765, 0, HandType{"One pair", 6}},
		{"KTJJT", 220, 0, HandType{"Two pair", 5}},
		{"KK677", 28,  0, HandType{"Two pair", 5}},
		{"T55J5", 684, 0, HandType{"Three of a kind", 4}},
		{"QQQJA", 483, 0, HandType{"Three of a kind", 4}},
	}
	hands := parse_hands(data)
	hands.DetermineHandTypes(part_one)
	hands.SortHands(part_one)
	for idx, actual := range hands {
		expected :=expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%d: expected: %+v, actual: %+v", idx, expected, actual)
		}
	}

	expecteds = Hands{
		{"32T3K", 765, 0, HandType{"One pair", 6}},
		{"KK677", 28,  0, HandType{"Two pair", 5}},
		{"T55J5", 684, 0, HandType{"Four of a kind", 2}},
		{"QQQJA", 483, 0, HandType{"Four of a kind", 2}},
		{"KTJJT", 220, 0, HandType{"Four of a kind", 2}},		
	}
	hands = parse_hands(data)
	hands.DetermineHandTypes(!part_one)
	hands.SortHands(!part_one)
	for idx, actual := range hands {
		expected :=expecteds[idx]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v", idx, expected, actual)
		}
	}
}

func TestDetermineHandTypes(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	part_one := true
	expected_types  := HandTypes{
		{"One pair", 6},
		{"Three of a kind", 4},
		{"Two pair", 5},
		{"Two pair", 5},
		{"Three of a kind", 4},
	}
	hands := parse_hands(data)
	hands.DetermineHandTypes(part_one)
	for idx, hand := range hands {
		expected := expected_types[idx]
		actual := hand.Type

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 1 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}

	expected_types  = HandTypes{
		{"One pair", 6},
		{"Four of a kind", 2},
		{"Two pair", 5},
		{"Four of a kind", 2},
		{"Four of a kind", 2},
	}
	hands = parse_hands(data)
	hands.DetermineHandTypes(!part_one)
	for idx, hand := range hands {
		expected := expected_types[idx]
		actual := hand.Type

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v\n", idx, expected, actual)
		}
	}
}

func TestDetermineType(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_types  := HandTypes{
		{"One pair", 6},
		{"Three of a kind", 4},
		{"Two pair", 5},
		{"Two pair", 5},
		{"Three of a kind", 4},
	}
	for idx, line := range data {
		expected := expected_types[idx]
		actual := parse_hand(line)
		actual.DetermineType(true)
		if !reflect.DeepEqual(actual.Type, expected) {
			t.Errorf("Part 1 - %d: expected: %+v, actual: %+v\n", idx, expected, actual.Type)
		}
	}

	expected_types  = HandTypes{
		{"One pair", 6},
		{"Four of a kind", 2},
		{"Two pair", 5},
		{"Four of a kind", 2},
		{"Four of a kind", 2},
	}

	for idx, line := range data {
		expected := expected_types[idx]
		actual := parse_hand(line)
		actual.DetermineType(false)
		if !reflect.DeepEqual(actual.Type, expected) {
			t.Errorf("Part 2 - %d: expected: %+v, actual: %+v\n", idx, expected, actual.Type)
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
	actual := total_winnings(data, true)
	expected := 765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5 // 6440
	if expected != actual {
		t.Errorf("Part 1 - expected: %d, actual: %d\n", expected, actual)
	}

	actual = total_winnings(data, false)
	expected = 765 * 1 + 684 * 3 + 28 * 2 + 220 * 5 + 483 * 4
	if expected != actual {
		t.Errorf("Part 2 - expected: %d, actual: %d\n", expected, actual)
	}
}