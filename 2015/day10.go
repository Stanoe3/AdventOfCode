package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func lookAndSay(s string) string {
	var sb strings.Builder
	dupes := 1
	for i := range s {
		if i == len(s)-1 {
			sb.WriteString(strconv.Itoa(dupes) + string(s[i]))
			break
		}
		if s[i] == s[i+1] {
			dupes += 1
		} else {
			sb.WriteString(strconv.Itoa(dupes) + string(s[i]))
			dupes = 1
		}

	}
	return sb.String()
}

func lookAndSayNTimes(s string, n int) string {
	for i := 0; i < n; i++ {
		s = lookAndSay(s)
	}
	return s
}

func day10() {
	data := "1321131112"
	look := lookAndSayNTimes(data, 50)
	fmt.Printf("There are %v characters in the string", utf8.RuneCountInString(look))
}
