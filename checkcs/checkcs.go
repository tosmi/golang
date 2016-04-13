package main

import (
	"checkcs"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var testserver bool
	var testok bool
	var testfailed bool
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage %s: [options|url]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, " --testserver: Start a simple testserver (default false)\n")
		fmt.Fprintf(os.Stderr, " --testok: Test the ok service on the testserver (optional)\n")
		fmt.Fprintf(os.Stderr, " --testfailed: Test the ok service on the testserver (optional)\n")
	}

	flag.BoolVar(&testserver, "testserver", false, "Start a simple testserver")
	flag.BoolVar(&testok, "testok", false, "Test the ok service on the testserver")
	flag.BoolVar(&testfailed, "testfailed", false, "Test the failed service on the testserver")
	flag.Parse()

	if testserver {
		fmt.Printf("Starting the testserver...\n")
		checkcs.Startserver()
	} else if testok {
		checkurl("http://localhost:8000/ok")

	} else if testfailed {
		checkurl("http://localhost:8000/failed")
	} else {
		checkurl(os.Args[1])
	}
}

func checkurl(url string) {
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

	var checkdata checkcs.Check
	if err := json.Unmarshal(body, &checkdata); err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal: %v\n", err)
	}

	fmt.Printf("%v", checkdata.Code)
}
