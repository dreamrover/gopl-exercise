package main

import "fmt"

func main() {
	defer func() {
		e := recover()
		fmt.Println(e)
	}()

	f()
	//go f()
	//time.Sleep(time.Second)
}

func f() {
	panic(fmt.Errorf("test"))
}
