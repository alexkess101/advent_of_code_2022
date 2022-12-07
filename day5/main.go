package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	values []string
}

type Instruction struct {
	amount int
	from   string
	to     string
}

type Reader struct {
	stacks       map[int]*Stack
	instructions []Instruction
}

func (s *Stack) insert(value string) {
	s.values = append(s.values, value)
}

func (s *Stack) insertGroup(group []string) {
	s.values = append(s.values, group...)
}

func (s *Stack) pop() string {
	item := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return item
}

func (s *Stack) slice(amount int) []string {
	value := s.values
	if amount >= len(value) {
		s.values = []string{}
		return value
	} else {
		s.values = value[:len(value)-amount]
		return value[len(value)-amount : len(value)]
	}
}

func (r *Reader) decodeInstructions(input []string) {
	var instructions []Instruction
	for _, item := range input {
		parsedItem := strings.Split(item, " ")
		amount, _ := strconv.Atoi(parsedItem[1])
		from := parsedItem[3]
		to := parsedItem[5]
		instructions = append(instructions, Instruction{amount: amount, from: from, to: to})
	}

	r.instructions = instructions
}

func (r *Reader) decodeStacks(input []string) {
	stackNames := strings.Split(strings.ReplaceAll(input[len(input)-1], " ", ""), "")
	for _, key := range stackNames {
		key, _ := strconv.Atoi(key)
		r.stacks[key] = &Stack{}
	}
	input = input[:len(input)-1]

	var columnCount int
	for i := len(input) - 1; i >= 0; i-- {
		columnCount = 1
		for j := 0; j < len(input[i])-1; j++ {
			var space uint8
			space = 32
			char := string(input[i][j])
			nextChar := string(input[i][j+1])

			if char == "[" {
				r.stacks[columnCount].insert(nextChar)
				columnCount++
			} else if input[i][j] == space && input[i][j+3] == space {
				columnCount++
				j += 3
			}
		}
	}
}

func (r *Reader) decode(input string) {
	parser := strings.Split(input, "\n\n")
	//log.Println(regex.Split(parser[0], -1))
	r.decodeStacks(strings.Split(parser[0], "\n"))
	r.decodeInstructions(strings.Split(parser[1], "\n"))
}

func (r *Reader) print() {
	for key, stack := range r.stacks {
		log.Println(key, stack.values)
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := Reader{stacks: make(map[int]*Stack)}
	reader.decode(string(content))

	for _, command := range reader.instructions {
		//PART 1:
		//for i := 0; i < command.amount; i++ {
		//	from, _ := strconv.Atoi(command.from)
		//	to, _ := strconv.Atoi(command.to)
		//	crate := reader.stacks[from].pop()
		//	reader.stacks[to].insert(crate)
		//}

		//PART 2:
		from, _ := strconv.Atoi(command.from)
		to, _ := strconv.Atoi(command.to)
		crates := reader.stacks[from].slice(command.amount)

		reader.stacks[to].insertGroup(crates)
	}

	reader.print()
}
