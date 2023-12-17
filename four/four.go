package four

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	ids_equal := card.ID == other_card.ID
	matches_equal := card.Matches == other_card.Matches
	
	winning_digits_equal := true
	for idx, num_win := range card.Winning {
		if num_win != other_card.Winning[idx] {
			winning_digits_equal = false
			break
		}
	}
	
	current_digits_equal := true
	for idx, num_win := range card.Winning {
		if num_win != other_card.Winning[idx] {
			current_digits_equal = false
			break
		}
	}
	
	isEqual = ids_equal && matches_equal && winning_digits_equal && current_digits_equal
	return isEqual
}

type Cards []Card

func parse_numbers(line string) (current []int) {
	var found_digit bool
	var num_str string

	for idx, char := range line {
		if unicode.IsDigit(char) { // found a numeric digit
			if len(line) - 1 == idx {
				num_str += string(char)
				digit, err := strconv.Atoi(num_str)
				if err != nil {
					fmt.Printf("Error parsing \"%s\".", num_str)					
				}
				current = append(current, digit)
			}
			found_digit = true
			num_str += string(char)
		} else {
			if found_digit {
				digit, err := strconv.Atoi(num_str)
				if err != nil {
					fmt.Printf("Error parsing \"%s\".", num_str)
				}
				current = append(current, digit)
				num_str = ""
				found_digit = false
			}
		}
	}
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