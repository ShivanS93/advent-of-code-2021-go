package day04

import (
	"log"
	"reflect"
	"testing"
)

func TestOpenDrawData(t *testing.T) {
	got, err := OpenDrawData("sample_input_draw.txt")
	//	got := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	want := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	if err != nil {
		log.Fatal(err)
	}

	if reflect.DeepEqual(got, want) != true {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}

func TestOpenBingoData(t *testing.T) {
	got, err := OpenBingoData("sample_input.txt")
	want := [][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	/*
		22 13 17 11 0
		8 2 23 4 24
		21 9 14 16 7
		6 10 3 18 5
		1 12 20 15 19
	*/

	if err != nil {
		log.Fatal(err)
	}

	if reflect.DeepEqual(got, want) != true {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}

func TestCheckBingoLose(t *testing.T) {
	win := [][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	got := CheckBingo(win)
	want := false

	if got != want {
		t.Error("Got:%v \n Want: %v", got, want)
	}
}

func TestCheckBingoWinRow(t *testing.T) {
	win := [][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{-1, -1, -1, -1, -1},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}
	got := CheckBingo(win)
	want := true

	if got != want {
		t.Error("Got:%v \n Want: %v", got, want)
	}
}

func TestCheckBingoWinCol(t *testing.T) {
	win := [][5]int{
		{-1, 13, 17, 11, 0},
		{-1, 2, 23, 4, 24},
		{-1, 9, 14, 16, 7},
		{-1, 10, 3, 18, 5},
		{-1, 12, 20, 15, 19},
	}
	got := CheckBingo(win)
	want := true

	if got != want {
		t.Error("Got:%v \n Want: %v", got, want)
	}
}

func TestWinnerBingo(t *testing.T) {
	got := WinnerBingo("sample_input.txt", "sample_input_draw.txt")
	want := 4512

	if got != want {
		t.Error("Got:%v \n Want: %v", got, want)
	}
}
