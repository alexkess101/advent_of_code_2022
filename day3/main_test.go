package main

import (
	"testing"
)

func TestIsInSlice(t *testing.T) {
	slice := []string{"a", "b", "c"}

	if !Includes(slice, "a") {
		t.Fatal("expected true, got false")
	}
}

func TestGetCommonBadgeFromGroup(t *testing.T) {
	s1 := "adsfjkl;"
	s2 := "qu"
	s3 := "9"
	if len(GetCommonBadgeFromGroup(s1, s2, s3)) > 0 {
		t.Fatal("something went wrong")
	}

	s1 = "adsfjkl;"
	s2 = "qua"
	s3 = "9aaaaaaa"
	answer1 := GetCommonBadgeFromGroup(s1, s2, s3)
	if len(answer1) != 1 {
		t.Fatal("should be equal to 1 but got: ", len(answer1))
	}

	s1 = "adsfbjkl;"
	s2 = "qbua"
	s3 = "9aaabaaaa"
	answer2 := GetCommonBadgeFromGroup(s1, s2, s3)
	if len(answer2) != 2 {
		t.Fatal("should be equal to 2 but got: ", len(answer2))
	}
}

func TestThatTheyAreInAllGroups(t *testing.T) {
	s1 := "vJrwpWtwJgWrhcsFMMfFFhFp"
	s2 := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
	s3 := "PmmdzqPrVvPwwTWBwg"

	answer := GetCommonBadgeFromGroup(s1, s2, s3)
	if len(answer) != 1 {
		t.Fatal("should be equal to 1 but got: ", len(answer))
	}
}
