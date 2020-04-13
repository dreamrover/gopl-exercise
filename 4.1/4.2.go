package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var pc [256]byte = func() [256]byte {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return pc
}()

var a = flag.Int("a", 256, "Algorithm")

func main() {
	var diff int
	//var sha1, sha2 []byte
	var s1, s2 []byte
	//var s = [4]byte{'a', 'b', 0, 'd'}
	//var d []byte = []byte(s[1:])
	//fmt.Printf("%x\n", d)
	if len(os.Args) < 3 {
		fmt.Println("usage:", os.Args[0], "[-a] [256|384|512] string1 string2")
		return
	}
	flag.Parse()
	//sha1 = string(d)
	//sha2 = string(s[:])
	switch *a {
	case 256:
		sha1 := sha256.Sum256([]byte(flag.Arg(0)))
		sha2 := sha256.Sum256([]byte(flag.Arg(1)))
		//fmt.Printf("%T:%x\n%T%x\n", sha1, sha1[1], sha2, sha2[1])
		diff = shaDiff(sha1[:], sha2[:])
		s1 = []byte(sha1[:])
		s2 = []byte(sha2[:])
	case 384:
		sha1 := sha512.Sum384([]byte(flag.Arg(0)))
		s1 = []byte(sha1[:])
		sha2 := sha512.Sum384([]byte(flag.Arg(1)))
		s2 = []byte(sha2[:])
	case 512:
		sha1 := sha512.Sum512([]byte(flag.Arg(0)))
		s1 = []byte(sha1[:])
		sha2 := sha512.Sum512([]byte(flag.Arg(1)))
		s2 = []byte(sha2[:])
	}
	fmt.Printf("%x\n%x\n%d\n", s1, s2, diff)
	fmt.Printf("%x\n%x\n%d\n", s1, s2, shaDiff(s1, s2))
}

func shaDiff(sha1, sha2 []byte) int {
	var diff int
	fmt.Printf("%x\n%x\n", sha1, sha2)
	if string(sha1) == string(sha2) {
		return 0
	}
	for i := 0; i < len(sha1); i++ {
		diff += int(pc[sha1[i]^sha2[i]])
	}
	return diff
}
