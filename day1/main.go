package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	var amounts []int
	if err != nil {
		log.Println(err)
	}

	for _, list := range strings.Split(string(content), "\n\n") {
		var number int
		number = 0
		for _, item := range strings.Split(list, "\n") {
			value, _ := strconv.Atoi(item)
			number += value
		}
		amounts = append(amounts, number)
	}

	sort.Ints(amounts)

	var total int
	for _, subTotal := range amounts[len(amounts)-3:] {
		total += subTotal
	}
	log.Println(total)
}
