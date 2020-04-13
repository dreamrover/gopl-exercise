package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(join(""))
		return
	}
	if len(os.Args) < 3 {
		fmt.Println(join(os.Args[1]))
		return
	}
	fmt.Println(join(os.Args[1], os.Args[2:]...))
}

func join(sep string, a ...string) string {
	if a == nil {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	b := a[0]
	for _, s := range a[1:] {
		b += sep + s
	}
	return b
}
