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

	cases := []struct {
		name   string
		l1, l2 *linkedlist.List
		want   []int
	}{
		{"empty l1 and l2", nil, nil, []int{}},
		{"empty l1", nil, l2, []int{5, 9, 2}},
		{"empty l2", l1, nil, []int{7, 1, 6}},
		{"happy path", l1, l2, []int{2, 1, 9}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sumLists(c.l1, c.l2)

			if result == nil && len(c.want) == 0 {
				t.Skip()
			}

			i := 0
			current := result.Head
			for current != nil {
				if current.Value != c.want[i] {
					t.Errorf("value at node %d == %d; want %d", i+1, current.Value, c.want[i])
				}
				current = current.Next
				i++
			}
		})
	}
}

func TestSumForwardLists(t *testing.T) {
	l1 := &linkedlist.List{}
	l1.Append(6)
	l1.Append(1)
	l1.Append(7)

	l2 := &linkedlist.List{}
	l2.Append(2)
	l2.Append(9)
	l2.Append(5)

	l3 := &linkedlist.List{}
	l3.Append(9)
	l3.Append(1)
	l3.Append(7)

	cases := []struct {
		name   string
		l1, l2 *linkedlist.List
		want   []int
	}{
		{"empty l1 and l2", nil, nil, []int{}},
		{"empty l1", nil, l2, []int{2, 9, 5}},
		{"empty l2", l1, nil, []int{6, 1, 7}},
		{"same length", l1, l2, []int{9, 1, 2}},
		{"first digit carry", l3, l2, []int{1, 2, 1, 2}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := sumForwardLists(c.l1, c.l2)

			if result == nil && len(c.want) == 0 {
				t.Skip()
			}

			i := 0
			current := result.Head
			for current != nil {
				if current.Value != c.want[i] {
					t.Fatalf("value at node %d == %d; want %d", i+1, current.Value, c.want[i])
				}
				current = current.Next
				i++
			}
		})
	}
}
