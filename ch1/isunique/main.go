package main

func allUnique(s string) bool {
	set := make(map[rune]struct{})
	for _, c := range s {
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = struct{}{}
	}
	return true
}
