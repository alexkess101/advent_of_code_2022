package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileStructure := NewItem(Dir)
	var counter int64
	var totalSum int64
	counter = 0
	totalSum = 0
	r, err := os.Open("input.txt")
	defer r.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewReader(r)
	Analyzer(fileStructure, scanner)
	Counter(fileStructure, &counter)
	TotalSumOfSizeLessThanOrEqualTo(100000, fileStructure, &totalSum)
	log.Println(totalSum)
}

type FileType string

const (
	File FileType = "FILE"
	Dir           = "DIR"
)

type Item struct {
	t        FileType
	name     string
	children map[string]*Item
	size     int64
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

func ProcessItems(dir *Item, scanner *bufio.Reader) {
	for {
		command, _, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}
		line := Parse(string(command))
		if line[0] == "dir" {
			newDir := NewItem(Dir)
			dir.children[line[1]] = newDir
		} else {
			file := NewItem(File)
			size, err := strconv.ParseInt(line[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			file.size = size
			file.name = line[1]
			dir.children[line[1]] = file
		}
		peek, _ := scanner.Peek(1)
		if string(peek) == "$" {
			break
		}
	}
}

func Analyzer(dir *Item, scanner *bufio.Reader) {
	for {
		command, _, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}
		line := Parse(string(command))
		dir.t = Dir

		if line[1] == "ls" {
			ProcessItems(dir, scanner)
		} else {
			if line[2] != ".." {
				Analyzer(dir.children[line[2]], scanner)
			} else {
				return
			}
		}
	}
}

func Counter(dir *Item, counter *int64) {
	for _, item := range dir.children {
		if item.t == File {
			*counter += item.size
		} else {
			Counter(item, counter)
		}
	}
	dir.size = *counter
}

func TotalSumOfSizeLessThanOrEqualTo(size int64, dir *Item, totalSize *int64) {
	for _, item := range dir.children {
		if item.t == Dir && item.size <= size {
			*totalSize += item.size
			TotalSumOfSizeLessThanOrEqualTo(size, item, totalSize)
		}
	}
}
