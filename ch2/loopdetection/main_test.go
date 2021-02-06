package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestFindLoop(t *testing.T) {
	l1 := &linkedlist.List{}
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	l1.Head.Next.Next.Next.Next.Next = l1.Head.Next.Next

	l2 := &linkedlist.List{}
	l2.Append(1)
	l2.Append(2)
	l2.Append(3)

	cases := []struct {
		name string
		in   *linkedlist.List
		want *linkedlist.Node
	}{
		{"loop at node 3", l1, l1.Head.Next.Next},
		{"no loop", l2, nil},
		{"empty list", nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := findLoop(c.in)
			if result != c.want {
				t.Errorf("findLoop(%v) == %v; want %v", c.in, result, c.want)
			}
		})
	}
}
