package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	n := 0
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		n++
	}
	*c += WordCounter(n)
	return n, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	n := 0
	input := bufio.NewScanner((bytes.NewReader(p)))
	input.Split(bufio.ScanLines)
	for input.Scan() {
		n++
	}
	*c += LineCounter(n)
	return n, nil
}

func main() {
	var w WordCounter
	var l LineCounter
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(b)
	l.Write(b)
	fmt.Println(w, "words,", l, "lines")
	w, l = 0, 0
	fmt.Fprint(&w, string(b))
	fmt.Fprintf(&l, "%s", b)
	fmt.Println(w, "words,", l, "lines")
}
