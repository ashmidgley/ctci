package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestDeleteMiddle(t *testing.T) {
	l := linkedlist.List{}
	l.Append(1)
	l.Append(2)
	l.Append(3)

	want := []int{1, 3}
	deleteMiddle(l.Head.Next)

	i := 0
	current := l.Head
	for current != nil {
		if current.Value != want[i] {
			t.Fatalf("value at node %d == %d; want %d", i+1, current.Value, want[i])
		}
		current = current.Next
		i++
	}
}
