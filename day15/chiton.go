package day15

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// read the file
func readDataFromFile(file string) ([][]int, error) {
	content, err := os.Open(file)
	scanner := bufio.NewScanner(content)
	var intSlice [][]int
	for scanner.Scan() {
		strLine := strings.Split(scanner.Text(), "")
		var intLine []int
		for _, s := range strLine {
			var n int
			n, err = strconv.Atoi(s)
			intLine = append(intLine, n)
		}
		intSlice = append(intSlice, intLine)
	}
	return intSlice, err
}

func PathFinder(file string) (int, error) {
	riskMap, err := readDataFromFile(file)
	// start at bottom right and work backwards
	// iterate through the current list
	// also look at square in bottom position
	// riskmap[row][col]
	var currentRow int = 0
	var currentCol int = 0
	var pathScore int

	maxMapRows := len(riskMap) - 1
	maxMapCols := len(riskMap[0]) - 1
	nextRowValue := riskMap[currentRow+1][currentCol]
	nextColValue := riskMap[currentRow][currentCol+1]

	for {
		// at bottom right "end"
		if (currentRow == maxMapRows) && (currentCol == maxMapCols) {
			break
		} else if currentRow == maxMapRows {
			currentCol++
		} else if currentCol == maxMapCols {
			currentRow++
		} else if nextRowValue < nextColValue {
			currentRow++
		} else {
			currentCol++
		}
		pathScore += riskMap[currentRow][currentCol]
	}
	return pathScore, err
}
