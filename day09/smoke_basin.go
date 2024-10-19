package day09

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
		var d int
		for _, n := range strSlice {
			d, _ = strconv.Atoi(n)
			intSlice = append(intSlice, d)
		}
		output = append(output, intSlice)
	}
	return output
}

// loop through slice
// check "above", "below", "left", "right" for minima?finder
func FindMinima(file string) int {
	input := readDataFromFile(file)
	var minima [][2]int // index is row x, value is column y
	for i, row := range input {
		for j, num := range row {
			// top
			if i == 0 {
				if input[i+1][j] > num {
					if j == 0 {
						if row[j+1] > num {
							minima = append(minima, [2]int{i, j})
						}
					} else if j == (len(row) - 1) {
						if row[j-1] > num {
							minima = append(minima, [2]int{i, j})
						}
					} else {
						if (row[j+1] > num) || (row[j-1] > num) {
							minima = append(minima, [2]int{i, j})
						}
					}
				}
				// bottom
			} else if i == len(input)-1 {
				if input[i-1][j] > num {
					if j == 0 {
						if row[j+1] > num {
							minima = append(minima, [2]int{i, j})
						}
					} else if j == (len(row) - 1) {
						if row[j-1] > num {
							minima = append(minima, [2]int{i, j})
						}
					} else {
						if (row[j+1] > num) || (row[j-1] > num) {
							minima = append(minima, [2]int{i, j})
						}
					}
					// middle
				} else {
					if (input[i-1][j] > num) || (input[i+1][j] > num) {
						if j == 0 {
							if row[j+1] > num {
								minima = append(minima, [2]int{i, j})
							}
						} else if j == (len(row) - 1) {
							if row[j-1] > num {
								minima = append(minima, [2]int{i, j})
							}
						} else {
							if (row[j+1] > num) || (row[j-1] > num) {
								minima = append(minima, [2]int{i, j})
							}
						}
					}
				}
			}
		}
	}
	var output int
	for _, coord := range minima {
		output += input[coord[0]][coord[0]]
	}

	return output
}
