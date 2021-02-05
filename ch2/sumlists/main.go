package main

import (
	"github.com/ashmidgley/ctci/ch2/linkedlist"
)

func sumLists(l1 *linkedlist.List, l2 *linkedlist.List) *linkedlist.List {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var v1 []int
	var v2 []int
	current := l1.Head
	for current != nil {
		v1 = append(v1, current.Value)
		current = current.Next
	}

	current = l2.Head
	for current != nil {
		v2 = append(v2, current.Value)
		current = current.Next
	}

	carry := 0
	result := &linkedlist.List{}
	for i, j := 0, 0; i < len(v1) && j < len(v2); {
		result.Append((v1[i] + v2[j] + carry) % 10)
		carry = (v1[i] + v2[j]) / 10
		i++
		j++
	}

	return result
}
