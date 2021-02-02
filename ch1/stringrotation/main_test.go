package main

import "testing"

func TestIsRotation(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbottle", "erbottleduh", false},
		{"waterbottle", "erbottle", false},
		{"waterbot", "erbottlewat", false},
		{"", "", false},
	}

	for _, c := range cases {
		result := isRotation(c.s1, c.s2)
		if result != c.want {
			t.Errorf("isRotation(%q, %q) == %t; want %t", c.s1, c.s2, result, c.want)
		}
	}
}

func BenchmarkIsRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isRotation("waterbottle", "erbottlewat")
	}
}
