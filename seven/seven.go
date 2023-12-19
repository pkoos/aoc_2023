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

func (h *Hand) DetermineType() {
	card_count := make(map[string]int)
	for _, card := range h.Cards {
		card_count[string(card)] ++
	}
	var max_cards int
	for _, count := range card_count {
		if count > max_cards {
			max_cards = count
		}
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
	{"Four pair", 2},
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

type Hands []Hand

func (h Hands) DetermineHandTypes() {
	for idx := range h {
		var hand *Hand= &h[idx]
		hand.DetermineType()
	}
}

func (h Hands) SortHands() { // this needs to be tested
	slices.SortFunc[Hands](h, compare_hands)
}

func (first Hand) Compare(second Hand) (result int) {
	
	// first hand is worse than second hand
	if first.Type.Rank > second.Type.Rank { return -1 }
	
	// first hand is better than second hand
	if first.Type.Rank < second.Type.Rank { return 1 }
	
	var first_type CardType
	var second_type CardType
	for i := range first.Cards {
		first_type = cards[string(first.Cards[i])]
		second_type = cards[string(second.Cards[i])]
		
		// first card is worse than second card
		if first_type.Rank > second_type.Rank { return -1 }
		
		// first card is better than second card
		if first_type.Rank < second_type.Rank { return 1 }
	}
	return result
}

func (h Hands) RankHands() {
	// stuff and things
}

func (h Hands) CalculateWinnings() (result int) {
	return result
}

func compare_hands(first Hand, second Hand) (result int) {
	return first.Compare(second)
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

func total_winnings(data []string) (result int) {
	hands := parse_hands(data)
	for _, hand := range hands {
		hand.DetermineType()
	}
	hands.SortHands() // if rank worked correctly, then all hands should be in order
	hands.RankHands()
	result = hands.CalculateWinnings()

	return result
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	part_one := total_winnings(data)
	fmt.Println("==== Day 7 - Part 1 ====")
	fmt.Printf("Answer: %d\n", part_one)

	fmt.Println("==== Day 7 - Part 2 ====")
}