package day04

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// a way to parse the draw data
func OpenDrawData(file string) ([]int, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	str_values := strings.Split(string(content), ",")
	var values []int
	for _, value := range str_values {
		str_value, _ := strconv.Atoi(value)
		values = append(values, str_value)
	}

	return values, err
}

// open and parse the bingo data
// slice for all the bingo boards
// nx5 int arrays
func OpenBingoData(file string) ([][5]int, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var values [][5]int
	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			line_str_slice := strings.Fields(line_str) // 22 13 17 11 0
			var line_slice [5]int
			for i, str_num := range line_str_slice {
				num, _ := strconv.Atoi(str_num)
				line_slice[i] = num
			}
			values = append(values, line_slice)
		}
	}
	return values, scanner.Err()
}

func CheckBingo(bingoboard [][5]int) bool {
	// scan row for -1
	complete := [5]int{-1, -1, -1, -1, -1}
	for _, row := range bingoboard {
		if reflect.DeepEqual(row, complete) {
			return true // complete row
		}
	}
	// scan columns
	var col [5]int
	for i := 0; i < 5; i++ {
		for j, row := range bingoboard {
			col[j] = row[i]
		}
		if reflect.DeepEqual(col, complete) {
			return true // complete column
		}
	}
	return false
}

func WinnerBingo(file string, draw_file string) int {
	draw, _ := OpenDrawData(draw_file)
	bingo, _ := OpenBingoData(file)

	// look at 1 board 5 x 5, replace with -1
	// write a checker?
	draw_counter := 0 // number of draws
	var bingoboard [][5]int
	var board_count int = len(bingo) / 5 // number of boards
	for _, lot := range draw {
		// scan the board and replaces matches
		draw_counter++
		for j, row := range bingo {
			for k, num := range row {
				if num == lot {
					bingo[j][k] = -1 // sets match to -1
				}
			}
		}

		if draw_counter >= 5 {
			sum := 0
			// scan every board, split bingo into 5s and pass through function
			for i := 0; i < board_count; i++ {
				bingoboard = bingo[5*i : 5*i+5]
				if CheckBingo(bingoboard) {
					// match, stop and calculate!
					// find non matches
					for _, row := range bingoboard {
						for _, num := range row {
							if num != -1 {
								sum = sum + num
							}
						}
					}
					product := sum * lot
					return product
				}
			}
		}
	}
	return -1 // no matches at all
}
