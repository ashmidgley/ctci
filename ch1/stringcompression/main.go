package main

import (
	"fmt"
	"strings"
)

func compress(s string) string {
	if len(s) == 0 {
		return s
	}

	builder := strings.Builder{}
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			builder.WriteString(fmt.Sprintf("%s%d", string(s[i]), count+1))
			count = 0
		}
	}
	builder.WriteString(fmt.Sprintf("%s%d", string(s[len(s)-1]), count+1))

	if len(s) < builder.Len() {
		return s
	}
	return builder.String()
}
