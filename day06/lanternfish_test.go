package day06

import (
	"reflect"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	got := readDataFromFile("sample_input.txt")
	// [0 1 1 2 1 0 0 0]
	want := [9]int{0, 1, 1, 2, 1, 0, 0, 0}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestSimulateFish(t *testing.T) {
	got := SimulateFish("sample_input.txt", 256)
	// want := 5934 // 80 days
	want := 26984457539 // 256 days

	if got != want {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
