package utils

import (
	"os"
	"strings"
)

func Slice_file(file_path string) ([]string, error) {
	raw_data, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(raw_data), "\n"), nil
}