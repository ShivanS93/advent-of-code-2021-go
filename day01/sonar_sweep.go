// day 01

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// problemOne()
	problemTwo()
}

func problemOne() {
	// opening file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	counter := -1 // set to -1 because counter will be off by one
	var prev int64
	prev = 0
	// iterating through lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		if prev < num {
			counter = counter + 1
		}
		prev = num
	}
	fmt.Println("Answer to problem one: ")
	fmt.Println(counter)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func problemTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// turning text into slice of intergers
	scanner := bufio.NewScanner(file)
	var scans []int64
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		scans = append(scans, num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// turning rolling sum into another slice
	var rolling_values []int64
	for i, val := range scans {
		if i < len(scans)-2 {
			sum := val + scans[i+1] + scans[i+2]
			rolling_values = append(rolling_values, sum)
		}
	}

	counter := 0 // count the number of increases
	for i, val := range rolling_values {
		if i > 0 {
			if val > rolling_values[i-1] {
				counter++
			}
		}
	}
	fmt.Println("Answer to problem two: ")
	fmt.Println(counter)
}
