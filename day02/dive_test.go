package day02

import "testing"

func TestProductDepthHorizontalPosition(t *testing.T) {
	got, _ := ProductDepthHorizontalPosition("sample_input_01.txt")
	want := 150

	if got != want {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}

func TestProductDepthAimHorizontalPosition(t *testing.T) {
	got, _ := ProductDepthAimHorizontalPosition("sample_input_01.txt")
	//	got := 901 // force fail
	// got := 900 // force pass
	want := 900

	if got != want {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}
