package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s url tagname1 [tagname2] ...\n", os.Args[0])
		return
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, node := range ElementsByTagName(doc, os.Args[2:]...) {
		fmt.Println(*node)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if tag == doc.Data {
				nodes = append(nodes, doc)
				break
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, name...)...)
	}
	return nodes
}
