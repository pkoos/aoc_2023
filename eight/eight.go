package eight

import (
	"aoc_2023/utils"
	"fmt"
)

const INPUT = "eight/input"

type Node struct {
	L string
	R string
}

type Nodes map[string]Node

func travel_nodes(instructions string, nodes Nodes) (result int) {
	return result
}

func parse_file(data []string) (instructions string, nodes Nodes) {

	return instructions, nodes
}

func find_steps(data []string) (result int) {
	instructions, nodes := parse_file(data)
	result = travel_nodes(instructions, nodes)
	return result
}

func Run() {
	data, err := utils.String_slice_file(INPUT)
	if err != nil {
		panic(err)
	}
	part_one := find_steps(data)
	fmt.Println("==== Day 8 - Part 1 ====")
	fmt.Printf("Answer: %d\n", part_one)
	// fmt.Println("==== Day 8 - Part 2 ====")
}