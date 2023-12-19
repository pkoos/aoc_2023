package eight

import (
	"aoc_2023/utils"
	"fmt"
	"strings"
)

const INPUT = "eight/input"

type Node struct {
	Name string
	L string
	R string
}

type Nodes map[string]Node

func travel_nodes(instructions string, nodes Nodes) (result int) {
	var path []string
	start := nodes["AAA"]
	current := start
	var next Node
	end := nodes["ZZZ"]
	path = append(path, start.Name)
	var finished bool = false
	for {
		for idx, direction := range instructions {
			_ = idx
			result ++
			if direction == rune('L') {
				next = nodes[current.L]
			} else if direction == rune('R') {
				next = nodes[current.R]
			}
			
			current = next
			path = append(path, current.Name)
			
			if next == end {
				finished = true
				break
			}

		}
		if finished {
			break
		}
	}
	fmt.Printf("path: %+v\n", path)
	return result
}

func parse_node(line string) (node Node) {
	line_data := strings.Split(line, "=")
	node.Name = strings.TrimSpace(line_data[0])
	node_data := strings.Split(line_data[1], ",")
	node.L = strings.TrimSpace(node_data[0][2:])
	node.R = strings.TrimSpace(node_data[1][:len(node_data[1]) - 1])

	return node
}

func parse_file(data []string) (instructions string, nodes Nodes) {
	nodes = make(Nodes)
	for idx, line := range data {
		if idx == 0 {
			instructions = line
			continue
		}
		if line == "" { continue }
		node := parse_node(line)
		nodes[node.Name] = node
	}
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