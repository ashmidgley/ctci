package main

import "strings"

func urlify(s string, length int) string {
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}
