package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func do_range() {
	fmt.Println(os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}

}

func do_join() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func do_loop() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	start := time.Now()
	do_range()
	duration := time.Now().Sub(start)
	fmt.Printf("range took %v\n", duration)

	start = time.Now()
	do_join()
	duration = time.Now().Sub(start)
	fmt.Printf("join took %v\n", duration)

	start = time.Now()
	do_loop()
	duration = time.Now().Sub(start)
	fmt.Printf("loop took %v\n", duration)
}
