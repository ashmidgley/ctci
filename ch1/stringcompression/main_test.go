package main

import "testing"

func TestCompress(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"aaBcccccAAa", "a2B1c5A2a1"},
		{"abcd", "abcd"},
		{"aabcccccaad", "a2b1c5a2d1"},
		{"", ""},
	}

	for _, c := range cases {
		result := compress(c.in)
		if result != c.want {
			t.Errorf("compress(%q) == %q; want %q", c.in, result, c.want)
		}
	}
}
