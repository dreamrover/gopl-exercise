package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s \"string\"\n", os.Args[0])
		return
	}
	fmt.Printf("%s\n", reverse([]byte(os.Args[1])))
}

func reverse(a []byte) []byte {
	var r rune
	var size int
	length := len(a)
	if length == 0 || length == 1 {
		return a
	}
	tmp := make([]byte, length, length)
	copy(tmp, a)
	for i := 0; i < length; i += size {
		r, size = utf8.DecodeRune(tmp[i:])
		utf8.EncodeRune(a[length-i-size:], r)
	}
	return a
}
