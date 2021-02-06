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

func sumForwardLists(l1 *linkedlist.List, l2 *linkedlist.List) *linkedlist.List {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var result []int
	current1 := l1.Head
	current2 := l2.Head
	for current1 != nil || current2 != nil {
		carry := (current1.Value + current2.Value) / 10
		sum := (current1.Value + current2.Value) % 10

		if carry == 1 {
			length := len(result)
			if length > 0 {
				result[length-1] += carry
			} else {
				result = append(result, carry)
			}
		}

		result = append(result, sum)
		current1 = current1.Next
		current2 = current2.Next
	}

	l := &linkedlist.List{}
	for _, val := range result {
		l.Append(val)
	}
	return l
}
