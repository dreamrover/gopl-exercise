package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	cw, c := CountingWriter(os.Stdout)
	fmt.Fprint(cw, "abcdefg")
	fmt.Println("\n", *c)
	fmt.Fprint(cw, "hijklmnopq")
	fmt.Println("\n", *c)
	fmt.Fprint(cw, "rstuvwxyz")
	fmt.Println("\n", *c)
}

type CWriter struct {
	cnt    int64
	writer io.Writer
}

func (cw *CWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	cw.cnt += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CWriter{0, w}
	return &cw, &cw.cnt
}
