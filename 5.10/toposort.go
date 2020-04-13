// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

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
	order := topoSort(prereqs)
	sorted := make([]string, len(order), len(order))
	for k, v := range order {
		sorted[v] = k
		delete(order, k)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
		for _, pre := range prereqs[course] {
			found := false
			for _, p := range sorted[:i] {
				if p == pre {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("%s is not before %s!\n", pre, course)
			}
		}
	}
}

func topoSort(m map[string][]string) map[string]int {
	seen := make(map[string]bool)
	var visitAll func(items []string)
	var i int
	order := make(map[string]int)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[item] = i
				i++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}

//!-main
