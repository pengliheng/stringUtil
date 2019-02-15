package stringUtil

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func _comma(s string) string {
	var buf bytes.Buffer
	var tempString string
	n := len(s)
	index := strings.Index(s, ".")
	if index > 0 {
		s, tempString = s[:index], s[index:]
		n = len(s)
	}
	for i := 0; i < n; i++ {
		buf.WriteString(s[i : i+1])
		if (n-i-1)%3 == 0 && n != i+1 {
			buf.WriteString(",")
		}
	}
	return buf.String() + tempString
}
