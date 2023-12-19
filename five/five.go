package five

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"aoc_2023/utils"
)

const INPUT = "five/input"

var maps = map[string][]int{
	"seed-to-soil": {1, 2},
	"soil-to-fertilizer": {2, 3},
	"fertilizer-to-water": {3, 4},
	"water-to-light": {4, 5},
	"light-to-temperature": {5, 6},
	"temperature-to-humidity": {6, 7},
	"humidity-to-location": {7, 0}}

	type Seed_location struct {
		Seed int `json:"seed"`
		Location int `json:"location"`
		Path []int `json:"path"`
	}

	type Seeds []Seed_location

type Farm struct {
	Seeds []int
	Maps FarmMaps
}

type Farm_map struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Destination int `json:"destination"`
	Mapping map[int]int `json:"mapping"`
}

func (farm Farm) MapSeeds() (seed_locations Seeds) {
	for _, seed := range farm.Seeds {
		location := farm.TraverseMapping(seed)
		seed_locations = append(seed_locations, location)
	}
	return seed_locations
}

func (farm Farm) TraverseMapping(seed_num int) (location Seed_location) {
	location.Seed = seed_num
	var path []int
	path = append(path, seed_num)
	source := seed_num
	destination := -1
	for _, farm_map := range farm.Maps {
		destination = farm_map.Mapping[source]
		if destination == 0 {
			destination = source
		}
		path = append(path, destination)
		source = destination
	}
	location.Path = path
	location.Location = destination

	return location
}

func (current_map Farm_map) Equals(other_map Farm_map) (isEqual bool) {
	ids_equal := current_map.ID == other_map.ID
	names_equal := current_map.Name == other_map.Name
	destinations_equal := current_map.Destination == other_map.Destination
	mappings_equal := reflect.DeepEqual(current_map.Mapping, other_map.Mapping)

	isEqual = ids_equal && names_equal && destinations_equal && mappings_equal
	return isEqual
}

type FarmMaps []Farm_map

func parse_map_data(data [] string) (farm_map Farm_map) {
	farm_map.Mapping = make(map[int]int)
	for idx, line := range data {
		if idx == 0 {
			for key, value := range maps {
				if strings.Contains(line, key) {
					farm_map.ID = value[0]
					farm_map.Destination = value[1]
					farm_map.Name = key
					break
				}
			}
		} else {
			map_str := strings.Split(line, " ")
			map_source, _ := strconv.Atoi(map_str[1])
			map_destination, _ := strconv.Atoi(map_str[0])
			map_range, _ := strconv.Atoi(map_str[2])
			for i := 0; i < map_range;i++ {
				farm_map.Mapping[map_source + i] = map_destination + i
			}
		}
	}

	return farm_map
}

func parse_seeds(line string) (seeds []int) {
	var seed_str string
	part_of_digit := false
	for idx, char := range line {
		if unicode.IsDigit(char) {
			seed_str += string(char)
			part_of_digit = true
			if idx + 1 == len(line) {
				val, _ := strconv.Atoi(seed_str)
				seeds = append(seeds, val)
			}
		} else {
			if part_of_digit {
				val, _ := strconv.Atoi(seed_str)
				seeds = append(seeds, val)
				part_of_digit = false
				seed_str = ""
			}
		}
	}
	return seeds
}

func parse_file(data [] string) (farm Farm) {
	var map_data []string
	in_map_data := false
	for idx, line := range data {
		if strings.Contains(line, "seeds:") {
			farm.Seeds = parse_seeds(line)
			continue
		}

		if strings.Contains(line, "map:") {
			in_map_data = true
		}
		if in_map_data {
			if line == "" || idx + 1 == len(data) {
				farm_map := parse_map_data(map_data)
				farm.Maps = append(farm.Maps, farm_map)
				in_map_data = false
				map_data = nil
			} else {
				map_data = append(map_data, line)
			}
			
		}
	}
	return farm
}

func find_lowest_location(data []string) (lowest_location int) {
	farm := parse_file(data)
	lowest_location = -1
	for _, seed_location := range farm.MapSeeds() {
		if lowest_location < 0 {
			lowest_location = seed_location.Location
			continue
		}
		if seed_location.Location < lowest_location {
			lowest_location = seed_location.Location
		}
	}
	return lowest_location
}

func Run() {
	data, err := utils.String_slice_file("five/input")
	if err != nil {
		panic(err)
	}

	min_loc_part_one := find_lowest_location(data)
	fmt.Println("==== Day 5 - Part 1 ====")
	fmt.Printf("Answer: %d\n", min_loc_part_one)
	
	fmt.Println("==== Day 5 - Part 2 ====")
}