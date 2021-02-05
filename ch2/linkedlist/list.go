package linkedlist

import (
	"fmt"
)

// Node represents a node in a linked list.
type Node struct {
	Next  *Node
	Value int
}

// List is a singly linked list.
type List struct {
	Head *Node
}

// Append adds a new node to the end of the list.
func (l *List) Append(value int) {
	end := &Node{
		Value: value,
	}

	if l.Head == nil {
		l.Head = end
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = end
}

// Delete removes a corresponding node from the list.
func (l *List) Delete(value int) *Node {
	if l.Head == nil {
		return nil
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return l.Head
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return current.Next
		}
		current = current.Next
	}
	return l.Head
}

// Print outputs your list in a readable manner.
func (l *List) Print() {
	if l.Head == nil {
		fmt.Println("[]")
		return
	}

	fmt.Print("[ ")
	current := l.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println("]")
}
