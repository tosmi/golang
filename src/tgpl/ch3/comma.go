package main

import (
	"bytes"
	"fmt"
	"os"
)

// comma inserts commas in a non-negative decimal integer string.
func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func commaNonRecursive(s string) string {
	var buf bytes.Buffer
	// 38,388,293


	n := len(s)

	fmt.Printf("mod %v\n", n%3)	
	if n <= 3 {
		return s
	}

	i := 0
	for c := n-1; c >= 0; c-- {

		
		if c >= 3 {
			buf.WriteByte(s[c])
			i++
		} else {
			buf.WriteByte(s[c])
		}

		if i > 2 {
			buf.WriteString(",")
			i = 0
		}
	}
	fmt.Printf("after comma [%v]\n", buf.String())
	r := revert(buf.String())
	return r
}

func revert(s string) string {
	var buf bytes.Buffer
	for i := len(s) - 1;  i >= 0; i-- {
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	//fmt.Printf("%[v]\n", os.Args)
	fmt.Printf("hello\n")
	fmt.Printf("[%v]\n", len(os.Args))
	if len(os.Args) < 2 {
		panic("please provide a number")
	}
	fmt.Printf("[%v]\n", commaRecursive(os.Args[1]))
	fmt.Printf("[%v]\n", commaNonRecursive(os.Args[1]))
}
