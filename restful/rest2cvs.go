package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Service struct {
	Project     string
	Application string
	Lane        string
	CIID        int
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
