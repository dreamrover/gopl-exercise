// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(addr string) []string {
	fmt.Println(addr)

	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	host := u.Host
	path := u.Path
	if path != "" {
		path = u.Path[1:]
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
		}
	}
	dir := filepath.Dir(path)
	if dir != "." {
		if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}
	}
	name := filepath.Base(path)
	if name != "." {
		resp, err := http.Get(addr)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		err = ioutil.WriteFile(path, body, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	tmp, err := links.Extract(addr)
	if err != nil {
		log.Print(err)
		return nil
	}
	var list []string
	for _, link := range tmp {
		u, err = url.Parse(link)
		if err != nil {
			continue
		}
		if u.Host == host && !strings.ContainsAny(u.Path, "#?") {
			list = append(list, link)
		}
	}
	return list
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
