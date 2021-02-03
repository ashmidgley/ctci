package main

import "strings"

func isPermutation(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	occurrences := make(map[rune]int)
	for _, r := range s {
		occurrences[r]++
	}

	singleOddHit := false
	for _, val := range occurrences {
		if val != 2 {
			if !singleOddHit {
				singleOddHit = true
			} else {
				return false
			}
		}
	}
	return true
}
