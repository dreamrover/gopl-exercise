// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type TabWriter struct {
	n int
	m int
}

func (t *TabWriter) Write(p []byte) (int, error) {
	for i, b := range p {
		if b == '\t' && t.n%t.m == 0 {
			p[i] = '\n'
		}
		t.n++
	}
	return os.Stdout.Write(p)
}

func main() {
	tw := &TabWriter{0, len(os.Args) - 1}

	for _, arg := range os.Args[1:] {
		s := strings.Split(arg, "=")
		tw.Write([]byte(s[0] + "\t"))
		go func() {
			time.Sleep(time.Second)
			conn, err := net.Dial("tcp", s[1])
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(tw, conn)
		}()
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Minute)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
