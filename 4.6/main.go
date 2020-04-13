package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s \"string\"\n", os.Args[0])
		return
	}
	fmt.Printf("%s\n", stripSpace([]byte(os.Args[1])))
}

func stripSpace(in []byte) []byte {
	out := in[:0]
	var r0 rune
	s := string(in)
	for _, r := range s {
		if unicode.IsSpace(r0) && unicode.IsSpace(r) {
			r0 = r
			continue
		}
		out = append(out, []byte(string(r))...)
		r0 = r
	}
	return out
}
