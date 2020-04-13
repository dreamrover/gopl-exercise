package main

import (
	"io"
	"os"

	"exercise/7.4/my"
)

func main() {
	rd := my.NewReader("alirrgklreorjgele;orij\n")
	io.Copy(os.Stdout, rd)
}
