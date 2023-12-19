package seven

import (
	"aoc_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const INPUT_FILE = "seven/input"

type Hand struct {
	Cards string
	Bid int
	Rank int
	Type HandType
}

func (h *Hand) DetermineType(part_one bool) {
	card_count := make(map[string]int)
	for _, card := range h.Cards {
		card_count[string(card)] ++
	}
	var max_cards int
	var jokers int
	for card, count := range card_count {
		if count > max_cards {
			max_cards = count
		}
		if !part_one {
			if card == "J" {
				jokers = count
			}
		}
	}
	if !part_one {
		if jokers > 0 { max_cards += jokers }
	}
	if max_cards == 5 {
		h.Type = types[6]
		return
	}
	if max_cards == 4 {
		h.Type = types[5]
		return
	}
	if max_cards == 3 {
		if len(card_count) == 2 {
			h.Type = types[4]
			return
		}
		h.Type = types[3]
		return
	}
	if max_cards == 2 {
		if len(card_count) == 3 {
			h.Type = types[2]
			return
		}
		h.Type = types[1]
		return
	}
	h.Type = types[0]
}

type HandType struct {
	Name string
	Rank int
}

type HandTypes []HandType

var types HandTypes = HandTypes{
	{"High card", 7},
	{"One pair", 6},
	{"Two pair", 5},
	{"Three of a kind", 4},
	{"Full House", 3},
	{"Four of a kind", 2},
	{"Five of a kind", 1},
}

type CardType struct {
	Face string
	Rank int
}

type CardTypes []CardType

var cards map[string]CardType = map[string]CardType{
	"A": {"A", 1},
	"K": {"K", 2},
	"Q": {"Q", 3},
	"J": {"J", 4},
	"T": {"T", 5},
	"9": {"9", 6},
	"8": {"8", 7},
	"7": {"7", 8},
	"6": {"6", 9},
	"5": {"5", 10},
	"4": {"4", 11},
	"3": {"3", 12},
	"2": {"2", 13},
}

var cards_joker map[string]CardType = map[string]CardType{
	"A": {"A", 1},
	"K": {"K", 2},
	"Q": {"Q", 3},
	"T": {"T", 5},
	"9": {"9", 6},
	"8": {"8", 7},
	"7": {"7", 8},
	"6": {"6", 9},
	"5": {"5", 10},
	"4": {"4", 11},
	"3": {"3", 12},
	"2": {"2", 13},
	"J": {"J", 14},
}

type Hands []Hand

func (h Hands) DetermineHandTypes(part_one bool) {
	for idx := range h {
		var hand *Hand= &h[idx]
		hand.DetermineType(part_one)
	}
}

func (h Hands) SortHands(part_one bool) { // this needs to be tested
	var sort_func func(Hand, Hand) int
	if part_one {
		sort_func = compare_hands
	} else {
		sort_func = compare_hands_jokers
	}
	slices.SortFunc[Hands](h, sort_func)
	
}

func (first Hand) Compare(second Hand, part_one bool) (result int) {
	
	// first hand is worse than second hand
	if first.Type.Rank > second.Type.Rank { return -1 }
	
	// first hand is better than second hand
	if first.Type.Rank < second.Type.Rank { return 1 }
	
	var first_type CardType
	var second_type CardType
	var working_cards map[string]CardType
	if part_one {
		working_cards = cards
	} else {
		working_cards = cards_joker
	}
	for i := range first.Cards {
		first_type = working_cards[string(first.Cards[i])]
		second_type = working_cards[string(second.Cards[i])]
		
		// first card is worse than second card
		if first_type.Rank > second_type.Rank { return -1 }
		
		// first card is better than second card
		if first_type.Rank < second_type.Rank { return 1 }
	}
	return result
}

func (h Hands) RankHands() {
	for i := range h {
		var hand *Hand = &h[i]
		hand.Rank = i + 1
	}
}

func (h Hands) CalculateWinnings(part_one bool) (result int) {
		h.DetermineHandTypes(part_one)
		h.SortHands(part_one)
		h.RankHands()
		for _, hand := range h {
			result += hand.Rank * hand.Bid
		}	

	return result
}

func compare_hands(first Hand, second Hand) (result int) {
	return first.Compare(second, true)
}

func compare_hands_jokers(first Hand, second Hand) (result int) {
	return first.Compare(second, false)
}

func parse_hand(line string) (result Hand) {
	hand_data := strings.Split(line, " ")
	result.Cards = hand_data[0]
	val, _ := strconv.Atoi(hand_data[1])
	result.Bid = val

	return result
}

func parse_hands(data []string) (result Hands) {
	for _, line := range data {
		result = append(result, parse_hand(line))
	}
	return result
}

func total_winnings(data []string, part_one bool) (result int) {
	if part_one {
		result = parse_hands(data).CalculateWinnings(true)
	} else {
		result = parse_hands(data).CalculateWinnings(false)
	}


	return result
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	part_one := total_winnings(data, true)
	fmt.Println("==== Day 7 - Part 1 ====")
	fmt.Printf("Answer: %d\n", part_one)

	part_two := total_winnings(data, false)
	fmt.Println("==== Day 7 - Part 2 ====")
	fmt.Printf("Answer: %d\n", part_two)

}