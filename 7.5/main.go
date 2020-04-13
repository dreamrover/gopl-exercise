package main

import (
	"io"
	"os"
	"strings"

	"exercise/7.5/limit"
)

func main() {
	r1 := strings.NewReader("abcdefghijklmnopqrstuvwxyz\n")
	lr1 := limit.LimitReader(r1, 30)
	io.Copy(os.Stdout, lr1)

	r2 := strings.NewReader("abcdefghijklmnopqrstuvwxyz\n")
	lr2 := limit.LimitReader(r2, 27)
	io.Copy(os.Stdout, lr2)

	r3 := strings.NewReader("abcdefghijklmnopqrstuvwxyz\n")
	lr3 := limit.LimitReader(r3, 10)
	io.Copy(os.Stdout, lr3)
}
