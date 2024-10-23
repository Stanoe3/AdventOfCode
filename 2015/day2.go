package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitLineToArray(line string) []string {
	return strings.Split(line, "x")

}

func splitAllLinesToArrays(lines []string) [][]string {
	newArray := [][]string{}
	for _, line := range lines {
		newArray = append(newArray, splitLineToArray(line))
	}
	return newArray
}

func convertStringArrayToInt(array []string) []int {
	intArr := []int{}
	for _, str := range array {
		if str == "" {
			continue
		}
		newInt, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, newInt)
	}
	return intArr
}

func convertStringArraysToInts(arrays [][]string) [][]int {
	newIntArrays := [][]int{{}}
	for _, arr := range arrays {
		intArr := convertStringArrayToInt(arr)
		newIntArrays = append(newIntArrays, intArr)
	}
	return newIntArrays
}

func calculatePresent(dimensions []int) int {
	l := dimensions[0]
	w := dimensions[1]
	h := dimensions[2]

	smallest := min(l, w, h)
	highest := max(l, w, h)
	second := l + w + h - smallest - highest

	return 2*l*w + 2*w*h + 2*h*l + smallest*second
}

func calculateAllPresents(dimensions [][]int) int {
	total := 0
	for _, dimension := range dimensions {
		total += calculatePresent(dimension)
	}
	return total
}

func day2() {
	data := ReadFile("./inputs/day2.txt")
	splitLines := splitDataByLine(data)
	arrays := splitAllLinesToArrays(splitLines)
	intArrays := convertStringArraysToInts(arrays)[1:len(arrays)]
	fmt.Println(intArrays[0])
	fmt.Println(intArrays[len(intArrays)-1])
	fmt.Println(calculateAllPresents(intArrays))

}
