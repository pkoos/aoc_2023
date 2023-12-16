package four

import (
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "four/input"

type Card struct {
	ID int `json:"id"`
	Winning []int `json:"winning"`
	Current []int `json:"current"`
	Matches int `json:"matches"`
}

func (card Card) CalculatePoints() (points int) {
	if card.Matches == 0 {
		return 0
	}
	points += 1
	for i := 1; i<card.Matches; i++ {
		points *= 2
	}
	return points
}

func (card Card) CalculateMatches() {
	// loop through winning
	// check to see if individual number is in current
	// for every match, add to card.Matches field
}

func (card Card) Equals(other_card Card) (isEqual bool) {
	// compare ids
	// compare matches
	// loop to compare winning
	// loop to compare current
	return isEqual
}

type Cards []Card

func parse_numbers(line string) (current []int) {
	return current
}


func parse_id(card_str string) (id int) {
	id_data := strings.Split(card_str, " ")
	id_str := id_data[1]
	id_val, _ := strconv.Atoi(id_str)
	id = id_val
	return id
}

func parse_scratchcard_input(line string) (card Card) {
	split_line := strings.Split(line, ":")
	card.ID = parse_id(split_line[0])

	numbers_data := strings.Split(split_line[1], "|")
	card.Winning = parse_numbers(numbers_data[0])
	card.Current =  parse_numbers(numbers_data[1])
	return card
}

func scratchcard_points(line string) (points int) {
	scratch_card := parse_scratchcard_input(line)
	scratch_card.CalculateMatches()
	points = scratch_card.CalculatePoints()
	return points
}

func sum_points(data []string) (total int) {
	for _, line := range data {
		total += scratchcard_points(line)
	}
	return total
}

func Run() {
	fmt.Println("==== Day 4 ====")
}