package day09

import (
	"reflect"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	got := readDataFromFile("sample_input.txt")
	want := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestFindMinima(t *testing.T) {
	got := FindMinima("sample_input.txt")
	want := 15

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
