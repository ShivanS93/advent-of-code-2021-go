package day16

import (
	"reflect"
	"testing"
)

func TestReadDataFromFile(t *testing.T) {
	got, err := readDataFromFile("sample_input.txt")
	want := []string{"D", "2", "F", "E", "2", "8"}

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestConvertHexadecimalBinary(t *testing.T) {
	input, err := readDataFromFile("sample_input.txt")
	got := convertHexadecimalBinary(input)
	want := "110100101111111000101000"

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}

func TestDecoderLiteral(t *testing.T) {
	input, err := readDataFromFile("sample_input.txt")
	got := Decoder(input)
	want := 6

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("Got: %v \n Want: %v", got, want)
	}
}
