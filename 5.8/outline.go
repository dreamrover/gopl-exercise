// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s url id\n", os.Args[0])
		return
	}
	outline(os.Args[1], os.Args[2])
}

func outline(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	node := ElementByID(doc, id)
	//!-call
	fmt.Println(node.Data, node.Attr)

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	var node *html.Node
	if pre != nil {
		if pre(n, id) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				tmp := forEachNode(c, id, pre, post)
				if tmp != nil {
					node = tmp
					break
				}
			}
		} else {
			node = n
		}
	}

	if post != nil {
		post(n, id)
	}
	return node
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		//fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
}

//!-startend
