package main

import "testing"

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		a    string
		b    string
		want bool
	}{
		{"abc", "cab", true},
		{"AbC", "CAb", true},
		{"This is a permutation", "a permutation This is", true},
		{"abc", "ab", false},
		{"ABC", "abc", false},
		{"This is a permutation", "a permutation This id", false},
	}

	for _, c := range cases {
		got := isPermutation(c.a, c.b)
		if got != c.want {
			t.Errorf("isPermutation(%q, %q) == %t; want %t", c.a, c.b, got, c.want)
		}
	}
}

func BenchmarkIsPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermutation("abc", "cab")
	}
}
