package main

import "testing"

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"Tact Coa", true},
		{"Tact Coao", true},
		{"Tact Cod", false},
		{"Tact Coaa", false},
	}

	for _, c := range cases {
		result := isPermutation(c.in)
		if result != c.want {
			t.Errorf("isPermutation(%q) == %t; want %t", c.in, result, c.want)
		}
	}
}
