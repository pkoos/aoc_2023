package two

import (
	"aoc_2023/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_FILE = "two/input"

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

var MAX_ROUND = Game_round{MAX_RED, MAX_GREEN, MAX_BLUE}

type Game struct {
	ID int
	Rounds []Game_round
	Is_possible bool
}

func (this_game Game) Equals (that_game Game) (equals bool) {
	equals = true

	if this_game.ID != that_game.ID { return false }
	if len(this_game.Rounds) != len(that_game.Rounds) { return false }
	for idx, this_round := range this_game.Rounds {
		that_round := that_game.Rounds[idx]
		if this_round != that_round { return false }
	}
	if this_game.Is_possible != that_game.Is_possible { return false }
	
	return equals
}

type Game_round struct {
	Red int
	Green int
	Blue int
}

type Color_val struct {
	name string
	value int
}

func split_game_and_rounds(line string) (game_str string, rounds_str string) {
	data := strings.Split(line, ":")
	game_str = data[0]
	rounds_str = data[1]

	return game_str, rounds_str
}

func color_value(line string) (cv Color_val) {
	space_idx := strings.LastIndex(line, " ")
	val_str := strings.TrimSpace(line[: space_idx])
	value, _ := strconv.Atoi(val_str)
	name := strings.TrimSpace(line[space_idx:])

	return Color_val{name, value}
}

func process_one_round(line string) (round Game_round) {
	// fmt.Printf("process_one_round::line: '%s'\n", line)
	colors := strings.Split(line, ",")
	for _, color := range colors {
		cv := color_value(color)
		if cv.name == "red" {
			round.Red = cv.value
		}
		if cv.name == "green" {
			round.Green = cv.value
		}
		if cv.name == "blue" {
			round.Blue = cv.value
		}

	}
	return round
}

func process_all_rounds(line string) (rounds []Game_round) {
	data := strings.Split(line, ";")
	for _, line := range data {
		round := process_one_round(line)
		rounds = append(rounds, round)
	}
	return rounds
}

func game_id_int(line string) (id int) {
	space_idx := strings.Index(line, " ")
	id_string := strings.TrimSpace(line[space_idx:])
	id, _ = strconv.Atoi(id_string)
	
	return id
}

func is_possible(rounds []Game_round) bool {
	for _, round := range rounds {
		if round.Red > MAX_RED || round.Green > MAX_GREEN || round.Blue > MAX_BLUE {
			return false
		}
	}
	return true	
}

func process_game_string(line string) (game Game) {
	game_str, rounds_str := split_game_and_rounds(line)
	game.ID = game_id_int(game_str)
	game.Rounds = process_all_rounds(rounds_str)
	game.Is_possible = is_possible(game.Rounds)

	return game
}

func sum_ids(data []string) (sum int) {
	for _, line := range data {
		if game := process_game_string(line); game.Is_possible {
			sum += game.ID
		}
	}
	return sum
}

func min_game_product(game Game) (int) {
	var min_round Game_round

	for _, round := range game.Rounds {
		if min_round.Red == 0 || round.Red > min_round.Red {
			min_round.Red = round.Red
		}
		if min_round.Green == 0 || round.Green > min_round.Green {
			min_round.Green = round.Green
		}
		if min_round.Blue == 0 || round.Blue > min_round.Blue {
			min_round.Blue = round.Blue
		}
	}
	return min_round.Red * min_round.Green * min_round.Blue
}

func sum_products(data []string) (sum int) {
	for _, line := range data {
		game := process_game_string(line)
		sum += min_game_product(game)
	}
	return sum
}

func Run() {
	data, err := utils.String_slice_file(INPUT_FILE)
	if err != nil {
		panic(err)
	}
	sum := sum_ids(data)
	fmt.Println("==== Day 2 - Part 1 ====")
	fmt.Printf("Answer: %d\n", sum)

	sum_prod := sum_products(data)
	fmt.Println("==== Day 2 - Part 2 ====")
	fmt.Printf("Answer: %d\n", sum_prod)
}