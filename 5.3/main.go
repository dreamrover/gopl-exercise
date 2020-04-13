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
	for _, s := range visit(nil, doc) {
		fmt.Printf("%s", s)
	}
}

func visit(contents []string, n *html.Node) []string {
	//print(n.Type)
	if n.Type == html.TextNode {
		contents = append(contents, n.Data)
		//fmt.Println(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.TextNode && c.Data == "script" {
			continue
		}
		contents = visit(contents, c)
	}
	return contents
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
