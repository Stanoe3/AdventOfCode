package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInstructions(strs []string) [][]string {
	finalArray := [][]string{}
	for _, instr := range strs {
		instrArr := strings.Split(instr, " ")
		if instrArr[0] == "turn" {
			finalArray = append(finalArray, []string{instrArr[1], instrArr[2], instrArr[3], instrArr[4]})
		} else {
			finalArray = append(finalArray, instrArr)
		}
	}
	return finalArray
}

func executeInstructions(g *[1000][1000]bool, instructions [][]string) {
	for _, instr := range instructions {
		if instr[0] == "" {
			continue
		}
		fmt.Printf("Executing instructions: %v\n", instr)
		start := strings.Split(instr[1], ",")
		end := strings.Split(instr[3], ",")
		startX, err := strconv.Atoi(start[0])
		if err != nil {
			panic(err)
		}
		startY, err := strconv.Atoi(start[1])
		if err != nil {
			panic(err)
		}
		endX, err := strconv.Atoi(end[0])
		if err != nil {
			panic(err)
		}
		endY, err := strconv.Atoi(end[1])
		if err != nil {
			panic(err)
		}
		startArr := []int{startX, startY}
		endArr := []int{endX, endY}

		switch instr[0] {
		case "on":
			turnOnRange(g, [2][]int{startArr, endArr})
		case "off":
			turnOffRange(g, [2][]int{startArr, endArr})
		case "toggle":
			toggleRange(g, [2][]int{startArr, endArr})
		}

	}
}

func generateLightsGrid() [1000][1000]bool {
	grid := [1000][1000]bool{}
	return grid
}

func turnOn(g *[1000][1000]bool, p []int) {
	g[p[0]][p[1]] = true
}

func turnOff(g *[1000][1000]bool, p []int) {
	g[p[0]][p[1]] = false
}

func toggle(g *[1000][1000]bool, p []int) {
	if g[p[0]][p[1]] {
		g[p[0]][p[1]] = false
	} else {
		g[p[0]][p[1]] = true
	}
}

func turnOnRange(g *[1000][1000]bool, p [2][]int) {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			turnOn(g, []int{row, col})
		}
	}
}

func turnOffRange(g *[1000][1000]bool, p [2][]int) {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			turnOff(g, []int{row, col})
		}
	}
}

func toggleRange(g *[1000][1000]bool, p [2][]int) {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			toggle(g, []int{row, col})
		}
	}
}

func countLights(g *[1000][1000]bool) int {
	count := 0
	for i := range g {
		for j := range g {
			if g[i][j] {
				count++
			}
		}
	}
	return count
}

// ================= Part 2 =============================

func newTurnOn(g *[1000][1000]int, p []int) {
	g[p[0]][p[1]] += 1
}

func newTurnOff(g *[1000][1000]int, p []int) {
	g[p[0]][p[1]] -= 1
	if g[p[0]][p[1]] < 0 {
		g[p[0]][p[1]] = 0
	}
}

func newToggle(g *[1000][1000]int, p []int) {
	g[p[0]][p[1]] += 2
}

func newTurnOnRange(g *[1000][1000]int, p [2][]int) {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			newTurnOn(g, []int{row, col})
		}
	}
}

func newTurnOffRange(g *[1000][1000]int, p [2][]int) {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			newTurnOff(g, []int{row, col})
		}
	}
}

func newToggleRange(g *[1000][1000]int, p [2][]int) {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			newToggle(g, []int{row, col})
		}
	}
}

func newExecuteInstructions(g *[1000][1000]int, instructions [][]string) {
	for _, instr := range instructions {
		if instr[0] == "" {
			continue
		}
		fmt.Printf("Executing instructions: %v\n", instr)
		start := strings.Split(instr[1], ",")
		end := strings.Split(instr[3], ",")
		startX, err := strconv.Atoi(start[0])
		if err != nil {
			panic(err)
		}
		startY, err := strconv.Atoi(start[1])
		if err != nil {
			panic(err)
		}
		endX, err := strconv.Atoi(end[0])
		if err != nil {
			panic(err)
		}
		endY, err := strconv.Atoi(end[1])
		if err != nil {
			panic(err)
		}
		startArr := []int{startX, startY}
		endArr := []int{endX, endY}

		switch instr[0] {
		case "on":
			newTurnOnRange(g, [2][]int{startArr, endArr})
		case "off":
			newTurnOffRange(g, [2][]int{startArr, endArr})
		case "toggle":
			newToggleRange(g, [2][]int{startArr, endArr})
		}

	}
}

func generateNewLightsGrid() [1000][1000]int {
	grid := [1000][1000]int{}
	return grid
}

func newCountLights(g *[1000][1000]int) int {
	count := 0
	for i := range g {
		for j := range g {
			count += g[i][j]
		}
	}
	return count
}

func day6() {
	data := ReadFile("./inputs/day6.txt")
	splitData := splitDataByLine(data)
	instructions := getInstructions(splitData)
	grid := generateLightsGrid()
	executeInstructions(&grid, instructions)
	lights := countLights(&grid)
	fmt.Printf("There are %d lights on", lights)

	newGrid := generateNewLightsGrid()
	newExecuteInstructions(&newGrid, instructions)
	newLights := newCountLights(&newGrid)
	fmt.Printf("There are now %d lights on", newLights)
}
