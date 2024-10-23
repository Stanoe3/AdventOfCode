package main

import "fmt"

type Coord struct {
	x, y int
}

func calculateNumberOfHouses(data string) int {
	pointer := Coord{0, 0}
	houseMap := make(map[Coord]bool)
	houseMap[pointer] = true
	for _, move := range data {
		switch move {
		case '^':
			pointer.y++
		case 'v':
			pointer.y--
		case '<':
			pointer.x--
		case '>':
			pointer.x++
		}
		houseMap[pointer] = true
	}
	return len(houseMap)
}

func newCalcHouses(data string) int {
	pointer := Coord{0, 0}
	roboPointer := Coord{0, 0}
	houseMap := make(map[Coord]bool)
	houseMap[pointer] = true
	for i, move := range data {
		switch move {
		case '^':
			if i%2 == 0{
				pointer.y++
			} else {roboPointer.y++}
		case 'v':
			if i%2 == 0{
				pointer.y--
			} else {roboPointer.y--}
		case '<':
			if i%2 == 0{
				pointer.x--
			} else {roboPointer.x--}
		case '>':
			if i%2 == 0{
				pointer.x++
			} else {roboPointer.x++}
		}
		houseMap[pointer] = true
		houseMap[roboPointer] = true
	}
	return len(houseMap)
}

func day3() {
	data := ReadFile("./inputs/day3.txt")
	houses := calculateNumberOfHouses(data)
	fmt.Printf("Number of Houses: %d\n", houses)
	houses2 := newCalcHouses(data)
	fmt.Printf("New number of houses: %d\n", houses2)
}
