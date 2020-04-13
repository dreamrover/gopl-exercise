package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s \"string\"\n", os.Args[0])
		return
	}
	var array [18]byte
	copy(array[:], os.Args[1])
	fmt.Printf("%s\n", *reverse(&array))
	s := "asdf"
	var a [len("asdf")]byte
	copy(a[:], s)
}

func reverse(a *[18]byte) *[18]byte {
	length := len(a)
	if length == 0 || length == 1 {
		return a
	}
	for i := 0; i < length/2; i++ {
		a[i], a[length-1-i] = a[length-1-i], a[i]
	}
	return a
}
