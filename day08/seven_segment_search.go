package day08

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// extract and parse .txt file
func readDataFromFile(file string) [][]string {
	// turn string into string array by splting by space
	fileData, _ := os.Open(file)
	scanner := bufio.NewScanner(fileData)
	var output [][]string
	for scanner.Scan() {
		strSlice := strings.Split(scanner.Text(), " ")
		output = append(output, strSlice)

	}
	return output
}

// read the digits
func DigitReader(file string) int {
	input := readDataFromFile(file)
	// iterate through the slice of strings
	// find the length of the string slice
	// look at last 4 elements
	// count elements if only equal to below numbers
	// 1 segs 2
	// 4 segs 4
	// 7 segs 3
	// 8 segs 7

	var count int
	var observeSlice []string
	for _, strSlice := range input {
		observeSlice = strSlice[len(strSlice)-4:]
		for _, str := range observeSlice {
			if (len(str) == 2) || (len(str) == 4) || (len(str) == 3) || (len(str) == 7) {
				count++
			}
		}
	}
	return count
}

func alphaSort(str string) string {
	r := []rune(str)
	for x := range r {
		y := x + 1
		for y = range r {
			if r[x] < r[y] {
				r[x], r[y] = r[y], r[x]
			}
		}
	}
	var outputStr string
	for _, c := range r {
		outputStr += fmt.Sprintf("%c", c)
	}
	return outputStr
}

// take a string slice and return a slice string mapping the number to the combination of letters.
func mapNumbers(file string) []string {
	input := readDataFromFile(file)
	for _, strSlice := range input {
		// 1 is 1 | 7 is 7
		// 4 is 4 | 8 is 8
		// 2 is
		// 3 (includes 1)
		// 5 is
		// 6
		// 9 (includes
	}
}
