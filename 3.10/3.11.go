// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	var s1, s2 string
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(byte(s[0]))
		s1 = s[1:]
	} else {
		s1 = s
	}
	if i := strings.IndexByte(s1, '.'); i != -1 {
		s2 = s1[i:]
		s1 = s1[:i]
	}
	n := len(s1)
	if n <= 3 {
		return s
	}
	r := n % 3
	for i, b := range s1 {
		if i > 0 && i%3 == r {
			buf.WriteByte(',')
		}
		buf.WriteByte(byte(b))
	}
	buf.Write([]byte(s2))
	return buf.String()
}

//!-
