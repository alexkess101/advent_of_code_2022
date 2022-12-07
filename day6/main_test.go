package main

import "testing"

func TestCheckForDuplicates(t *testing.T) {
	if HasDuplicates("abcd") {
		t.Fatal("failed")
	}
	if !HasDuplicates("abbd") {
		t.Fatal("failed")
	}
	if !HasDuplicates("ajkla;sjdkflsjda") {
		t.Fatal("failed")
	}
}
