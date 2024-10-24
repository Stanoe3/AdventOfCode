package main

import (
	"fmt"
	"strconv"
)

func calculateCodeChars(strings []string) int {
	count := 0
	for _, line := range strings {
		if line == "" {
			continue
		}
		count += len(line)
	}
	return count
}

func calculateMemoryChars(strings []string) int {
	count := 0
	for _, line := range strings {
		if line == "" {
			continue
		}
		newStr, err := strconv.Unquote(line)
		if err != nil {
			panic(err)
		}
		count += len(newStr)
	}
	return count
}

func calculateNewChars(strings []string) int {
	count := 0
	for _, line := range strings {
		if line == "" {
			continue
		}
		count += len(strconv.Quote(line))
	}
	return count
}

func day8() {
	data := ReadFile("./inputs/day8.txt")
	splitData := splitDataByLine(data)
	memChars := calculateMemoryChars(splitData)
	codeChars := calculateCodeChars(splitData)
	fmt.Printf("The number of code characters is : %d\n", codeChars)
	fmt.Printf("The number of memory characters is : %d\n", memChars)
	fmt.Printf("The number difference is : %d\n", codeChars-memChars)
	fmt.Println("\n\n\n\n")
	fmt.Println("Part 2:\n")
	newChars := calculateNewChars(splitData)
	fmt.Printf("The number of code characters is : %d\n", codeChars)
	fmt.Printf("The number of memory characters is : %d\n", newChars)
	fmt.Printf("The number difference is : %d\n", newChars-codeChars)
}
