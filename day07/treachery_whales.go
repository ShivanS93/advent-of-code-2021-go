package day07

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// extract file
// find differences of each crap for each position
// store this in a alice / array

func readDataFromFile(file string) []int {
	content, _ := ioutil.ReadFile(file)
	str_values := strings.Split(string(content), ",")
	var values []int
	for _, str_value := range str_values {
		value, _ := strconv.Atoi(str_value)
		values = append(values, value)
	}
	return values
}

// find max, assume min is 0

func findMax(values []int) int {
	var max int
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func FindMoves(file string) int {
	values := readDataFromFile(file)
	max := findMax(values)
	var minimumDifference int
	var currDifference int
	var difference int
	for i := 0; i <= max; i++ { // i is position being tested
		currDifference = 0
		for _, value := range values {
			difference = i - value
			if difference < 0 {
				difference = difference * -1 // make positive
			}
			currDifference = currDifference + difference
		}
		if i == 0 {
			minimumDifference = currDifference
		}
		if currDifference < minimumDifference {
			minimumDifference = currDifference
		}
	}
	return minimumDifference
}

func FindMovesExpensive(file string) int {
	values := readDataFromFile(file)
	max := findMax(values)
	var minimumFuel int
	var currFuel int
	var difference int
	for i := 0; i <= max; i++ {
		currFuel = 0
		for _, value := range values {
			difference = i - value
			if difference < 0 {
				difference = difference * -1 // make positive
			}
			// work out the new fuel cost
			// 1 2 3 4 5 -> sum for 5 position move
			var quickSum int
			for n := 0; n <= difference; n++ {
				quickSum = quickSum + n
			}
			currFuel = currFuel + quickSum
		}
		if i == 0 {
			minimumFuel = currFuel
		}
		if currFuel < minimumFuel {
			minimumFuel = currFuel
		}
	}
	return minimumFuel
}
