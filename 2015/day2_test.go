package main

import "testing"

func TestCalculatePresent(t *testing.T) {
	present := []int{1,1,10}
	got := calculatePresent(present)
	if got != 43 {
		t.Errorf("Expected 43, got %d", got)
	}
}


func TestCalculateRibbon(t *testing.T) {
	present := []int{1,1,10}
	got := calculateRibbon(present)
	if got != 14 {
		t.Errorf("Expected 14, got %d", got)
	}
}