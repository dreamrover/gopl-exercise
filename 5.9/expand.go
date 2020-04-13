package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
)

const ss = "$foo"

func main() {
	var s string
	if len(os.Args) > 1 {
		fmt.Println(expand(os.Args[1], eval))
	} else {
		fmt.Scanf("%s", &s)
		fmt.Println(expand(s, eval))
	}
}

func eval(ss string) string {
	return fmt.Sprintf("%016x", md5.Sum([]byte(ss)))
}

func expand(s string, f func(string) string) string {
	var b []byte
	a := []byte(s)
	i := bytes.Index(a, []byte(ss))
	for i != -1 {
		b = append(b, a[:i]...)
		b = append(b, []byte(f(ss[1:]))...)
		a = a[i+len(ss):]
		i = bytes.Index(a, []byte(ss))
	}
	b = append(b, a...)
	return string(b)
}
