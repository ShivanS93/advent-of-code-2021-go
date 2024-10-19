package day03

import "testing"

func TestCalculatePowerConsumption(t *testing.T) {
	got := CalculatePowerConsumption("sample_input.txt")
	// got := 199 // force fail
	// got := 198 // force pass
	var want int64
	want = 198

	if got != want {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}

func TestCalculateOxygenCO2(t *testing.T) {
	got := CalculateOxygenCO2("sample_input.txt")
	var want int64
	want = 230

	if got != want {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}
