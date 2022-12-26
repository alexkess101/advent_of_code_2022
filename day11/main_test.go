package main

import (
	"log"
	"testing"
)

func TestNewMonkey(t *testing.T) {
	items0 := []int{79, 98}
	operation0 := func(old int) int { return old * 19 }
	condition0 := func(lvl int) (int, int) { return Condition(lvl, 23, 1, 2) }
	monkey0 := NewMonkey(
		items0,
		operation0,
		condition0,
	)
	var items1 []int
	operation1 := func(old int) int { return old * 5 }
	condition1 := func(lvl int) (int, int) { return Condition(lvl, 2, 0, 2) }
	monkey1 := NewMonkey(
		items1,
		operation1,
		condition1,
	)
	var items2 []int
	operation2 := func(old int) int { return old * 5 }
	condition2 := func(lvl int) (int, int) { return Condition(lvl, 2, 8, 9) }
	monkey2 := NewMonkey(
		items2,
		operation2,
		condition2,
	)
	monkeys := []*Monkey{monkey0, monkey1, monkey2}

	monkey0.run(monkeys)

	if monkey0.count != 2 {
		t.Fatal()
	}
	if len(monkey2.items) != 2 {
		t.Fatal()
	}
}

func TestInputExample(t *testing.T) {
	items0 := []int{79, 98}
	operation0 := func(old int) int { return old * 19 }
	condition0 := func(lvl int) (int, int) { return Condition(lvl, 23, 2, 3) }
	monkey0 := NewMonkey(
		items0,
		operation0,
		condition0,
	)
	items1 := []int{54, 65, 75, 74}
	operation1 := func(old int) int { return old + 6 }
	condition1 := func(lvl int) (int, int) { return Condition(lvl, 19, 2, 0) }
	monkey1 := NewMonkey(
		items1,
		operation1,
		condition1,
	)
	items2 := []int{79, 60, 97}
	operation2 := func(old int) int { return old * old }
	condition2 := func(lvl int) (int, int) { return Condition(lvl, 13, 1, 3) }
	monkey2 := NewMonkey(
		items2,
		operation2,
		condition2,
	)
	items3 := []int{74}
	operation3 := func(old int) int { return old + 3 }
	condition3 := func(lvl int) (int, int) { return Condition(lvl, 17, 0, 1) }
	monkey3 := NewMonkey(
		items3,
		operation3,
		condition3,
	)
	monkeys := []*Monkey{monkey0, monkey1, monkey2, monkey3}
	NoMoreMonkeys(20, monkeys)
	log.Println(monkey0.items, monkey0.count)
	log.Println(monkey1.items, monkey1.count)
	log.Println(monkey2.items, monkey2.count)
	log.Println(monkey3.items, monkey3.count)
}
