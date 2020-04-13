package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			counts["invalid"]++
		} else if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsNumber(r) {
			counts["number"]++
		} else if unicode.IsSpace(r) {
			counts["space"]++
		} else if unicode.IsControl(r) {
			counts["control"]++
		} else {
			counts["other"]++
		}
	}
	for k, v := range counts {
		fmt.Println(k, v)
	}
}
