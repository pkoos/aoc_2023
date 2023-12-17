package five

import (
	_ "aoc_2023/utils"
	"fmt"
	"testing"
)

func TestParseSeeds(t *testing.T) {
	actual := parse_seeds("seeds: 79 14 55 13")
	var expected = []int{ 79, 14, 55, 13}
	for idx, act := range actual {
		exp := expected[idx]
		if exp != act {
			t.Errorf("%d: expected: %d, actual: %d\n", idx, exp, act)
		}
	}
}

func TestParseFile(t *testing.T) {
	t.Skip("Not yet implemented")
}

func TestFindLowestLocation(t *testing.T) {
	t.Skip("Not yet implemented")
	fmt.Println("Hello, world!")
}