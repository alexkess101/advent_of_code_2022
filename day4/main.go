package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func IsContained(inputA, inputB string) bool {
	a := strings.Split(inputA, "-")
	b := strings.Split(inputB, "-")
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])

	if (a1-b1 <= 0 && a0-b0 >= 0) || (b1-a1 <= 0 && b0-a0 >= 0) {
		return true
	} else {
		return false
	}
}

func DoesOverlap(inputA, inputB string) bool {
	doesOverlap := false
	legend := make(map[int]int)
	a := strings.Split(inputA, "-")
	b := strings.Split(inputB, "-")
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])

	for i := a0; i <= a1; i++ {
		legend[i] = 0
	}
	for i := b0; i <= b1; i++ {
		if value, ok := legend[i]; ok {
			legend[i] = value + 1
		}
	}

	for _, value := range legend {
		if value > 0 {
			doesOverlap = true
		}
	}

	return doesOverlap
}

func main() {
	content, err := os.ReadFile("input.txt")
	count := 0
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range strings.Split(string(content), "\n") {
		inputs := strings.Split(item, ",")
		if DoesOverlap(inputs[0], inputs[1]) {
			count += 1
		}
	}
	log.Println(count)
}
