package main

import "testing"

var block = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestPerimeter(t *testing.T) {
	perimeter := GetPerimeterAmount(block)
	if perimeter != 16 {
		t.Fatal("failed")
	}
}

func TestCheckLeft(t *testing.T) {
	if CheckLeft(block, 1, 1) != 1 {
		t.Fatal("failed")
	}
	if CheckLeft(block, 1, 2) != 1 {
		t.Fatal("failed")
	}
	if CheckLeft(block, 3, 3) != 1 {
		t.Fatal("failed")
	}
	if CheckLeft(block, 3, 2) != 2 {
		t.Fatal("failed")
	}
}

func TestCheckRight(t *testing.T) {
	if CheckRight(block, 1, 2) != 2 {
		t.Fatal("failed ", CheckRight(block, 1, 2))
	}
	if CheckRight(block, 0, 4) != 0 {
		t.Fatal("failed")
	}
	if CheckRight(block, 1, 3) != 1 {
		t.Fatal("failed")
	}
	if CheckRight(block, 3, 1) != 1 {
		t.Fatal("failed")
	}
}

func TestCheckUp(t *testing.T) {
	if CheckUp(block, 1, 2) != 1 {
		t.Fatal("failed")
	}
	if CheckUp(block, 0, 3) != 0 {
		t.Fatal("failed")
	}
	if CheckUp(block, 3, 3) != 3 {
		t.Fatal("failed")
	}
	if CheckUp(block, 4, 3) != 4 {
		t.Fatal("failed")
	}
}

func TestCheckDown(t *testing.T) {
	if CheckDown(block, 4, 2) != 0 {
		t.Fatal("failed")
	}
	if CheckDown(block, 4, 4) != 0 {
		t.Fatal("failed")
	}
	if CheckDown(block, 0, 3) != 4 {
		t.Fatal("failed")
	}
	if CheckDown(block, 3, 3) != 1 {
		t.Fatal("failed")
	}
}
