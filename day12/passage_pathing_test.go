package day12

import (
	"reflect"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	got := readDataFromFile("sample_input.txt")
	want := [][]string{
		{"fs", "end"},
		{"he", "DX"},
		{"fs", "he"},
		{"start", "DX"},
		{"pj", "DX"},
		{"end", "zg"},
		{"zg", "sl"},
		{"zg", "pj"},
		{"pj", "he"},
		{"RW", "he"},
		{"fs", "DX"},
		{"pj", "RW"},
		{"zg", "RW"},
		{"start", "pj"},
		{"he", "WI"},
		{"zg", "he"},
		{"pj", "fs"},
		{"start", "RW"},
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestSetUpNodes(t *testing.T) {
	got := setUpNodes("sample_input.txt")
	want := make(map[string][]string)
	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
