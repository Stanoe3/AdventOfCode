package main

import "testing"

func TestNewCalcHouses(t *testing.T) {
	test := "^>v<"
	got := newCalcHouses(test)
	if got != 3 {
		t.Errorf("Expected 3 but got %v", got)
	}
}