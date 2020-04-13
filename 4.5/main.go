package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s string1 [string2] ...\n", os.Args[0])
		return
	}
	fmt.Println(norepeat(os.Args[1:]))
}

func norepeat(strings []string) []string {
	out := strings[:0]
	var s0 string
	for _, s := range strings {
		if s != s0 {
			out = append(out, s)
			s0 = s
		}
	}
	return out
}
