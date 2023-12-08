package utils

import (
	"os"
	"strconv"
	"strings"
)

func String_slice_file(file_path string) ([]string, error) {
	raw_data, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(raw_data), "\n"), nil
}

func Int_slice_file(file_path string) ([]int, error) {
	string_data, err := String_slice_file(file_path)
	if err != nil {
		return nil, err
	}
	
	var int_data []int
	for _, val := range string_data {
		int_val, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		
		int_data = append(int_data, int_val)
	}

	return int_data, nil
}