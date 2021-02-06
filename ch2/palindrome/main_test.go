package main

import (
	"testing"

	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func TestIsPalindrome(t *testing.T) {
	l1 := &linkedlist.List{}
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Append(2)
	l1.Append(1)

	l2 := &linkedlist.List{}
	l2.Append(1)
	l2.Append(2)
	l2.Append(3)
	l2.Append(3)
	l2.Append(2)
	l2.Append(1)

	l3 := &linkedlist.List{}
	l3.Append(1)
	l3.Append(2)
	l3.Append(3)

	cases := []struct {
		in   *linkedlist.List
		want bool
	}{
		{l1, true},
		{l2, true},
		{l3, false},
	}

	for _, c := range cases {
		result := isPalindrome(c.in)
		if result != c.want {
			t.Errorf("isPalindrome(%v) == %t; want %t", c.in, result, c.want)
		}
	}
}
