package main

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	cases := []struct {
		in, want [][]int
	}{
		{
			in:   [][]int{{1, 0, 1}, {1, 1, 1}, {1, 1, 1}},
			want: [][]int{{0, 0, 0}, {1, 0, 1}, {1, 0, 1}},
		},
		{
			in:   [][]int{{1, 0, 1, 0}, {1, 1, 1, 1}},
			want: [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}},
		},
		{
			in:   [][]int{},
			want: [][]int{},
		},
	}

	for _, c := range cases {
		result := setZeroes(c.in)
		if !reflect.DeepEqual(result, c.want) {
			t.Errorf("setZeroes(%v) == %v; want %v", c.in, result, c.want)
		}
	}
}
