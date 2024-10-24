package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateInstructions(rawInstructions []string) [][]string {
	instructions := [][]string{}
	op1, op2, operator := "", "", ""
	for _, instr := range rawInstructions {

		if len(instr) == 0 {
			continue
		}
		instArr := strings.Split(instr, " -> ")
		operation := instArr[0]
		target := instArr[1]
		operationArr := strings.Split(operation, " ")
		if len(operationArr) == 3 {
			op1 = operationArr[0]
			op2 = operationArr[2]
			operator = operationArr[1]
		} else if len(operationArr) == 2 {
			op1 = operationArr[1]
			op2 = ""
			operator = operationArr[0]
		} else {
			op1 = operationArr[0]
			op2 = ""
			operator = ""
		}

		instructions = append(instructions, []string{op1, operator, op2, target})
	}
	return instructions
}

func execute(op1 uint16, op2 uint16, operator string) uint16 {
	switch operator {
	case "AND":
		return op1 & op2
	case "OR":
		return op1 | op2
	case "LSHIFT":
		return op1 << op2
	case "RSHIFT":
		return op1 >> op2
	}
	fmt.Printf("There was an error executing %v %v %v\n\n\n", op1, operator, op2)
	return 0
}

func executeOnceInstructions(instructions [][]string, wireMap map[string]uint16, part2 bool) (map[string]uint16, [][]string) {
	waitingInstructions := [][]string{}
	for _, instruction := range instructions {
		op1 := instruction[0]
		operator := instruction[1]
		op2 := instruction[2]
		targetWire := instruction[3]
		if targetWire == "b" && part2 {
			continue
		}
		if operator == "" && op2 == "" {
			assignVal, err := strconv.Atoi(instruction[0])
			if err != nil {
				_, found := wireMap[op1]
				if found {
					assignVal = int(wireMap[op1])
				} else {
					waitingInstructions = append(waitingInstructions, instruction)
					continue
				}
			}
			targetWire := instruction[3]
			wireMap[targetWire] = uint16(assignVal)
		} else if op2 == "" {
			_, found := wireMap[op1]
			if found {
				wireMap[targetWire] = ^wireMap[op1]
				continue
			}
			op1Int, err := strconv.Atoi(op1)
			op1UInt := uint16(op1Int)
			// if op1 is a number
			if err == nil {
				wireMap[targetWire] = ^op1UInt
			} else {
				waitingInstructions = append(waitingInstructions, instruction)
			}
		} else {
			_, found1 := wireMap[op1]
			_, found2 := wireMap[op2]
			if found1 && found2 {
				op1UInt := wireMap[op1]
				op2UInt := wireMap[op2]
				wireMap[targetWire] = execute(op1UInt, op2UInt, operator)
				continue
			}

			op1Int, err := strconv.Atoi(op1)
			op1UInt := uint16(op1Int)
			// if op1 is a number
			if err == nil {
				if !found2 {
					waitingInstructions = append(waitingInstructions, instruction)
					continue
				} else {
					op2UInt := uint16(wireMap[op2])
					wireMap[targetWire] = execute(op1UInt, op2UInt, operator)
					continue
				}
			}
			op2Int, err := strconv.Atoi(op2)
			op2UInt := uint16(op2Int)

			// if op 2 is a number
			if err == nil {
				if !found1 {
					waitingInstructions = append(waitingInstructions, instruction)
					continue
				} else {
					op1UInt = uint16(wireMap[op1])
					wireMap[targetWire] = execute(op1UInt, op2UInt, operator)
					continue
				}

			}
			waitingInstructions = append(waitingInstructions, instruction)
		}
	}
	return wireMap, waitingInstructions
}

func executeAllInstructions(instructions [][]string, wireMap map[string]uint16, part2 bool) map[string]uint16 {
	for {
		//fmt.Println(len(instructions))
		wireMap, instructions = executeOnceInstructions(instructions, wireMap, part2)
		if len(instructions) == 0 {
			return wireMap
		}
	}
}

func day7() {
	data := ReadFile("./inputs/day7.txt")
	instructionLines := splitDataByLine(data)
	instructions := generateInstructions(instructionLines)
	fmt.Println(len(instructions))
	wireMap := map[string]uint16{}
	wireMap = executeAllInstructions(instructions, wireMap, false)
	fmt.Printf("The wire signal on a is %v\n", wireMap["a"])
	newWireMap := map[string]uint16{}
	newWireMap["b"] = uint16(3176)
	fmt.Println(newWireMap)
	newWireMap = executeAllInstructions(instructions, newWireMap, true)
	fmt.Printf("The wire signal on a is now: %v\n", newWireMap["a"])

}
