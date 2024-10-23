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

func executeInstructions(g [1000][1000]bool, instructions [][]string) [1000][1000]bool {
	for _, instr := range instructions {
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
	return g
}

func generateLightsGrid() [1000][1000]bool {
	grid := [1000][1000]bool{}
	return grid
}

func turnOn(g [1000][1000]bool, p []int) [1000][1000]bool {
	g[p[0]][p[1]] = true
	return g
}

func turnOff(g [1000][1000]bool, p []int) [1000][1000]bool {
	g[p[0]][p[1]] = false
	return g
}

func toggle(g [1000][1000]bool, p []int) [1000][1000]bool {
	if g[p[0]][p[1]] == true {
		g[p[0]][p[1]] = false
	} else {
		g[p[0]][p[1]] = true
	}

	return g
}

func turnOnRange(g [1000][1000]bool, p [2][]int) [1000][1000]bool {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			g = turnOn(g, []int{row, col})
		}
	}
	return g
}

func turnOffRange(g [1000][1000]bool, p [2][]int) [1000][1000]bool {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			g = turnOff(g, []int{row, col})
		}
	}
	return g
}

func toggleRange(g [1000][1000]bool, p [2][]int) [1000][1000]bool {
	start := p[0]
	end := p[1]
	for row := start[0]; row <= end[0]; row++ {
		for col := start[1]; col <= end[1]; col++ {
			g = toggle(g, []int{row, col})
		}
	}
	return g
}

func day6() {
	data := ReadFile("./inputs/day6.txt")
	splitData := splitDataByLine(data)
	instructions := getInstructions(splitData)
	grid := generateLightsGrid()
	grid = executeInstructions(grid, instructions)
	fmt.Println(grid)
}
