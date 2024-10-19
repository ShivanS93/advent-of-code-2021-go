package day05

import (
	"reflect"
	"testing"
)

func TestReadInputFromFile(t *testing.T) {
	got := readInputFromFile("sample_input.txt")
	want := [][4]int{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{6, 4, 2, 0},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
		{0, 0, 8, 8},
		{5, 5, 8, 2},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestFilterHVlines(t *testing.T) {
	test_input := readInputFromFile("sample_input.txt")
	got := filterHVlines(test_input)
	want := [][4]int{
		{0, 9, 5, 9},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestBuildField(t *testing.T) {
	test_input := readInputFromFile("sample_input.txt")
	got := buildField(test_input)
	want := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestDrawField(t *testing.T) {
	// test_input := readInputFromFile("sample_input.txt")
	test_input := [][4]int{
		{0, 9, 5, 9},
	}
	got := drawField(test_input)
	want := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
