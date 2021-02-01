package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		in, want [][]int
	}{
		{
			in:   [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
			want: [][]int{{3, 2, 1}, {3, 2, 1}, {3, 2, 1}},
		},
		{
			in:   [][]int{{1, 1}, {2, 2}, {3, 3}},
			want: [][]int{{3, 2, 1}, {3, 2, 1}},
		},
		{
			in:   [][]int{},
			want: [][]int{},
		},
	}

	for _, c := range cases {
		result := rotate(c.in)
		if !reflect.DeepEqual(result, c.want) {
			t.Errorf("rotate(%v) == %v; want %v", c.in, result, c.want)
		}
	}
}
