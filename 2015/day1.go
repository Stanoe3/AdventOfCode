package main

import (
	"fmt"
)

func findFloor(data string) int {
	floor := 0
	for _, v := range data {
		if v == 0 {
			continue
		}
		if v == '(' {
			floor++
		} else if v == ')' {
			floor--
		}
	}
	return floor
}

func findBasementEntry(data string) int {
	floor := 0
	for index, v := range data {
		if v == 0 {
			continue
		}
		if v == '(' {
			floor++
		} else if v == ')' {
			floor--
		}
		if floor == -1 {
			return index+1
		}
	}
	return 0
}

func day1() {
	data := ReadFile("./inputs/day1.txt")
	floorFound := findFloor(data)
	basementFound := findBasementEntry(data)
	fmt.Printf("Final floor found: %d\n", floorFound)
	fmt.Printf("Basement entered at : %d\n", basementFound)
}
