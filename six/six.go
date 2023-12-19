package six

import (
	"aoc_2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Race struct {
	Time int `json:"time"`
	Distance int `json:"distance"`
}

func (r Race) Possibilities() (result int) {
	var distance int
	for held := 0; held < r.Time; held++ {
		if distance = held * (r.Time - held); distance > r.Distance {
			result++
		}
	}
	return result
}

func (r Race) Equals(r2 Race) (isEqual bool) {
	isEqual = true
	isEqual = isEqual && r.Time == r2.Time
	isEqual = isEqual && r.Distance == r2.Distance

	return isEqual
}

type Races []Race

func parse_line(line string, part_one bool) (results []int) {
	var val_str string
	var found bool
	for idx, char := range line {
		_ = idx
		if unicode.IsDigit(char) {
			found = true
			val_str += string(char)
			if idx + 1 == len(line) {
				value, _ := strconv.Atoi(val_str)
				results = append(results, value)
			}
		} else {
			if part_one && found {
				value, _ := strconv.Atoi(val_str)
				results = append(results, value)
				found = false
				val_str = ""
			}
		}
	}
	return results
}

func parse_races(data []string, part_one bool) (races Races) {
	var times []int
	var distances []int
	for _, line := range data {
		if strings.Contains(line, "Time:") {
			times = parse_line(line, part_one)
		}
		if strings.Contains(line, "Distance:") {
			distances = parse_line(line, part_one)
		}
	}

	if len(times) != len(distances) {
		fmt.Println("There's something wrong, these should be the same!")
	}

	for i := range times {
		races = append(races, Race{times[i], distances[i]})
	}

	return races
}

func product_winning_possibilities(data []string, part_one bool) (result int) {
	races := parse_races(data, part_one)
	for _, race := range races {
		if result == 0 {
			result += race.Possibilities()
		} else {
			result *= race.Possibilities()
		}
	}
	return result
}

func Run() {
	data, err := utils.String_slice_file("six/input")
	if err != nil {
		panic(err)
	}
	fmt.Println("==== Day 6 - Part 1 ====")
	part_one := product_winning_possibilities(data, true)
	fmt.Printf("Answer: %d\n", part_one)
	fmt.Println("==== Day 6 - Part 2 ====")
	part_two := product_winning_possibilities(data, false)
	fmt.Printf("Answer: %d\n", part_two)
}