package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage:", os.Args[0], "string1 string2")
		return
	}
	fmt.Println(disturb(os.Args[1], os.Args[2]))
}

func disturb(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			return false
		}
	}
	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}
	return true
}
