package main

import (
	"log"
	"os"
)

func HasDuplicates(s string) bool {
	key := make(map[string]int)
	hasDuplicate := false
	for _, c := range s {
		if _, ok := key[string(c)]; ok {
			hasDuplicate = true
		} else {
			key[string(c)] = 0
		}
	}

	return hasDuplicate
}

func GetProcessedPosition(str string, size int) int {
	for i := 0; i < len(str)-size; i++ {
		substring := str[i : i+size]
		if !HasDuplicates(substring) {
			return i + size
		}
	}
	return -1
}

func main() {
	byteContent, _ := os.ReadFile("input.txt")
	content := string(byteContent)

	log.Println(GetProcessedPosition(content, 14))
}
