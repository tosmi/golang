package main

import (
	"flag"
	"fmt"
	"os"
	"checkcs"
)

func main() {
	var testserver bool
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage %s: [options|url]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, " --testserver: Start a simple testserver (default false)\n")

	}

	flag.BoolVar(&testserver, "testserver", false, "Start a simple testserver")

	flag.Parse()

	if testserver {
		fmt.Printf("Starting the testserver...\n")
		checkcs.Startserver()
	} else {
		checkurl()
	}
}

func checkurl() {
	// url := "http://localhost:8000"
	// res := Service{}

	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 	os.Exit(1)

	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "read: %v\n", err)
	// 	os.Exit(1)

	// }
	// resp.Body.Close()

	// fmt.Fprintf("%v", string(body))
	
}
