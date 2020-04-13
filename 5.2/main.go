package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "populate: %v\n", err)
		os.Exit(1)
	}
	names := make(map[string]int)
	fmt.Println(visit(names, doc))
}

func visit(names map[string]int, n *html.Node) map[string]int {
	for node := n; node != nil; node = node.NextSibling {
		if node.Type == html.ElementNode {
			names[node.Data]++
		}
		names = visit(names, node.FirstChild)
	}
	return names
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
