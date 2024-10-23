package main

import "fmt"

func hasThreeVowels(s string) bool {
	vowelsMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0
	for _, char := range s {
		_, contains := vowelsMap[char]
		if contains {
			count++
		}
	}
	if count >= 3 {
		return true
	} else {
		return false
	}
}

func hasDouble(s string) bool {
	for index := range s[:len(s)-1] {
		if s[index] == s[index+1] {
			return true
		}
	}
	return false
}

func notBadString(s string) bool {
	for index := range s[:len(s)-1] {
		switch s[index] {
		case 'a':
			if s[index+1] == 'b' {
				return false
			}
		case 'c':
			if s[index+1] == 'd' {
				return false
			}
		case 'p':
			if s[index+1] == 'q' {
				return false
			}
		case 'x':
			if s[index+1] == 'y' {
				return false
			}
		}
	}
	return true
}

func isNice(s string) bool {
	if hasThreeVowels(s) && hasDouble(s) && notBadString(s) {
		return true
	}
	return false
}

func countNice(strs []string) int {
	count := 0
	for _, str := range strs {
		if isNice(str) {
			count++
		}
	}
	return count
}

// --------- Part 2 functions ---------
func hasSandwich(s string) bool {
	for index := range s[:len(s)-2] {
		if s[index] == s[index+2] {
			return true
		}
	}
	return false
}

func hasDoublePair(s string) bool {
	pairMap := map[string]bool{}
	duplicate := false
	for index := range s[:len(s)-1] {
		testString := s[index : index+2]
		if duplicate {
			if testString[0] == testString[1] {
				duplicate = false
				continue
			} else {
				duplicate = false
			}

		}
		if testString[0] == testString[1] {
			duplicate = true
		}
		_, found := pairMap[testString]
		if found {
			return true
		} else {
			pairMap[testString] = true
		}
	}
	return false
}

func newIsNice(s string) bool {
	if hasSandwich(s) && hasDoublePair(s) {
		return true
	}
	return false
}

func newCountNice(strs []string) int {
	count := 0
	for _, str := range strs {
		if newIsNice(str) {
			count++
		}
	}
	return count
}

func day5() {
	data := ReadFile("./inputs/day5.txt")[:len(ReadFile("./inputs/day5.txt"))-1]
	splitData := splitDataByLine(data)
	nice := countNice(splitData)
	fmt.Printf("There are %v nice strings\n", nice)
	newNice := newCountNice(splitData)
	fmt.Printf("There are now %v nice strings\n", newNice)

}
