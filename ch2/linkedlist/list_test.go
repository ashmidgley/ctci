package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	want := []int{1, 2, 3, 4}
	l := List{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	index := 0
	current := l.Head
	for current != nil {
		if current.Value != want[index] {
			t.Errorf("value at node %d == %d; want %d", index+1, current.Value, want[index])
		}
		current = current.Next
		index++
	}
}

func TestDelete(t *testing.T) {
	l1 := List{}
	l1.Append(1)
	l1.Append(2)
	l1.Delete(1)

	l2 := List{}
	l2.Append(1)
	l2.Append(2)
	l2.Delete(2)

	l3 := List{}
	l3.Append(1)
	l3.Append(2)
	l3.Append(3)
	l3.Delete(3)

	cases := []struct {
		name   string
		result List
		want   []int
	}{
		{"remove head", l1, []int{2}},
		{"remove node 2", l2, []int{1}},
		{"remove node 3", l3, []int{1, 2}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			index := 0
			current := c.result.Head
			for current != nil {
				if current.Value != c.want[index] {
					t.Errorf("value at node %d == %d; want %d", index+1, current.Value, c.want[index])
				}
				current = current.Next
				index++
			}
		})
	}
}
