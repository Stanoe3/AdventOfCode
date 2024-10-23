package main

import "testing"

func TestHasSandwich(t *testing.T) {
	test:= "xxyxx"
	got := hasSandwich(test)
	if !got {
		t.Errorf("Expected true for %v but got %v", test, got)
	}
}

func TestHasDoublePair(t *testing.T) {
	test:= "qjhvhtzxzqqjkmpb"
	test2 := "asfasfawfs"
	test3 := "aaa"
	got1 := hasDoublePair(test)
	got2 := hasDoublePair(test2)
	got3 := hasDoublePair(test3)
	if !got1 {
		t.Errorf("Expected true for %v but got %v", test, got1)
	}
	if !got2 {
		t.Errorf("Expected true for %v but got %v",test2, got2)
	}
	if got3 {
		t.Errorf("Expected false for %v but got %v", test3, got3)
	}
}

func TestNewIsNice(t *testing.T) {
	test1:= "qjhvhtzxzqqjkmpb"
	test2:= "xxyxx"
	test3 := "uurcxstgmygtbstg"
	test4 := "ieodomkazucvgmuy"
	got1 := newIsNice(test1)
	got2 := newIsNice(test2)
	got3 := newIsNice(test3)
	got4 := newIsNice(test4)

	if !got1 {
		t.Errorf("Expected true from %v but got %v", test1, got1)
	}
	if !got2 {
		t.Errorf("Expected true from  %v but got %v", test2, got2)
	}
	if got3 {
		t.Errorf("Expected false from %v but got %v", test3, got3)
	}
	if got4 {
		t.Errorf("Expected false from %v but got %v", test4, got4)
	}
}