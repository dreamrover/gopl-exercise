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
		fmt.Printf("%s\n", s)
	}
}

func visit(links []string, n *html.Node) []string {
	//print(n.Type)
	loc := map[string]string{
		"a":      "href",
		"img":    "src",
		"script": "src",
		"link":   "href",
	}
	if n.Type == html.ElementNode && loc[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == loc[n.Data] {
				links = append(links, a.Val)
			}
		}
		//fmt.Println(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
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
