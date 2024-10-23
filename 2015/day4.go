package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func getMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func findLowestHashWithFiveZeroes(s string) int {
	for i := 0; ; i++ {
		test := getMD5(s+strconv.Itoa(i))
		// if i%10000 == 0 {
		// 	fmt.Printf("Test string: %v\n", s+strconv.Itoa(i))
		// 	fmt.Printf("Testing on i= %v\n", i)
		// 	fmt.Printf("Test result: %v\n", test)
		// }
		if string(test[:5]) == "00000" {
			return i
		}
	}
}

func findLowestHashWithSixZeroes(s string) int {
	for i := 0; ; i++ {
		test := getMD5(s+strconv.Itoa(i))
		// if i%10000 == 0 {
		// 	fmt.Printf("Test string: %v\n", s+strconv.Itoa(i))
		// 	fmt.Printf("Testing on i= %v\n", i)
		// 	fmt.Printf("Test result: %v\n", test)
		// }
		if string(test[:6]) == "000000" {
			return i
		}
	}
}
func day4() {
	data := "iwrupvqb"
	lowest := findLowestHashWithFiveZeroes(data)
	fmt.Println(lowest)
	lowest2 := findLowestHashWithSixZeroes(data)
	fmt.Println(lowest2)
}
