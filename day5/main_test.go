package main

import (
	"reflect"
	"testing"
)

func TestInsertIntoStack(t *testing.T) {
	stack := Stack{values: []string{}}

	stack.insert("A")
	stack.insert("B")
	stack.insert("C")

	if !reflect.DeepEqual(stack.values, []string{"A", "B", "C"}) {
		t.Fatal("slice contents do not match")
	}
}

func TestPopOffStack(t *testing.T) {
	stack := Stack{values: []string{"A", "B", "C"}}
	var value string
	value = stack.pop()

	if value != "C" {
		t.Fatal("incorrect value. expected 'C' but got: ", value)
	}
	if len(stack.values) != 2 {
		t.Fatal("incorrect slice length. expected '2' but got: ", len(stack.values))
	}
}

func TestInstructionDecoder(t *testing.T) {
	var r Reader

	r.decodeInstructions([]string{"move 1 from 3 to 5", "move 2 from 2 to 8"})

	if len(r.instructions) != 2 {
		t.Fatal("length err - should be '2' but got: ", len(r.instructions))
	}
	if r.instructions[0].amount != 1 {
		t.Fatal("set instruction wrong - should be '1' but got: ", r.instructions[0].amount)
	}
	if r.instructions[0].from != "3" {
		t.Fatal("set instruction wrong - should be '3' but got: ", r.instructions[0].amount)
	}
	if r.instructions[0].to != "5" {
		t.Fatal("set instruction wrong - should be '5' but got: ", r.instructions[0].amount)
	}
}

func TestGetSliceGroups(t *testing.T) {
	stack := Stack{values: []string{"A", "B", "C", "D", "E"}}

	if len(stack.slice(3)) != 3 {
		t.Fatal("failed - expected length 3, but got: ", len(stack.slice(3)))
	}
	if len(stack.values) != 2 {
		t.Fatal("failed - expected length 2, but got: ", stack.values)
	}
}

func TestStackDecoder(t *testing.T) {
	r := Reader{stacks: make(map[int]*Stack)}

	r.decodeStacks([]string{"[Z] [G] [V] [V] [Q] [M] [L] [N] [R]", " 1   2   3   4   5   6   7   8   9 "})
	if r.stacks[2].values[0] != "G" {
		t.Fatal("failed - expected 'G' but got: ", r.stacks[2])
	}

	r.decodeStacks([]string{"[W]     [L] [D] [D] [J] [W] [T] [C]", " 1   2   3   4   5   6   7   8   9 "})
	if len(r.stacks[2].values) != 0 {
		t.Fatal("failed - expected 0 length instead got: ", len(r.stacks[2].values))
	}

	r.decodeStacks([]string{"[I]     [F]         [L]     [H] [W]", " 1   2   3   4   5   6   7   8   9 "})
	if r.stacks[6].values[0] != "L" {
		t.Fatal("failed - expected 'L' but got: ", r.stacks[6])
	}

	r.decodeStacks([]string{"                    [L]     [H] [W]", " 1   2   3   4   5   6   7   8   9 "})
	if r.stacks[9].values[0] != "W" {
		t.Fatal("failed - expected 'W' but got: ", r.stacks[9])
	}

	r.decodeStacks([]string{"[Z] [G] [V] [V] [Q] [M] [L] [N] [R]", "                    [L]     [H] [W]", " 1   2   3   4   5   6   7   8   9 "})
	if len(r.stacks[9].values) != 2 {
		t.Fatal("failed - expected 2 length instead got: ", len(r.stacks[9].values))
	}
	if len(r.stacks[1].values) != 1 {
		t.Fatal("failed - expected 2 length instead got: ", len(r.stacks[1].values))
	}
}
