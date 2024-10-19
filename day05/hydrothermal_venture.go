package day05

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// read file
// turn into an slice of and array
// 0,9 -> 5,9
// [x1 y1 x2 y2]
func readInputFromFile(file string) [][4]int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)
	var output [][4]int
	for scanner.Scan() { // 0,9 -> 5,9
		// string
		// split by space into a slice
		// keep 0 and 2
		// split by comma
		// 0 will be 0 value etc
		strSlice := strings.Fields(scanner.Text()) // ["0,9" "->" "5,9"]
		start := strings.Split(strSlice[0], ",")
		end := strings.Split(strSlice[2], ",")
		strParseItem := append(start, end...)
		var parseItem [4]int
		for i, item := range strParseItem {
			num, _ := strconv.Atoi(item)
			parseItem[i] = num
		}
		output = append(output, parseItem)
	}
	return output
}

// filter horizontal and vertical lines
func filterHVlines(lines [][4]int) [][4]int {
	var output [][4]int
	for _, line := range lines {
		// for a line to be horizontal x1 == x2, vertical y1 == y2
		if (line[0] == line[2]) || (line[1] == line[3]) {
			output = append(output, line)
		}
	}
	return output
}

// give 2D array of a 'field'
// find max x (assume min x is 0)
// find max y (assume min y is 0)
func buildField(lines [][4]int) [][]int {
	maxX := 0
	maxY := 0
	for _, line := range lines {
		if line[0] > maxX { // 0
			maxX = line[0]
		}
		if line[2] > maxX { // 5
			maxX = line[2]
		}
		if line[1] > maxY { // 9
			maxY = line[1]
		}
		if line[3] > maxY { // 9
			maxY = line[3]
		}
	}
	// add +1 to maxX and maxY since 0-9 is 10 spots
	fieldX := make([]int, maxX+1)
	var output [][]int
	for i := 0; i < maxY+1; i++ {
		output = append(output, fieldX)
	}
	return output
}

// draw lines onto the field
// find if horizonal or vertical line
// create cooridinates for all points using endpoints as ranges
func drawField(lines [][4]int) [][]int {
	field := buildField(lines)
	HVLines := filterHVlines(lines)
	start := 0
	end := 0
	var currentField [][]int
	for _, line := range HVLines {
		// if vertical line
		if line[0] == line[2] {
			if line[1] > line[3] {
				start = line[3]
				end = line[1]
			} else {
				start = line[1]
				end = line[3]
			}
			currentField = field[:]
			for i, _ := range currentField {
				if (start <= i) && (i <= end) {
					field[i][line[0]]++
				}
			}
			// if horizontal line
		} else if line[1] == line[3] {
			if line[0] > line[2] {
				start = line[2]
				end = line[0]
			} else {
				start = line[0] // 0
				end = line[2]   // 5
			}
			for i := start; i <= end; i++ {
				field[line[1]][i]++
			}
		}
	}
	return field
}

// previousFieldState := field[:]
// for i, row := range previousFieldState {
// 	for j, _ := range row {
// 		if field[i][j] != 0 {
// 			num := field[i][j]
// 			num++
// 			field[i][j] = num
// 		}
// 	}
// }
