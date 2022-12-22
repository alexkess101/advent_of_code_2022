package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := 0
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	cycles := ParseInstructions(strings.Split(string(content), "\n"))
	xValues := CycleThrough(cycles, []int{20, 60, 100, 140, 180, 220})

	for key, value := range xValues {
		answer += key * value
	}

	log.Println(answer)
}

func ParseInstructions(input []string) (cycles []int) {
	for _, item := range input {
		c := strings.Split(item, " ")
		switch c[0] {
		case "noop":
			cycles = append(cycles, 0)
		case "addx":
			value, _ := strconv.Atoi(c[1])
			cycles = append(cycles, 0)
			cycles = append(cycles, value)
		}
	}

	return cycles
}

func CycleThrough(cycles, window []int) map[int]int {
	X := 1
	xValues := make(map[int]int)
	for i := 0; i < len(cycles); i++ {
		if isMatchFound(i+1, window) {
			xValues[i+1] = X
		}

		X += cycles[i]
	}

	return xValues
}

func isMatchFound(value int, list []int) bool {
	for _, item := range list {
		if value == item {
			return true
		}
	}
	return false
}
