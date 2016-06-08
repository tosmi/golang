package main

import (
	"fmt"
	"tgpl/ch2/popcount"
)

func main() {
	var x uint64

	x = 34534
	fmt.Printf("%v: %v\n",x, popcount.Popcount(x))
	fmt.Printf("%v: %v\n",x, popcount.PopcountLoop(x))
}
