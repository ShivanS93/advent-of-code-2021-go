package day02

import (
	"adventofcode2021/utils"
	"strconv"
	"strings"
)

// Note: down INCREASES depth
//		 up DECREASES depth

// works out depth
// split line into a string
// multiple depth by horizontal position and return this
func ProductDepthHorizontalPosition(file string) (int, error) {
	depth := 0
	horizontal_position := 0

	movements, err := utils.ReadDataFromFile(file)
	for _, move := range movements {
		s := strings.Fields(move) // [forward 5]
		direction := s[0]
		magnitude, _ := strconv.Atoi(s[1])
		switch direction {
		case "forward":
			horizontal_position = horizontal_position + magnitude
		case "up":
			depth = depth - magnitude
		case "down":
			depth = depth + magnitude
		}
	}
	product := depth * horizontal_position
	return product, err
}

func ProductDepthAimHorizontalPosition(file string) (int, error) {
	depth := 0
	horizontal_position := 0
	aim := 0

	movements, err := utils.ReadDataFromFile(file)
	for _, move := range movements {
		s := strings.Fields(move)          // [forward 5]
		direction := s[0]                  // "forward"
		magnitude, _ := strconv.Atoi(s[1]) // 5
		switch direction {
		case "forward":
			horizontal_position = horizontal_position + magnitude
			depth = depth + (aim * magnitude)
		case "up":
			aim = aim - magnitude
		case "down":
			aim = aim + magnitude
		}
	}
	product := depth * horizontal_position
	return product, err
}
