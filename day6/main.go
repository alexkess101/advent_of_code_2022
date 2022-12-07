package main

import "log"

func HasDuplicates(s string) bool {
	key := make(map[string]interface{})
	hasDuplicate := false
	for _, c := range s {
		if _, ok := key[string(c)]; ok {
			hasDuplicate = true
		}
	}

	return hasDuplicate
}

func main() {
	log.Println("hello")
}
