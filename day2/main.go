package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func generateFirstRoundKey() map[string]interface{} {
	column1 := [][]interface{}{{"A", 1}, {"B", 2}, {"C", 3}}
	column2 := [][]interface{}{{"X", 1}, {"Y", 2}, {"Z", 3}}
	legend := make(map[string]interface{})

	for _, you := range column1 {
		for _, me := range column2 {
			key := fmt.Sprintf("%s %s", you[0], me[0])

			if you[0] == "A" && me[0] == "X" || you[0] == "B" && me[0] == "Y" || you[0] == "C" && me[0] == "Z" {
				//draw
				legend[key] = me[1].(int) + 3
			} else if you[0] == "A" && me[0] == "Y" || you[0] == "B" && me[0] == "Z" || you[0] == "C" && me[0] == "X" {
				//I win
				legend[key] = me[1].(int) + 6
			} else if you[0] == "A" && me[0] == "Z" || you[0] == "B" && me[0] == "X" || you[0] == "C" && me[0] == "Y" {
				//I lose
				legend[key] = me[1]
			} else {
				log.Fatalf("something didn't catch: %s%s", you[0], me[0])
			}

		}
	}

	return legend
}

func generateSecondRoundKey() map[string]int {
	//column := [][]interface{}{{"A", 1}, {"B", 2}, {"C", 3}}
	youList := []string{"A", "B", "C"}
	meList := []string{"X", "Y", "Z"}
	legend := make(map[string]int)
	//key := map[string]int{"X": 1, "Y": 2, "Z": 3}

	for _, you := range youList {
		for _, me := range meList {
			key := fmt.Sprintf("%s %s", you, me)

			if you == "A" {
				if me == "Y" {
					//draw with rock+1
					legend[key] = 4
				} else if me == "X" {
					//lose with scissors+3
					legend[key] = 3
				} else if me == "Z" {
					//win with paper+2
					legend[key] = 8
				} else {
					log.Fatalf("something didn't catch: %s%s", you, me)
				}
			} else if you == "B" {
				if me == "Y" {
					//draw with paper+2
					legend[key] = 5
				} else if me == "X" {
					//lose with rock+1
					legend[key] = 1
				} else if me == "Z" {
					//win with scissors+3
					legend[key] = 9
				} else {
					log.Fatalf("something didn't catch: %s%s", you, me)
				}
			} else if you == "C" {
				if me == "Y" {
					//draw with scissors+3
					legend[key] = 6
				} else if me == "X" {
					//lose with paper+2
					legend[key] = 2
				} else if me == "Z" {
					//win with rock+1
					legend[key] = 7
				} else {
					log.Fatalf("something didn't catch: %s%s", you, me)
				}
			} else {
				log.Fatalf("something didn't catch: %s%s", you, me)
			}
		}
	}

	return legend
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	legend := generateSecondRoundKey()
	total := 0

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputList := strings.Split(string(content), "\n")

	for _, item := range inputList {
		if item != "" {
			total += legend[item]
		}
	}

	log.Println("total: ", total)
}
