package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func findLoop(l *linkedlist.List) *linkedlist.Node {
	if l == nil {
		return nil
	}

	set := make(map[*linkedlist.Node]struct{})
	current := l.Head
	for current != nil {
		if _, ok := set[current]; ok {
			return current
		} else {
			set[current] = struct{}{}
		}
		current = current.Next
	}
	return nil
}
