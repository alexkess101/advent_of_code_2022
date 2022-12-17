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
				tail.move(head.pos)
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
	name string
	pos  Pos
	prev Pos
	path map[string]int
}

type Pos struct {
	x int
	y int
}

func NewHead() Head {
	return Head{pos: Pos{x: 0, y: 0}}
}

func NewTail() Tail {
	return Tail{pos: Pos{x: 0, y: 0}, path: make(map[string]int)}
}

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

func (this *Tail) move(pos Pos) {
	if this.isDiagonal(pos) {
		switch this.findQuadrant(pos) {
		case 1:
			this.pos.x = this.pos.x + 1
			this.pos.y = this.pos.y - 1
		case 2:
			this.pos.x = this.pos.x - 1
			this.pos.y = this.pos.y - 1
		case 3:
			this.pos.x = this.pos.x + 1
			this.pos.y = this.pos.y + 1
		case 4:
			this.pos.x = this.pos.x - 1
			this.pos.y = this.pos.y + 1
		}
	} else {
		xDif := pos.x - this.pos.x
		yDif := pos.y - this.pos.y

		if xDif == 0 {
			if yDif > 0 {
				this.pos.y = this.pos.y + 1
			} else {
				this.pos.y = this.pos.y - 1
			}
		} else {
			if xDif > 0 {
				this.pos.x = this.pos.x + 1
			} else {
				this.pos.x = this.pos.x - 1
			}
		}
	}
}

func (this *Tail) findQuadrant(pos Pos) int {
	xDif := pos.x - this.pos.x
	yDif := pos.y - this.pos.y

	if xDif > 0 && yDif < 0 {
		return 1
	} else if xDif < 0 && yDif < 0 {
		return 2
	} else if xDif > 0 && yDif > 0 {
		return 3
	} else {
		return 4
	}
}

func (this *Tail) shouldMove(head Head) bool {
	xDif := head.pos.x - this.pos.x
	yDif := head.pos.y - this.pos.y
	return math.Abs(float64(xDif)) > 1 || math.Abs(float64(yDif)) > 1
}

func (this *Tail) shouldMoveT(tail Tail) bool {
	xDif := tail.pos.x - this.pos.x
	yDif := tail.pos.y - this.pos.y
	return math.Abs(float64(xDif)) > 1 || math.Abs(float64(yDif)) > 1
}

func (this *Tail) isDiagonal(pos Pos) bool {
	xDif := pos.x - this.pos.x
	yDif := pos.y - this.pos.y
	return (math.Abs(float64(xDif)) == 2 || math.Abs(float64(yDif)) == 2) && (xDif != 0 && yDif != 0)
}

//func (this *Tail) isDiagonalT(tail Tail) bool {
//	xDif := tail.pos.x - this.pos.x
//	yDif := tail.pos.y - this.pos.y
//	return xDif == -2 || xDif == 2 || yDif == -2 || yDif == 2
//}

func (this *Tail) addToPath() {
	if value, ok := this.path[this.pos.stringify()]; ok {
		this.path[this.pos.stringify()] = value + 1
	} else {
		this.path[this.pos.stringify()] = 1
	}
}
