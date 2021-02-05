package main

import "testing"

func TestAllUnique(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"abcdefg", true},
		{"世界", true},
		{"aaaaaaa", false},
		{"世世", false},
	}

	for _, c := range cases {
		got := allUnique(c.in)
		if got != c.want {
			t.Errorf("allUnique(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
