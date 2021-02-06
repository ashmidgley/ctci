package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestIntersect(t *testing.T) {
	l1 := &linkedlist.List{}
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)

	l2 := &linkedlist.List{}
	l2.Append(1)
	l2.Append(2)
	l2.Append(3)

	l3 := &linkedlist.List{}
	l3.Append(1)
	l3.Append(2)

	l4 := &linkedlist.List{}
	l4.Append(1)
	l4.Head.Next = l3.Head.Next

	cases := []struct {
		name   string
		l1, l2 *linkedlist.List
		want   *linkedlist.Node
	}{
		{"no intersect", l1, l2, nil},
		{"intersect at 2", l3, l4, l3.Head.Next},
		{"l1 nil", nil, l2, nil},
		{"l2 nil", l1, nil, nil},
		{"both inputs nil", nil, nil, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := intersect(c.l1, c.l2)
			if result != c.want {
				t.Errorf("intersect(%v, %v) == %v; want %v", c.l1, c.l2, result, c.want)
			}
		})
	}
}
