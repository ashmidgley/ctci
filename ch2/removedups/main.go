package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func removeDupes(l linkedlist.List) linkedlist.List {
	if l.Head == nil {
		return l
	}

	result := linkedlist.List{}
	seen := make(map[int]struct{})
	current := l.Head
	for current != nil {
		if _, ok := seen[current.Value]; !ok {
			result.Append(current.Value)
			seen[current.Value] = struct{}{}
		}
		current = current.Next
	}
	return result
}
