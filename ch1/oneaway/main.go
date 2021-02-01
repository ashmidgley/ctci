package main

import (
	"sort"
	"strings"
)

func isOneEditAway(a, b string) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen && aLen+1 != bLen && aLen != bLen+1 {
		return false
	}

	if aLen > bLen {
		//Remove
		return offByOne(b, a)
	}
	if aLen < bLen {
		// Insert
		return offByOne(a, b)
	}

	// Replace
	replacementHit := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if !replacementHit {
				replacementHit = true
			} else {
				return false
			}
		}
	}
	return true
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func offByOne(a, b string) bool {
	a = sortString(a)
	b = sortString(b)

	extraHit := false
	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] != b[j] {
			if !extraHit {
				extraHit = true
				j++
			} else {
				return false
			}
		} else {
			i++
			j++
		}
	}
	return true
}
