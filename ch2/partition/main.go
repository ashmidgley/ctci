package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func partition(l *linkedlist.List, x int) (*linkedlist.List, *linkedlist.List) {
	if l == nil {
		return nil, nil
	}

	l1 := &linkedlist.List{}
	l2 := &linkedlist.List{}

	current := l.Head
	for current != nil {
		if current.Value < x {
			l1.Append(current.Value)
		} else {
			l2.Append(current.Value)
		}
		current = current.Next
	}
	return l1, l2
}
