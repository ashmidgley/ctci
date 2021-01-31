package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func isPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	sa := sortString(a)
	sb := sortString(b)
	for i := 0; i < len(a); i++ {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}
