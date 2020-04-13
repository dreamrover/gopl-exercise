package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
