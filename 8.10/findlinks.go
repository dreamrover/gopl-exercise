// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"exercise/8.10/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

var depth = flag.Int("depth", 0, "depth")

//!+
func main() {
	flag.Parse()
	fmt.Printf("total: %d\n", *depth)
	for i, arg := range os.Args {
		if arg == "-depth" {
			copy(os.Args[i:], os.Args[i+2:])
			os.Args = os.Args[:len(os.Args)-2]
		}
	}

	worklist := make(chan []string)
	var n int // number of pending sends to worklist
	var wg sync.WaitGroup
	var m int

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		fmt.Printf("depth: %d\n", m)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				wg.Add(1)
				go func(link string, m int) {
					tmp := crawl(link)
					wg.Done()
					if m < *depth {
						worklist <- tmp
					}
				}(link, m)
			}
		}
		m++
		if m > *depth {
			break
		}
	}
	fmt.Println("waiting...")
	wg.Wait()
}

//!-
