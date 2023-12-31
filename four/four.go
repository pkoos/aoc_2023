package four

import (
	"aoc_2023/utils"
	"fmt"
	"slices"
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

func (card *Card) CalculateMatches() {
	// loop through winning
	for _, digit := range card.Winning {
		if slices.Contains[[]int, int](card.Current, digit) {
			card.Matches += 1
		}
	}
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

func calculate_scratchcards(data []string) (num_scratchcards int) {
	results := make(map[int]int)
	total_cards := len(data)
	for _, line := range data {
		card := parse_scratchcard_input(line)
		card.CalculateMatches() // this gets the number of matches
		results[card.ID] += 1 // This is the original card being added
		for i := 0; i < results[card.ID]; i++ {
			for j:= 0; j < card.Matches; j++ {
				copy_card_id := 1 + card.ID + j
				if copy_card_id <= total_cards {
					results[copy_card_id] += 1
				}
			}
		}
	}
	for _, val := range results {
		num_scratchcards += val
	}
	return num_scratchcards
}

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
	var id_str string
	for _, char := range card_str {
		if unicode.IsDigit(char) {
			id_str += string(char)
		}
	}
	val, _ := strconv.Atoi(id_str)
	id = val
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
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}

	sum_part_one := sum_points(data)
	fmt.Println("==== Day 4 - Part 1 ====")
	fmt.Printf("Answer: %d\n", sum_part_one)

	sum_part_two := calculate_scratchcards(data)
	fmt.Println("==== Day 4 - Part 2 ====")
	fmt.Printf("Answer: %d\n", sum_part_two)
}