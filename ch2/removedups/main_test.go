package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestRemoveDupes(t *testing.T) {
	l1 := linkedlist.List{}
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)

	l2 := linkedlist.List{}
	l2.Append(1)
	l2.Append(1)
	l2.Append(1)

	cases := []struct {
		name string
		in   linkedlist.List
		want []int
	}{
		{"no dupes", l1, []int{1, 2, 3}},
		{"all dupes", l2, []int{1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := removeDupes(c.in)

			index := 0
			current := result.Head
			for current != nil {
				if current.Value != c.want[index] {
					t.Fatalf("value at head == %d; want %d", current.Value, c.want[index])
				}
				current = current.Next
				index++
			}
		})
	}
}
