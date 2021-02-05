package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestSumLists(t *testing.T) {
	l1 := &linkedlist.List{}
	l1.Append(7)
	l1.Append(1)
	l1.Append(6)

	l2 := &linkedlist.List{}
	l2.Append(5)
	l2.Append(9)
	l2.Append(2)

	want := []int{2, 1, 9}
	result := sumLists(l1, l2)

	i := 0
	current := result.Head
	for current != nil {
		if current.Value != want[i] {
			t.Errorf("value at node %d == %d; want %d", i+1, current.Value, want[i])
		}
		current = current.Next
		i++
	}
}
