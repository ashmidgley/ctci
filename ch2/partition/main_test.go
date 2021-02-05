package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestPartition(t *testing.T) {
	l := linkedlist.List{}
	l.Append(3)
	l.Append(5)
	l.Append(8)
	l.Append(5)
	l.Append(10)
	l.Append(2)
	l.Append(1)

	want1 := []int{3, 2, 1}
	want2 := []int{5, 8, 5, 10}

	l1, l2 := partition(l, 5)
	i := 0
	current := l1.Head
	for current != nil {
		if current.Value != want1[i] {
			t.Errorf("value at node %d == %d; want %d", i+1, current.Value, want1[i])
		}
		i++
		current = current.Next
	}

	i = 0
	current = l2.Head
	for current != nil {
		if current.Value != want2[i] {
			t.Errorf("value at node %d == %d; want %d", i+1, current.Value, want2[i])
		}
		i++
		current = current.Next
	}
}
