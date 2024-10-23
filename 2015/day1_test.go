package main

import "testing"

func TestFindFloor(t *testing.T) {
	test := "(((((((())"
	got := findFloor(test)
	if got != 6 {
		t.Errorf("Expected 6, got %v", got)
	}
}

func TestFindBasementEntry(t *testing.T) {
	test:= "()())"
	got := findBasementEntry(test)
	if got != 5 {
		t.Errorf("Expected 5, got %v", got)
	}
}
