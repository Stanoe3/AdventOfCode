package main

import (
	"fmt"
	"strconv"
	"strings"
)

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func generateDistMap(lines []string) map[string]int {
	distMap := map[string]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineArr := strings.Split(line, " ")
		a := lineArr[0]
		b := lineArr[2]
		dist, err := strconv.Atoi(lineArr[4])
		if err != nil {
			panic(err)
		}
		distMap[a+b] = dist
	}
	return distMap
}

func getDistance(a string, b string, distMap map[string]int) int {
	_, found := distMap[a+b]
	if found {
		return distMap[a+b]
	} else {
		return distMap[b+a]
	}
}

func getDistances(towns []string, distMap map[string]int) int {
	dist := 0
	for index := range len(towns) - 1 {
		dist += getDistance(towns[index], towns[index+1], distMap)
	}
	return dist
}

func getShortestDistance(townPerms [][]string, distMap map[string]int) (int, []string) {
	shortest := 1000000000
	shortestTowns := []string{}
	for _, towns := range townPerms {
		dist := getDistances(towns, distMap)
		if dist < shortest {
			shortest = dist
			shortestTowns = towns
		}
	}
	return shortest, shortestTowns
}

func getLongestDistance(townPerms [][]string, distMap map[string]int) (int, []string) {

	longest := 0
	longestTowns := []string{}
	for _, towns := range townPerms {
		dist := getDistances(towns, distMap)
		if dist > longest {
			longest = dist
			longestTowns = towns
		}
	}
	return longest, longestTowns
}

func getTowns(lines []string) []string {
	townSet := map[string]bool{"Straylight": true}
	for _, line := range lines {
		if line == "" {
			continue
		}
		townSet[strings.Split(line, " ")[0]] = true
	}
	towns := []string{}
	for town := range townSet {
		towns = append(towns, town)
	}
	return towns
}

func day9() {
	data := ReadFile("./inputs/day9.txt")
	lines := splitDataByLine(data)
	towns := getTowns(lines)
	townsPerms := permutations(towns)
	distMap := generateDistMap(lines)
	shortest, shortestTowns := getShortestDistance(townsPerms, distMap)
	fmt.Printf("The shortest distance is %d\n", shortest)
	fmt.Printf("Found by going %v", shortestTowns)

	longest, longestTowns := getLongestDistance(townsPerms, distMap)
	fmt.Printf("The longest distance is %d\n", longest)
	fmt.Printf("Found by going %v", longestTowns)
}
