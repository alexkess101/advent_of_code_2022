package main

import (
	"strconv"
	"strings"
	"testing"
)

var instructions = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
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

func TestHeadPreviousMove(t *testing.T) {
	head := NewHead()
	head.move("R")
	head.move("U")
	if head.prev.sum() != 1 {
		t.Fatal("failed")
	}
}

func TestIfTailShouldMove(t *testing.T) {
	head := NewHead()
	tail := NewTail()
	head.move("R")
	if tail.shouldMove(head) {
		t.Fatal("failed")
	}
	head.move("R")
	if !tail.shouldMove(head) {
		t.Fatal("failed")
	}
	tail.move("U")
	if !tail.shouldMove(head) {
		t.Fatal("failed")
	}
	tail.move("R")
	if tail.shouldMove(head) {
		t.Fatal("failed")
	}
	head.move("L")
	head.move("U")
	//here they are on top of each other
	if tail.shouldMove(head) {
		t.Fatal("failed")
	}
	head.move("U")
	head.move("R")
	if tail.shouldMove(head) {
		t.Fatal("failed")
	}
}

func TestAddValueToTailPath(t *testing.T) {
	tail := NewTail()
	tail.move("R")
	tail.addToPath()
	tail.move("U")
	tail.addToPath()
	tail.move("D")
	tail.addToPath()
	tail.move("R")
	tail.move("R")
	tail.move("U")
	tail.addToPath()
	result := map[string]int{
		"1-0": 2,
		"1-1": 1,
		"3-1": 1,
	}
	if result["1-0"] != tail.path["1-0"] {
		t.Fatal("failed")
	}
}

func TestInstructions(t *testing.T) {
	head := NewHead()
	tail := NewTail()
	tail.addToPath()
	for _, i := range instructions {
		rule := strings.Split(i, " ")
		amount, _ := strconv.Atoi(rule[1])
		for i := 0; i < amount; i++ {
			head.move(rule[0])
			if tail.shouldMove(head) {
				tail.pos = head.prev
				tail.addToPath()
			}
		}
	}

	if len(tail.path) != 13 {
		t.Fatal("failed")
	}
}
