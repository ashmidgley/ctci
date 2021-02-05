package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func removeDupes(l *linkedlist.List) {
	if l == nil || l.Head == nil {
		return
	}

	seen := make(map[int]struct{})
	var previous *linkedlist.Node
	current := l.Head
	for current != nil {
		if _, ok := seen[current.Value]; ok {
			previous.Next = current.Next
		} else {
			seen[current.Value] = struct{}{}
			previous = current
		}
		current = current.Next
	}
}
