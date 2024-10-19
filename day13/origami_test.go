package day13

import (
	"reflect"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	gotCoords, gotFolds := readDataFromFile("sample_input.txt")
	wantCoords :=
	wantFolds := 

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
