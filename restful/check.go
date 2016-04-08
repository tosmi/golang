package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func flags() {
	flag.Bool("server", false, "start the test server")
	flag.Parse()
}


func main() {
	url := "http://localhost:8000"
	res := Service{}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)

	}
	resp.Body.Close()

	fmt.Fprintf("%v", string(body))
}
