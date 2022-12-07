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

func main() {
	byteContent, _ := os.ReadFile("input.txt")
	content := string(byteContent)

	for i := 0; i < len(content)-4; i++ {
		substring := content[i : i+4]
		log.Println(substring)
	}
}
