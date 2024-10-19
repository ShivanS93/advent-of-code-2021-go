package day16

import (
	"io/ioutil"
	"strings"
)

// read file as hexadecimal
func readDataFromFile(file string) ([]string, error) {
	content, err := ioutil.ReadFile(file)
	contentStr := string(content)
	return strings.Split(contentStr, ""), err
}

// convert hexadecimal into binary as a string
func convertHexadecimalBinary(hexString []string) string {
	hexToBinaryMap := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	var binaryOutput string
	for _, hexChar := range hexString {
		binaryChar := hexToBinaryMap[string(hexChar)]
		binaryOutput += binaryChar
	}
	return binaryOutput
}

// first three is packet version
// second three packet ID
// packet ID 4 = literal
// literal has a single binary number

// for non packet ID 6, operator
// binary number after version and ID is length

// packet ID other than 4 is an operator
func Decoder(hexString []string) int {
	binaryStream := convertHexadecimalBinary(hexString)
	var versionNumberSum int
	// literal
	for i, binary := range binaryStream {
		packetVersion := binaryStream[i : i+3] // 110
		// packetType := binaryStream[i+3:i+6]   // 100

		// packetVersion 6
		if packetVersion == "110" {
			versionNumberSum += 6
			// store packet
			var packetStream []string
			for j, binaryP := range binaryStream[6:] {
				if binaryP == "1" {
					j + 5
				}
			}
		}

	}

	return versionNumberSum
}
