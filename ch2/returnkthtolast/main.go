package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func kthToLast(l *linkedlist.List, k int) *linkedlist.Node {
	if l == nil {
		return nil
	}

	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}

	i := 0
	current = l.Head
	for current != nil {
		if i+k == count {
			return current
		}
		i++
		current = current.Next
	}
	return current
}
