package day12

import (
	"bufio"
	"os"
	"strings"
)

func readDataFromFile(file string) [][]string {
	fileData, _ := os.Open(file)
	scanner := bufio.NewScanner(fileData)
	var output [][]string
	for scanner.Scan() {
		output = append(output, strings.Split(scanner.Text(), "-"))
	}
	return output
}

// create a map key is point, and value are linked 'nodes'

func setUpNodes(file string) map[string][]string {
	links := readDataFromFile(file)
	output := make(map[string][]string)
	for _, link := range links {
		_, ok := output[link[0]]
		if !ok {
			output[link[0]] = make([]string, 0)
		}
		output[link[0]] = append(output[link[0]], link[1])
		_, ok2 := output[link[1]]
		if !ok2 {
			output[link[1]] = make([]string, 0)
		}
		output[link[1]] = append(output[link[1]], link[0])
	}
	return output
}

// traverse the map
