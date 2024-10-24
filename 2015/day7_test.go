package main

import (
	"fmt"
	"testing"
)

func TestExecuteOnceInstructions(t *testing.T) {
	testInstructions := [][]string{{"a", "AND", "b", "c"}, {"b", "NOT", "", "a"}, {"10", "NOT", "", "b"}}
	wireMap := map[string]uint16{}
	part2 := false
	gotMap, gotInstr := executeOnceInstructions(testInstructions, wireMap, part2)
	fmt.Println(gotMap)
	fmt.Println(gotInstr)
	if len(gotInstr) != 2 {
		t.Errorf("Expected 2 instructions to fail, but got %d", len(gotInstr))
	}

}
