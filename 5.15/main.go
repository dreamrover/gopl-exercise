package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s integer1 [integer2] ...\n", os.Args[0])
		return
	}
	var ints []int
	for _, a := range os.Args[1:] {
		i, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Input error:", a, "is not an integer!")
			return
		}
		ints = append(ints, i)
	}
	fmt.Println("max: ", max(ints[0], ints[1:]...))
	fmt.Println("min: ", min(ints[0], ints[1:]...))
}

func max(val1 int, vals ...int) int {
	ret := val1
	for _, val := range vals {
		if val > ret {
			ret = val
		}
	}
	return ret
}

func min(val1 int, vals ...int) int {
	ret := val1
	for _, val := range vals {
		if val < ret {
			ret = val
		}
	}
	return ret
}
