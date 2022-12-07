package main

import "testing"

func TestIsContainedFromLeftSide(t *testing.T) {
	inputA := "1-10"
	inputB := "2-5"

	if !IsContained(inputA, inputB) {
		t.Fatal("should be true but got false")
	}
}

func TestIsContainedFromRightSide(t *testing.T) {
	inputA := "2-5"
	inputB := "1-10"

	if !IsContained(inputA, inputB) {
		t.Fatal("should be true but got false")
	}
}

func TestIsContainedIfNotContained(t *testing.T) {
	inputA := "2-5"
	inputB := "6-10"

	if IsContained(inputA, inputB) {
		t.Fatal("should be false but got true")
	}
}

func TestDoesOverlapReturnFalse(t *testing.T) {
	if DoesOverlap("2-4", "6-8") {
		t.Fatal("should be false but got true")
	}
	if DoesOverlap("2-3", "4-5") {
		t.Fatal("should be false but got true")
	}
	if DoesOverlap("0-0", "1-1") {
		t.Fatal("should be false but got true")
	}
	if DoesOverlap("11-72", "2-5") {
		t.Fatal("should be false but got true")
	}
	if DoesOverlap("34-34", "37-45") {
		t.Fatal("should be false but got true")
	}
	if DoesOverlap("37-45", "34-34") {
		t.Fatal("should be false but got true")
	}
}

func TestDoesOverlapReturnTrue(t *testing.T) {
	if !DoesOverlap("5-7", "7-9") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("2-8", "3-7") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("6-6", "4-6") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("2-6", "4-8") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("5-42", "5-41") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("7-9", "5-7") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("38-38", "37-45") {
		t.Fatal("should be true but got false")
	}
	if !DoesOverlap("37-45", "38-38") {
		t.Fatal("should be true but got false")
	}
}
