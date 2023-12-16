package four

import (
	// "aoc_2023/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "four/input"

type Card struct {
	ID int
	Winning []int
	Current []int
	Matches int
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

}

type Cards []Card

func parse_current_numbers(line string) (current []int) {
	return current
}

func parse_winning_numbers(line string) (winning []int) {
	return winning
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
	card.Winning = parse_winning_numbers(line)
	card.Current =  parse_current_numbers(line)
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