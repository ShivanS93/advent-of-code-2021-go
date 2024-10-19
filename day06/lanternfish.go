package day06

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// func oldreadDataFromFile(file string) []int {
// 	content, _ := ioutil.ReadFile(file)
// 	str_values := strings.Split(string(content), ",")
// 	var values []int
// 	for _, str_value := range str_values {
// 		value, _ := strconv.Atoi(str_value)
// 		values = append(values, value)
// 	}
// 	return values
// }

// count number of 0's
// loop through state, check if above 0, then subject by 1
// if 0 then count 0++ and reset to 6
// also keep trace of the days = 80s
// func OldSimulateFish(file string) int {
// 	initialState := readDataFromFile(file)
// 	const days = 256 // after 80 days (cycles)
// 	var currentState []int
// 	var numReproduce int
// 	currentState = initialState
// 	for i := 0; i < days; i++ {
// 		numReproduce = 0
// 		for j, fish := range currentState {
// 			if fish == 0 {
// 				currentState[j] = 6 // resetting fish
// 				numReproduce++
// 			} else {
// 				currentState[j]-- // reducing fish by one
// 			}
// 		}
// 		// adding new fish
// 		for j := 0; j < numReproduce; j++ {
// 			currentState = append(currentState, 8)
// 		}
// 	}
// 	// count number of fish
// 	return len(currentState)
// }

func readDataFromFile(file string) [9]int {
	content, _ := ioutil.ReadFile(file)
	str_values := strings.Split(string(content), ",")
	var count [9]int
	for _, str_value := range str_values {
		value, _ := strconv.Atoi(str_value)
		count[value]++
	}
	return count
}

// convert state into a array counter to keep count of all fishe
func SimulateFish(file string, days int) int {
	count := readDataFromFile(file)

	// i = day, j = fish state
	var countSlice []int
	countSlice = count[:]
	for i := 0; i < days; i++ {
		newFish := countSlice[0]
		// shifting slice
		countSlice = append(countSlice[1:])      // removes 0 index
		countSlice[6] = countSlice[6] + newFish  // 0th fish regenerated
		countSlice = append(countSlice, newFish) // new fish
	}
	// sum slice
	var total int
	for _, num := range countSlice {
		total = total + num
	}
	return total
}
