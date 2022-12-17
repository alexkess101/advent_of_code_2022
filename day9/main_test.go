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

//func TestInstructionsPart2Small(t *testing.T) {
//	head := NewHead()
//	knots := []Tail{NewTail(), NewTail(), NewTail(), NewTail(), NewTail(), NewTail(), NewTail(), NewTail(), NewTail()}
//	knots[8].addToPath()
//	knots[8].name = "last"
//
//	for _, i := range instructionsSmall {
//		rule := strings.Split(i, " ")
//		amount, _ := strconv.Atoi(rule[1])
//		for i := 0; i < amount; i++ {
//			head.move(rule[0])
//			if knots[0].shouldMove(head) {
//				knots[0].prev = knots[0].pos
//				knots[0].pos = head.prev
//
//				for i := 1; i < len(knots)-2; i++ {
//					if knots[i].shouldMoveT(knots[i-1]) {
//						knots[i].prev = knots[i].pos
//						knots[i].pos = knots[i-1].prev
//					}
//				}
//
//				if knots[8].shouldMoveT(knots[7]) {
//					knots[8].pos = knots[7].prev
//					knots[8].addToPath()
//				}
//			}
//		}
//	}
//
//	for _, item := range knots {
//		log.Println(item.pos)
//	}
//}
