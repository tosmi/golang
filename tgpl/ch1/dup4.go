package main

import (
	"bufio"
	"fmt"
	"os"
)

type counter struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]counter)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(f, counts, arg)
			f.Close()
		}
	}

	for line, counter := range counts {
		if counter.Count > 1 {
			fmt.Printf("%d\t%s: %q\n", counter.Count, line, counter.Filenames)
		}

	}
}

func countlines(f *os.File, counts map[string]counter, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if count,ok := counts[line]; ok {
			count.Count++
			if ok := haselement(count.Filenames, filename); !ok {
				count.Filenames = append(count.Filenames, filename)
			}
			counts[line] = count

		} else {
			var count counter
			count.Count++
			count.Filenames = append(count.Filenames, filename)
			counts[line] = count
		}
	}
}

func haselement(strings []string, element string) (bool) {
	for _, string := range strings {
		fmt.Printf("%v == %v\n", string, element)
		if string == element {
			return true
		}
	}
	return false
}
