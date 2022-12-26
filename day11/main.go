package main

import "log"

func main() {
	items0 := []int{59, 65, 89, 56, 74, 57, 56}
	operation0 := func(old int) int { return old * 17 }
	condition0 := func(lvl int) (int, int) { return Condition(lvl, 3, 3, 6) }
	monkey0 := NewMonkey(
		items0,
		operation0,
		condition0,
	)
	items1 := []int{63, 83, 50, 63, 56}
	operation1 := func(old int) int { return old + 2 }
	condition1 := func(lvl int) (int, int) { return Condition(lvl, 13, 3, 0) }
	monkey1 := NewMonkey(
		items1,
		operation1,
		condition1,
	)
	items2 := []int{93, 79, 74, 55}
	operation2 := func(old int) int { return old + 1 }
	condition2 := func(lvl int) (int, int) { return Condition(lvl, 2, 0, 1) }
	monkey2 := NewMonkey(
		items2,
		operation2,
		condition2,
	)
	items3 := []int{86, 61, 67, 88, 94, 69, 56, 91}
	operation3 := func(old int) int { return old + 7 }
	condition3 := func(lvl int) (int, int) { return Condition(lvl, 11, 6, 7) }
	monkey3 := NewMonkey(
		items3,
		operation3,
		condition3,
	)
	items4 := []int{76, 50, 51}
	operation4 := func(old int) int { return old * old }
	condition4 := func(lvl int) (int, int) { return Condition(lvl, 19, 2, 5) }
	monkey4 := NewMonkey(
		items4,
		operation4,
		condition4,
	)
	items5 := []int{77, 76}
	operation5 := func(old int) int { return old + 8 }
	condition5 := func(lvl int) (int, int) { return Condition(lvl, 17, 2, 1) }
	monkey5 := NewMonkey(
		items5,
		operation5,
		condition5,
	)
	items6 := []int{74}
	operation6 := func(old int) int { return old * 2 }
	condition6 := func(lvl int) (int, int) { return Condition(lvl, 5, 4, 7) }
	monkey6 := NewMonkey(
		items6,
		operation6,
		condition6,
	)
	items7 := []int{86, 85, 52, 86, 91, 95}
	operation7 := func(old int) int { return old + 6 }
	condition7 := func(lvl int) (int, int) { return Condition(lvl, 7, 4, 5) }
	monkey7 := NewMonkey(
		items7,
		operation7,
		condition7,
	)
	monkeys := []*Monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}
	NoMoreMonkeys(20, monkeys)
	for _, monkey := range monkeys {
		log.Println(monkey.count)
	}
}
