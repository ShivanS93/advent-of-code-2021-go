package day13

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// read data from file
// return slice of intergers
// return map of folds
func readDataFromFile(file string) ([][]int, map[string]int) {
	fileData, _ := os.Open(file)
	scanner := bufio.NewScanner(fileData)
	var sliceCoords [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && !strings.Contains(line, "fold") {
			strDigits := strings.Split(line, ",")
			var intDigits []int
			for _, d := range strDigits {
				n, _ := strconv.Atoi(d)
				intDigits = append(intDigits, n)
			}
		} else if strings.Contains(line, "fold") {
		}
	}
	return sliceCoords, mapFolds
}
