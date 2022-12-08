package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func main() {
	//fileStructure := make(map[string]Item)
	//r, err := os.Open("input.txt")
	//defer r.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//scanner := bufio.NewScanner(r)
	//scanner.Split(bufio.ScanLines)
	//next := scanner.Scan
	//text := scanner.Text
	//
	//next()

}

type FileType string

const (
	File FileType = "FILE"
	Dir           = "DIR"
	None          = "NONE"
)

type Item struct {
	t        FileType
	name     string
	children map[string]*Item
	size     int
}

func NewItem(t FileType) *Item {
	return &Item{
		t:        t,
		name:     "",
		children: make(map[string]*Item),
		size:     0,
	}
}

func Parse(s string) []string {
	return strings.Split(s, " ")
}

func ProcessItems(dir *Item, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := Parse(scanner.Text())
		if line[0] == "$" {
			break
		}
		if line[0] == "dir" {
			newDir := NewItem(Dir)
			dir.children[line[1]] = newDir
		} else {
			file := NewItem(File)
			size, err := strconv.Atoi(line[0])
			if err != nil {
				log.Fatal(err)
			}
			file.size = size
			file.name = line[1]
			dir.children[line[1]] = file
		}
	}
}

func cdCommand(dir *Item, scanner *bufio.Scanner) {
	scanner.Scan()
	command := Parse(scanner.Text())
	dir.t = Dir

	if command[0] == "$" {
		if command[1] == "ls" {
			scanner.Scan()
			ProcessItems(dir, scanner)
		} else {
			if command[1] == ".." {
				scanner.Scan()
				return
			} else {
				cdCommand(dir.children[command[1]], scanner)
			}
		}
	}
}
