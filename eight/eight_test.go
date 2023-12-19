package eight

import (
	"aoc_2023/utils"
	"reflect"
	"testing"
)
const TEST_INPUT_FILE = "test_files/test_input"
const TEST_INPUT2_FILE = "test_files/test_input2"

func TestTravelNodes(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 2
	actual := travel_nodes(parse_file(data))
	if expected != actual {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}

	data2, _ := utils.String_slice_file(TEST_INPUT2_FILE)
	expected2 := 6
	actual2 := travel_nodes(parse_file(data2))
	if expected != actual {
		t.Errorf("expected: %d, actual %d\n", expected2, actual2)
	}
	t.Skip("Not yet implemented.")
}

func TestParseNode(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_nodes := []Node{
		{"AAA", "BBB", "CCC"},
		{"BBB", "DDD", "EEE"},
		{"CCC", "ZZZ", "GGG"},
		{"DDD", "DDD", "DDD"},
		{"EEE", "EEE", "EEE"},
		{"GGG", "GGG", "GGG"},
		{"ZZZ", "ZZZ", "ZZZ"},
	}

	for idx, line := range data {
		if idx < 2 { continue }
		actual := parse_node(line)
		expected := expected_nodes[idx - 2]
		if !reflect.DeepEqual(actual, expected){
			t.Errorf("%d: expected: %+v, actual: %+v\n", (idx - 2), expected, actual)
		}
	}

	data2, _ := utils.String_slice_file(TEST_INPUT2_FILE)
	expected_nodes2 := []Node{
		{"AAA", "BBB", "BBB"},
		{"BBB", "AAA", "ZZZ"},
		{"ZZZ", "ZZZ", "ZZZ"},
	}

	for idx, line := range data2 {
		if idx < 2 { continue }
		actual := parse_node(line)
		expected := expected_nodes2[idx - 2]
		if !reflect.DeepEqual(actual,expected) {
			t.Errorf("%d: expected: %+v, actual: %+v\n", (idx - 2), expected, actual)
		}
	}
}

func TestParseFile(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected_instructions := "RL"
	expected_nodes := Nodes{
		"AAA":{"AAA", "BBB", "CCC"},
		"BBB":{"BBB", "DDD", "EEE"},
		"CCC":{"CCC", "ZZZ", "GGG"},
		"DDD":{"DDD", "DDD", "DDD"},
		"EEE":{"EEE", "EEE", "EEE"},
		"GGG":{"GGG", "GGG", "GGG"},
		"ZZZ":{"ZZZ", "ZZZ", "ZZZ"},
	}
	var actual_instructions string
	var actual_nodes Nodes = make(Nodes)
	actual_instructions, actual_nodes = parse_file(data)
	
	if expected_instructions != actual_instructions {
		t.Errorf("instructions - expected: %s, actual: %s\n", expected_instructions, actual_instructions)
	}
	if !reflect.DeepEqual(actual_nodes, expected_nodes) {
		t.Errorf("nodes - expected: %+v, actual: %+v\n", expected_nodes, actual_nodes)
	}

	data2, _ := utils.String_slice_file(TEST_INPUT2_FILE)
	expected_instructions2 := "LLR"
	expected_nodes2 := Nodes{
		"AAA":{"AAA", "BBB", "BBB"},
		"BBB":{"BBB", "AAA", "ZZZ"},
		"ZZZ":{"ZZZ", "ZZZ", "ZZZ"},
	}

	var actual_instructions2 string
	var actual_nodes2 Nodes = make(Nodes)
	actual_instructions2, actual_nodes2 = parse_file(data2)
	if expected_instructions2 != actual_instructions2 {
		t.Errorf("instructions - expected: %s, actual: %s\n", expected_instructions2, actual_instructions2)
	}
	if !reflect.DeepEqual(actual_nodes2, expected_nodes2) {
		t.Errorf("nodes - expected: %+v, actual: %+v\n", expected_nodes2, actual_nodes2)
	}
}

func TestFindSteps(t *testing.T) {
	data, _ := utils.String_slice_file(TEST_INPUT_FILE)
	expected := 2
	actual := find_steps(data)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d\n", expected, actual)
	}

	data2, _ := utils.String_slice_file(TEST_INPUT2_FILE)
	expected2 := 6
	actual2 := find_steps(data2)
	if actual2 != expected2 {
		t.Errorf("expected: %d, actual: %d\n", expected2, actual2)
	}
}