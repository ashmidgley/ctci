package main

import "testing"

func TestIsOneEditAway(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"pale", "ple", true},
		{"pale", "pales", true},
		{"pale", "bale", true},
		{"pale", "paless", false},
		{"pale", "pde", false},
		{"pale", "palds", false},
		{"pale", "bake", false},
	}

	for _, c := range cases {
		result := isOneEditAway(c.a, c.b)
		if result != c.want {
			t.Errorf("isOneEditAway(%q, %q) == %t; want %t", c.a, c.b, result, c.want)
		}
	}
}
