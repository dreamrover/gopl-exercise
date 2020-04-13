package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	//"exercise/7.4/my"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s URL\n", os.Args[0])
		return
	}
	words, images, _ := CountWordsAndImages(os.Args[1])
	fmt.Println(words, "words,", images, "images")
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		//input := bufio.NewScanner(bytes.NewBufferString(n.Data))
		//input := bufio.NewScanner(my.NewReader(n.Data))
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wds, imgs := countWordsAndImages(c)
		words += wds
		images += imgs
	}
	return
}
