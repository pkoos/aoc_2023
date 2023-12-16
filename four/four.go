package four

import (
	// "aoc_2023/utils"
	"fmt"
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

type Cards []Card

func sum_points(data []string) (total int) {
	return total
}

func Run() {
	fmt.Println("==== Day 4 ====")
}