package main

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

var instructionsSmall = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

var instructionsBig = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func TestMoveHead(t *testing.T) {
	head := NewHead()
	head.move("R")
	head.move("R")
	head.move("R")
	head.move("R")
	head.move("U")
	head.move("U")
	head.move("U")
	if head.pos.x != 4 && head.pos.y != 3 {
		t.Fatal("failed")
	}
	head.move("L")
	head.move("L")
	head.move("D")
	if head.pos.x != 2 && head.pos.y != 2 {
		t.Fatal("failed")
	}
}

func TestIfTailShouldMove(t *testing.T) {
	head := NewHead()
	tail := NewTail()

	head.move("R")
	head.move("R")
	if !tail.shouldMove(head) {
		t.Fatal("failed")
	}
	if tail.isDiagonal(head.pos) {
		t.Fatal("failed")
	}
	head.move("L")
	head.move("U")
	if tail.shouldMove(head) {
		t.Fatal("failed")
	}
	head.move("R")
	if !tail.shouldMove(head) {
		t.Fatal()
	}
	if !tail.isDiagonal(head.pos) {
		t.Fatal()
	}
}

func TestAddValueToTailPath(t *testing.T) {
	head := NewHead()
	tail := NewTail()

	head.move("R")
	head.move("R")
	tail.move(head.pos)
	tail.addToPath()
	head.move("U")
	head.move("U")
	tail.move(head.pos)
	tail.addToPath()
	log.Println(tail.path)
}

func TestQuadrants(t *testing.T) {
	tail := NewTail()
	tail.pos = Pos{2, 4}
	if tail.findQuadrant(Pos{4, 3}) != 1 {
		t.Fatal()
	}
}

func TestInstructions(t *testing.T) {
	head := NewHead()
	tail := NewTail()
	tail.addToPath()
	for _, i := range instructionsSmall {
		rule := strings.Split(i, " ")
		amount, _ := strconv.Atoi(rule[1])
		for i := 0; i < amount; i++ {
			head.move(rule[0])
			if tail.shouldMove(head) {
				tail.move(head.pos)
				tail.addToPath()
			}
		}
	}

	if len(tail.path) != 13 {
		t.Fatal("failed")
	}
}

func TestPullingTheLastKnot(t *testing.T) {
	var input = []string{
		"R 20",
	}

	if len(SimulatePart2(input)) != 18 {
		t.Fatal()
	}
}

func TestInstructionsPart2Small(t *testing.T) {
	log.Println(SimulatePart2(instructionsSmall))
}

func TestInstructionsPart2Big(t *testing.T) {
	log.Println("here ", len(SimulatePart2(instructionsBig)))
}
