// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"linear algebra"},
	"intro to programming": {"networks"}, // for loop

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string, map[string]bool) bool

	visitAll = func(items []string, posts map[string]bool) (loop bool) {
		for _, item := range items {
			if posts[item] {
				loop = true
				fmt.Printf("loop: %s", item)
				break
			}
			if !seen[item] {
				seen[item] = true
				p := make(map[string]bool)
				for k, v := range posts {
					p[k] = v
				}
				p[item] = true
				loop = visitAll(m[item], p)
				if loop {
					fmt.Printf(" <- %s", item)
					break
				}
				order = append(order, item)
			}
		}
		return
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	//sort.Strings(keys)
	visitAll(keys, nil)
	fmt.Println()
	return order
}

//!-main
