package seven

import (
	"aoc_2023/utils"
	"cmp"
	"fmt"
	"slices"
)

const INPUT_FILE = "seven/input"

type Hand struct {
	Cards string
	Bid int
	Rank int
	Type HandType
}

type HandType struct {
	Name string
	Rank int
}

type HandTypes []HandType

var types HandTypes = HandTypes{
	{"High card", 1},
	{"One pair", 2},
	{"Two pair", 3},
	{"Three of a kind", 4},
	{"Full House", 5},
	{"Four pair", 6},
	{"Five of a kind", 7},
}


type Hands []Hand

func (h Hands) SortHands() { // this needs to be tested
	slices.SortFunc[Hands](h, func(a, b Hand) (n int) {
		if n = cmp.Compare(a.Type.Rank, b.Type.Rank); n != 0 {
			return n
		}
		return cmp.Compare(a.Cards, b.Cards)

	})
}

func (h Hands) Rank() {
	// stuff and things
}

func (h Hands) CalculateWinnings() (result int) {
	return result
}

func parse_hand(line string) (result Hand) {
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
	hands.SortHands() // if rank worked correctly, then all hands should be in order
	hands.Rank()
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