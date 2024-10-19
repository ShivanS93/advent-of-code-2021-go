package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Read data from .txt file line by line
func ReadDataFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values, scanner.Err()
}

func ReadDataFromFileInt(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var values []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		values = append(values, num)
	}
	return values, scanner.Err()
}
