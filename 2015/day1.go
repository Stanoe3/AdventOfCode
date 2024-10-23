package main

import (
	"fmt"
	"unicode/utf8"
)

func findFloor(data string) int {
	up := 0
	down := 0
	fmt.Println(utf8.RuneCountInString(data))
	for _, v := range data {
		if v == 0 {
			continue
		}
		if v == '(' {
			up++
		} else if v == ')' {
			down++
		}
	}
	fmt.Println(up)
	fmt.Println(down)
	return up - down
}

func main() {
	data := ReadFile("./inputs/day1.txt")
	floorFound := findFloor(data)
	fmt.Println(floorFound)
}
