package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	content, _ := os.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")

	dirTotals := MapValuesToDirectory(ProcessInput(input))
	log.Println(GetTotalCount(dirTotals))
}

func ProcessInput(s []string) []string {
	var key []string
	var dir string
	for _, item := range s {
		line := strings.Split(item, " ")

		if item == "$ cd /" {
			dir = "/"
		} else if item == "$ ls" {
			continue
		} else if item == "$ cd .." {
			dir = RemoveDir(dir)
		} else if strings.HasPrefix(item, "$ cd ") {
			dir += fmt.Sprintf("/%s", line[2])
		} else if strings.HasPrefix(item, "dir ") {
			continue
		} else {
			key = append(key, fmt.Sprintf("%s %s", dir, line[0]))
		}
	}

	return key
}

func RemoveDir(dir string) string {
	return dir[:strings.LastIndex(dir, "/")]
}

func MapValuesToDirectory(key []string) map[string]int {
	dirTotals := make(map[string]int)
	for _, item := range key {
		line := strings.Split(item, " ")

		value, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		for _, path := range strings.Split(line[0], "/") {
			if path != "" {
				dirTotals[path] += value
			}
		}
	}

	return dirTotals
}

func GetTotalCount(dirTotals map[string]int) int {
	totalValue := 0

	for _, item := range dirTotals {
		if item <= 100000 {
			totalValue += item
		}
	}

	return totalValue
}
