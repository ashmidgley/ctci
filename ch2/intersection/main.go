package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func intersect(l1 *linkedlist.List, l2 *linkedlist.List) *linkedlist.Node {
	if l1 == nil || l2 == nil {
		return nil
	}

	current1 := l1.Head
	for current1 != nil {
		current2 := l2.Head
		for current2 != nil {
			if current1 == current2 {
				return current1
			}
			current2 = current2.Next
		}
		current1 = current1.Next
	}

	return nil
}
