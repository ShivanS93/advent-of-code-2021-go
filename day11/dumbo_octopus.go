package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readDataFromFile(file string) [][]int {
	fileData, _ := os.Open(file)
	scanner := bufio.NewScanner(fileData)
	var output [][]int
	for scanner.Scan() {
		strSlice := strings.SplitAfter(scanner.Text(), "")
		var intSlice []int
		for _, str := range strSlice {
			n, _ := strconv.Atoi(str)
			intSlice = append(intSlice, n)
		}
		output = append(output, intSlice)
	}
	return output
}

func CheckFlashes(file string) int {
	octopusState := readDataFromFile(file)

	// add 1 to adjacents cells
	//	diagnols, ups, downs
	// check for 9s
	// if left
	// if left-top
	// if top
	// if right top
	// if right
	// if right bottom
	// if left bottom
	// if bottom
	var flashCount int

	for i, row := range octopusState {
		for j, n := range row {
			if n == 9 {
				flashCount++
				if (i == 0) && (j == 0) { // left-top hand side
					octopusState[i][j]++
					if 
				}
			}
		}
	}
}

func checkFlash(pos_x int, pos_y int, octopusState [][]int) {
	if num == 9 {
		newNum := 0
		// adjacent numbers
	}
}
