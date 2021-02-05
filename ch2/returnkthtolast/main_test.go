package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestFindKToLast(t *testing.T) {
	l := &linkedlist.List{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	cases := []struct {
		name string
		k    int
		want int
	}{
		{"last", 1, 5},
		{"second to last", 2, 4},
		{"first", 5, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := kthToLast(l, c.k)

			if result.Value != c.want {
				t.Errorf("value at node %d places from last == %d; want %d", c.k, result.Value, c.want)
			}
		})
	}
}
