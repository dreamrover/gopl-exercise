package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s string n\n", os.Args[0])
	}
	s := []byte(os.Args[1])
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n%s\n", os.Args[1], rotate(s, n))
}

func rotate(s []byte, n int) []byte {
	l := len(s)
	if l == 0 || n%l == 0 {
		return s
	}
	m := n % l
	if m < 0 {
		m += l
	}
	//s = append(s, make([]byte, m)...)
	s1 := string(s[l-m:])
	for i := l - 1; i >= m; i-- {
		s[i] = s[i-m]
	}
	for i := 0; i < len(s1); i++ {
		s[i] = s1[i]
	}
	return s
}
