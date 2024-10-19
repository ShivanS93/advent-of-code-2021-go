package day07

import (
	"reflect"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	got := readDataFromFile("sample_input.txt")
	want := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestFindMax(t *testing.T) {
	inputData := readDataFromFile("sample_input.txt")

	got := findMax(inputData)
	want := 16

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestFindMoves(t *testing.T) {
	got := FindMoves("sample_input.txt")
	want := 37

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestFindMovesExpensive(t *testing.T) {
	got := FindMovesExpensive("sample_input.txt")
	want := 168

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
