package four

import (
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

func Run() {
	fmt.Println("==== Day 4 ====")
}