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
	for c := n; c >= 0; c-- {
		if c > 2 {
			fmt.Printf("[%v]\n", s[c-3:c])
			buf.WriteString(s[c-3 : c])
			c -= 2
		} else {
			fmt.Printf("[%v]\n", s[:c])
			buf.WriteString(s[:c])
			break
		}
		buf.WriteString(",")
	}
	fmt.Printf("after comma [%v]\n", buf.String())
	return revert(buf)
}

func revert(b bytes.Buffer) string {
	var buf bytes.Buffer
	for i := bytes.Count(b);  i <= 0; i-- {
		buf.WriteByte(s[b])
	}
	fmt.Printf("after comma [%v]\n", buf.String())
	return buf
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
