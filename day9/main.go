package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")

	head := NewHead()
	tail := NewTail()
	tail.addToPath()
	for _, i := range input {
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

	log.Println(len(tail.path))
}

type Head struct {
	pos  Pos
	prev Pos
}

type Tail struct {
	pos  Pos
	path map[string]int
}

type Pos struct {
	x int
	y int
}

type History struct {
	storage Pos
	prevMov Pos
}

func NewHead() Head {
	return Head{pos: Pos{x: 0, y: 0}}
}

func NewTail() Tail {
	return Tail{pos: Pos{x: 0, y: 0}, path: make(map[string]int)}
}

//func (h *History) store(way string) {
//	if h.storage.sum() == 0 {
//		h.storage = h
//	} else {
//		h.prevMov = h.storage
//		h.storage = way
//	}
//}

func (p Pos) stringify() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

func (p Pos) sum() int {
	return p.x + p.y
}

func (this *Head) move(way string) {
	this.prev = this.pos
	switch way {
	case "R":
		this.pos.x = this.pos.x + 1
	case "U":
		this.pos.y = this.pos.y + 1
	case "L":
		this.pos.x = this.pos.x - 1
	case "D":
		this.pos.y = this.pos.y - 1
	}

}

func (this *Tail) move(way string) {
	switch way {
	case "R":
		this.pos.x = this.pos.x + 1
	case "U":
		this.pos.y = this.pos.y + 1
	case "L":
		this.pos.x = this.pos.x - 1
	case "D":
		this.pos.y = this.pos.y - 1
	}
}

func (this *Tail) shouldMove(head Head) bool {
	xDif := head.pos.x - this.pos.x
	yDif := head.pos.y - this.pos.y
	return math.Abs(float64(xDif)) > 1 || math.Abs(float64(yDif)) > 1
}

func (this *Tail) addToPath() {
	if value, ok := this.path[this.pos.stringify()]; ok {
		this.path[this.pos.stringify()] = value + 1
	} else {
		this.path[this.pos.stringify()] = 1
	}
}
