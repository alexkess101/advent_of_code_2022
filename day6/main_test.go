package main

import "testing"

func TestCheckForDuplicates(t *testing.T) {
	if HasDuplicates("abcd") {
		t.Fatal("failed")
	}
	if HasDuplicates("wftq") {
		t.Fatal("failed")
	}
	if !HasDuplicates("abbd") {
		t.Fatal("failed")
	}
	if !HasDuplicates("ajkla;sjdkflsjda") {
		t.Fatal("failed")
	}
}

func TestGetProcessedPosition(t *testing.T) {
	if GetProcessedPosition("bvwbjplbgvbhsrlpgdmjqwftvncz", 4) != 5 {
		t.Fatal("failed - expected '5' but got: ", GetProcessedPosition("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	}
	if GetProcessedPosition("nppdvjthqldpwncqszvftbrmjlhg", 4) != 6 {
		t.Fatal("failed - expected '6' but got: ", GetProcessedPosition("nppdvjthqldpwncqszvftbrmjlhg", 4))
	}
	if GetProcessedPosition("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4) != 10 {
		t.Fatal("failed - expected '10' but got: ", GetProcessedPosition("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	}
	if GetProcessedPosition("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4) != 11 {
		t.Fatal("failed - expected '11' but got: ", GetProcessedPosition("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
	}
	if GetProcessedPosition("ababbabababababababababababababa", 4) != -1 {
		t.Fatal("failed - expected '-1' but got: ", GetProcessedPosition("ababbabababababababababababababa", 4))
	}
	if GetProcessedPosition("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14) != 19 {
		t.Fatal("failed - expected '19' but got: ", GetProcessedPosition("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	}
	if GetProcessedPosition("bvwbjplbgvbhsrlpgdmjqwftvncz", 14) != 23 {
		t.Fatal("failed - expected '23' but got: ", GetProcessedPosition("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	}
}
