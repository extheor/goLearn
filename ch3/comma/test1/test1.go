package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	r := comma("+12345678.2")
	fmt.Println(r)
}

func comma(s string) string {
	var symbol string
	var buf bytes.Buffer

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		symbol = s[:1]
		s = s[1:]
		buf.WriteString(symbol)
	}

	dotIndex := strings.IndexByte(s, '.')
	intnum := s
	var decnum string

	if dotIndex > -1 {
		intnum = s[:dotIndex]
		decnum = s[dotIndex+1:]
	}
	n := len(intnum)

	if n <= 3 {
		return s
	}

	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	if dotIndex > -1 {
		buf.WriteByte('.')
		buf.WriteString(decnum)
	}

	return buf.String()
}
