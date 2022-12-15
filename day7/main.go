package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	content, _ := os.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")

	dirTotals := MapValuesToDirectory(ProcessInput(input))
	log.Println(SizeOfDirectoryToCleanUpSpace(dirTotals))
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
			dir = filepath.Join(dir, line[2])
		} else if strings.HasPrefix(item, "dir ") {
			continue
		} else {
			key = append(key, fmt.Sprintf("%s %s", dir, line[0]))
		}
	}

	return key
}

func RemoveDir(dir string) string {
	return filepath.Join(filepath.Dir(dir))
}

func MapValuesToDirectory(key []string) (directories map[string]int) {
	directories = make(map[string]int)
	for _, line := range key {
		sections := strings.Split(line, " ")
		path := sections[0]
		value, err := strconv.Atoi(sections[1])
		if err != nil {
			log.Fatal(err)
		}
		directories[path] += value
		for path != "/" {
			path = filepath.Dir(path)
			directories[path] += value
		}
	}

	return directories
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

func GetValueOfMainDirectory(dirTotals map[string]int) int {
	return dirTotals["/"]
}

func SizeOfDirectoryToCleanUpSpace(dirTotals map[string]int) int {
	unusedSpace := TotalCapacity - GetValueOfMainDirectory(dirTotals)
	mustDeleteAtLease := SizeOfUpdate - unusedSpace
	var candidate int
	candidate = 0xFFFFFFFF
	for _, dirValue := range dirTotals {
		if dirValue >= mustDeleteAtLease && dirValue < candidate {
			candidate = dirValue
		}
	}

	return candidate
}

const (
	TotalCapacity = 70_000_000
	SizeOfUpdate  = 30_000_000
)
