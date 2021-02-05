package main

import "github.com/ashmidgley/ctci/ch2/linkedlist"

func deleteMiddle(n *linkedlist.Node) {
	if n == nil || n.Next == nil {
		return
	}

	next := n.Next
	n.Value = next.Value
	n.Next = next.Next
}
