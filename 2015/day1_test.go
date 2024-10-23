package main

import "testing"

func TestFindFloor(t *testing.T) {
	test := "(((((((())"
	got := findFloor(test)
	if got != 6 {
		t.Errorf("Expected 6, got %v", got)
	}
}
