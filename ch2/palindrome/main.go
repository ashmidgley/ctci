package main

import (
	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func isPalindrome(l *linkedlist.List) bool {
	occurences := make(map[int]int)
	current := l.Head
	for current != nil {
		occurences[current.Value]++
		current = current.Next
	}

	singleOddHit := false
	for _, val := range occurences {
		if val%2 == 1 {
			if singleOddHit {
				return false
			}
			singleOddHit = true
		}
	}
	return true
}
