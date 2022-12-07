package main

import (
	"log"
	"os"
	"strings"
)

func Includes(s []string, c string) bool {
	isFound := false
	for _, item := range s {
		if item == c {
			isFound = true
		}
	}
	return isFound
}

func getDuplicatesFromBothCompartments(compartments string) []string {
	var duplicates []string
	finder := make(map[string]int)
	middleIndex := len(compartments) / 2
	comp1 := strings.Split(compartments[0:middleIndex], "")
	comp2 := strings.Split(compartments[middleIndex:], "")

	for i := 0; i < len(comp1); i++ {
		finder[comp1[i]] = 0
	}
	for i := 0; i < len(comp2); i++ {
		if _, ok := finder[comp2[i]]; ok && !Includes(duplicates, comp2[i]) {
			duplicates = append(duplicates, comp2[i])
		}
	}

	return duplicates
}

func createKey(s string) map[string]int {
	key := make(map[string]int)
	for i := 0; i < len(s); i++ {
		key[string(s[i])] = 0
	}

	return key
}

func IsFound(item string, comp string) bool {
	key := createKey(comp)

	if _, ok := key[item]; ok {
		return true
	}
	return false
}

func GetCommonBadgeFromGroup(s1, s2, s3 string) []string {
	var matches []string

	for _, item := range strings.Split(s1, "") {
		if IsFound(item, s2) && IsFound(item, s3) && !Includes(matches, item) {
			matches = append(matches, item)
		}
	}

	return matches

	//var count int
	//if len(s2) > len(s3) {
	//	count = len(s2)
	//} else {
	//	count = len(s3)
	//}

	//for i := 0; i < count-1; i++ {
	//	if len(s2)-1 > i {
	//		value := string(s2[i])
	//		if _, ok := key1[value]; ok && !Includes(matches, value) {
	//			matches = append(matches, value)
	//		}
	//	}
	//	if len(s3)-1 > i {
	//		value := string(s3[i])
	//		if _, ok := key1[value]; ok && !Includes(matches, value) {
	//			matches = append(matches, value)
	//		}
	//	}
	//}

	return matches
}

func main() {
	contents, err := os.ReadFile("input.txt")
	total := 0
	if err != nil {
		log.Fatal(err)
	}
	contentList := strings.Split(string(contents), "\n")

	//for _, item := range contentList {
	//	doops := getDuplicatesFromBothCompartments(item)
	//	for _, doop := range doops {
	//		total += Priorities[doop]
	//	}
	//}

	for i := 0; i < len(contentList)-3; i += 3 {
		s1 := contentList[i]
		s2 := contentList[i+1]
		s3 := contentList[i+2]

		commonValues := GetCommonBadgeFromGroup(s1, s2, s3)

		for _, value := range commonValues {
			total += Priorities[value]
		}

		log.Println(commonValues)
	}

	log.Println(total)
}
