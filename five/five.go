package five

import (
	"fmt"
	"reflect"
	"slices"
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

type DSR struct {
	Dest int
	Src int
	Rng int
}

type Seed_location struct {
	Seed int `json:"seed"`
	Location int `json:"location"`
	// Path []int `json:"path"`
}

type Seeds []Seed_location

type Seed_range struct {
	Start int
	Length int
}

type Seedranges []Seed_range

type Farm struct {
	Seeds []int
	Maps FarmMaps
}

type Farm_map struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Next int `json:"next"`
	Destinations []int `json:"destination"`
	Sources []int `json:"source"`
	Ranges []int `json:"range"`
}

func (farm Farm) MapSeedRanges(seeds Seedranges) (seed_locations Seeds) {
	for _, seedrange := range seeds {
		start := seedrange.Start
		end := start + seedrange.Length
		for i := seedrange.Start; i < end; i++ {
			seed_locations = append(seed_locations, farm.TraverseMapping(i))
		}
	}
	return seed_locations
}

func (farm Farm) MapSeeds() (seed_locations Seeds) {
	for _, seed := range farm.Seeds {
		location := farm.TraverseMapping(seed)
		seed_locations = append(seed_locations, location)
	}
	return seed_locations
}

func (farm Farm) LowestLocationRange() (location DSR) {
	l_map := farm.Maps[len(farm.Maps) - 1]
	min := slices.Index(l_map.Destinations, slices.Min(l_map.Destinations))

	return DSR{l_map.Destinations[min], l_map.Sources[min], l_map.Ranges[min]}
}

func (farm Farm) LowestLocation(loc DSR) (lowest_location int) {
	rev_maps := slices.Clone[FarmMaps](farm.Maps)
	slices.Reverse[FarmMaps](rev_maps)

	// found_destination := true
	var c_source int
	var c_destination  int

	_ = c_source
	_ = c_destination
	for i := 0; i <= loc.Rng; i++ {
		source := loc.Src + i
		destination := 0
		// i is the specific value of loc.Rng we are on right now
Next_Rev_Map:
		for map_idx, current_map := range rev_maps {
			// found_destination = false
			if map_idx == 0 { continue } // we have the lowest range already
			for dest_idx, current_destination := range current_map.Destinations {
				current_range := current_map.Ranges[dest_idx]
				max_destination := current_destination + current_range
				if source <= max_destination && source >= current_destination {
					// found_destination = true
					// fmt.Printf("%d: We found a match! source(%d), dest(%d), max_destination(%d)\n", dest_idx, source, current_destination, max_destination)
					diff := source - current_destination
					destination = current_destination + diff
					if destination != source {
						fmt.Printf("Sanity check, we should never fall within here EVER")
					}
					source = current_map.Sources[dest_idx] + diff
					// fmt.Printf("Name: %s, destination: %d, source: %d\n", current_map.Name, destination, source)
					continue Next_Rev_Map
				}
			}
		}
		seed_ranges := parse_seed_range(farm.Seeds)
		for seed_idx, seed_range := range seed_ranges {
			_ = seed_idx
			start := seed_range.Start
			end := start + seed_range.Length
			if source >= start && source <= end {
				difference := source - start
				lowest_location = start + difference
				// fmt.Printf("%d: We found a seed! lowest_location: %d\n", seed_idx, lowest_location)
				break
			}
		}
		if lowest_location > 0 {
			break
		}
	}

	return lowest_location
}

func (farm Farm) LocationToSeedMapping(dest int, src int) (seed int) {
	
	return seed
}

func (farm Farm) TraverseMapping(seed_num int) (location Seed_location) {
	location.Seed = seed_num
	// var path []int
	// path = append(path, seed_num)
	source := seed_num
	destination := -1
Next_Map:
	for idx, farm_map := range farm.Maps {
		_ = idx // not sure if I need this yet
		for i, src := range farm_map.Sources {
			rng := farm_map.Ranges[i]
			src_high := src + rng - 1
			if source >= src && source <= src_high {
				dest := farm_map.Destinations[i]
				diff := source - src
				destination = dest + diff
				// path = append(path, destination)
				source = destination
				continue Next_Map
			}
		}
	}
	// location.Path = path
	location.Location = destination

	return location
}

func (current_map Farm_map) Equals(other_map Farm_map) (isEqual bool) {
	isEqual = true // start at true, and hope to fall out the other end still true
	// IDs
	isEqual = isEqual && current_map.ID == other_map.ID
	// Names
	isEqual = isEqual && current_map.Name == other_map.Name
	// Next
	isEqual = isEqual && current_map.Next == other_map.Next
	// Destinations
	isEqual = isEqual && reflect.DeepEqual(current_map.Destinations, other_map.Destinations)
	// Sources
	isEqual = isEqual && reflect.DeepEqual(current_map.Sources, other_map.Sources)
	// Ranges
	isEqual = isEqual && reflect.DeepEqual(current_map.Ranges, other_map.Ranges)

	return isEqual
}

type FarmMaps []Farm_map

func parse_map_data(data [] string) (farm_map Farm_map) {
	for idx, line := range data {
		if idx == 0 {
			for key, value := range maps {
				if strings.Contains(line, key) {
					farm_map.ID = value[0]
					farm_map.Next = value[1]
					farm_map.Name = key
					break
				}
			}
		} else {
			map_str := strings.Split(line, " ")
			map_source, _ := strconv.Atoi(map_str[1])
			map_destination, _ := strconv.Atoi(map_str[0])
			map_range, _ := strconv.Atoi(map_str[2])
			farm_map.Sources = append(farm_map.Sources, map_source)
			farm_map.Destinations = append(farm_map.Destinations, map_destination)
			farm_map.Ranges = append(farm_map.Ranges, map_range)
		}
	}

	return farm_map
}

func parse_seed_range(vals []int) (seeds Seedranges) {
	for i := 0;i < len(vals); i += 2 {
		seeds = append(seeds, Seed_range{vals[i], vals[i + 1]})
	}

	return seeds
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
				if idx + 1 == len(data) {
					map_data = append(map_data, line)
				}
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

func find_lowest_location_seedrange(data []string) (lowest_location int) {
	farm := parse_file(data)
	lowest_location_range := farm.LowestLocationRange()
	lowest_location = farm.LowestLocation(lowest_location_range)
	
	return lowest_location
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
	
	min_loc_part_two := find_lowest_location_seedrange(data)
	fmt.Println("==== Day 5 - Part 2 ====")
	fmt.Printf("Answer: %d\n", min_loc_part_two)

}