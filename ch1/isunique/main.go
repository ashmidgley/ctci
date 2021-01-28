package main

func allUnique(s string) bool {
	set := make(map[rune]bool)
	for _, c := range s {
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = true
	}
	return true
}
