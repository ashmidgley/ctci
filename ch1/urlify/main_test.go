package main

import "testing"

func TestUrlify(t *testing.T) {
	cases := []struct {
		in   string
		len  int
		want string
	}{
		{"test test  ", 9, "test%20test"},
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"", 0, ""},
		{" ", 1, "%20"},
	}

	for _, c := range cases {
		result := urlify(c.in, c.len)
		if result != c.want {
			t.Errorf("urlify(%q, %d) === %q; want %q", c.in, c.len, result, c.want)
		}
	}
}

func BenchmarkUrlify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlify("Mr John Smith    ", 13)
	}
}
